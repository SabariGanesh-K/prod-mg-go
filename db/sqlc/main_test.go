package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
 _ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
)

var testQueries *Queries;
var testDB  *sql.DB;
func TestMain(m *testing.M) {
	// config,Configerr:= util.LoadConfig("../../")
	// if Configerr!= nil {
	// 	log.Fatal("error loading config",Configerr)
	// }
	var err error
	testDB,err= sql.Open("postgres","postgresql://root:secret@localhost:5432/prod-mgm?sslmode=disable")
	log.Print("hello")
	if err!=nil {
		log.Fatal("Cannot connect to DB: ",err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}