package Services

import Database.EmbarcadoDatabase
import Database.GpsRepository
import Datas.GpsData
import Datas.GyroscopeData
import Devices.GPSSensor
import Network.ApiService
import android.content.Context
import android.util.Log
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

class GPSService(private val context: Context, database: EmbarcadoDatabase) : IService {
    private val gpsSensor = GPSSensor(context)
    private val gpsRepository = GpsRepository(database.gpsDataDao())
    private val api = ApiService()

    override fun Execute() {
        gpsSensor.read {
            gpsData ->

            if(gpsData != null){
                Log.d(gpsData.javaClass.simpleName, "gps data: ${gpsData.toString()}")
                saveOnDatabase(gpsData)
                sendToAPI()
            }
        }
    }

    private fun saveOnDatabase(data: GpsData) {
        CoroutineScope(Dispatchers.IO).launch {
            gpsRepository.insertGpsData(data)
            Log.d(gpsRepository.javaClass.simpleName, "Inserted gps data ${data.toString()}")
        }
    }

    private fun sendToAPI() {
        CoroutineScope(Dispatchers.Default).launch {
            var gpsDatas: List<GpsData> = gpsRepository.GetAllGpsData()
            Log.d("Total GPS Data in database", gpsDatas.count().toString())
            for(data in gpsDatas){
                try {
                    //api.sendGpsData(data)
                    Log.d(javaClass.simpleName, "data: ${data.toString()}")
                }
                catch(ex: Exception){
                    Log.e(ex.javaClass.simpleName, "Erro ao enviar dados ${ex.message}")
                }
            }
        }
    }
}