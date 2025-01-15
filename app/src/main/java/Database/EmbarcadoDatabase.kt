package Database

import Datas.GyroscopeData
import android.content.Context
import androidx.room.Database
import androidx.room.Room
import androidx.room.RoomDatabase

@Database(entities = [GyroscopeData::class], version = 1, exportSchema = false)
abstract class EmbarcadoDatabase : RoomDatabase() {
    abstract fun gyroscopeDataDao(): GyroscopeDataDao

    companion object {
        @Volatile
        private var INSTANCE: EmbarcadoDatabase? = null

        fun getInstance(context: Context): EmbarcadoDatabase {

            return INSTANCE ?: synchronized(this) {
                val instance = Room.databaseBuilder(
                    context.applicationContext,
                    EmbarcadoDatabase::class.java,
                    "telemetry"
                ).build()
                INSTANCE = instance
                instance
            }
        }
    }
}