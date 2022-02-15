package main

import (
	dbutil "github.com/AOrps/empty-room/src/back-end/DB"
)


func main() {

	cfg := dbutil.ParseCreds("creds.toml")

	conn := dbutil.MysqlConnection(cfg)

	statement := fmt.Sprintf("SELECT * from %s;", cfg.Table)

	query, err := conn.Query(statement)
	if err != nil {
		log.Fatal(err.Error())
	}

	for query.Next() {
		var tag dbutil.Test
		err = query.Scan(&tag.ID, &tag.Title)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("%d: %s",tag.ID, tag.Title )
	}
}
