package com.example.desafioandroidembarcado

import Database.EmbarcadoDatabase
import Services.OrchestratorService
import android.app.Activity
import android.hardware.SensorManager
import android.os.Bundle
import android.os.Handler
import android.os.Looper
import kotlinx.coroutines.Runnable

class MainActivity : Activity() {
    private lateinit var orchestratorServices: OrchestratorService
    private lateinit var database: EmbarcadoDatabase
    private val handler = Handler(Looper.getMainLooper())
    private val intervalMillis: Long = 10_000

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        database = EmbarcadoDatabase.getInstance(applicationContext)
        orchestratorServices = OrchestratorService(getSystemService(SENSOR_SERVICE) as SensorManager, database)

        startSensorExecute()

    }

    private fun startSensorExecute() {
        handler.postDelayed(object: Runnable {
            override fun run() {
                orchestratorServices.collectAndProcessData()
                handler.postDelayed(this, intervalMillis)
            }
        }, intervalMillis)
    }

    override fun onDestroy(){
        super.onDestroy()

        handler.removeCallbacksAndMessages(this)
    }
}

