package Datas

import java.sql.Timestamp

open class BaseData(
    open val id: Long,
    open val timestamp: Timestamp,
    open val deviceId: String,
)