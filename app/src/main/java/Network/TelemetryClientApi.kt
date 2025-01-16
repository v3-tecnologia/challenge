package Network

import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory

object TelemetryClientApi {
    private val telemetryApi = Retrofit.Builder()
        .baseUrl("http://localhost:8085")
        .addConverterFactory(GsonConverterFactory.create())
        .build()

    val api: IApiService = telemetryApi.create(IApiService::class.java)
}