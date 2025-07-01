package com.alvarotobita.coletadados.sensors;

import android.annotation.SuppressLint;
import android.content.Context;
import android.location.Location;
import android.util.Log;

import com.google.android.gms.location.*;

public class GpsManager {

    private final FusedLocationProviderClient fusedLocationClient;
    private Location lastLocation;

    public GpsManager(Context context) {
        fusedLocationClient = LocationServices.getFusedLocationProviderClient(context);
    }

    @SuppressLint("MissingPermission")
    public void getLocation(LocationCallback callback) {
        LocationRequest locationRequest = LocationRequest.create()
                .setPriority(Priority.PRIORITY_HIGH_ACCURACY)
                .setInterval(10000) // 10s
                .setFastestInterval(5000);

        fusedLocationClient.requestLocationUpdates(locationRequest, callback, null);
    }

    public void stop(LocationCallback callback) {
        fusedLocationClient.removeLocationUpdates(callback);
    }
}
