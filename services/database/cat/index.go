package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"restful-mvc/config"
)

var db_provider = config.DB_PROVIDER
var db_auth = "host=" + config.DB_HOST + 
				" port=" + config.DB_PORT +
				" user=" + config.DB_USER +
				" dbname=" + config.DB_NAME +
				" password=" + config.DB_PASSWORD +
				" sslmode=" + config.SSL_MODE

func GetAll() {
	db, err := gorm.Open(db_provider, db_auth)
	if err != nil {
        fmt.Println(err)
    }
	fmt.Println("connect database success!!!")
  	defer db.Close()
}