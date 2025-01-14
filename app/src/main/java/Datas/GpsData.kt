package Datas

import androidx.room.Entity
import androidx.room.PrimaryKey
import java.sql.Timestamp

@Entity(tableName = "gps")
data class GpsData(
    @PrimaryKey(autoGenerate = true) override val id: Long,
    override val timestamp: Timestamp,
    override val deviceId: String,
    val latitude: Double,
    val longitude: Double,
) : BaseData(id, timestamp, deviceId)