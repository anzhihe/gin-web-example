package model

const (
	// 由于需要连接多个数据库，创建事务时，通过数据库类型来区分不同的数据库
	DBTypeODB = "odb"
	DBTypeIDB = "Idb"
	DBTypeRDB = "rdb"
)
