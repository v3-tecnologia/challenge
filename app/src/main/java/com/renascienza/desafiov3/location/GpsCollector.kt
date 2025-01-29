package com.renascienza.desafiov3.location

import android.content.Context
import android.location.Location
import com.google.android.gms.location.FusedLocationProviderClient
import com.google.android.gms.location.LocationServices
import com.google.android.gms.tasks.Task

class GpsCollector(private val context: Context) {
    private val fusedLocationClient: FusedLocationProviderClient =
        LocationServices.getFusedLocationProviderClient(context)

    fun getLocation(callback: (GpsData?) -> Unit) {
        val task: Task<Location> = fusedLocationClient.lastLocation
        task.addOnSuccessListener { location: Location? ->
            location?.let {
                val gpsData = GpsData(
                    latitude = it.latitude,
                    longitude = it.longitude,
                    altitude = it.altitude
                )
                callback(gpsData)
            } ?: run {
                callback(null)
            }
        }
    }
}

data class GpsData(
    val latitude: Double,
    val longitude: Double,
    val altitude: Double
)