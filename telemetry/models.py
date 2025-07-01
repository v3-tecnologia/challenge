from django.db import models
from . import fields

class Device(models.Model):
    mac = fields.EncryptedTextField(max_length=32)

class Gyroscope(models.Model):
    x = models.FloatField()
    y = models.FloatField()
    z = models.FloatField()
    moment = models.DateTimeField(auto_now_add=True)
    device = models.ForeignKey(Device, on_delete=models.CASCADE)

class GPSData(models.Model):
    latitude = models.FloatField()
    longitude = models.FloatField()
    moment = models.DateTimeField(auto_now_add=True)
    device = models.ForeignKey(Device, on_delete=models.CASCADE)

class Photo(models.Model):
    device = models.ForeignKey(Device, on_delete=models.CASCADE)
    photo = models.ImageField('telemetry', upload_to='telemetry/photos')
    moment = models.DateTimeField(auto_now_add=True)
    hash = models.CharField(max_length=64, null=True)
    face_contains = models.BooleanField(default=False)