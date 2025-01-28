package com.example.mvp

import android.annotation.SuppressLint
import android.content.Context
import android.hardware.Sensor
import android.hardware.SensorManager
import android.location.Location
import androidx.arch.core.executor.testing.InstantTaskExecutorRule
import androidx.test.core.app.ApplicationProvider
import com.google.android.gms.location.FusedLocationProviderClient
import com.google.ar.core.Config
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.test.StandardTestDispatcher
import kotlinx.coroutines.test.resetMain
import kotlinx.coroutines.test.runTest
import kotlinx.coroutines.test.setMain
import org.junit.After
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.mockito.Mock
import org.mockito.Mockito.`when`
import org.mockito.MockitoAnnotations
import org.mockito.junit.MockitoJUnitRunner
import org.robolectric.annotation.Config


@RunWith(MockitoJUnitRunner::class)
@Config(manifest= Config.NONE)
class DataCollectorTest {

    @get:Rule
    var instantTaskExecutorRule = InstantTaskExecutorRule()

    @OptIn(ExperimentalCoroutinesApi::class)
    private val testDispatcher = StandardTestDispatcher()

    @Mock
    private lateinit var context: Context

    @Mock
    private lateinit var sensorManager: SensorManager

    @Mock
    private lateinit var fusedLocationProviderClient: FusedLocationProviderClient

    @Mock
    private lateinit var gyroscopeSensor: Sensor

    private lateinit var dataCollector: DataCollector

    @Before
    fun setup() {
        MockitoAnnotations.openMocks(this)
        Dispatchers.setMain(testDispatcher)
        `when`(context.getSystemService(Context.SENSOR_SERVICE)).thenReturn(sensorManager)
        `when`(sensorManager.getDefaultSensor(Sensor.TYPE_GYROSCOPE)).thenReturn(gyroscopeSensor)
        `when`(context.applicationContext).thenReturn(ApplicationProvider.getApplicationContext())
        dataCollector = DataCollector(context)
    }

    @After
    fun tearDown() {
        Dispatchers.resetMain()
    }

    @Test
    fun collectGyroData_returnsEmptyListWhenNoSensor() = runTest {
        // Arrange
        `when`(sensorManager.getDefaultSensor(Sensor.TYPE_GYROSCOPE)).thenReturn(null)

        // Act
        val gyroData = dataCollector.collectGyroData()

        // Assert
        assertEquals(emptyList<Float>(), gyroData)
    }

    @SuppressLint("MissingPermission")
    @Test
    fun collectGpsData_returnsLocation() = runTest {
        // Arrange
        val expectedLocation = Location("test")
        expectedLocation.latitude = 37.7749
        expectedLocation.longitude = -122.4194

        `when`(fusedLocationProviderClient.lastLocation).thenReturn(com.google.android.gms.tasks.Tasks.forResult(expectedLocation))

        // Act
        val location = dataCollector.collectGpsData()

        // Assert
        assertEquals(expectedLocation, location)
    }

    @Test
    fun capturePhoto_returnsPhotoPath() {
        // Arrange
        val expectedPhotoPath = context.filesDir.absolutePath + "/photo.jpg"

        // Act
        val photoData = dataCollector.capturePhoto()

        // Assert
        assertEquals(expectedPhotoPath, photoData)
    }

    @Test
    fun detectFaceInPhoto_returnsFaceData() {
        // Arrange
        val photoPath = context.filesDir.absolutePath + "/photo.jpg"

        // Act
        val faceData = dataCollector.detectFaceInPhoto(photoPath)

        // Assert
        assertEquals("_", faceData) // Assuming no face is detected in the test image
    }

    @Test
    fun collectData_returnsData() = runTest {
        // Arrange
        val expectedGyroData = emptyList<Float>()
        val expectedGpsData = 0.0 // Assuming latitude and longitude are 0.0
        val expectedPhotoData = context.filesDir.absolutePath + "/photo.jpg"
        val expectedFaceData = "_" // Assuming no face is detected in the test image

        // Act
        val data = dataCollector.collectData()

        // Assert
        assertEquals(expectedGyroData, data.gyroData)
        assertEquals(expectedGpsData, data.gpsData)
        assertEquals(expectedPhotoData, data.photoData)
        assertEquals(expectedFaceData, data.faceData)
    }
}