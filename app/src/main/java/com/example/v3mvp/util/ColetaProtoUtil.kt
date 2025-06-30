package com.example.v3mvp.util

import com.example.v3mvp.model.Coleta
import com.example.v3mvp.proto.Coleta.ColetaMsg
import com.google.protobuf.ByteString

fun coletaRoomToProto(coleta: Coleta, fotoEmBytes: ByteArray?): ColetaMsg {
    val builder = ColetaMsg.newBuilder()
        .setTimestamp(coleta.timestamp)
        .setLatitude(coleta.latitude ?: 0.0)
        .setLongitude(coleta.longitude ?: 0.0)
        .setGyroX(coleta.gyroX ?: 0f)
        .setGyroY(coleta.gyroY ?: 0f)
        .setGyroZ(coleta.gyroZ ?: 0f)
        .setDeviceId(coleta.deviceId ?: "")
        .setStatus(coleta.status ?: "")
    if (fotoEmBytes != null) builder.setFoto(ByteString.copyFrom(fotoEmBytes))
    return builder.build()
}
