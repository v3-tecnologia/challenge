package com.example.demo.controller

import org.springframework.web.bind.annotation.*
import org.springframework.http.ResponseEntity

// 1. DTO dos dados recebidos
data class ColetaRequest(
    val idDispositivo: String,
    val timestamp: Long,
    val giroscopio: Giroscopio,
    val gps: GPS,
    val fotoBase64: String // foto como string base64
)

data class Giroscopio(
    val x: Double,
    val y: Double,
    val z: Double
)

data class GPS(
    val latitude: Double,
    val longitude: Double
)

@RestController
@RequestMapping("/api/coletas")
class ColetaController {

    @PostMapping
    fun receberColeta(@RequestBody coleta: ColetaRequest): ResponseEntity<String> {
        // Aqui tu pode salvar no banco, processar, etc
        println("Recebeu coleta do dispositivo: ${coleta.idDispositivo}")
        println("Giroscópio: x=${coleta.giroscopio.x}, y=${coleta.giroscopio.y}, z=${coleta.giroscopio.z}")
        println("GPS: lat=${coleta.gps.latitude}, long=${coleta.gps.longitude}")
        println("Timestamp: ${coleta.timestamp}")
        println("Foto (tamanho base64): ${coleta.fotoBase64.length}")

        // Retorna um OK
        return ResponseEntity.ok("Coleta recebida com sucesso!")
    }
}
