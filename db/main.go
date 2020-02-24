package db

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func Connect() (*gorm.DB, error) {
	dialect, connString := dbConnectionString()

	if dialect == "" {
		dialect = "mysql"
	}

	db, err := gorm.Open(dialect, connString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()

	return db, nil
}

func dbConnectionString() (string, string) {
	dir, err := os.Getwd()
	dir = path.Join(dir, "configuration")
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(dir)
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	dialect := viper.GetString("database.dialect")

	username, password, host, dbName := viper.GetString("database.db_username"), viper.GetString("database.db_password"), viper.GetString("database.db_host"), viper.GetString("database.db_database")
	port := viper.GetString("database.db_port")
	chartset := viper.GetString("database.charset")
	if chartset == "" {
		chartset = "utf8"
	}

	if port == "" {
		return dialect, fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True", username, password, host, dbName, chartset)
	}

	return dialect, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True", username, password, host, port, dbName, chartset)
}
