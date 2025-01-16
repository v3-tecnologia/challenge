package Datas

import androidx.room.Entity
import androidx.room.PrimaryKey
import org.jetbrains.annotations.NotNull

@Entity(tableName = "gyroscope")
data class GyroscopeData(
    @PrimaryKey(autoGenerate = true)
    @NotNull
    override val id: Long,
    override val timestamp: Long,
    override val deviceId: String,
    val x: Float,
    val y: Float,
    val z: Float
): BaseData(id, timestamp, deviceId) {
    override fun toString(): String {
        return "TimeStamp: ${timestamp} - DeviceID: ${deviceId} - X: ${x}, Y: ${x}, Z: ${x} "
    }
}