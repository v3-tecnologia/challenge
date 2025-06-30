package com.example.v3mvp.adapter

import android.graphics.Color
import android.icu.text.SimpleDateFormat
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.recyclerview.widget.DiffUtil
import androidx.recyclerview.widget.ListAdapter
import androidx.recyclerview.widget.RecyclerView
import com.example.v3mvp.R
import com.example.v3mvp.model.Coleta
import java.util.Date
import java.util.Locale

class ColetaAdapter : ListAdapter<Coleta, ColetaAdapter.ColetaViewHolder>(DIFF_CALLBACK) {

    companion object {
        private val DIFF_CALLBACK = object : DiffUtil.ItemCallback<Coleta>() {
            override fun areItemsTheSame(oldItem: Coleta, newItem: Coleta): Boolean {
                return oldItem.id == newItem.id // Usa o id, que é a chave!
            }

            override fun areContentsTheSame(oldItem: Coleta, newItem: Coleta): Boolean {
                return oldItem == newItem
            }
        }
    }

    class ColetaViewHolder(itemView: View) : RecyclerView.ViewHolder(itemView) {
        private val txtDados: TextView = itemView.findViewById(R.id.txtDados)
        private val txtStatus: TextView = itemView.findViewById(R.id.txtStatus)

        fun bind(coleta: Coleta) {
            val latitude = coleta.latitude ?: 0.0
            val longitude = coleta.longitude ?: 0.0
            val x = coleta.gyroX ?: 0.0f
            val y = coleta.gyroY ?: 0.0f
            val z = coleta.gyroZ ?: 0.0f
            val enviadoTexto = if (coleta.enviado) "✔️ Enviado" else "❌ Não enviado"

            val sdf = SimpleDateFormat("dd/MM/yyyy HH:mm", Locale.getDefault())

            val textoFormatado = buildString {
                appendLine("🗓 Data: ${sdf.format(Date(coleta.timestamp))}")
                appendLine("📍 Localização:")
                appendLine("  Lat: %.6f".format(latitude))
                appendLine("  Long: %.6f".format(longitude))
                appendLine("🌀 Giroscópio:")
                appendLine("  X: %.4f".format(x))
                appendLine("  Y: %.4f".format(y))
                appendLine("  Z: %.4f".format(z))
                appendLine("📦 Status envio: $enviadoTexto")
            }

            txtDados.text = textoFormatado

            when (coleta.status) {
                "FOTO SEM ROSTO" -> {
                    txtStatus.text = "Foto sem rosto"
                    txtStatus.setTextColor(Color.RED)
                }
                else -> {
                    txtStatus.text = "OK"
                    txtStatus.setTextColor(Color.parseColor("#008000"))
                }
            }
        }
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ColetaViewHolder {
        val view = LayoutInflater.from(parent.context).inflate(R.layout.item_coleta, parent, false)
        return ColetaViewHolder(view)
    }

    override fun onBindViewHolder(holder: ColetaViewHolder, position: Int) {
        holder.bind(getItem(position))
    }
}
