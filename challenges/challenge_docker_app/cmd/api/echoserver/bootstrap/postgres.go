package bootstrap

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

const (
	user   = "postgres_user"
	pass   = "postgres_pass"
	dbName = "learning_go_db"
)

type postgresCtx struct {
	user   string `default:"postgres_user"`
	pass   string `default:"postgres_pass"`
	dbName string `default:"learning_go_db"`
}

func newPostgresCtx() *postgresCtx {
	return &postgresCtx{
		user:   user,
		pass:   pass,
		dbName: dbName,
	}
}

func (postgresCtx *postgresCtx) getConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@database:5432/%s?sslmode=disable", postgresCtx.user, postgresCtx.pass, postgresCtx.dbName)
	// return fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", postgresCtx.user, postgresCtx.pass, postgresCtx.dbName)
}

func GetDBConnection() *sql.DB {
	postgresCtx := newPostgresCtx()
	db, err := sql.Open("postgres", postgresCtx.getConnectionString())
	if err != nil {
		panic(" Myads: Error, We can't open a connection to the database")
	}
	testConnection(db, 0)
	return db
}

func testConnection(db *sql.DB, tries int) {
	if tries == 10 {
		panic("Error, We can't query to database")
	}
	var result string
	err := db.QueryRow("SELECT version()").Scan(&result)
	if err != nil {
		fmt.Println(">>>  Myads: Trying to connect to Postgres, try #", tries, err)
		tries++
		time.Sleep(3 * time.Second)
		testConnection(db, tries)
	}
	fmt.Println(">>> Myads: Connecting to Postgres with success!!. PostgreSQL Version:", result)
}
