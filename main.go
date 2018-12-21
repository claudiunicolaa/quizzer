package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"

	"github.com/calinpristavu/quizzer/webapp"
)

func main() {
	// TODO: ADD PROPPER CORS HANDLING!!!!!
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	appPort := flag.String("appPort", "8000", "app port")
	dbHost := flag.String("dbHost", "127.0.0.1", "db host")
	dbUser := flag.String("dbUser", "root", "db user")
	dbPass := flag.String("dbPass", "", "db password")
	flag.Parse()

	db := initDb(*dbHost, *dbUser, *dbPass)

	r := mux.NewRouter()

	webapp.Init(db, r)

	fmt.Printf("App running on: %s\n", *appPort)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", *appPort),
		c.Handler(r),
	)
	if err != nil {
		panic(err)
	}
}

func initDb(host, user, pass string) *gorm.DB {
	var err error
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/quizzer?charset=utf8&parseTime=True", user, pass, host))

	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	return db.Debug()
}
