package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	instance *sql.DB
}

func Open(filename string, scheme string) (*DB, error) {
	instance, err := sql.Open("sqlite3", filename)

	if err != nil {
		return nil, err
	}

	_, err = instance.Exec(scheme)

	if err != nil {
		return nil, err
	}

	return &DB{instance: instance}, nil
}

func (db *DB) Close() {
	err := db.instance.Close()

	if err != nil {
		log.Fatal(err)
	}
}

type Table struct {
	db         *DB
	table      string
	attributes []string
}

func (db *DB) Table(table string, attributes ...string) *Table {
	return &Table{
		db:         db,
		table:      table,
		attributes: attributes,
	}
}

func (t *Table) Insert(values ...interface{}) error {
	if len(values) == 0 {
		panic("no values to insert into table " + t.table)
	}

	attrs := strings.Join(t.attributes, ",")

	valuesStr := strings.Repeat("?,", len(values))
	valuesStr = valuesStr[:len(valuesStr)-1]
	query := fmt.Sprintf("INSERT OR IGNORE INTO %v (%v) VALUES (%v)", t.table, attrs, valuesStr)

	_, err := t.db.instance.Exec(query, values...)
	return err
}

func (t *Table) Delete(condition string, values ...interface{}) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE %v", t.table, condition)
	_, err := t.db.instance.Exec(query, values...)
	return err
}

func (t *Table) HasRow(condition string, values ...interface{}) bool {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE %v", t.table, condition)
	row := *t.db.instance.QueryRow(query, values...)
	var count int
	err := row.Scan(&count)

	if err != nil {
		log.Fatal(err)
	}

	return count != 0
}

func (t *Table) Rows(scan func(func(...interface{})), condition string, values ...interface{}) error {
	selectValues := "*"

	if t.attributes != nil {
		selectValues = strings.Join(t.attributes, ",")
	}

	if condition == "" {
		condition = "1 = 1"
	}

	query := fmt.Sprintf("SELECT %v FROM %v WHERE %v", selectValues, t.table, condition)
	rows, err := t.db.instance.Query(query, values...)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		scan(func(values ...interface{}) {
			err = rows.Scan(values...)
		})

		if err != nil {
			return err
		}

	}

	return nil
}
