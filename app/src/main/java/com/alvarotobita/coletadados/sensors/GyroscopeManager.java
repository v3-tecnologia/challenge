package com.alvarotobita.coletadados.sensors;

import android.content.Context;
import android.hardware.Sensor;
import android.hardware.SensorEvent;
import android.hardware.SensorEventListener;
import android.hardware.SensorManager;

public class GyroscopeManager implements SensorEventListener {

    private final SensorManager sensorManager;
    private final Sensor gyroSensor;
    private float[] lastGyro = new float[3];

    public GyroscopeManager(Context context) {
        sensorManager = (SensorManager) context.getSystemService(Context.SENSOR_SERVICE);
        gyroSensor = sensorManager.getDefaultSensor(Sensor.TYPE_GYROSCOPE);
    }

    public void start() {
        sensorManager.registerListener(this, gyroSensor, SensorManager.SENSOR_DELAY_NORMAL);
    }

    public void stop() {
        sensorManager.unregisterListener(this);
    }

    public float[] getLastGyro() {
        return lastGyro;
    }

    @Override
    public void onSensorChanged(SensorEvent event) {
        lastGyro[0] = event.values[0];
        lastGyro[1] = event.values[1];
        lastGyro[2] = event.values[2];
    }

    @Override
    public void onAccuracyChanged(Sensor sensor, int accuracy) {
    }
}
