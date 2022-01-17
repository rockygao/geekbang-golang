package database

import (
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Data 数据库连接包装对象
type Data struct {
	database *gorm.DB
}

// ServerHttpAddr: viper.GetString("server.http.addr"),
// DatabaseDriver: viper.GetString("data.database.driver"),
// DatabaseDsn:    viper.GetString("data.database.source"),

// Gdb 返回一个数据库连接
func (p *Data) Gdb(debug ...bool) *gorm.DB {
	if len(debug) > 0 && debug[0] {
		return p.database.Debug()
	}
	return p.database
}

func InitDB(v *viper.Viper) (*Data, error) {
	db, err := gorm.Open(mysql.Open(v.GetString("data.database.source")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// return db, nil
	return &Data{
		database: db,
	}, nil
}
