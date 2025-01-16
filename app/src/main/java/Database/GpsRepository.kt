package Database

import Datas.GpsData

class GpsRepository(private val gpsDataDao: GpsDataDao) {

    suspend fun GetAllGpsData(): List<GpsData> {
        return gpsDataDao.GetAll() as List<GpsData>
    }

    suspend fun insertGpsData(gpsData: GpsData){
        gpsDataDao.insert(gpsData)
    }
}