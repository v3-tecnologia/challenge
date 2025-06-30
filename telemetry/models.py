from django.db import models

class Device(models.Model):
    mac = models.CharField(max_length=17)

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
    photo = models.ImageField('telemetry', upload_to='telemetry/photos')
    moment = models.DateTimeField(auto_now_add=True)
    device = models.ForeignKey(Device, on_delete=models.CASCADE)
