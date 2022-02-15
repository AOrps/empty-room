package backend

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	toml "github.com/pelletier/go-toml/v2"
)

type Config struct {
	Username string
	Passwd   string
	Protocol string
	Address  string
	DBNAME   string
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

	dbInfo := fmt.Sprintf("%s:%s@%s(%s)/%s", cfg.Username, cfg.Passwd, cfg.Protocol, cfg.Address, cfg.DBNAME)

	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	query := fmt.Sprintf("SELECT * from %s", cfg.DBNAME)

	ha, err := db.Query(query)

	fmt.Println(ha)
}
