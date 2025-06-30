package com.example.v3mvp.util

import android.content.Context
import android.widget.Toast
import com.example.v3mvp.model.Coleta
import java.io.File
import java.io.FileWriter

object Exportador {
    fun exportarColetas(context: Context, lista: List<Coleta>) {
        try {
            println("Exportador: Iniciando exportação de ${lista.size} coletas")

            val dir = File(context.getExternalFilesDir(null), "export")
            if (!dir.exists()) dir.mkdirs()

            val file = File(dir, "coletas.csv")
            val writer = FileWriter(file)

            writer.append("ID,Timestamp,Latitude,Longitude,GyroX,GyroY,GyroZ\n")

            for (c in lista) {
                println("Exportador: Coleta -> ${c.id}, ${c.timestamp}, ${c.latitude}, ${c.longitude}, ${c.gyroX}, ${c.gyroY}, ${c.gyroZ}")
                writer.append("${c.id},${c.timestamp},${c.latitude},${c.longitude},${c.gyroX},${c.gyroY},${c.gyroZ}\n")
            }

            writer.flush()
            writer.close()

            Toast.makeText(context, "Exportado em: ${file.absolutePath}", Toast.LENGTH_LONG).show()
        } catch (e: Exception) {
            println("Exportador: Erro ao exportar - ${e.message}")
            Toast.makeText(context, "Erro ao exportar: ${e.message}", Toast.LENGTH_LONG).show()
        }
    }
}

