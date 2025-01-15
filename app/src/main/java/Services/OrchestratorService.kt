package Services

import Database.EmbarcadoDatabase
import android.hardware.SensorManager

class OrchestratorService(private val sensorManager: SensorManager, private val database: EmbarcadoDatabase) {
    private var sensorServices: List<IService> = listOf()

    fun collectAndProcessData() {
        loadSensorServices()

        for(sensor in sensorServices){
            sensor.Execute()
        }
    }

    fun loadSensorServices(){
        sensorServices = listOf(
            GyroscopeService(sensorManager, database)
        )
    }
}