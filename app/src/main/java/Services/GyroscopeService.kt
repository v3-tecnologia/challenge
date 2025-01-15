package Services

import Database.EmbarcadoDatabase
import Database.GyroscopeRepository
import Datas.GyroscopeData
import Devices.GyroscopeSensor
import android.hardware.SensorManager
import android.util.Log
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

class GyroscopeService(private val sensorManager: SensorManager, database: EmbarcadoDatabase) : IService {
    private val gyroscopeSensor = GyroscopeSensor(sensorManager)
    private val gyroscopeRepository = GyroscopeRepository(database.gyroscopeDataDao())

    override fun Execute() {
        gyroscopeSensor.read{
            gyroscopeData ->

            saveOnDatabase(gyroscopeData)
            sendToAPI(gyroscopeData)
        }
    }

    fun saveOnDatabase(data: GyroscopeData) {
        if(data is GyroscopeData){
            CoroutineScope(Dispatchers.IO).launch {
                gyroscopeRepository.insertGyroscopeData(data)
                Log.d(gyroscopeRepository.javaClass.simpleName, "Inserted data ${data.toString()}")
            }
        }
    }

    fun sendToAPI(data: GyroscopeData) {
        CoroutineScope(Dispatchers.Default).launch {
            var gyroscopeDatas: List<GyroscopeData> = gyroscopeRepository.GetAllGyroscope()
            Log.d("Total in database", gyroscopeDatas.count().toString())
            for(data in gyroscopeDatas){
                Log.d("In Database", data.toString())
            }
        }

    }
}