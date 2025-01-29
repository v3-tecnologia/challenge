package com.renascienza.desafiov3

import android.app.Application
import com.renascienza.desafiov3.db.Repository
import com.renascienza.desafiov3.schedule.AlarmInitializer
import org.koin.android.ext.koin.androidContext
import org.koin.core.context.startKoin
import org.koin.dsl.module

val appModule = module {
    single { Repository() }
}

class App : Application() {
    private lateinit var alarmInitializer: AlarmInitializer

    override fun onCreate() {
        super.onCreate()

        // Inicializando o Koin
        startKoin {
            androidContext(this@App)
            modules(appModule)
        }

        // Iniciar o alarme
        alarmInitializer = AlarmInitializer(this)
        alarmInitializer.startAlarm()
    }

    override fun onTerminate() {
        // Parar o alarme quando a aplicação for encerrada
        alarmInitializer.stopAlarm()
        super.onTerminate()
    }
}