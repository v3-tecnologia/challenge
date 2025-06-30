from django.core.files.uploadedfile import SimpleUploadedFile
from rest_framework.test import APITestCase
from . import models

class DeviceModelTest(APITestCase):
    def test_device_creation(self):
        device = models.Device.objects.create(mac='00:1B:44:11:3A:B7')
        self.assertEqual(device.mac, '00:1B:44:11:3A:B7')
        self.assertEqual(models.Device.objects.count(), 1)

class GyroscopeModelTest(APITestCase):
    def setUp(self):
        self.device = models.Device.objects.create(mac='00:1B:44:11:3A:B7')

    def test_gyroscope_creation(self):
        gyro = models.Gyroscope.objects.create(
            x=1.1,
            y=2.2,
            z=3.3,
            device=self.device
        )
        self.assertEqual(gyro.x, 1.1)
        self.assertEqual(gyro.y, 2.2)
        self.assertEqual(gyro.z, 3.3)
        self.assertEqual(gyro.device, self.device)
        self.assertIsNotNone(gyro.moment)  # auto_now_add

    def test_gyroscope_device_relationship(self):
        gyro = models.Gyroscope.objects.create(x=0, y=0, z=0, device=self.device)
        self.assertEqual(gyro.device.mac, '00:1B:44:11:3A:B7')


class GPSDataModelTest(APITestCase):
    def setUp(self):
        self.device = models.Device.objects.create(mac='00:1B:44:11:3A:B7')

    def test_gpsdata_creation(self):
        gps = models.GPSData.objects.create(
            latitude=-23.55052,
            longitude=-46.633308,
            device=self.device
        )
        self.assertEqual(gps.latitude, -23.55052)
        self.assertEqual(gps.longitude, -46.633308)
        self.assertEqual(gps.device, self.device)
        self.assertIsNotNone(gps.moment)

class PhotoModelTest(APITestCase):
    def setUp(self):
        self.device = models.Device.objects.create(mac='00:1B:44:11:3A:B7')

    def test_photo_upload(self):
        test_image = SimpleUploadedFile(
            name='test_image.jpg',
            content=b'\x47\x49\x46\x38\x89\x61',  # Simula conteúdo binário de imagem
            content_type='image/jpeg'
        )
        photo = models.Photo.objects.create(
            photo=test_image,
            device=self.device
        )
        self.assertTrue(photo.photo.name.startswith('telemetry/photos/test_image'))
        self.assertEqual(photo.device, self.device)
        self.assertIsNotNone(photo.moment)
