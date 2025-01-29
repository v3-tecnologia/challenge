package com.renascienza.desafiov3.network

import android.content.Context
import android.net.wifi.WifiManager
import android.text.format.Formatter

class MacAddressCollector(private val context: Context) {
    fun getMacAddress(): String? {
        val wifiManager = context.applicationContext.getSystemService(Context.WIFI_SERVICE) as WifiManager
        return wifiManager.connectionInfo.macAddress
    }
}