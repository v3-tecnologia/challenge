package Database

import Datas.GyroscopeData
import androidx.lifecycle.LiveData
import androidx.room.Dao
import androidx.room.Insert
import androidx.room.OnConflictStrategy
import androidx.room.Query

@Dao
interface GyroscopeDataDao {

    @Insert(onConflict = OnConflictStrategy.IGNORE)
    fun insert(data: GyroscopeData): Long

    @Query("SELECT * FROM gyroscope ORDER BY timestamp DESC")
    suspend fun GetAll(): List<GyroscopeData>
}