package com.example.mvp

import android.content.Intent
import android.os.Bundle
import android.util.Log
import androidx.appcompat.app.AppCompatActivity

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        val myIntent = Intent(applicationContext, DataCollectionService::class.java)
        try {
            startService(myIntent)
        } catch (e: Exception) {
            Log.e("MainActivity", "Erro na inicialização", e)
        }

    }
}