package com.example.v3mvp.service

import android.app.*
import android.content.*
import android.content.pm.ServiceInfo
import android.hardware.*
import android.os.*
import android.provider.Settings
import android.util.Log
import android.widget.Toast
import com.example.v3mvp.FotoActivity
import com.example.v3mvp.data.AppDatabase
import com.example.v3mvp.model.Coleta
import com.example.v3mvp.util.FotoHelper
import com.example.v3mvp.data.repository.ColetaRepository
import com.example.v3mvp.util.FaceDetectorUtil
import com.google.android.gms.location.LocationServices
import kotlinx.coroutines.*
import kotlinx.coroutines.tasks.await
import kotlinx.coroutines.suspendCancellableCoroutine

import com.example.v3mvp.proto.Coleta.ColetaMsg
import com.example.v3mvp.remote.RetrofitInstance
import com.example.v3mvp.util.coletaRoomToProto
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import java.io.IOException

import java.io.File


class ColetaService : Service(), SensorEventListener {

    companion object {
        const val ACTION_COLETAR_AGORA = "ACTION_COLETAR_AGORA"
        const val ACTION_UPDATE_INTERVAL = "ACTION_UPDATE_INTERVAL"
        const val EXTRA_INTERVAL = "EXTRA_INTERVAL"
    }

    private val scope = CoroutineScope(Dispatchers.IO + SupervisorJob())
    private lateinit var sensorManager: SensorManager
    private var lastGyro: FloatArray? = null
    private var intervalo: Long = 10_000L
    private var coletaJob: Job? = null

    private val repository by lazy {
        ColetaRepository(AppDatabase.getInstance(applicationContext).coletaDao())
    }

