package com.renascienza.desafiov3.camera

import android.content.Context
import android.graphics.Bitmap
import android.graphics.BitmapFactory
import android.util.Base64
import androidx.camera.core.*
import androidx.camera.lifecycle.ProcessCameraProvider
import androidx.core.content.ContextCompat
import java.io.ByteArrayOutputStream
import java.nio.ByteBuffer
import java.util.concurrent.ExecutorService
import java.util.concurrent.Executors

class PhotoCollector(private val context: Context) {

    private val cameraExecutor: ExecutorService = Executors.newSingleThreadExecutor()

    fun capturePhoto(callback: (String) -> Unit) {
        val cameraProviderFuture = ProcessCameraProvider.getInstance(context)
        
        cameraProviderFuture.addListener({
            val cameraProvider: ProcessCameraProvider = cameraProviderFuture.get()

            val imageCapture = ImageCapture.Builder()
                .setCaptureMode(ImageCapture.CAPTURE_MODE_MINIMIZE_LATENCY)
                .build()

            val cameraSelector = CameraSelector.DEFAULT_BACK_CAMERA

            try {
                cameraProvider.unbindAll()
                val camera = cameraProvider.bindToLifecycle(
                    null, // Não precisa de uma Activity
                    cameraSelector,
                    imageCapture
                )

                val outputOptions = ImageCapture.OutputFileOptions.Builder(createTempFile()).build()

                imageCapture.takePicture(outputOptions, ContextCompat.getMainExecutor(context),
                    object : ImageCapture.OnImageSavedCallback {
                        override fun onImageSaved(outputFileResults: ImageCapture.OutputFileResults) {
                            val bitmap = BitmapFactory.decodeFile(outputFileResults.savedUri?.path)

                            if (bitmap != null && hasFace(bitmap)) {
                                val base64Photo = encodeToBase64(bitmap)
                                callback(base64Photo)
                            } else {
                                callback("")
                            }
                        }

                        override fun onError(exception: ImageCaptureException) {
                            exception.printStackTrace()
                            callback("")
                        }
                    })

            } catch (exc: Exception) {
                exc.printStackTrace()
                callback("")
            }
        }, ContextCompat.getMainExecutor(context))
    }

    private fun createTempFile(): java.io.File {
        return java.io.File.createTempFile("captured_photo", ".jpg", context.cacheDir)
    }

    private fun hasFace(bitmap: Bitmap): Boolean {
        val maxFaces = 1
        val faceDetector = android.media.FaceDetector(bitmap.width, bitmap.height, maxFaces)
        val faces = arrayOfNulls<android.media.FaceDetector.Face>(maxFaces)
        val numFaces = faceDetector.findFaces(bitmap, faces)
        return numFaces > 0
    }

    private fun encodeToBase64(bitmap: Bitmap): String {
        val outputStream = ByteArrayOutputStream()
        bitmap.compress(Bitmap.CompressFormat.JPEG, 100, outputStream)
        return Base64.encodeToString(outputStream.toByteArray(), Base64.DEFAULT)
    }
}