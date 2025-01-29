package com.renascienza.desafiov3.db

import com.renascienza.desafiov3.api.*
import com.renascienza.desafiov3.camera.PhotoCollector
import com.renascienza.desafiov3.location.GpsCollector
import com.renascienza.desafiov3.network.MacAddressCollector
import com.renascienza.desafiov3.sensor.GyroscopeCollector
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

class Repository(private val context: Context) {
    private val gyroscopeCollector = GyroscopeCollector(context)
    private val gpsCollector = GpsCollector(context)
    private val photoCollector = PhotoCollector(context)
    private val macAddressCollector = MacAddressCollector(context)

    fun collectData() {
        gyroscopeCollector.startCollecting()

        val timestamp = System.currentTimeMillis()
        val macAddress = macAddressCollector.getMacAddress() ?: return

        // Coletar dados do giroscópio
        val gyroscopeData = gyroscopeCollector.getGyroscopeData()
        if (gyroscopeData != null) {
            val gyroscopeModel = GyroscopeDataModel(
                macAddress = macAddress,
                timestamp = timestamp,
                x = gyroscopeData.x,
                y = gyroscopeData.y,
                z = gyroscopeData.z
            )
            CoroutineScope(Dispatchers.IO).launch {
                Api.sendGyroscopeData(gyroscopeModel)
            }
        }

        // Coletar dados do GPS
        gpsCollector.getLocation { gpsData ->
            if (gpsData != null) {
                val gpsModel = GpsDataModel(
                    macAddress = macAddress,
                    timestamp = timestamp,
                    latitude = gpsData.latitude,
                    longitude = gpsData.longitude,
                    altitude = gpsData.altitude
                )
                CoroutineScope(Dispatchers.IO).launch {
                    Api.sendGpsData(gpsModel)
                }
            }
        }

        // Coletar foto
        photoCollector.capturePhoto { photoBase64 ->
            val photoModel = PhotoDataModel(
                macAddress = macAddress,
                timestamp = timestamp,
                photoBase64 = photoBase64
            )
			
			//Se não houver Photo de rosto capturada, não enviar
			if(!photoModel.photoBase64.isEmpty()){
				CoroutineScope(Dispatchers.IO).launch {
                	Api.sendPhotoData(photoModel)
            	}
			}
            
        }
    }
}