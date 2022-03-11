package dao

import (
	"context"
	"log"
	"os"
	"thor-backend/internal/model"
	"thor-backend/internal/setting"
	"time"

	"go.uber.org/zap"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"

	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"github.com/go-redis/redis/v8"
)

type Dao struct {
	odb *gorm.DB
	idb *gorm.DB
	rdb *gorm.DB
	rds *redis.Client
}

var c = setting.Conf

func Init() *Dao {

	odb, err := initDB(c.ODB)
	if err != nil {
		panic(err)
	}

	idb, err := initDB(c.IDB)
	if err != nil {
		panic(err)
	}

	rdb, err := initDB(c.RDB)
	if err != nil {
		panic(err)
	}

	rds, err := newPool(c.RDS)
	if err != nil {
		panic(err)
	}

	defer rds.Close()

	return &Dao{
		odb: odb,
		idb: idb,
		rdb: rdb,
		rds: rds,
	}
}

// 初始化
func initDB(c *setting.DBConfig) (*gorm.DB, error) {

	var dbLogger logger.Interface
	if c.Debug {
		dbLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // 禁用彩色打印
			},
		)
	}

	db, err := gorm.Open(mysql.Open(c.Addr), &gorm.Config{Logger: dbLogger})
	if err != nil {
		return nil, errors.Wrap(err, "new mysql engine. addr:"+c.Addr)
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "get sql.DB")
	}
	defer sqlDB.Close()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func (d *Dao) Close() {
	zap.L().Info("dao shutdown...")
}

// 新建事务
func (d *Dao) NewTx(dbType ...string) (*gorm.DB, error) {
	dbType = append(dbType, model.DBTypeODB)
	var tx *gorm.DB
	switch dbType[0] {
	case model.DBTypeODB:
		tx = d.odb.Begin()
	case model.DBTypeIDB:
		tx = d.idb.Begin()
	case model.DBTypeRDB:
		tx = d.rdb.Begin()
	default:
		return nil, errors.New("illegal db type: " + dbType[0])
	}
	return tx, tx.Error
}

// 关闭事务
func (d *Dao) CloseTx(tx *gorm.DB, txErr *error) {
	if txErr != nil && *txErr != nil {
		tx.Rollback()
		return
	}
	if err := tx.Commit().Error; err != nil {
		zap.L().Error("关闭事务失败", zap.Error(err))
	}
}

// 初始化redis连接
func newPool(c *setting.RedisConfig) (*redis.Client, error) {
	rds := redis.NewClient(&redis.Options{
		Addr:       c.Addr,
		Password:   c.Password, // no password set
		DB:         0,          // use default DB
		PoolSize:   c.PoolSize, // 连接池大小
		MaxRetries: 3,          // 连接重试次数
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rds.Ping(ctx).Result()
	return rds, err
}
