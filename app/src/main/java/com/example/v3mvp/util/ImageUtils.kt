package com.example.v3mvp.util

import android.graphics.BitmapFactory
import java.io.File

object ImageUtils {
    fun isImageMostlyBlackOrWhite(fotoPath: String): Boolean {
        val file = File(fotoPath)
        if (!file.exists() || file.length() < 1024) return true
        val bitmap = BitmapFactory.decodeFile(fotoPath) ?: return true
        var blackOrWhitePixels = 0
        val totalPixels = 20 * 20
        for (x in 0 until bitmap.width step bitmap.width / 20) {
            for (y in 0 until bitmap.height step bitmap.height / 20) {
                val pixel = bitmap.getPixel(x, y)
                val r = (pixel shr 16) and 0xff
                val g = (pixel shr 8) and 0xff
                val b = pixel and 0xff
                if ((r < 10 && g < 10 && b < 10) || (r > 245 && g > 245 && b > 245)) {
                    blackOrWhitePixels++
                }
            }
        }
        return (blackOrWhitePixels.toFloat() / totalPixels) > 0.8
    }
}
