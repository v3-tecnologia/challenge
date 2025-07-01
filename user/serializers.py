from rest_framework.serializers import ModelSerializer

from user import models

class UserSerializer(ModelSerializer):
    class Meta:
        model = models.User
        fields = ['username', 'email', 'is_staff', 'is_active', 'date_joined', 'last_login']

class HistoricalSerializer(ModelSerializer):
    class Meta:
        model = models.Historical
        fields = '__all__'