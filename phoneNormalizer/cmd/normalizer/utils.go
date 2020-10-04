package main

import (
	"database/sql"
	"fmt"
)

func createDB(db *sql.DB, dbName string) error {
	_, err := db.Exec("CREATE DATABASE " + dbName)
	return err
}

func resetDB(dbName string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS " + dbName)
	if err != nil {
		return err
	}
	return createDB(db, dbName)
}

func createTable(db *sql.DB) error {
	statement := `
    DROP TABLE phone_numbers;
    CREATE TABLE phone_numbers (
      id SERIAL,
      value VARCHAR(255)
    )`
	_, err := db.Exec(statement)
	return err
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	var id int
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func initDB(phones []string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable dbname=%s", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	defer db.Close()
	if err != nil {
		return err
	}
	if err = createTable(db); err != nil {
		return err
	}
	for _, phone := range phones {
		if _, err = insertPhone(db, phone); err != nil {
			return err
		}
	}
	return nil
}
