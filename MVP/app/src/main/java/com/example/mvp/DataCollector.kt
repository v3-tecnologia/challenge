package com.example.mvp

import android.annotation.SuppressLint
import android.content.Context
import android.hardware.Sensor
import android.hardware.SensorEvent
import android.hardware.SensorEventListener
import android.hardware.SensorManager
import android.location.Location
import android.net.Uri
import android.util.Log
import androidx.core.content.ContextCompat
import com.example.mvp.entity.TelemetryData
import com.google.android.gms.location.LocationServices
import kotlinx.coroutines.suspendCancellableCoroutine
import java.io.File
import java.util.ArrayList
import kotlin.coroutines.resume
import kotlin.coroutines.resumeWithException

class DataCollector(private val context: Context)  {
    private lateinit var fusedLocationClient: com.google.android.gms.location.FusedLocationProviderClient

    suspend fun collectData(): TelemetryData {
        fusedLocationClient = com.google.android.gms.location.LocationServices.getFusedLocationProviderClient(context)

        // Coletar dados do sensor de giroscópio
        val gyroData = collectGyroData()
        Log.i("DataCollectionService gyro", gyroData.toString())

        // Coletar dados do GPS
        val location = collectGpsData();
        val gpsData = (location?.latitude ?: 0.0) + (location?.longitude ?: 0.0);
        Log.i("DataCollectionService latitude", gpsData.toString());
        //val gpsData = Any();

        // Capturar uma foto
        val photoData = capturePhoto()
        Log.i("DataCollectionService photo", photoData.toString())
        // Detectar rosto na foto (opcional, para o nível 4)

        val faceData = detectFaceInPhoto(photoData)
        Log.i("DataCollectionService face", faceData.toString())
        //return Data(gyroData, gpsData, photoData, faceData)

        return TelemetryData(
            0,
            0F,//gyroData.getX(),
            0F,//gyroData.getY(),
            0F,//gyroData.getZ(),
            (location?.latitude ?: 0.0),
            (location?.longitude ?: 0.0),
            photoData.toString(),
            faceData.toString()
        );
    }

    fun detectFaceInPhoto(photoData: Any): Any {
        // Verificar se photoData é um caminho de arquivo válido
        if (photoData is String && File(photoData).exists()) {
            val image = com.google.mlkit.vision.common.InputImage.fromFilePath(context, Uri.fromFile(File(photoData)))

            val options = com.google.mlkit.vision.face.FaceDetectorOptions.Builder()
                .setPerformanceMode(com.google.mlkit.vision.face.FaceDetectorOptions.PERFORMANCE_MODE_FAST)
                .setLandmarkMode(com.google.mlkit.vision.face.FaceDetectorOptions.LANDMARK_MODE_ALL)
                .setClassificationMode(com.google.mlkit.vision.face.FaceDetectorOptions.CLASSIFICATION_MODE_ALL)
                .build()

            val detector = com.google.mlkit.vision.face.FaceDetection.getClient(options)
            val faceList = ArrayList<Any>();

            detector.process(image)
                .addOnSuccessListener { faces ->
                    // faces é uma lista de objetos Face
                    // Você pode acessar informações sobre cada rosto, como:
                    // - Bounding box: faces[i].boundingBox
                    // - Pontos de referência: faces[i].allLandmarks
                    // - Probabilidade de sorriso: faces[i].smilingProbability
                    // - Probabilidade de olhos abertos: faces[i].leftEyeOpenProbability, faces[i].rightEyeOpenProbability
                    // ...

                    // Retornar as informações desejadas sobre os rostos
                    // Por exemplo, você pode retornar uma lista de bounding boxes
                    faceList.add(faces.get(0).boundingBox);
                }
            if(faceList.size > 0){
                return faceList.get(0);
            } else {
                return "_";
            }

        } else {
            // Lidar com photoData inválido
            Log.i("DataCollectionService ", "Caminho de arquivo inválido para a foto")
            return "_"; // Retornar uma lista vazia
        }

        return "_";
    }

    fun capturePhoto(): Any {
        val imageCapture = androidx.camera.core.ImageCapture.Builder()
            .setTargetRotation(android.view.Surface.ROTATION_0)
            .build()

        // Criar um arquivo para salvar a foto
        val photoFile = File(context.filesDir, "photo.jpg")

        // Criar um OutputFileOptions
        val outputOptions =
            androidx.camera.core.ImageCapture.OutputFileOptions.Builder(photoFile).build()

        try{
            imageCapture.takePicture(
                outputOptions,
                ContextCompat.getMainExecutor(context),
                object : androidx.camera.core.ImageCapture.OnImageSavedCallback {
                    override fun onImageSaved(outputFileResults: androidx.camera.core.ImageCapture.OutputFileResults) {
                        // Foto salva com sucesso
                        Log.i("DataCollectionService", "Foto salva em: ${photoFile.absolutePath}")
                    }

                    override fun onError(exception: androidx.camera.core.ImageCaptureException) {
                        // Erro ao salvar a foto
                        Log.e("DataCollector", "Erro ao salvar a foto", exception)
                    }
                }
            )
        }catch (e: Exception){
            Log.e("DataCollector", "Erro ao capturar a foto", e)
        }
        // Tirar a foto

        Log.i("DataCollectionService", "saindo da captura")
        return photoFile.absolutePath // Retornar o caminho da foto
    }

    @SuppressLint("MissingPermission")
    suspend fun collectGpsData(): Location? {
        val fusedLocationClient = LocationServices.getFusedLocationProviderClient(context)
        return suspendCancellableCoroutine { continuation ->
            fusedLocationClient.lastLocation
                .addOnSuccessListener { location: Location? ->
                    continuation.resume(location)
                }
                .addOnFailureListener { e ->
                    continuation.resumeWithException(e)
                }
        }
    }

    fun collectGyroData(): Any {
        val sensorManager = context.getSystemService(Context.SENSOR_SERVICE) as SensorManager
        val gyroscopeSensor = sensorManager.getDefaultSensor(Sensor.TYPE_GYROSCOPE)

        val gyroData = mutableListOf<Float>()

        val sensorEventListener = object : SensorEventListener {
            override fun onSensorChanged(event: SensorEvent) {
                if (event.sensor.type == Sensor.TYPE_GYROSCOPE) {
                    gyroData.addAll(event.values.toList())
                    // Desregistrar o listener após obter os dados
                    sensorManager.unregisterListener(this)
                }
            }

            override fun onAccuracyChanged(sensor: Sensor, accuracy: Int) {
                // Lidar com mudanças na precisão do sensor
            }
        }

        sensorManager.registerListener(sensorEventListener, gyroscopeSensor, SensorManager.SENSOR_DELAY_NORMAL)

        return gyroData // Retornar os dados do giroscópio
    }

}