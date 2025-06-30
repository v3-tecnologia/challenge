package com.example.v3mvp.remote

import com.example.v3mvp.api.ColetaRequest
import okhttp3.ResponseBody
import retrofit2.Call
import retrofit2.http.Body
import retrofit2.http.POST

interface ApiService {
    @POST("coleta")
    fun enviarColeta(@Body request: ColetaRequest): Call<ResponseBody>
}
