package com.alvarotobita.coletadados.storage;

import androidx.room.*;
import com.alvarotobita.coletadados.model.DataSample;
import java.util.List;

@Dao
public interface DataSampleDao {
    @Insert
    void insert(DataSample sample);

    @Query("SELECT * FROM DataSample ORDER BY timestamp DESC")
    List<DataSample> getAll();
}