import hashlib
import json

from django.contrib.contenttypes.models import ContentType

from rest_framework.permissions import IsAuthenticated
from rest_framework import status, parsers
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.request import  Request

from drf_yasg.utils import swagger_auto_schema
from botocore.exceptions import ClientError

from . import serializers, models
from user import models as user_models

from core.settings import REKOGNATION_CLIENT, FACE_DATABASE_CLIENT_ID, LOGGER


class GyroscopeApiView(APIView):
    permission_classes = [IsAuthenticated]
    @swagger_auto_schema(
        operation_summary="Enviar dados de giroscopio",
        operation_description="Endpoint para enviar dados do giroscópio.",
        request_body=serializers.GyroscopeSerializer,
        responses={
            201: serializers.GyroscopeSerializer,
            400: 'Bad Request - dados inválidos'
        }
    )
    def post(self, request: Request) -> Response:
        gyroscope_serializer = serializers.GyroscopeSerializer(data=request.data)
        gyroscope_serializer.is_valid(raise_exception=True)
        data_model:models.Gyroscope = gyroscope_serializer.save(moment=None)
        user_models.Historical.objects.create(
            user=request.user,
            action="register gyroscope data",
            content_type=ContentType.objects.get_for_model(data_model),
            object_id=data_model.id,
        )

        return Response(gyroscope_serializer.data, status=status.HTTP_201_CREATED)

class GPSDataApiView(APIView):
    permission_classes = [IsAuthenticated]
    @swagger_auto_schema(
        operation_summary="Enviar dados de GPS",
        operation_description="Endpoint para enviar dados de GPS",
        request_body=serializers.GPSDataSerializer,
        responses={
            201: serializers.GPSDataSerializer,
            400: 'Bad Request - dados inválidos'
        }
    )
    def post(self, request: Request) -> Response:
        gps_serializer = serializers.GPSDataSerializer(data=request.data)
        gps_serializer.is_valid(raise_exception=True)
        data_model = gps_serializer.save(moment=None)

        user_models.Historical.objects.create(
            user=request.user,
            action="register gyroscope data",
            content_type=ContentType.objects.get_for_model(data_model),
            object_id=data_model.id,
        )

        return Response(gps_serializer.data, status=status.HTTP_201_CREATED)

class PhotoApiView(APIView):
    permission_classes = [IsAuthenticated]

    @swagger_auto_schema(
        operation_summary="Enviar Fotos",
        operation_description="Endpoint para enviar foto",
        request_body=serializers.PhotoSerializer,
        responses={
            201: serializers.PhotoSerializer,
            400: 'Bad Request - dados inválidos'
        }
    )
    def post(self, request: Request) -> Response:
        device = request.data.get('device')
        photo = request.data.get('photo')

        if not device or not photo:
            return Response({"error": ["requisicao mal formada"]}, status=status.HTTP_400_BAD_REQUEST)

        photo_serializer = serializers.PhotoSerializer(
            data={'device': json.loads(device), 'photo': photo}
        )

        photo_serializer.is_valid(raise_exception=True)

        photo.seek(0)
        photo_bytes = photo.read()

        LOGGER.info("Iniciando processamento da imagem.")

        hash_photo = hashlib.md5(photo_bytes).hexdigest()

        if models.Photo.objects.filter(hash=hash_photo).exists():
            cached_model: models.Photo = models.Photo.objects.filter(hash=hash_photo)[0]

            data_model = photo_serializer.save(moment=None, hash=hash_photo, photo=photo, face_contains=cached_model.face_contains)
            user_models.Historical.objects.create(
                user=request.user,
                action="register gyroscope data",
                content_type=ContentType.objects.get_for_model(data_model),
                object_id=data_model.id,
            )

            return Response({**photo_serializer.data, **{"find": "cache"}}, status=status.HTTP_201_CREATED)


        try:
            REKOGNATION_CLIENT.search_faces_by_image(
                CollectionId=FACE_DATABASE_CLIENT_ID,
                Image={'Bytes': photo_bytes},
                FaceMatchThreshold=90,
                MaxFaces=1
            )
            data_model = photo_serializer.save(moment=None, hash=hash_photo, photo=photo, face_contains=True)

        except ClientError:
            LOGGER.error('Rosto não encontrado')
            data_model = photo_serializer.save(moment=None, hash=hash_photo, photo=photo, face_contains=False)

        user_models.Historical.objects.create(
            user=request.user,
            action="register gyroscope data",
            content_type=ContentType.objects.get_for_model(data_model),
            object_id=data_model.id,
        )

        return Response(photo_serializer.data, status=status.HTTP_201_CREATED)