package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	_ "github.com/lib/pq"
	"github.com/raihaninfo/hackernews/models"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
	view    *jet.Set
	session *scs.SessionManager
	model   models.Model
}
type server struct {
	host string
	port string
	url  string
}

func main() {

	migrate := flag.Bool("migrate", false, "should migrate - drop all tables")

	flag.Parse()

	server := server{
		host: "localhost",
		port: "8080",
		url:  "http://localhost:8080",
	}

	db2, err := openDB("postgres://dev:secret@localhost/hacker?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db2.Close()

	upper, err := postgresql.New(db2)
	if err != nil {
		log.Fatal(err)
	}
	defer func(upper db.Session) {
		err := upper.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(upper)

	// run migration
	if *migrate {
		fmt.Println("Running migration")
		err = runMigrate(upper)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Done running migration")
	}

	app := &application{
		appName: "Hacker News",
		server:  server,
		debug:   true,
		errLog:  log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog: log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime|log.Lshortfile),
		model:   models.New(upper),
	}

	// init jet template
	if app.debug {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"), jet.InDevelopmentMode())
	} else {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"))
	}

	// init session
	app.session = scs.New()
	app.session.Lifetime = 24 * time.Hour
	app.session.Cookie.Persist = true
	app.session.Cookie.Name = app.appName
	app.session.Cookie.Domain = app.server.host
	app.session.Cookie.SameSite = http.SameSiteDefaultMode
	app.session.Store = postgresstore.New(db2)

	err = app.listenServer()
	if err != nil {
		app.errLog.Fatal(err)
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func runMigrate(db db.Session) error {
	script, err := os.ReadFile("./migrations/table.sql")
	if err != nil {
		return err
	}
	_, err = db.SQL().Exec(string(script))
	if err != nil {
		return err
	}
	return err
}
