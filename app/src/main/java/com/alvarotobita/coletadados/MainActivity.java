package com.alvarotobita.coletadados;

import android.Manifest;
import android.content.Intent;
import android.content.pm.PackageManager;
import android.os.Build;
import android.os.Bundle;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AppCompatActivity;
import androidx.core.app.ActivityCompat;
import androidx.core.content.ContextCompat;

import com.alvarotobita.coletadados.service.DataCollectorService;
import android.widget.Toast;

public class MainActivity extends AppCompatActivity {

    private static final int PERMISSION_REQUEST_CODE = 1001;

    // Array de permissões necessárias
    private final String[] REQUIRED_PERMISSIONS = {
            Manifest.permission.ACCESS_FINE_LOCATION,
            Manifest.permission.ACCESS_COARSE_LOCATION,
    };

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        checkAndRequestPermissions();
    }

    /**
     * Verifica se as permissões necessárias já foram concedidas.
     * @return true se todas as permissões forem concedidas, false caso contrário.
     */
    private boolean hasAllPermissions() {
        for (String permission : REQUIRED_PERMISSIONS) {
            if (ContextCompat.checkSelfPermission(this, permission) != PackageManager.PERMISSION_GRANTED) {
                return false;
            }
        }
        return true;
    }

    /**
     * Verifica as permissões e as solicita se necessário.
     */
    private void checkAndRequestPermissions() {
        if (hasAllPermissions()) {
            // Se as permissões já foram concedidas, inicia o serviço.
            iniciarServico();
        } else {
            // Solicita as permissões que ainda não foram concedidas.
            ActivityCompat.requestPermissions(this, REQUIRED_PERMISSIONS, PERMISSION_REQUEST_CODE);
        }
    }

    /**
     * Inicia o DataCollectorService como um Foreground Service.
     */
    private void iniciarServico() {
        Intent intent = new Intent(this, DataCollectorService.class);
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            ContextCompat.startForegroundService(this, intent);
        } else {
            startService(intent);
        }
    }

    @Override
    public void onRequestPermissionsResult(int requestCode, @NonNull String[] permissions,
                                           @NonNull int[] grantResults) {
        super.onRequestPermissionsResult(requestCode, permissions, grantResults);

        if (requestCode == PERMISSION_REQUEST_CODE) {
            if (hasAllPermissions()) {
                // Se todas as permissões foram concedidas, inicia o serviço.
                iniciarServico();
            } else {
                // Permissões negadas
                Toast.makeText(this, "Não foi possível iniciar o app", Toast.LENGTH_SHORT).show();
            }
        }
    }
}
