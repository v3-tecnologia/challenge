package com.renascienza.desafiov3.sensor

import android.content.Context
import android.hardware.Sensor
import android.hardware.SensorEvent
import android.hardware.SensorEventListener
import android.hardware.SensorManager

class GyroscopeCollector(private val context: Context) : SensorEventListener {
    private var sensorManager: SensorManager? = null
    private var gyroscopeSensor: Sensor? = null
    private var gyroscopeData: GyroscopeData? = null

    init {
        sensorManager = context.getSystemService(Context.SENSOR_SERVICE) as SensorManager
        gyroscopeSensor = sensorManager?.getDefaultSensor(Sensor.TYPE_GYROSCOPE)
    }

    fun startCollecting() {
        sensorManager?.registerListener(this, gyroscopeSensor, SensorManager.SENSOR_DELAY_NORMAL)
    }

    fun stopCollecting() {
        sensorManager?.unregisterListener(this)
    }

    override fun onSensorChanged(event: SensorEvent?) {
        event?.let {
            if (it.sensor.type == Sensor.TYPE_GYROSCOPE) {
                gyroscopeData = GyroscopeData(
                    x = it.values[0],
                    y = it.values[1],
                    z = it.values[2]
                )
            }
        }
    }

    override fun onAccuracyChanged(sensor: Sensor?, accuracy: Int) {}

    fun getGyroscopeData(): GyroscopeData? {
        return gyroscopeData
    }
}

data class GyroscopeData(
    val x: Float,
    val y: Float,
    val z: Float
)