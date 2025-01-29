package com.renascienza.desafiov3.location

import android.content.Context
import androidx.test.core.app.ApplicationProvider
import org.junit.Assert.assertNotNull
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RobolectricTestRunner

@RunWith(RobolectricTestRunner::class)
class GpsCollectorTest {
    private lateinit var context: Context
    private lateinit var gpsCollector: GpsCollector

    @Before
    fun setup() {
        context = ApplicationProvider.getApplicationContext()
        gpsCollector = GpsCollector(context)
    }

    @Test
    fun testGpsDataCollection() {
        gpsCollector.getLocation { data ->
            assertNotNull(data)
        }
    }
}