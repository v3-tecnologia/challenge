package Datas

import androidx.room.Entity
import androidx.room.PrimaryKey
import java.sql.Timestamp

@Entity(tableName = "photo")
data class PhotoData(
    @PrimaryKey(autoGenerate = true) override val id: Long,
    override val timestamp: Timestamp,
    override val deviceId: String,
    val photoPath: String
) : BaseData(id, timestamp, deviceId)