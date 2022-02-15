package main

import (
	"strconv"
	"net/http"
	"fmt"
	"log"

	dbutil "github.com/AOrps/empty-room/src/back-end/DB"
)

/* Global Variable Declaration*/
var PORT int = 9990


func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "Send POST request")
	} else {
		// fmt.Fprintf(w, "POST request successful")

		// Credentials for MYSQL Service
		cfg := dbutil.ParseCreds("creds.toml")

		conn := dbutil.MysqlConnection(cfg)

		defer conn.Close()

		statement := fmt.Sprintf("SELECT * from %s", cfg.Table)

		query, err := conn.Query(statement)

		if err != nil {
			fmt.Println("Here")
			log.Fatal(err.Error())
		}

		for query.Next() {
			var tag dbutil.Test
			err = query.Scan(&tag.ID, &tag.Title)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Printf("%d: %s\n",tag.ID, tag.Title )
			fmt.Fprintf(w,"%d: %s\n", tag.ID, tag.Title)
		}

	}
}


func main() {

	port := strconv.Itoa(PORT) 

	http.HandleFunc("/", rootHandler)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
