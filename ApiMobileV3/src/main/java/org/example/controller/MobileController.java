package org.example.controller;

import org.example.model.*;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/telemetry")
public class MobileController {

    @PostMapping("/gyroscope")
    public ResponseEntity saveGyroscope(@RequestBody Gyroscrope body){
        //TODO aqui o backend faz alguma coisa
        //Eu criei esse cara somente para testar a parte android
        return ResponseEntity.ok("success");
    }

    @PostMapping("/gps")
    public ResponseEntity saveGps(@RequestBody Location body){
        //TODO aqui o backend faz alguma coisa
        //Eu criei esse cara somente para testar a parte android
        return ResponseEntity.ok("success");
    }

    @PostMapping("/photo")
    public ResponseEntity savePhoto(@RequestBody Photo body){
        //TODO aqui o backend faz alguma coisa
        //Eu criei esse cara somente para testar a parte android
        return ResponseEntity.ok("success");
    }

    @PostMapping("/face")
    public ResponseEntity saveFace(@RequestBody Face body){
        //TODO aqui o backend faz alguma coisa
        //Eu criei esse cara somente para testar a parte android
        return ResponseEntity.ok("success");
    }

}
