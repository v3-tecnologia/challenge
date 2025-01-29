package com.renascienza.desafiov3.sensor

import android.content.Context
import androidx.test.core.app.ApplicationProvider
import org.junit.Assert.assertNotNull
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RobolectricTestRunner

@RunWith(RobolectricTestRunner::class)
class GyroscopeCollectorTest {
    private lateinit var context: Context
    private lateinit var gyroscopeCollector: GyroscopeCollector

    @Before
    fun setup() {
        context = ApplicationProvider.getApplicationContext()
        gyroscopeCollector = GyroscopeCollector(context)
    }

    @Test
    fun testGyroscopeDataCollection() {
        gyroscopeCollector.startCollecting()
        val data = gyroscopeCollector.getGyroscopeData()
        assertNotNull(data)
    }
}