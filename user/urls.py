from django.urls import path

from . import views

urlpatterns = [
    path('token', views.DecoratedTokenObtainPairView.as_view(), name='token_obtain_pair'),
    path('token/refresh', views.DecoratedTokenRefreshView.as_view(), name='token_refresh'),
    path('token/blacklist', views.DecoratedTokenBlacklistView.as_view(), name='token_blacklist'),
    path('historical', views.HistoricalAPIView.as_view(), name='historical'),
]