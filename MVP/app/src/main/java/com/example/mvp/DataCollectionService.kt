package com.example.mvp

import android.app.Service
import android.content.Intent
import android.os.Binder
import android.os.IBinder
import android.util.Log
import com.example.mvp.entity.TelemetryData
import com.example.mvp.model.Face
import com.example.mvp.model.Gyroscrope
import com.example.mvp.model.Location
import com.example.mvp.model.Photo
import com.example.mvp.repository.AppDatabase
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch
import retrofit2.Response
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import retrofit2.http.Body
import retrofit2.http.POST

interface ApiService {
    @POST("/telemetry/gyroscope")
    suspend fun sendGyroscopeData(@Body data: Gyroscrope): Response<Void>

    @POST("/telemetry/gps")
    suspend fun sendGpsData(@Body data: Location): Response<Void>

    @POST("/telemetry/photo")
    suspend fun sendPhoto(@Body data: Photo): Response<Void>

    @POST("/telemetry/face")
    suspend fun sendFace(@Body data: Face): Response<Void>
}

class DataCollectionService : Service() {

    private val dataCollector = DataCollector(this)
    private val binder = LocalBinder()

    val retrofit = Retrofit.Builder()
        .baseUrl("http://192.168.1.100:8080")
        .addConverterFactory(GsonConverterFactory.create())
        .build()

    val apiService = retrofit.create(ApiService::class.java)

    inner class LocalBinder : Binder() {
        // Retorna a instância do serviço para que os clientes possam chamar métodos públicos
        fun getService(): DataCollectionService = this@DataCollectionService
    }

    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        Log.i("DataCollectionService", "entrou no metodo da coleta de dados")

        CoroutineScope(Dispatchers.IO).launch {
            while (true) {
                try {
                    Log.i("DataCollectionService entrou dentro do loop da data", "entrou no metodo da coleta de dados")
                    val data = dataCollector.collectData()
                    Log.i("DataCollectionService tem data", "tem data agora")
                    sendDataToServer(data)
                    saveDataToDatabase(data);
                    Log.i("DataCollectionService terminou de enviar os dados", "terminou de enviar os dados")
                } catch (e: Exception) {
                    Log.e("DataCollectionService", "Erro na coleta de dados 2", e)
                    Log.e("Exception", e.toString(), e)
                }
                delay(2000) // Aguardar 10 segundos antes da próxima iteração
            }
        }

        return START_REDELIVER_INTENT
    }

    override fun onBind(p0: Intent?): IBinder? {
        return binder
    }

    private fun logRequest(response: Response<Void>, requestName: String){
        if (response.isSuccessful) {
            // Dados enviados com sucesso
            Log.d("DataCollector", "Dados enviados com sucesso, ${requestName}")
        } else {
            // Erro ao enviar os dados
            Log.e("DataCollector", "Erro ao enviar os dados, ${requestName}: ${response.code()} - ${response.message()}")
        }
    }

    private suspend fun saveDataToDatabase(data: TelemetryData) {
        val dao = AppDatabase.getDatabase(this).telemetryDataDao()
        dao.insert(data)
    }

    private suspend fun sendDataToServer(telemetryData: TelemetryData) {
        try {

            logRequest(apiService.sendGyroscopeData(Gyroscrope(
                gyroX = telemetryData.gyroX,
                gyroY = telemetryData.gyroY,
                gyroZ = telemetryData.gyroZ
            )),"gyroscope");

            logRequest(apiService.sendGpsData(Location(
                latitude = telemetryData.latitude,
                longitude = telemetryData.longitude
            )),"gps");

            logRequest(apiService.sendPhoto(Photo(
                photoPath = telemetryData.photoPath
            )),"photo");

            if(telemetryData.faceData != "_"){
                logRequest(apiService.sendFace(Face(
                    faceData = telemetryData.faceData
                )),"face");
            }

        } catch (e: Exception) {
            // Erro na chamada à API
            Log.e("DataCollector", "Erro na chamada à API", e)
        }
    }
}

