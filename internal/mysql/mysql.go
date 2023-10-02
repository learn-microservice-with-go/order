package mysql

import (
	"fmt"

	config "github.com/learn-microservice-with-go/user_microservice/internal/config"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Provider = wire.NewSet(NewMySQL)

func NewMySQL(config *config.Config) (*gorm.DB, error) {
	dbAddr := fmt.Sprintf("%s:%s", config.MySQLHost, config.MySQLPort)
	mysqlClient, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.MySQLUser, config.MySQLPassword, dbAddr, config.MySQLDbName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return mysqlClient, nil
}
