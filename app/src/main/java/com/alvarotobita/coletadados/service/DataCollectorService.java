package com.alvarotobita.coletadados.service;

import android.app.*;
import android.content.Intent;
import android.os.*;
import androidx.annotation.Nullable;
import androidx.core.app.NotificationCompat;
import android.util.Log;
import java.util.concurrent.*;
import com.alvarotobita.coletadados.sensors.GpsManager;
import com.alvarotobita.coletadados.sensors.GyroscopeManager;
import android.location.Location;
import com.google.android.gms.location.LocationCallback;
import com.google.android.gms.location.LocationResult;


public class DataCollectorService extends Service {

    private ScheduledExecutorService scheduler;
    private GpsManager gpsManager;
    private GyroscopeManager gyroManager;
    private Location lastKnownLocation;
    private LocationCallback locationCallback;

    @Override
    public void onCreate() {
        super.onCreate();
        createNotificationChannel();
        startForeground(1, buildNotification());
        scheduler = Executors.newSingleThreadScheduledExecutor();

        gpsManager = new GpsManager(this);
        gyroManager = new GyroscopeManager(this);
        gyroManager.start();

// Callback do GPS
        locationCallback = new LocationCallback() {
            @Override
            public void onLocationResult(LocationResult locationResult) {
                if (locationResult != null && !locationResult.getLocations().isEmpty()) {
                    lastKnownLocation = locationResult.getLastLocation();
                }
            }
        };
        gpsManager.getLocation(locationCallback);


        scheduler.scheduleAtFixedRate(this::coletarDados, 0, 10, TimeUnit.SECONDS);
    }

    private void coletarDados() {
        float[] gyro = gyroManager.getLastGyro();

        if (lastKnownLocation != null) {
            Log.d("GPS", "Lat: " + lastKnownLocation.getLatitude() +
                    ", Lon: " + lastKnownLocation.getLongitude());
        }

        Log.d("GYRO", "X: " + gyro[0] + ", Y: " + gyro[1] + ", Z: " + gyro[2]);
    }


    @Override
    public int onStartCommand(Intent intent, int flags, int startId) {
        return START_STICKY;
    }

    @Override
    public void onDestroy() {
        super.onDestroy();
        if (scheduler != null) scheduler.shutdownNow();
        gyroManager.stop();
        gpsManager.stop(locationCallback);
    }

    @Nullable
    @Override
    public IBinder onBind(Intent intent) {
        return null;
    }

    private Notification buildNotification() {
        return new NotificationCompat.Builder(this, "coleta_canal")
                .setContentTitle("📡 Coletando dados")
                .setContentText("Serviço de coleta em execução.")
                .setSmallIcon(android.R.drawable.ic_menu_mylocation)
                .build();
    }

    private void createNotificationChannel() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            NotificationChannel channel = new NotificationChannel(
                    "coleta_canal",
                    "Canal de Coleta",
                    NotificationManager.IMPORTANCE_LOW
            );
            NotificationManager manager = getSystemService(NotificationManager.class);
            manager.createNotificationChannel(channel);


        }
    }
}