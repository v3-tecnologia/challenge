from rest_framework import status
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.request import  Request

from drf_yasg.utils import swagger_auto_schema
from drf_yasg import openapi

from . import serializers

class GyroscopeApiView(APIView):
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
        gyroscope_serializer.save(moment=None)
        return Response(gyroscope_serializer.data, status=status.HTTP_201_CREATED)

class GPSDataApiView(APIView):
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
        gps_serializer.save(moment=None)
        return Response(gps_serializer.data, status=status.HTTP_201_CREATED)

class PhotoApiView(APIView):
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
        photo_serializer = serializers.PhotoSerializer(data=request.data)
        photo_serializer.is_valid(raise_exception=True)
        photo_serializer.save(moment=None)
        return Response(photo_serializer.data, status=status.HTTP_201_CREATED)
