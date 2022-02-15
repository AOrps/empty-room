package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	toml "github.com/pelletier/go-toml/v2"
)


// Structures 
type Config struct {
	Username string
	Passwd   string
	DBNAME   string
	Table 	 string
}


type Test struct {
	ID    int `json:"id"`
	Title string `json:"title"`
}

func ParseCreds(path string) Config {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	toml_obj := string(f)

	var cfg Config
	err = toml.Unmarshal([]byte(toml_obj), &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func MysqlConnection(conf Config) *sql.DB {

	dbInfo := fmt.Sprintf("%s:%s@/%s", conf.Username, conf.Passwd, conf.DBNAME)

	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	return db
}


