package com.example.v3mvp.util

import android.util.Base64
import java.io.File

object FotoHelper {
    fun toBase64(path: String?): String {
        if (path.isNullOrEmpty()) return ""
        val file = File(path)
        if (!file.exists()) return ""
        val bytes = file.readBytes()
        return Base64.encodeToString(bytes, Base64.NO_WRAP)
    }

    fun toByteArray(path: String?): ByteArray? {
        if (path.isNullOrEmpty()) return null
        val file = File(path)
        if (!file.exists()) return null
        return file.readBytes()
    }
}
