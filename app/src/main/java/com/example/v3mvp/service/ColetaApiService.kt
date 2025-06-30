package com.example.v3mvp.service

import okhttp3.RequestBody
import okhttp3.ResponseBody
import retrofit2.Response
import retrofit2.http.Body
import retrofit2.http.Headers
import retrofit2.http.POST

interface ColetaApiService {
    @Headers("Content-Type: application/x-protobuf")
    @POST("/coleta/sync")
    suspend fun enviarColetaBinaria(@Body requestBody: RequestBody): Response<ResponseBody>
}
