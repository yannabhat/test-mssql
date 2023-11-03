package database

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"test-mssql/load"

	_ "github.com/microsoft/go-mssqldb"
)

var Conn *sql.DB

var (
	server   = flag.String("server", "192.168.1.31\\SQL2019", "the database server")
	user     = flag.String("user", "sa", "the database user")
	password = flag.String("password", "Warehousedbo2023", "the database password")
	database = flag.String("database", "similanwls_mahajak", "the database name")
)

func Connection() {
	var err error
	load.LoadEnv()
	flag.Parse()
	// connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;", *server, *user, *password, *database)
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;", *flag.String("server", os.Getenv("server"), "the database server"), *flag.String("user", os.Getenv("user"), "the database user"), *flag.String("password", os.Getenv("password"), "the database password"), *flag.String("database", os.Getenv("database"), "the database name"))
	Conn, err = sql.Open("mssql", connString)
	if err != nil {
		log.Fatal(err)
	}

	// defer Conn.Close()

}
