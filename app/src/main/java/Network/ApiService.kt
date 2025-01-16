package Network

import Datas.GpsData
import Datas.GyroscopeData
import Datas.PhotoData
import android.util.Log
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.launch
import okhttp3.MultipartBody
import okhttp3.RequestBody

class ApiService {

    private val api = TelemetryClientApi.api
    private val scope = CoroutineScope(Dispatchers.IO + SupervisorJob())

    fun sendGyroscopeData(data: GyroscopeData){
        try{
            scope.launch {
                api.sendGyroscopeData(data)
                Log.d(javaClass.simpleName, "Success sending gyroscope data: ${data.toString()}")
            }
        }
        catch(ex: Exception){
            Log.e(javaClass.simpleName, "Error sending gyroscope data: ${data.toString()}")
            Log.e(ex.javaClass.simpleName, "Error message: ${ex.message}")
        }
    }

    fun sendGpsData(data: GpsData){
        try {
            scope.launch{
                api.sendGPSData(data)
                Log.d(javaClass.simpleName, "Success sending gps data: ${data.toString()}")
            }
        }
        catch(ex: Exception){
            Log.e(javaClass.simpleName, "Error sending gps data: ${data.toString()}")
            Log.e(ex.javaClass.simpleName, "Error message: ${ex.message}")
        }
    }

    fun sendPhoto(data: PhotoData) {
        try {
            scope.launch{
                val photoPart = MultipartBody.Part.createFormData(
                    data.photoPath,
                    "image.jpg",
                    RequestBody.create(MultipartBody.FORM, ByteArray(1024))
                )
                api.sendPhoto(photoPart)
                Log.d(javaClass.simpleName, "Success sending photo data: ${data.toString()}")
            }
        }
        catch(ex: Exception){
            Log.e(javaClass.simpleName, "Error sending photo data: ${data.toString()}")
            Log.e(ex.javaClass.simpleName, "Error message: ${ex.message}")
        }
    }
}