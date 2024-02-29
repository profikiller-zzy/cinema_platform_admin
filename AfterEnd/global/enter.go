package global

// global 全局变量
import (
	"AfterEnd/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config   *config.Config
	Log      *logrus.Logger
	MysqlLog logger.Interface
	Db       *gorm.DB
)
