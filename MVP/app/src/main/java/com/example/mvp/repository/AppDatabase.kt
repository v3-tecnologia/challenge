package com.example.mvp.repository

import android.content.Context
import androidx.room.Database
import androidx.room.Room
import androidx.room.RoomDatabase
import androidx.room.TypeConverters
import com.example.mvp.entity.TelemetryData

@Database(entities = [TelemetryData::class], version = 1, exportSchema = false)
@TypeConverters(androidx.databinding.adapters.Converters::class)
abstract class AppDatabase : RoomDatabase() {
    abstract fun telemetryDataDao(): TelemetryDataDao

    companion object {
        @Volatile
        private var INSTANCE: AppDatabase? = null

        fun getDatabase(context: Context): AppDatabase {
            return INSTANCE ?: synchronized(this) {
                val instance = Room.databaseBuilder(
                    context.applicationContext,
                    AppDatabase::class.java,
                    "telemetry_database"
                ).build()
                INSTANCE = instance
                instance
            }
        }
    }
}