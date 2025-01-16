package Devices

import Datas.GpsData
import android.Manifest
import android.content.Context
import android.content.pm.PackageManager
import android.location.Location
import android.util.Log
import androidx.core.app.ActivityCompat
import com.google.android.gms.location.FusedLocationProviderClient
import com.google.android.gms.location.LocationServices
import com.google.android.gms.tasks.OnSuccessListener

class GPSSensor(private val context: Context) {

    private val fusedLocationClient: FusedLocationProviderClient =
        LocationServices.getFusedLocationProviderClient(context)

    fun read(callback: (gpsData: GpsData?) -> Unit) {
        if (ActivityCompat.checkSelfPermission(
                context,
                Manifest.permission.ACCESS_FINE_LOCATION
            ) != PackageManager.PERMISSION_GRANTED &&
            ActivityCompat.checkSelfPermission(
                context,
                Manifest.permission.ACCESS_COARSE_LOCATION
            ) != PackageManager.PERMISSION_GRANTED
        ) {
            Log.e(javaClass.simpleName, "Location permissions are not granted")
            return
        }

        fusedLocationClient.lastLocation.addOnSuccessListener(object: OnSuccessListener<Location> {
            override fun onSuccess(location: Location?) {
                if(location != null){
                    val gpsData = GpsData(
                        id = 0,
                        timestamp = System.currentTimeMillis(),
                        deviceId = "",
                        longitude = location.longitude,
                        latitude = location.latitude
                    )
                    callback(gpsData)
                }
                else {
                    Log.e(javaClass.simpleName, "Error reading gps data")
                    callback(null)
                }
            }
        })

    }
}