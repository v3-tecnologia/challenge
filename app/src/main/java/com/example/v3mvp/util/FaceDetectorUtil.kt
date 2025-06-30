package com.example.v3mvp.util

import android.content.Context
import android.net.Uri
import com.google.mlkit.vision.common.InputImage
import com.google.mlkit.vision.face.FaceDetection
import com.google.mlkit.vision.face.FaceDetectorOptions
import java.io.File

object FaceDetectorUtil {
    fun validarFotoContemRosto(
        context: Context,
        fotoPath: String,
        onResult: (Boolean) -> Unit
    ) {
        try {
            val image = InputImage.fromFilePath(context, Uri.fromFile(File(fotoPath)))
            val options = FaceDetectorOptions.Builder()
                .setPerformanceMode(FaceDetectorOptions.PERFORMANCE_MODE_FAST)
                .build()
            val detector = FaceDetection.getClient(options)
            detector.process(image)
                .addOnSuccessListener { faces ->
                    android.util.Log.d("FaceDetectorUtil", "Validando foto: $fotoPath")
                    android.util.Log.d("FaceDetectorUtil", "Rostos detectados: ${faces.size}")
                    onResult(faces.isNotEmpty())
                }
                .addOnFailureListener { e ->
                    android.util.Log.e("FaceDetectorUtil", "Erro ao detectar rosto: ${e.message}")
                    onResult(false)
                }
        } catch (e: Exception) {
            android.util.Log.e("FaceDetectorUtil", "Exceção: ${e.message}")
            onResult(false)
        }
    }
}

