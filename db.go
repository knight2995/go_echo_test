package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	const tmpl = (`duckbo:@{{.}}@tcp(139.150.64.36:3306)/echo_test_db`)
	t := template.New("db_access")
	t, err := t.Parse(tmpl)

	dbPassword, _ := ioutil.ReadFile("config")
	var tpl bytes.Buffer

	t.Execute(&tpl, string(dbPassword))

	fmt.Println(tpl.String())

	db, err := sql.Open("mysql", tpl.String())

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//...(db 사용)....
	/*
		result, err := db.Exec("INSERT INTO test(name) VALUES (?)", "wowa2")
		fmt.Println(result)
		if err != nil {
			log.Fatal(err)
		}
	*/
	// 복수 Row를 갖는 SQL 쿼리
	var id int
	var name string
	rows, err := db.Query("SELECT idx, name FROM test")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("{? ?}", id, name)
	}

}
