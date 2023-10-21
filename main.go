package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/proullon/ramsql/driver"
)

func main() {
	db, err := sql.Open("ramsql", "goimdb")
	if err != nil {
		log.Fatal(err)
		return
	}

	createTable := `CREATE TABLE IF NOT EXISTS goimdb (
	id INT AUTO_INCREMENT,
	imdbID TEXT NOT NULL UNIQUE,
	title TEXT NOT NULL,
	year int NOT NULL,
	rating float NOT NULL,
	isSuperHero BOOLEAN NOT NULL,
	PRIMARY KEY (id)
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Create table error:", err)
		return
	}

	fmt.Println("Database created successfully")
	insert := `
	INSERT INTO goimdb (imdbID, title, year, rating, isSuperHero) 
	VALUES (?, ?, ?, ?, ?);
	`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal("Prepare statement error:", err)
		return
	}

	r, err := stmt.Exec("tt4154796", "Avengers: Infinity War", 2018, 8.5, true)
	if err != nil {
		log.Fatal("insert statement error:", err)
		return
	}

	l, err := r.LastInsertId()
	fmt.Println("Last insert id:", l, "err:", err)
	ef, err := r.RowsAffected()
	fmt.Println("Rows affected:", ef, "err:", err)

	row, err := db.Query("SELECT * FROM goimdb")
	if err != nil {
		log.Fatal("Query error:", err)
		return
	}

	for row.Next() {
		var id int
		var imdbID string
		var title string
		var year int
		var rating float32
		var isSuperHero bool
		err = row.Scan(&id, &imdbID, &title, &year, &rating, &isSuperHero)
		if err != nil {
			log.Fatal("Scan error:", err)
			return
		}
		fmt.Println(id, imdbID, title, year, rating, isSuperHero)
	}

	stm2, err := db.Prepare(`
	UPDATE goimdb 
	SET rating=? 
	WHERE imdbID = ?
	`)
	_, err = stm2.Exec(9.0, "tt4154796")
	if err != nil {
		log.Fatal("Update error:", err)
		return
	}
}

//tt4154796
//tt4154756
//tt4154664
