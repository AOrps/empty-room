package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	toml "github.com/pelletier/go-toml/v2"
)

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

func main() {

	f, err := ioutil.ReadFile("creds.toml")
	if err != nil {
		log.Fatal(err)
	}

	toml_obj := string(f)

	var cfg Config
	err = toml.Unmarshal([]byte(toml_obj), &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("username: [%s]", cfg.Username)
	// fmt.Println("passwd:", cfg.Passwd)

	dbInfo := fmt.Sprintf("%s:%s@/%s", cfg.Username, cfg.Passwd, cfg.DBNAME) // cfg.Protocol, cfg.Address, cfg.DBNAME)

	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	statement := fmt.Sprintf("SELECT * from %s;", cfg.Table)

	query, err := db.Query(statement)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(query)

	for query.Next() {
		var tag Test
		err = query.Scan(&tag.ID, &tag.Title)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("%d: %s",tag.ID, tag.Title )
	}
}
