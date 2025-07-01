from django.urls import path, include

from rest_framework import permissions
from drf_yasg.views import get_schema_view
from drf_yasg import openapi

schema_view = get_schema_view(
   openapi.Info(
      title="Documentação do desafio",
      default_version='v1',
      description="desafio tecnico: desenvolvimento backend de API ",
      contact=openapi.Contact(email="ogirdo.sant@gmail.com"),
      license=openapi.License(name="BSD License"),
   ),
   public=True,
   permission_classes=(permissions.AllowAny,),
)

from telemetry.urls import urlspatterns as telemetry_patterns
from user.urls import urlpatterns as user_patterns
urlpatterns = [
    path('swagger/', schema_view.with_ui('swagger', cache_timeout=0), name='schema-swagger-ui'),
    path('redoc/', schema_view.with_ui('redoc', cache_timeout=0), name='schema-redoc'),

    path('telemetry/', include(telemetry_patterns)),
    path('user/', include(user_patterns)),
]
