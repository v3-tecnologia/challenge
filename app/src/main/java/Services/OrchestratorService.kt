package Services

import Database.EmbarcadoDatabase
import android.content.Context
import android.util.Log
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.runBlocking

class OrchestratorService(private val context: Context, private val database: EmbarcadoDatabase) {
    private var sensorServices: List<IService> = listOf()

    fun collectAndProcessData() {
        loadSensorServices()

        runBlocking {
            sensorServices.map { sensor ->
                async {
                    Log.d(sensor.javaClass.simpleName, "Iniciando o Execute")
                    sensor.Execute()
                }
            }.awaitAll()
        }

    }

    fun loadSensorServices(){
        sensorServices = listOf(
            GyroscopeService(context, database),
            GPSService(context, database)
        )
    }
}