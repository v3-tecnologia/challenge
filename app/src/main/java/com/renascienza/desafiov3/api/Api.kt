package com.renascienza.desafiov3.api

import android.util.Log
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import retrofit2.http.Body
import retrofit2.http.POST

object Api {
    private val retrofit = Retrofit.Builder()
        .baseUrl("http://server.com/") // Substitua pela URL base do servidor
        .addConverterFactory(GsonConverterFactory.create())
        .build()

    private val service = retrofit.create(ApiService::class.java)

    suspend fun sendGyroscopeData(data: GyroscopeDataModel) {
        try {
            withContext(Dispatchers.IO) {
                val response = service.sendGyroscopeData(data)
                if (response.isSuccessful) {
                    Log.d("Api", "Dados do giroscópio enviados com sucesso!")
                } else {
                    Log.e("Api", "Erro ao enviar dados do giroscópio: ${response.message()}")
                }
            }
        } catch (e: Exception) {
            Log.e("Api", "Erro na comunicação com a API (Giroscópio): ${e.message}")
        }
    }

    suspend fun sendGpsData(data: GpsDataModel) {
        try {
            withContext(Dispatchers.IO) {
                val response = service.sendGpsData(data)
                if (response.isSuccessful) {
                    Log.d("Api", "Dados do GPS enviados com sucesso!")
                } else {
                    Log.e("Api", "Erro ao enviar dados do GPS: ${response.message()}")
                }
            }
        } catch (e: Exception) {
            Log.e("Api", "Erro na comunicação com a API (GPS): ${e.message}")
        }
    }

    suspend fun sendPhotoData(data: PhotoDataModel) {
        try {
            withContext(Dispatchers.IO) {
                val response = service.sendPhotoData(data)
                if (response.isSuccessful) {
                    Log.d("Api", "Foto enviada com sucesso!")
                } else {
                    Log.e("Api", "Erro ao enviar foto: ${response.message()}")
                }
            }
        } catch (e: Exception) {
            Log.e("Api", "Erro na comunicação com a API (Foto): ${e.message}")
        }
    }

    interface ApiService {
        @POST("telemetry/gyroscope")
        suspend fun sendGyroscopeData(@Body data: GyroscopeDataModel): retrofit2.Response<Unit>

        @POST("telemetry/gps")
        suspend fun sendGpsData(@Body data: GpsDataModel): retrofit2.Response<Unit>

        @POST("telemetry/photo")
        suspend fun sendPhotoData(@Body data: PhotoDataModel): retrofit2.Response<Unit>
    }
}

data class GyroscopeDataModel(
    val macAddress: String, // Endereço MAC
    val timestamp: Long, // Timestamp da coleta
    val x: Float, // Dado do eixo X
    val y: Float, // Dado do eixo Y
    val z: Float // Dado do eixo Z
)

data class GpsDataModel(
    val macAddress: String, // Endereço MAC
    val timestamp: Long, // Timestamp da coleta
    val latitude: Double, // Latitude
    val longitude: Double, // Longitude
    val altitude: Double // Altitude
)

data class PhotoDataModel(
    val macAddress: String, // Endereço MAC
    val timestamp: Long, // Timestamp da coleta
    val photoBase64: String // Foto codificada em Base64
)