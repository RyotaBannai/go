package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/k0kubun/pp"
	"log"
	_ "github.com/lib/pq"
)

/*
	・go method chain
		https://qiita.com/roba4coding/items/4180923bf7a44364e414
*/

type Slicer interface {
	toSlice() []interface{}
}

type Person struct {
	id   string
	name string
	age  int
}

func (p *Person) toSlice() []interface{} {
	return []interface{}{&p.id, &p.name, &p.age}
}

type queries interface {
	selectAll() queries
	insert() queries
	update() queries
	delete() queries
	run(db *sql.DB) (interface{}, error)
}

type PeopleQueries struct {
	currentQuery string
	person       Person
}

func (p *PeopleQueries) selectAll() queries {
	p.currentQuery = `SELECT * FROM people;`
	return p
}

func (p *PeopleQueries) insert() queries {
	p.currentQuery = `INSERT INTO people(id, name, age) VALUES($1, $2, $3)`
	return p
}

func (p *PeopleQueries) update() queries {
	p.currentQuery = `UPDATE people where...`
	return p
}

func (p *PeopleQueries) delete() queries {
	p.currentQuery = `DELETE people where...`
	return p
}

func (p *PeopleQueries) run(db *sql.DB) (interface{}, error) {
	if p.currentQuery == "" {
		return nil, errors.New("Needs to query string before run.\n")
	}
	result, err := db.Exec(p.currentQuery, p.person.toSlice()...)
	//ri, err := result.LastInsertId()
	//ra, err := result.RowsAffected()
	//fmt.Printf("%v\n%v\n", ri, ra)
	return result, err
}

func test(db *sql.DB) {
	var people queries
	people = &PeopleQueries{person: Person{"1006", "Mark", 35}}
	result, err := people.insert().run(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func testSelect(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM people")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		p := Person{}
		err = rows.Scan(p.toSlice()...)
		pp.Println(p)
	}
}

func main() {
	connStr := "postgres://ryota:bannai@localhost/gotest?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	test(db)
}
