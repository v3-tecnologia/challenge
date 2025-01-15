package Devices

import Datas.GyroscopeData
import android.hardware.Sensor
import android.hardware.SensorEvent
import android.hardware.SensorEventListener
import android.hardware.SensorManager
import android.util.Log

class GyroscopeSensor(private val sensorManager: SensorManager) {
    private val gyroscope: Sensor? = sensorManager.getDefaultSensor(Sensor.TYPE_GYROSCOPE)

    fun read(callback: (gyroscopedata: GyroscopeData) -> Unit) {
        if(gyroscope == null){
            Log.e(javaClass.simpleName, "Gyroscope sensor is not avaiable")
            return
        }

        val sensorEventListener = object : SensorEventListener {
            override fun onSensorChanged(event: SensorEvent){
                if(event != null){
                    callback(
                        GyroscopeData(
                            id = 0,
                            timestamp = System.currentTimeMillis(),
                            deviceId = "",
                            x = event.values[0],
                            y = event.values[1],
                            z = event.values[2]
                        )
                    )
                    sensorManager.unregisterListener(this)
                }
            }

            override fun onAccuracyChanged(sensor: Sensor?, accuracy: Int) { }
        }

        sensorManager.registerListener(sensorEventListener, gyroscope, SensorManager.SENSOR_DELAY_NORMAL)
    }

}