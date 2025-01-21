package com.example.teste;

import android.content.pm.PackageManager;
import android.os.Bundle;

import androidx.activity.EdgeToEdge;
import androidx.appcompat.app.AppCompatActivity;
import androidx.core.app.ActivityCompat;
import androidx.core.content.ContextCompat;
import androidx.core.graphics.Insets;
import androidx.core.view.ViewCompat;
import androidx.core.view.WindowInsetsCompat;
import android.location.Location;
import android.location.LocationManager;
import android.content.Context;

import android.app.AlertDialog;
import android.content.DialogInterface;
import android.content.Intent;
import android.provider.Settings;

import android.hardware.Sensor;
import android.hardware.SensorEvent;
import android.hardware.SensorEventListener;
import android.hardware.SensorManager;

public class MainActivity extends AppCompatActivity {

    private SensorManager sensorManager;
    private LocationManager lm;
    private SensorService sense;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        EdgeToEdge.enable(this);
        setContentView(R.layout.activity_main);

        sensorManager = (SensorManager) getSystemService(Context.SENSOR_SERVICE);
        lm = (LocationManager)getSystemService(Context.LOCATION_SERVICE);
        /*\/ habilitar sensores; */
        sense = new SensorService(sensorManager);

        try {
            if (ContextCompat.checkSelfPermission(getApplicationContext(), android.Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED ) {
                ActivityCompat.requestPermissions(this, new String[]{android.Manifest.permission.ACCESS_FINE_LOCATION}, 101);
            }

//            getLocation();
            getGyros();
        } catch (Exception e){
            e.printStackTrace();
        }


        ViewCompat.setOnApplyWindowInsetsListener(findViewById(R.id.main), (v, insets) -> {
            Insets systemBars = insets.getInsets(WindowInsetsCompat.Type.systemBars());
            v.setPadding(systemBars.left, systemBars.top, systemBars.right, systemBars.bottom);
            return insets;
        });

    }

    public void getLocation(){
        MyGpsService gpsTracker = new MyGpsService(MainActivity.this);
        if(gpsTracker.canGetLocation()){
            double latitude = gpsTracker.getLatitude();
            double longitude = gpsTracker.getLongitude();
//            String latlon = String.valueOf(latitude) + String.valueOf(longitude);
//            System.out.println( latlon );
            gpsTracker.showAlert();
        }else{
            gpsTracker.showSettingsAlert();
        }
    }

    public void getGyros(){
        SensorService sen = new SensorService();
//        MyGpsService gpsTracker = new MyGpsService(MainActivity.this);
//        System.out.println( sen.getCordinates() );
//        gpsTracker.showAlert( sen.getCordinates() );
    }

}