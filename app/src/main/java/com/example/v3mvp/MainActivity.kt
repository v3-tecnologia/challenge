package com.example.v3mvp

import android.Manifest
import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.content.pm.PackageManager
import android.os.Build
import android.os.Bundle
import android.os.Environment
import android.provider.MediaStore
import android.widget.Button
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat
import androidx.core.content.FileProvider
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.example.v3mvp.adapter.ColetaAdapter
import com.example.v3mvp.service.ColetaService
import com.example.v3mvp.util.Exportador
import com.example.v3mvp.viewmodel.ColetaViewModel
import java.io.File
import java.text.SimpleDateFormat
import java.util.*
import android.os.Handler
import android.os.Looper
import android.content.IntentFilter


class MainActivity : AppCompatActivity() {

    private lateinit var viewModel: ColetaViewModel
    private lateinit var recyclerView: RecyclerView
    private lateinit var btnExportar: Button
    private lateinit var btnLimpar: Button
    private lateinit var btnColetarAgora: Button
    private lateinit var adapter: ColetaAdapter

    private var fotoPath: String? = null
    private val REQUEST_FOTO = 1002
    private var pedirFotoDepoisPermissao = false

    private val erroReceiver = object : BroadcastReceiver() {
        override fun onReceive(context: Context?, intent: Intent?) {
            val msg = intent?.getStringExtra("mensagem") ?: "Erro desconhecido na coleta"
            Toast.makeText(this@MainActivity, msg, Toast.LENGTH_LONG).show()
        }
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        // Inicia componentes
        viewModel = ColetaViewModel(application)
        recyclerView = findViewById(R.id.recyclerColetas)
        btnExportar = findViewById(R.id.btnExportar)
        btnLimpar = findViewById(R.id.btnLimpar)
        btnColetarAgora = findViewById(R.id.btnColetarAgora)

        adapter = ColetaAdapter()
        recyclerView.layoutManager = LinearLayoutManager(this)
        recyclerView.adapter = adapter

        // Observa o LiveData das coletas (reativo)
        viewModel.coletas.observe(this) {
            adapter.submitList(it)
        }

        // Botão Exportar
        btnExportar.setOnClickListener {
            Exportador.exportarColetas(this, viewModel.coletas.value ?: emptyList())
        }
        // Botão Limpar
        btnLimpar.setOnClickListener {
            viewModel.limparColetas()
        }

        // Botão Coletar Agora
        btnColetarAgora.setOnClickListener {
            val precisaDePermissaoLocalizacao =
                ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED

            val precisaDePermissaoCamera =
                ActivityCompat.checkSelfPermission(this, Manifest.permission.CAMERA) != PackageManager.PERMISSION_GRANTED

            val permissoesFaltando = mutableListOf<String>()
            if (precisaDePermissaoLocalizacao) {
                permissoesFaltando.add(Manifest.permission.ACCESS_FINE_LOCATION)
            }
            if (precisaDePermissaoCamera) {
                permissoesFaltando.add(Manifest.permission.CAMERA)
            }

            if (permissoesFaltando.isNotEmpty()) {
                pedirFotoDepoisPermissao = true
                ActivityCompat.requestPermissions(this, permissoesFaltando.toTypedArray(), 100)
            } else {
                abrirCameraParaFoto()

            }

        }
        registerReceiver(
            erroReceiver,
            IntentFilter("com.example.v3mvp.LOCATION_ERROR"),
            Context.RECEIVER_NOT_EXPORTED
        );
        checarPermissoes()
    }

    // ----- Abrir câmera, receber foto, enviar pro Service -----

    private fun abrirCameraParaFoto() {
        val nomeArquivo = "FOTO_${SimpleDateFormat("yyyyMMdd_HHmmss", Locale.getDefault()).format(Date())}.jpg"
        val arquivo = File(getExternalFilesDir(Environment.DIRECTORY_PICTURES), nomeArquivo)
        fotoPath = arquivo.absolutePath

        val fotoUri = FileProvider.getUriForFile(
            this,
            "${packageName}.provider",
            arquivo
        )
        val intent = Intent(MediaStore.ACTION_IMAGE_CAPTURE)
        intent.putExtra(MediaStore.EXTRA_OUTPUT, fotoUri)
        startActivityForResult(intent, REQUEST_FOTO)
    }

    override fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?) {
        super.onActivityResult(requestCode, resultCode, data)
        if (requestCode == REQUEST_FOTO && resultCode == RESULT_OK && fotoPath != null) {
            coletarAgoraComFoto(fotoPath!!)
        }
    }

    private fun coletarAgoraComFoto(fotoPath: String) {
        val intent = Intent(this, ColetaService::class.java)
        intent.action = ColetaService.ACTION_COLETAR_AGORA
        intent.putExtra("fotoPath", fotoPath)
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            startForegroundService(intent)
        } else {
            startService(intent)
        }
    }

    // ------------- Permissões e utilidades ---------------

    private fun checarPermissoes() {
        val permissoes = mutableListOf(
            Manifest.permission.ACCESS_FINE_LOCATION,
            Manifest.permission.ACCESS_COARSE_LOCATION,
            Manifest.permission.FOREGROUND_SERVICE,
            Manifest.permission.FOREGROUND_SERVICE_LOCATION,
            Manifest.permission.CAMERA
        )

        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.Q) {
            permissoes.add(Manifest.permission.ACCESS_BACKGROUND_LOCATION)
        }

        val faltando = permissoes.filter {
            ActivityCompat.checkSelfPermission(this, it) != PackageManager.PERMISSION_GRANTED
        }

        if (faltando.isNotEmpty()) {
            ActivityCompat.requestPermissions(this, faltando.toTypedArray(), 100)
        } else {
            iniciarServico()
        }
    }
    private fun iniciarServico() {
        val intent = Intent(this, ColetaService::class.java)
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            startForegroundService(intent)
        } else {
            startService(intent)
        }
    }

    override fun onRequestPermissionsResult(
        requestCode: Int,
        permissions: Array<out String>,
        grantResults: IntArray
    ) {
        super.onRequestPermissionsResult(requestCode, permissions, grantResults)

        if (requestCode == 100 && grantResults.all { it == PackageManager.PERMISSION_GRANTED }) {
            Handler(Looper.getMainLooper()).postDelayed({
                iniciarServico()
                if (pedirFotoDepoisPermissao) {
                    pedirFotoDepoisPermissao = false
                    abrirCameraParaFoto()
                }
            }, 1000) // Espera 1 segundo antes de iniciar o serviço
        } else {
            Toast.makeText(this, "Permissões necessárias não concedidas", Toast.LENGTH_LONG).show()
        }
    }

    override fun onDestroy() {
        super.onDestroy()
        unregisterReceiver(erroReceiver)
    }

}
