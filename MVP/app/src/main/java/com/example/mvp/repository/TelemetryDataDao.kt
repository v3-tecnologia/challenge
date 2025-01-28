package com.example.mvp.repository

import androidx.room.Dao
import androidx.room.Insert
import androidx.room.Query
import com.example.mvp.entity.TelemetryData

@Dao
interface TelemetryDataDao {
    @Insert
    suspend fun insert(telemetryData: TelemetryData)

    @Query("SELECT * FROM telemetry_data")
    suspend fun getAll(): List<TelemetryData>
}