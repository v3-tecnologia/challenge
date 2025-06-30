package com.example.v3mvp.remote

import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory

object RetrofitInstance {
    val retrofit: Retrofit by lazy {
        Retrofit.Builder()
            .baseUrl("http://192.168.0.12:8080/") 
            .addConverterFactory(GsonConverterFactory.create())
            .build()
    }
}
