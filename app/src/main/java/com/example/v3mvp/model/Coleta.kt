package com.example.v3mvp.model

import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity
data class Coleta(
    @PrimaryKey(autoGenerate = true) val id: Int = 0,
    val timestamp: Long,
    val latitude: Double?,
    val longitude: Double?,
    val gyroX: Float?,
    val gyroY: Float?,
    val gyroZ: Float?,
    val deviceId: String?,
    val fotoPath: String?,
    val status: String?,
    val enviado: Boolean = false
)

