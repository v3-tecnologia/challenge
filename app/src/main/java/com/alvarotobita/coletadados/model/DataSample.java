package com.alvarotobita.coletadados.model;

import androidx.room.Entity;
import androidx.room.PrimaryKey;

@Entity
public class DataSample {
    @PrimaryKey(autoGenerate = true)
    public int id;

    public long timestamp;
    public double latitude, longitude;
    public float gyroX, gyroY, gyroZ;
    public String photoPath;
    public String deviceId;
}
