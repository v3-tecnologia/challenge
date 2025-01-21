package com.example.teste;

import android.app.Service;
import android.content.Intent;
import android.os.IBinder;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.HashMap;
import java.util.Objects;

/* Decorator para procedimentos de requisições HTTP;
* */
public class RequestHTTPService extends Service {
    private final DateTimeFormatter ISO_FORMATTER = DateTimeFormatter.ISO_DATE_TIME;
    private final MyMacService mac = new MyMacService();

    public RequestHTTPService() {
    }

    @Override
    public IBinder onBind(Intent intent) {
        // TODO: Return the communication channel to the service.
        throw new UnsupportedOperationException("Not yet implemented");
    }

    public boolean sendPost(String url, HashMap<String, Object> campos){
        return reqPOST(url, campos);
    }

    private boolean reqPOST(String url, HashMap<String, Object> campos){
        /* OBS:
        * Inibindo nesta áreaimplementações puramente técnicas
        * onde não existem proporções algorítmicas a serem resolutas;
        * P.s. Implementaçẽos puramente técnicas para esta resolução.
        * */

        /* \/ tempo de envio da requisição; */
        String tempo = getTime();
        /* \/ obter Mac do usuário; */
        String myMacDevice = mac.getMac();
        return false;
    }

    private String getTime(){
        LocalDateTime ldt = LocalDateTime.now();
        return ldt.format(ISO_FORMATTER);  //2022-12-09T18:25:58.6037597
    }
}