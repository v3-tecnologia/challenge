package com.example.v3mvp.remote

import com.example.v3mvp.api.ColetaRequest
import com.example.v3mvp.api.GPS
import com.example.v3mvp.api.Giroscopio
import com.example.v3mvp.model.Coleta
import okhttp3.ResponseBody
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response

object ColetaRemoteDataSource {
    fun enviar(
        coleta: Coleta,
        fotoBase64: String,
        onSuccess: () -> Unit,
        onError: (String) -> Unit
    ) {
        val request = ColetaRequest(
            idDispositivo = coleta.deviceId ?: "",
            timestamp = coleta.timestamp,
            giroscopio = Giroscopio(
                x = coleta.gyroX?.toDouble() ?: 0.0,
                y = coleta.gyroY?.toDouble() ?: 0.0,
                z = coleta.gyroZ?.toDouble() ?: 0.0
            ),
            gps = GPS(
                latitude = coleta.latitude ?: 0.0,
                longitude = coleta.longitude ?: 0.0
            ),
            fotoBase64 = fotoBase64
        )

        val service = RetrofitInstance.retrofit.create(ApiService::class.java)

        service.enviarColeta(request)
            .enqueue(object : Callback<ResponseBody> {
                override fun onResponse(call: Call<ResponseBody>, response: Response<ResponseBody>) {
                    if (response.isSuccessful) {
                        onSuccess()
                    } else {
                        onError("Erro ${response.code()} - ${response.errorBody()?.string()}")
                    }
                }

                override fun onFailure(call: Call<ResponseBody>, t: Throwable) {
                    onError(t.message ?: "Erro desconhecido")
                }
            })
    }
}
