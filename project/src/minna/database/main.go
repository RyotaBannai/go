package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/k0kubun/pp"
	"log"
	_ "github.com/lib/pq"
)

type Slicer interface {
	toSlice() []interface{}
}

type Person struct {
	id   string
	name string
	age  int
}

func (p *Person) toSlice() []interface{} {
	return []interface{}{p.id, p.name, p.age}
}

type queries interface {
	selectAll() queries
	insert() queries
	update() queries
	delete() queries
	run() (interface{}, error)
}

type PeopleQueries struct {
	currentQuery string
	person       Person
	db           *sql.DB
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

func (p *PeopleQueries) run() (interface{}, error) {
	if p.currentQuery == "" {
		return nil, errors.New("Needs to query string before run.\n")
	}
	result, err := p.db.Exec(p.currentQuery, p.person.toSlice()...)
	//ri, err := result.LastInsertId()
	//ra, err := result.RowsAffected()
	//fmt.Printf("%v\n%v\n", ri, ra)
	return result, err
}

func test(db *sql.DB) {
	peopleQuery := PeopleQueries{
		person:       Person{"1006", "Mark", 35},
		db:           db,
		currentQuery: ""}
	result, err := peopleQuery.insert().run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func main() {
	connStr := "postgres://ryota:bannai@localhost/gotest?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM people")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		p := Person{}
		err = rows.Scan(&p.id, &p.name, &p.age)
		pp.Println(p)
	}
}
