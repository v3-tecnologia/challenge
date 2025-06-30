package com.example.v3mvp.api

import okhttp3.ResponseBody
import retrofit2.Call
import retrofit2.http.Body
import retrofit2.http.POST

data class Giroscopio(val x: Double, val y: Double, val z: Double)
data class GPS(val latitude: Double, val longitude: Double)
data class ColetaRequest(
    val idDispositivo: String,
    val timestamp: Long,
    val giroscopio: Giroscopio,
    val gps: GPS,
    val fotoBase64: String
)

interface ApiService {
    @POST("/api/coletas")
    fun enviarColeta(@Body coleta: ColetaRequest): Call<ResponseBody>
}
