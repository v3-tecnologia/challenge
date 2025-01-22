package com.example.teste;

import android.app.Service;
import android.content.Intent;
import android.os.IBinder;

import java.util.HashMap;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Callable;
import java.util.Set;
import java.util.HashSet;
import java.util.concurrent.Executors;

/*
 * Mediator de serviços de controlador de procedimentos concorrentes,
 * e terminador dos procedimentos concorrentes;
 *
 * invoca todos os procedimentos a serem realizados concorrentemente,
 * e executa o finalizador das threads ao término de todos os procedimentos;
 * */
public class TasksService extends Service {

    private ExecutorService executorService = Executors.newSingleThreadExecutor();
    private Set<Callable<String>> callables = new HashSet<Callable<String>>();
    private int nRequestsSends = 0;

    private MyGpsService myGpsService;
    private SensorService sensorService;

    public TasksService() {
    }

    public TasksService(MyGpsService myGpsService, SensorService sensorService){
        this.myGpsService = myGpsService;
        this.sensorService = sensorService;
    }

    @Override
    public IBinder onBind(Intent intent) {
        // TODO: Return the communication channel to the service.
        throw new UnsupportedOperationException("Not yet implemented");
    }

    /* \/:
     * introduzir listas de procedimentos a serem realizados concorrentemente;
     * */
    void addCallsRequests(){
        RequestHTTPService request = new RequestHTTPService();

        /*\/ */
        double lat = myGpsService.getLatitude();
        double lon = myGpsService.getLongitude();
        /*\/ */
        String cord = sensorService.getCordinates();

        callables.add(new Callable<String>() {
            public String call() throws Exception {
                /*\/ envio de dados para a API; */
                HashMap<String, Object> campos = new HashMap<>();
                if(request.sendPost("", campos)) {
                    nRequestsSends++;
                }
                return "Request 1";
            }
        });
        callables.add(new Callable<String>() {
            public String call() throws Exception {
                /*\/ envio de dados para a API; */
                HashMap<String, Object> campos = new HashMap<>();
                if(request.sendPost("", campos)) {
                    nRequestsSends++;
                }
                return "Request 2";
            }
        });
        callables.add(new Callable<String>() {
            public String call() throws Exception {
                /*\/ envio de dados para a API; */
                HashMap<String, Object> campos = new HashMap<>();
                if(request.sendPost("", campos)) {
                    nRequestsSends++;
                }
                return "Request 3";
            }
        });
    }

    /* \/:
     * invocar todos os procedimentos a serem realizados concorrentemente;
     * */
    public void invokeAllTasks() throws ExecutionException, InterruptedException {
        if(callables.size() > 0) {
            String result = executorService.invokeAny(callables);
            System.out.println("result = " + result);

            /* finalizador dos procedimentos concorrentes; */
            executorService.shutdown();
        }
    }

    /* Saber se todas as requisições foram enviadas com sucesso;
    * */
    public boolean getTodosEnviados(){
        return (nRequestsSends == callables.size());
    }
}