package database

import (
	"../DAO/database/entities"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbFile = "executor.sqlite3"
	driver = "sqlite3"
)

func InitDB() {
	ExecuteTransactionalDDL(Schema)
}

func GetContainerRun(id int) (*entities.ContainerRunEntity, error) {
	contRun := &entities.ContainerRunEntity{}
	// "SELECT id, dockerfile, script, environmentVariables FROM ContainerRun WHERE ID=?"
	err := ExecuteTransactionalSingleRowQuery(
		GetContainerRunStatement,
		[]interface{}{id},
		&contRun.ID,
		&contRun.Dockerfile,
		&contRun.Script,
		&contRun.EnvironmentVariables)
	return contRun, err
}

func InsertContainerRun(contRun *entities.ContainerRunEntity) error {
	return ExecuteTransactionalDDL(
		InsertContainerRunStatement,
		contRun.ID,
		contRun.Dockerfile,
		contRun.Script,
		contRun.EnvironmentVariables)
}

func DeleteContainerRun(id int) error {
	return ExecuteTransactionalDDL(DeleteContainerRunStatement, id)
}

func ExecuteTransactionalDDL(query string, args ...interface{}) error {
	transaction, err := getDB().Begin()
	defer transaction.Commit()
	if err != nil {
		return err
	}
	stmt, err := transaction.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(args...); err != nil {
		transaction.Rollback()
		return err
	}
	return nil
}

func ExecuteTransactionalSingleRowQuery(query string, selection []interface{}, targets ...interface{}) error {
	transaction, err := getDB().Begin()
	if err != nil {
		return err
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(query)
	if err != nil {
		return err
	}
	row := statement.QueryRow(selection...)
	if err := row.Scan(targets...); err != nil {
		transaction.Rollback()
		return err
	}
	return nil
}

var getDB = func() func() *sql.DB {
	db, err := sql.Open(driver, fmt.Sprintf("%v%v", utils.DatabaseDirectory(), dbFile))
	if err != nil {
		panic(err)
	}
	return func() *sql.DB {
		return db
	}
}()
