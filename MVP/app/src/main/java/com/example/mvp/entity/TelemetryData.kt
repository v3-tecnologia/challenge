package com.example.mvp.entity

import androidx.room.ColumnInfo
import androidx.room.Entity
import androidx.room.PrimaryKey
import java.util.Date

@Entity(tableName = "telemetry_data")
data class TelemetryData(
    @PrimaryKey(autoGenerate = true) val id: Int = 0,
    @ColumnInfo(name = "gyro_x") public val gyroX: Float,
    @ColumnInfo(name = "gyro_y") public val gyroY: Float,
    @ColumnInfo(name = "gyro_z") public val gyroZ: Float,
    @ColumnInfo(name = "latitude") val latitude: Double,
    @ColumnInfo(name = "longitude") val longitude: Double,
    @ColumnInfo(name = "photo_path") val photoPath: String,
    @ColumnInfo(name = "face_data") val faceData: String,
    @ColumnInfo(name = "timestamp") val timestamp: Date = Date()
)