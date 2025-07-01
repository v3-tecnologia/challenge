package com.alvarotobita.coletadados.service;

// Importo as classes necessárias para trabalhar com serviços, notificações, localização e sensores
import android.app.*;
import android.content.Intent;
import android.os.*;
import androidx.annotation.Nullable;
import androidx.core.app.NotificationCompat;
import android.util.Log;
import java.util.concurrent.*;
import com.alvarotobita.coletadados.sensors.GpsManager;
import com.alvarotobita.coletadados.sensors.GyroscopeManager;
import android.location.Location;
import com.google.android.gms.location.LocationCallback;
import com.google.android.gms.location.LocationResult;

public class DataCollectorService extends Service {

    // Declaro o agendador para executar tarefas periódicas
    private ScheduledExecutorService scheduler;
    // Mantenho referências para os gerenciadores de GPS e giroscópio
    private GpsManager gpsManager;
    private GyroscopeManager gyroManager;
    // Guardo a última localização conhecida do GPS
    private Location lastKnownLocation;
    // Callback para ser chamado quando houver nova localização
    private LocationCallback locationCallback;

    @Override
    public void onCreate() {
        super.onCreate();

        // Crio o canal de notificação para Android 8.0+ (necessário para rodar serviços em foreground)
        createNotificationChannel();

        // Inicio o serviço em primeiro plano com uma notificação
        startForeground(1, buildNotification());

        // Crio um agendador que executa tarefas a cada intervalo definido
        scheduler = Executors.newSingleThreadScheduledExecutor();

        // Inicializo os gerenciadores de sensores
        gpsManager = new GpsManager(this);
        gyroManager = new GyroscopeManager(this);

        // Começo a coletar dados do giroscópio continuamente
        gyroManager.start();

        // Configuro o callback que será chamado quando o GPS obtiver uma nova localização
        locationCallback = new LocationCallback() {
            @Override
            public void onLocationResult(LocationResult locationResult) {
                if (locationResult != null && !locationResult.getLocations().isEmpty()) {
                    // Atualizo a última localização conhecida
                    lastKnownLocation = locationResult.getLastLocation();
                }
            }
        };

        // Instruo o GPS manager a começar a escutar atualizações de localização
        gpsManager.getLocation(locationCallback);

        // Agendo a execução do método coletarDados() a cada 10 segundos
        scheduler.scheduleAtFixedRate(this::coletarDados, 0, 10, TimeUnit.SECONDS);
    }

    // Neste método, faço a coleta dos dados dos sensores
    private void coletarDados() {
        // Obtenho os valores atuais do giroscópio (X, Y, Z)
        float[] gyro = gyroManager.getLastGyro();

        // Se houver uma última localização conhecida, exibo os dados no log
        if (lastKnownLocation != null) {
            Log.d("GPS", "Lat: " + lastKnownLocation.getLatitude() +
                    ", Lon: " + lastKnownLocation.getLongitude());
        }

        // Exibo no log os dados do giroscópio
        Log.d("GYRO", "X: " + gyro[0] + ", Y: " + gyro[1] + ", Z: " + gyro[2]);
    }

    // Esse método garante que o serviço continue rodando mesmo se for interrompido pelo sistema
    @Override
    public int onStartCommand(Intent intent, int flags, int startId) {
        return START_STICKY;
    }

    // Quando o serviço for destruído, paro os sensores e encerro o agendador
    @Override
    public void onDestroy() {
        super.onDestroy();
        if (scheduler != null) scheduler.shutdownNow();
        gyroManager.stop();
        gpsManager.stop(locationCallback);
    }

    // Este serviço não se conecta a nenhuma activity, então retorno null no bind
    @Nullable
    @Override
    public IBinder onBind(Intent intent) {
        return null;
    }

    // Aqui construo a notificação que ficará visível enquanto o serviço estiver ativo
    private Notification buildNotification() {
        return new NotificationCompat.Builder(this, "coleta_canal")
                .setContentTitle("Coletando dados")
                .setContentText("Serviço de coleta em execução.")
                .setSmallIcon(android.R.drawable.ic_menu_mylocation)
                .build();
    }

    // Crio o canal de notificação exigido pelo Android O+ para serviços em foreground
    private void createNotificationChannel() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            NotificationChannel channel = new NotificationChannel(
                    "coleta_canal",                 // ID do canal
                    "Canal de Coleta",              // Nome visível do canal
                    NotificationManager.IMPORTANCE_LOW  // Prioridade da notificação
            );
            NotificationManager manager = getSystemService(NotificationManager.class);
            manager.createNotificationChannel(channel);
        }
    }
}