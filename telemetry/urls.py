from django.urls import path
from . import views
from .views import PhotoApiView

urlspatterns = [
    path('gyroscope', views.GyroscopeApiView.as_view(), name='gyroscope'),
    path('gps', views.GPSDataApiView.as_view(), name='gps'),
    path('photo', PhotoApiView.as_view(), name='photo'),
]