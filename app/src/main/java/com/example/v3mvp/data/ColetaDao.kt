package com.example.v3mvp.data

import androidx.room.Dao
import androidx.room.Insert
import androidx.room.Query
import com.example.v3mvp.model.Coleta

@Dao
interface ColetaDao {
    @Insert
    suspend fun inserir(coleta: Coleta)

    @Query("SELECT * FROM coleta ORDER BY timestamp DESC")
    suspend fun buscarTodas(): List<Coleta>

    @Query("SELECT * FROM coleta ORDER BY timestamp DESC")
    fun observarTodas(): androidx.lifecycle.LiveData<List<Coleta>> // <--- AQUI

    @Query("DELETE FROM coleta")
    suspend fun deletarTodas()

    @Query("SELECT * FROM coleta WHERE enviado = 0")
    suspend fun buscarNaoEnviadas(): List<Coleta>

    @Query("UPDATE coleta SET enviado = :enviado WHERE id = :id")
    suspend fun atualizarEnvio(id: Long, enviado: Boolean)

    @Query("UPDATE coleta SET enviado = 1 WHERE id = :id")
    suspend fun marcarComoEnviado(id: Long)

    @Query("SELECT * FROM coleta WHERE enviado = 0")
    suspend fun listarNaoEnviados(): List<Coleta>

}


