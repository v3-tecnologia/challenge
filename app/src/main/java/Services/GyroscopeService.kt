package Services

import Database.EmbarcadoDatabase
import Database.GyroscopeRepository
import Datas.GyroscopeData
import Devices.GyroscopeSensor
import Network.ApiService
import android.content.Context
import android.hardware.SensorManager
import android.util.Log
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

class GyroscopeService(private val context: Context, database: EmbarcadoDatabase) : IService {
    private val gyroscopeSensor = GyroscopeSensor(context.getSystemService(Context.SENSOR_SERVICE) as SensorManager)
    private val gyroscopeRepository = GyroscopeRepository(database.gyroscopeDataDao())
    private val api = ApiService()

    override fun Execute() {
        gyroscopeSensor.read{
            gyroscopeData ->

            saveOnDatabase(gyroscopeData)
            sendToAPI()
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

    fun sendToAPI() {
        CoroutineScope(Dispatchers.Default).launch {
            var gyroscopeDatas: List<GyroscopeData> = gyroscopeRepository.GetAllGyroscope()
            Log.d("Total Gyroscope data in database", gyroscopeDatas.count().toString())
            for(data in gyroscopeDatas){
                try {
                    //api.sendGyroscopeData(data)
                    Log.d(javaClass.simpleName, "data: ${data.toString()}")
                }
                catch(ex: Exception){
                    Log.e(ex.javaClass.simpleName, "Erro ao enviar dados ${ex.message}")
                }
            }
        }
    }
}