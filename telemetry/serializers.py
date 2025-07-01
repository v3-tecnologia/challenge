from rest_framework import serializers

from telemetry import models

class DeviceSerializer(serializers.ModelSerializer):
    class Meta:
        fields = ['mac']
        model = models.Device

class GyroscopeSerializer(serializers.ModelSerializer):
    device = DeviceSerializer()

    class Meta:
        model = models.Gyroscope
        exclude = ['id']


    def create(self, validated_data):
        device_data = validated_data.pop('device')
        device, _ = models.Device.objects.get_or_create(mac=device_data['mac'])
        return models.Gyroscope.objects.create(device=device, **validated_data)

class GPSDataSerializer(serializers.ModelSerializer):
    device = DeviceSerializer()

    class Meta:
        model = models.GPSData
        exclude = ['id']

    def create(self, validated_data):
        device_data = validated_data.pop('device')
        device, _ = models.Device.objects.get_or_create(mac=device_data['mac'])
        return models.GPSData.objects.create(device=device, **validated_data)

class PhotoSerializer(serializers.ModelSerializer):
    device = DeviceSerializer()

    class Meta:
        model = models.Photo
        exclude = ['id']

    def create(self, validated_data):
        device_data = validated_data.pop('device')
        device, _ = models.Device.objects.get_or_create(mac=device_data['mac'])
        return models.Photo.objects.create(device=device, **validated_data)