package Datas

import androidx.room.Entity
import androidx.room.PrimaryKey
import java.sql.Timestamp

@Entity(tableName = "gyroscope")
data class GyroscopeData(
    @PrimaryKey(autoGenerate = true) override val id: Long,
    override val timestamp: Timestamp,
    override val deviceId: String,
    val x: Float,
    val y: Float,
    val z: Float
): BaseData(id, timestamp, deviceId)