package database

import (
	"fmt"
	"os"
	"strconv"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Log      *zap.Logger
}

func New() *Database {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	log := zap.NewNop()

	return &Database{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_DATABASE"),
		Log:      log,
	}
}

func (database *Database) Connect() (*gorm.DB, error) {
	database.Log.Info("Connecting to database with dsn:", zap.String("dsn", database.databaseSourceName()))

	con, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/go_projee?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	con.Debug()

	return con, err
}

func (database *Database) RunMigrations(connection *gorm.DB, models ...interface{}) error {
	err := connection.AutoMigrate(models...)
	return err
}

func (database *Database) databaseSourceName() string {
	//the format for the database source name values are represented below
	//username:password@protocol(address)/dbname?param=value
	return fmt.Sprintf("%v:@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", database.User, database.Host, database.Port, database.Name)
}
