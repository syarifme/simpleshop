package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func Connect() (*sql.DB, error) {
	connString := dbConnectionString()
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

func dbConnectionString() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("../configuration")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	username, password, host, dbName := viper.GetString("database.db_username"), viper.GetString("database.db_password"), viper.GetString("database.db_host"), viper.GetString("database.db_database")
	port := viper.GetInt("database.db_port")

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbName)
}
