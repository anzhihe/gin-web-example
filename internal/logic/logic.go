package logic

import (
	"fmt"
	"thor-backend/internal/dao"
	"thor-backend/internal/logger"
	"thor-backend/internal/setting"
	"thor-backend/pkg/log"

	"github.com/go-playground/validator/v10"
)

type Logic struct {
	dao      *dao.Dao // 全局数据库驱动
	validate *validator.Validate
}

func Init(d *dao.Dao) *Logic {

	// 日志加载
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Env); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		panic(err)
	}

	// 校验器加载
	v := validator.New()
	v.SetTagName("binding")

	l := &Logic{
		dao:      d,
		validate: v,
	}

	// 同步任务加载

	// 异步任务加载

	return l
}

func (l *Logic) Close() {
	log.Info("logic shutdown...")
}

// 处理查询数据库时分页的偏移量
func fillOffset(page, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = setting.Conf.DefaultPageSize
	}
	if limit > setting.Conf.MaxPageSize {
		limit = setting.Conf.MaxPageSize
	}

	return page, limit
}
