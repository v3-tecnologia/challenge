package Database

import Datas.GyroscopeData

class GyroscopeRepository(private val gyroscopeDataDao: GyroscopeDataDao) {

    suspend fun GetAllGyroscope(): List<GyroscopeData> {
        return gyroscopeDataDao.GetAll() as List<GyroscopeData>
    }

    suspend fun insertGyroscopeData(gyroscopeData: GyroscopeData){
        gyroscopeDataDao.insert(gyroscopeData)
    }
}