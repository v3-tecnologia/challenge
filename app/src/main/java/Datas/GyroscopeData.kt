package Datas

import androidx.room.Entity
import androidx.room.PrimaryKey
import org.jetbrains.annotations.NotNull

@Entity(tableName = "gyroscope")
data class GyroscopeData(
    @PrimaryKey(autoGenerate = true)
    @NotNull
    val id: Long,
    val timestamp: Long,
    val deviceId: String,
    val x: Float,
    val y: Float,
    val z: Float
) {
    override fun toString(): String {
        return "TimeStamp: ${timestamp} - DeviceID: ${deviceId} - X: ${x}, Y: ${x}, Z: ${x} "
    }
}