    override fun onCreate() {
        super.onCreate()
        criarNotificacaoForeground()

        sensorManager = getSystemService(Context.SENSOR_SERVICE) as SensorManager
        val gyro = sensorManager.getDefaultSensor(Sensor.TYPE_GYROSCOPE)
        sensorManager.registerListener(this, gyro, SensorManager.SENSOR_DELAY_UI)

        scope.launch {
            delay(1000)
            iniciarLoopAutomatico()
        }

        tentarReenviarPendentes()
    }

    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        when (intent?.action) {
            ACTION_COLETAR_AGORA -> {
                val fotoPath = intent.getStringExtra("fotoPath")
                scope.launch { salvarColeta(fotoPath) }
            }
            ACTION_UPDATE_INTERVAL -> {
                intervalo = intent.getLongExtra(EXTRA_INTERVAL, 10_000L)
                reiniciarLoopAutomatico()
            }
        }
        return START_STICKY
    }

    private fun criarNotificacaoForeground() {
        val canalId = "canal_coleta"
        val manager = getSystemService(NOTIFICATION_SERVICE) as NotificationManager
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            val canal = NotificationChannel(canalId, "Coleta de Dados", NotificationManager.IMPORTANCE_LOW)
            manager.createNotificationChannel(canal)
        }

        val notificacao = Notification.Builder(this, canalId)
            .setContentTitle("Coleta de dados ativa")
            .setContentText("O app está coletando dados em segundo plano.")
            .setSmallIcon(android.R.drawable.ic_menu_mylocation)
            .build()

        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.Q) {
            startForeground(1, notificacao, ServiceInfo.FOREGROUND_SERVICE_TYPE_LOCATION)
        } else {
            startForeground(1, notificacao)
        }
    }

    private fun iniciarLoopAutomatico() {
        coletaJob?.cancel()
        coletaJob = scope.launch {
            while (isActive) {
                withContext(Dispatchers.Main) {
                    val intent = Intent(this@ColetaService, FotoActivity::class.java)
                    intent.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK)
                    startActivity(intent)
                }
                delay(intervalo)
            }
        }
    }

    private fun reiniciarLoopAutomatico() {
        coletaJob?.cancel()
        iniciarLoopAutomatico()
    }

    private fun gerarBinarioDaColeta(coleta: Coleta): ByteArray {
        val fotoBytes = coleta.fotoPath?.let { fotoPath ->
            val fotoFile = File(fotoPath)
            if (fotoFile.exists()) fotoFile.readBytes() else null
        }

        return coletaRoomToProto(coleta, fotoBytes).toByteArray()
    }

    private fun salvarBinarioNoDownloads(coleta: Coleta, binario: ByteArray) {
        try {
            val nomeArquivo = "coleta-${coleta.deviceId}-${coleta.timestamp}.bin"
            val downloadsDir = Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_DOWNLOADS)
            val arquivo = File(downloadsDir, nomeArquivo)
            arquivo.writeBytes(binario)
            Log.i("ColetaBinaria", "Arquivo binário salvo em: ${arquivo.absolutePath}")
        } catch (e: Exception) {
            Log.e("ColetaBinaria", "Erro ao salvar binário localmente", e)
        }
    }

    private fun enviarColetaBinariaParaApi(coleta: Coleta) {
        val binario = gerarBinarioDaColeta(coleta)

        // Salva localmente no Downloads
        salvarBinarioNoDownloads(coleta, binario)

        val requestBody = binario.toRequestBody("application/octet-stream".toMediaType())

        val request = Request.Builder()
            .url("http://10.0.2.2:8080/api/binario")
            .post(requestBody)
            .build()

        OkHttpClient().newCall(request).enqueue(object : Callback {
            override fun onFailure(call: Call, e: IOException) {
                Log.e("ColetaBinaria", "Erro ao enviar binário", e)
            }

            override fun onResponse(call: Call, response: Response) {
                Log.i("ColetaBinaria", "Resposta da API binária: ${response.code} - ${response.body?.string()}")
            }
        })
    }


    private suspend fun salvarColeta(fotoPath: String?) {
        try {
            val fused = LocationServices.getFusedLocationProviderClient(this)
            val location = withContext(Dispatchers.Main) { fused.lastLocation.await() }

            // Validação 1: Localização válida (não pode ser null nem zerada)
            if (location == null || (location.latitude == 0.0 && location.longitude == 0.0)) {
                Log.e("ColetaService", "Localização inválida")
                emitirErro("Localização inválida")
                return
            }

            // Validação 2: Giroscópio válido (precisa estar coletado e com 3 eixos)
            if (lastGyro == null || lastGyro?.size != 3) {
                Log.e("ColetaService", "Giroscópio não coletado corretamente")
                emitirErro("Giroscópio não coletado corretamente")
                return
            }

            // Validação 3: Foto precisa existir, ter tamanho razoável (>10kb) e não pode ser toda preta/branca
            if (!fotoPath.isNullOrBlank()) {
                val fotoFile = java.io.File(fotoPath)
                if (!fotoFile.exists() || fotoFile.length() < 10_000) {
                    Log.e("ColetaService", "Foto não foi salva corretamente ou é muito pequena")
                    emitirErro("Foto não foi salva corretamente")
                    return
                }
                // Verifica se é toda preta ou branca (camera tampada/desligada)
                if (com.example.v3mvp.util.ImageUtils.isImageMostlyBlackOrWhite(fotoPath)) {
                    Log.e("ColetaService", "Foto inválida: toda preta ou branca")
                    emitirErro("Foto inválida: toda preta ou branca")
                    return
                }
            }

            val deviceId = Settings.Secure.getString(contentResolver, Settings.Secure.ANDROID_ID)

            // Padrão OK, mas pode mudar caso a foto não tenha rosto
            var status = "OK"
            var finalFotoPath = fotoPath

            // Validação 4: Verifica rosto na foto (só se foto existir e passou nos filtros acima)
            if (!fotoPath.isNullOrBlank()) {
                val temRosto = suspendCancellableCoroutine<Boolean> { continuation ->
                    FaceDetectorUtil.validarFotoContemRosto(applicationContext, fotoPath) { resultado ->
                        continuation.resume(resultado) {}
                    }
                }
                if (!temRosto) {
                    status = "FOTO SEM ROSTO"
                    finalFotoPath = null // não salva a foto
                }
            }

            // Criação da coleta
            val coleta = Coleta(
                timestamp = System.currentTimeMillis(),
                latitude = location.latitude,
                longitude = location.longitude,
                gyroX = lastGyro?.getOrNull(0),
                gyroY = lastGyro?.getOrNull(1),
                gyroZ = lastGyro?.getOrNull(2),
                deviceId = deviceId,
                fotoPath = finalFotoPath,
                status = status,
                enviado = false
            )

            repository.inserir(coleta)

            // Feedback de erro ou envio para nuvem
            if (status == "FOTO SEM ROSTO") {
                emitirErro("Nenhum rosto detectado na foto")
            } else {
                enviarColetaBinariaParaApi(coleta)
            }

        } catch (e: Exception) {
            emitirErro("Erro inesperado ao salvar coleta: ${e.message}")
            Log.e("ColetaService", "Erro ao salvar: ${e.message}", e)
        }
    }


    private fun enviarColetaParaApi(coleta: Coleta) {
        val fotoBytes = coleta.fotoPath?.let { FotoHelper.toByteArray(it) }
        val msgProto: ColetaMsg = coletaRoomToProto(coleta, fotoBytes)
        val payloadBinario: ByteArray = msgProto.toByteArray()
        val requestBody = payloadBinario.toRequestBody("application/x-protobuf".toMediaType())

        val service = RetrofitInstance.retrofit.create(ColetaApiService::class.java)
        scope.launch {
            try {
                val response = service.enviarColetaBinaria(requestBody)
                if (response.isSuccessful) {
                    repository.marcarComoEnviado(coleta.id)
                } else {
                    notificarUsuario("Erro ao enviar coleta: ${response.code()}")
                }
            } catch (e: Exception) {
                notificarUsuario("Falha ao enviar coleta: ${e.message}")
            }
        }
    }

    private fun tentarReenviarPendentes() {
        scope.launch {
            repository.listarNaoEnviados().forEach { enviarColetaBinariaParaApi(it) }
        }
    }

    private fun emitirErro(msg: String) {
        Handler(Looper.getMainLooper()).post {
            Toast.makeText(applicationContext, msg, Toast.LENGTH_LONG).show()
        }
    }

    private fun notificarUsuario(msg: String) {
        val canalId = "erros_coleta"
        val manager = getSystemService(NOTIFICATION_SERVICE) as NotificationManager
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            val canal = NotificationChannel(canalId, "Erros de Coleta", NotificationManager.IMPORTANCE_HIGH)
            manager.createNotificationChannel(canal)
        }
        val notificacao = Notification.Builder(this, canalId)
            .setContentTitle("Erro na Coleta")
            .setContentText(msg)
            .setSmallIcon(android.R.drawable.ic_dialog_alert)
            .build()
        manager.notify(System.currentTimeMillis().toInt(), notificacao)
    }

    override fun onDestroy() {
        coletaJob?.cancel()
        scope.cancel()
        sensorManager.unregisterListener(this)
        super.onDestroy()
    }

    override fun onBind(intent: Intent?): IBinder? = null

    override fun onSensorChanged(event: SensorEvent?) {
        if (event?.sensor?.type == Sensor.TYPE_GYROSCOPE) {
            lastGyro = event.values.clone()
        }
    }

    override fun onAccuracyChanged(sensor: Sensor?, accuracy: Int) {}
}
