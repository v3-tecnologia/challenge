package com.example.teste;

import android.app.Service;
import android.content.Intent;
import android.os.IBinder;
import android.net.wifi.WifiInfo;
import android.net.wifi.WifiManager;

/* Service para procedimento de obter Mac do usuário;
 * */
public class MyMacService extends Service {
    public MyMacService() {
    }

    @Override
    public IBinder onBind(Intent intent) {
        // TODO: Return the communication channel to the service.
        throw new UnsupportedOperationException("Not yet implemented");
    }

    public String getMac(){
        WifiManager wifiMgr = (WifiManager) getApplicationContext().getSystemService(WIFI_SERVICE);
        WifiInfo wifiInfo = wifiMgr.getConnectionInfo();
        final String MAC_ADDRESS = wifiInfo.getMacAddress();
        return MAC_ADDRESS;
    }
}