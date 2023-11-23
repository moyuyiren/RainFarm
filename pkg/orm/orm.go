package orm

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Config struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  int
}

type DB struct {
	*gorm.DB
}

// todo :自定义gorm日志管理
type ormLog struct {
	Loglevel logger.LogLevel
}

func (l *ormLog) LogMode(level logger.LogLevel) logger.Interface {
	l.Loglevel = level
	return l
}

func (l *ormLog) Info(ctx context.Context, format string, v ...interface{}) {
	if l.Loglevel < logger.Info {
		return
	}
	logx.WithContext(ctx).Infof(format, v...)
}

func (l *ormLog) Warn(ctx context.Context, fromat string, v ...interface{}) {
	if l.Loglevel < logger.Warn {
		return
	}
	logx.WithContext(ctx).Infof(fromat, v...)
}

func (l *ormLog) Error(ctx context.Context, format string, v ...interface{}) {
	if l.Loglevel < logger.Error {
		return
	}
	logx.WithContext(ctx).Errorf(format, v...)
}

// todo: Trace追踪处理
func (l *ormLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	logx.WithContext(ctx).WithDuration(elapsed).Infof("[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
}

// todo :NewDbClient
func NewMysql(conf *Config) (*DB, error) {
	if conf.MaxIdleConns == 0 {
		conf.MaxIdleConns = 10
	}
	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = 100
	}
	if conf.MaxLifetime == 0 {
		conf.MaxLifetime = 3600
	}
	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{
		Logger: &ormLog{},
	})
	if err != nil {
		return nil, err
	}
	sdb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sdb.SetMaxOpenConns(conf.MaxOpenConns)
	sdb.SetMaxIdleConns(conf.MaxIdleConns)
	sdb.SetConnMaxLifetime(time.Second * time.Duration(conf.MaxLifetime))

	err = db.Use(NewCustomePlugin())
	if err != nil {
		return nil, err
	}
	return &DB{DB: db}, nil
}

func MustNewMysql(conf *Config) *DB {
	db, err := NewMysql(conf)
	if err != nil {
		panic(err)
	}

	return db
}
