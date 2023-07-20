package bootstrap

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type postgresCtx struct {
	user   string
	pass   string
	dbName string
	dbHost string
}

func newPostgresCtx(envHandler EnvHandler) *postgresCtx {
	return &postgresCtx{
		user:   envHandler.DbUser(),
		pass:   envHandler.DbPass(),
		dbName: envHandler.DbName(),
		dbHost: envHandler.DbHost(),
	}
}

func (postgresCtx *postgresCtx) getConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", postgresCtx.user, postgresCtx.pass, postgresCtx.dbHost, postgresCtx.dbName)
}

func GetDBConnectionOrNil(envHandler EnvHandler) *sql.DB {
	if !envHandler.EnableDatabase() {
		fmt.Println(">>> Running my-app Without POSTGRES DB...")
		return nil
	}
	fmt.Println(">>> Running my-app With POSTGRES DB...")
	return GetDBConnection(envHandler)
}

func GetDBConnection(envHandler EnvHandler) *sql.DB {
	postgresCtx := newPostgresCtx(envHandler)
	db, err := sql.Open("postgres", postgresCtx.getConnectionString())
	if err != nil {
		panic(" Myads: Error, We can't open a connection to the database")
	}
	testConnection(db, 0)
	return db
}

func testConnection(db *sql.DB, retryNumber int) {
	if retryNumber == 10 {
		panic("Error, We can't query to database")
	}
	var result string
	err := db.QueryRow("SELECT version()").Scan(&result)
	if err != nil {
		fmt.Println(">>>  Myads: Trying to connect to Postgres, try #", retryNumber, err)
		retryNumber++
		time.Sleep(3 * time.Second)
		testConnection(db, retryNumber)
	}
	fmt.Println(">>> Myads: Connecting to Postgres with success!!. PostgreSQL Version:", result)
}
