package com.example.v3mvp.viewmodel

import android.app.Application
import androidx.lifecycle.AndroidViewModel
import androidx.lifecycle.LiveData
import androidx.lifecycle.viewModelScope
import com.example.v3mvp.data.AppDatabase
import com.example.v3mvp.model.Coleta
import kotlinx.coroutines.launch

class ColetaViewModel(app: Application) : AndroidViewModel(app) {

    private val coletaDao = AppDatabase.getInstance(app).coletaDao()

    val coletas: LiveData<List<Coleta>> = coletaDao.observarTodas()

    fun limparColetas() {
        viewModelScope.launch {
            coletaDao.deletarTodas()
        }
    }
}

