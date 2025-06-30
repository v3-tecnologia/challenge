package com.example.v3mvp

import android.os.Bundle
import android.util.Log
import androidx.appcompat.app.AppCompatActivity
import androidx.camera.core.*
import androidx.camera.lifecycle.ProcessCameraProvider
import androidx.core.content.ContextCompat
import com.example.v3mvp.data.AppDatabase
import com.example.v3mvp.model.Coleta
import kotlinx.coroutines.*
import kotlinx.coroutines.tasks.await
import java.io.File
import java.text.SimpleDateFormat
import java.util.*

class FotoActivity : AppCompatActivity() {

    private val scope = CoroutineScope(Dispatchers.IO + SupervisorJob())

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        capturarFoto()
    }

    private fun capturarFoto() {
        val cameraProviderFuture = ProcessCameraProvider.getInstance(this)

        cameraProviderFuture.addListener({
            val cameraProvider = cameraProviderFuture.get()
            val imageCapture = ImageCapture.Builder().build()
            val cameraSelector = CameraSelector.DEFAULT_FRONT_CAMERA

            cameraProvider.unbindAll()
            cameraProvider.bindToLifecycle(this, cameraSelector, imageCapture)

            // Agora salva na pasta externa Pictures
            val fotoFile = File(
                getExternalFilesDir(android.os.Environment.DIRECTORY_PICTURES),
                "FOTO_${SimpleDateFormat("yyyyMMdd_HHmmss", Locale.US).format(Date())}.jpg"
            )
            val outputOptions = ImageCapture.OutputFileOptions.Builder(fotoFile).build()

            imageCapture.takePicture(
                outputOptions,
                ContextCompat.getMainExecutor(this),
                object : ImageCapture.OnImageSavedCallback {
                    override fun onImageSaved(output: ImageCapture.OutputFileResults) {
                        salvarColeta(fotoFile.absolutePath)
                    }

                    override fun onError(exc: ImageCaptureException) {
                        Log.e("FotoActivity", "Erro ao capturar foto", exc)
                        finish()
                    }
                }
            )
        }, ContextCompat.getMainExecutor(this))
    }

    private fun salvarColeta(fotoPath: String) {
        scope.launch {
            try {
                val fused = com.google.android.gms.location.LocationServices.getFusedLocationProviderClient(this@FotoActivity)
                val location = withContext(Dispatchers.Main) { fused.lastLocation.await() }

                if (location == null || (location.latitude == 0.0 && location.longitude == 0.0)) {
                    Log.e("FotoActivity", "Localização inválida. Coleta descartada.")
                    finish()
                    return@launch
                }

                // Detectar rosto na foto
                withContext(Dispatchers.Main) {
                    com.example.v3mvp.util.FaceDetectorUtil.validarFotoContemRosto(
                        this@FotoActivity, fotoPath
                    ) { temRosto ->
                        val status = if (temRosto) "OK" else "FOTO SEM ROSTO"

                        scope.launch {
                            val deviceId = android.provider.Settings.Secure.getString(contentResolver, android.provider.Settings.Secure.ANDROID_ID)
                            val coleta = Coleta(
                                timestamp = System.currentTimeMillis(),
                                latitude = location.latitude,
                                longitude = location.longitude,
                                gyroX = null, gyroY = null, gyroZ = null,
                                deviceId = deviceId,
                                fotoPath = if (temRosto) fotoPath else null, // só salva path se tiver rosto
                                status = status,
                                enviado = false
                            )
                            val db = AppDatabase.getInstance(applicationContext)
                            db.coletaDao().inserir(coleta)
                            Log.d("FotoActivity", "Coleta com foto salva: $coleta")
                            finish()
                        }
                    }
                }
            } catch (e: Exception) {
                Log.e("FotoActivity", "Erro ao salvar coleta com foto", e)
                finish()
            }
        }
    }

}
