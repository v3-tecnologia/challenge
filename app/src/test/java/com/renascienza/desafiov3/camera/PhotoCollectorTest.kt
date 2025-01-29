package com.renascienza.desafiov3.camera

import android.content.Context
import android.graphics.Bitmap
import android.graphics.BitmapFactory
import androidx.test.core.app.ApplicationProvider
import org.junit.Assert.assertTrue
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RobolectricTestRunner

@RunWith(RobolectricTestRunner::class)
class PhotoCollectorTest {
    private lateinit var context: Context
    private lateinit var photoCollector: PhotoCollector

    @Before
    fun setup() {
        context = ApplicationProvider.getApplicationContext()
        photoCollector = PhotoCollector(context)
    }

    @Test
    fun testPhotoWithFace() {
        // Carregar uma imagem com rosto
        val bitmap: Bitmap = BitmapFactory.decodeResource(context.resources, R.drawable.photo_with_face)
        assertTrue(photoCollector.hasFace(bitmap))
    }

    @Test
    fun testPhotoWithoutFace() {
        // Carregar uma imagem sem rosto
        val bitmap: Bitmap = BitmapFactory.decodeResource(context.resources, R.drawable.photo_without_face)
        assertTrue(!photoCollector.hasFace(bitmap))
    }
}