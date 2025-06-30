package com.example.demo.controller

import com.example.demo.proto.Coleta.ColetaMsg // <- Importação correta da classe gerada
import org.springframework.web.bind.annotation.*
import org.springframework.http.ResponseEntity
import java.io.File
import java.nio.file.Files

@RestController
@RequestMapping("/coleta-binaria")
class ColetaBinariaController {

    @PostMapping
    fun receberBinario(@RequestBody body: ByteArray): ResponseEntity<String> {
        return try {
            val coleta = ColetaMsg.parseFrom(body)

            println("ID: ${coleta.deviceId}")
            println("GPS: ${coleta.latitude}, ${coleta.longitude}")
            println("Gyro: ${coleta.gyroX}, ${coleta.gyroY}, ${coleta.gyroZ}")
            println("Status: ${coleta.status}")
            println("Timestamp: ${coleta.timestamp}")

            // Salvar a imagem (se houver dados)
            if (!coleta.foto.isEmpty) {
                val fotoFile = File("foto-${coleta.deviceId}.jpg")
                Files.write(fotoFile.toPath(), coleta.foto.toByteArray())
                println("Foto salva em ${fotoFile.absolutePath}")
            }

            ResponseEntity.ok("Coleta binária recebida com sucesso")
        } catch (e: Exception) {
            e.printStackTrace()
            ResponseEntity.badRequest().body("Erro ao processar dados binários: ${e.message}")
        }
    }
}
