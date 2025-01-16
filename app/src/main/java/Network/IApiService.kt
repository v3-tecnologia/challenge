package Network

import Datas.GpsData
import Datas.GyroscopeData
import okhttp3.MultipartBody
import retrofit2.http.Body
import retrofit2.http.POST
import retrofit2.http.Part

interface IApiService {
    @POST("/telemetry/gyroscope")
    suspend fun sendGyroscopeData(@Body gyroscopeData: GyroscopeData)

    @POST("/telemetry/gps")
    suspend fun sendGPSData(@Body gpsData: GpsData)

    @POST("/telemetry/photo")
    suspend fun sendPhoto(@Part photo: MultipartBody.Part)
}