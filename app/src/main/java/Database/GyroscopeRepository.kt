package Database

import Datas.GyroscopeData

class GyroscopeRepository(private val gyroscopeDataDao: GyroscopeDataDao) {

    suspend fun GetAllGyroscope(): List<GyroscopeData>? {
        return gyroscopeDataDao.GetAll().value
    }

    suspend fun insertGyroscopeData(gyroscopeData: GyroscopeData){
        gyroscopeDataDao.insert(gyroscopeData)
    }
}