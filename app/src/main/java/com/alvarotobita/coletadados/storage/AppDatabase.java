package com.alvarotobita.coletadados.storage;

import androidx.room.Database;
import androidx.room.RoomDatabase;
import com.alvarotobita.coletadados.model.DataSample;

@Database(entities = {DataSample.class}, version = 1)
public abstract class AppDatabase extends RoomDatabase {
    public abstract DataSampleDao dataSampleDao();
}