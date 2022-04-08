package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func GetConnectionStringFromJSON() string {
	// Open our jsonFile
	jsonFile, err := os.Open("database/dbconfig.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Error reading dbconfig file")
	}
	fmt.Println("Successfully Opened dbconfig.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	//Read all json file
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//Instance to configuration struct
	var connection ConnectionConfig

	json.Unmarshal(byteValue, &connection)

	//Concatenate connection string
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		connection.Host, connection.User, connection.Password, connection.Dbname, connection.Port, connection.Sslmode)

	return connectionString
}

func ConnectDatabase() {
	connectionString := GetConnectionStringFromJSON()
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error connecting to database")
	}
	DB.AutoMigrate()
}

type ConnectionConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     string `json:"port"`
	Sslmode  string `json:"sslmode"`
}
