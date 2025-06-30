package com.example.v3mvp.data.repository

import com.example.v3mvp.data.ColetaDao
import com.example.v3mvp.model.Coleta

class ColetaRepository(private val dao: ColetaDao) {

    suspend fun inserir(coleta: Coleta) = dao.inserir(coleta)

    suspend fun marcarComoEnviado(id: Int) = dao.marcarComoEnviado(id.toLong())

    suspend fun listarNaoEnviados(): List<Coleta> = dao.listarNaoEnviados()
}