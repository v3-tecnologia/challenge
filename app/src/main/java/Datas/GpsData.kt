package Datas

import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "gps")
data class GpsData(
    @PrimaryKey(autoGenerate = true) override val id: Long,
    override val timestamp: Long,
    override val deviceId: String,
    val latitude: Double,
    val longitude: Double,
) : BaseData(id, timestamp, deviceId){
    override fun toString(): String {
        return "Timestamp: ${timestamp} - Latitude: ${latitude}, Longitude: ${longitude}"
    }
}