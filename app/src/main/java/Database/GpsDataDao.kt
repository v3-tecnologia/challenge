package Database

import Datas.GpsData
import androidx.room.Dao
import androidx.room.Insert
import androidx.room.OnConflictStrategy
import androidx.room.Query

@Dao
interface GpsDataDao {
    @Insert(onConflict = OnConflictStrategy.IGNORE)
    fun insert(data: GpsData): Long

    @Query("SELECT * FROM gps ORDER BY timestamp DESC")
    suspend fun GetAll(): List<GpsData>
}