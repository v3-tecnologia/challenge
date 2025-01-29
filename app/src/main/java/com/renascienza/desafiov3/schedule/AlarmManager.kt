package com.renascienza.desafiov3.schedule

import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.util.Log
import com.renascienza.desafiov3.db.Repository
import org.koin.java.KoinJavaComponent.inject

class AlarmReceiver : BroadcastReceiver() {
    private val repository: Repository by inject(Repository::class.java)

    override fun onReceive(context: Context?, intent: Intent?) {
        try {
            Log.d("AlarmReceiver", "Coletando dados...")
            repository.collectData()
        } catch (e: Exception) {
            Log.e("AlarmReceiver", "Erro ao coletar dados: ${e.message}")
        }
    }
}

class AlarmInitializer(private val context: Context) {
	
	fun isAlarmActive(): Boolean {
        val intent = Intent(context, AlarmReceiver::class.java)
        val pendingIntent = PendingIntent.getBroadcast(
            context,
            0,
            intent,
            PendingIntent.FLAG_NO_CREATE or PendingIntent.FLAG_IMMUTABLE
        )
        return pendingIntent != null
    }

    fun startAlarm() {
		
		if (isAlarmActive()) return
		
        val alarmManager = context.getSystemService(Context.ALARM_SERVICE) as AlarmManager
        val intent = Intent(context, AlarmReceiver::class.java)
        val pendingIntent = PendingIntent.getBroadcast(
            context,
            0,
            intent,
            PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE
        )

        // Definir o alarme para se repetir a cada 10 segundos
        alarmManager.setRepeating(
            AlarmManager.ELAPSED_REALTIME_WAKEUP,
            SystemClock.elapsedRealtime(),
            10000, // Intervalo em milissegundos (10 segundos)
            pendingIntent
        )
    }

    fun stopAlarm() {
        val alarmManager = context.getSystemService(Context.ALARM_SERVICE) as AlarmManager
        val intent = Intent(context, AlarmReceiver::class.java)
        val pendingIntent = PendingIntent.getBroadcast(
            context,
            0,
            intent,
            PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE
        )
        alarmManager.cancel(pendingIntent)
    }
}