package main

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type dbSqLite struct {
	handler *sql.DB
}

func (db *dbSqLite) Create(t Task) error {

	sql_additem := `
    INSERT OR REPLACE INTO tasks(
    alias,
    description,
    task_type,
    tags,
    timestamp,
    estimate_time,
    real_time,
    reminders,
    InsertedDatetime
    ) values(?, ?, ?,?,?,?,?,?, CURRENT_TIMESTAMP)
    `

	stmt, err := db.handler.Prepare(sql_additem)
	if err != nil {
		panic(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Alias, t.Description, t.Task_type, strings.Join(t.Tags, ","), t.Timestamp, t.Estimate_time, t.Real_time, strings.Join(t.Reminders, ","))
	if err != nil {
		panic(err)
		return err
	}
	return nil

}

func (db *dbSqLite) ReadById(id *int64) (TaskList, error) {

	sql_readall := `
    SELECT id, alias, description, task_type, timestamp, estimate_time, real_time, tags, reminders FROM tasks
    WHERE id=?
    ORDER BY datetime(InsertedDatetime) DESC
    `
	stmt, err := db.handler.Prepare(sql_readall)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sql_readall, id)

	if err != nil {
		panic(err)
		return nil, err
	}

	defer rows.Close()

	var result TaskList
	var dbTags, dbReminders string

	for rows.Next() {
		t := Task{}
		err = rows.Scan(&t.Id, &t.Alias, &t.Description, &t.Task_type, &t.Timestamp, &t.Estimate_time, &t.Real_time, &dbTags, &dbReminders)
		if err != nil {
			panic(err)
			return nil, err
		}
		t.Tags = strings.Split(dbTags, ",")
		t.Reminders = strings.Split(dbReminders, ",")
		result = append(result, t)
	}
	return result, err

}

func (db *dbSqLite) ReadByAlias(alias *string) (TaskList, error) {
	sql_readall := `
    SELECT id, alias, description, task_type, timestamp, estimate_time, real_time, tags, reminders FROM tasks
    WHERE alias=?
    ORDER BY datetime(InsertedDatetime) DESC
    `
	stmt, err := db.handler.Prepare(sql_readall)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sql_readall, alias)

	if err != nil {
		panic(err)
		return nil, err
	}

	defer rows.Close()

	var result TaskList
	var dbTags, dbReminders string

	for rows.Next() {
		t := Task{}
		err = rows.Scan(&t.Id, &t.Alias, &t.Description, &t.Task_type, &t.Timestamp, &t.Estimate_time, &t.Real_time, &dbTags, &dbReminders)
		if err != nil {
			panic(err)
			return nil, err
		}
		t.Tags = strings.Split(dbTags, ",")
		t.Reminders = strings.Split(dbReminders, ",")
		result = append(result, t)
	}
	return result, err
}

func (db *dbSqLite) Update(t Task) error {
	sql_updateitem := `
    INSERT OR REPLACE INTO tasks(
    id,
    alias,
    description,
    task_type,
    tags,
    timestamp,
    estimate_time,
    real_time,
    reminders,
    InsertedDatetime
    ) values(?, ?, ?,?,?,?,?,?,?, CURRENT_TIMESTAMP)
    `

	stmt, err := db.handler.Prepare(sql_updateitem)
	if err != nil {
		panic(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Id, t.Alias, t.Description, t.Task_type, strings.Join(t.Tags, ","), t.Timestamp, t.Estimate_time, t.Real_time, strings.Join(t.Reminders, ","))
	if err != nil {
		panic(err)
		return err
	}
	return nil
}
func (db *dbSqLite) Delete(t Task) error {
	sql_delete := `DELETE FROM tasks WHERE id=?`
	stmt, err := db.handler.Prepare(sql_delete)
	if err != nil {
		panic(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Query(sql_delete, t.Id)

	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func (db *dbSqLite) ReadAll() (TaskList, error) {
	sql_readall := `
    SELECT id, alias, description, task_type, timestamp, estimate_time, real_time, tags, reminders FROM tasks
    ORDER BY datetime(InsertedDatetime) DESC
    `

	rows, err := db.handler.Query(sql_readall)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()

	var result TaskList
	var dbTags, dbReminders string

	for rows.Next() {
		t := Task{}

		err = rows.Scan(&t.Id, &t.Alias, &t.Description, &t.Task_type, &t.Timestamp, &t.Estimate_time, &t.Real_time, &dbTags, &dbReminders)
		if err != nil {
			panic(err)
			return nil, err
		}
		t.Tags = strings.Split(dbTags, ",")
		t.Reminders = strings.Split(dbReminders, ",")
		result = append(result, t)
	}
	return result, err
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
    CREATE TABLE IF NOT EXISTS tasks(
    Id TEXT NOT NULL PRIMARY KEY,
    alias TEXT,
    description TEXT,
    tags Text,
    timestamp int,
    estimate_time DATETIME,
    real_time DATETIME,
    reminders TEXT,
    InsertedDatetime DATETIME
    );
    `

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}
