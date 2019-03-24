package main

import (
	"flag"
	"fmt"
	"log"
	//	"time"

	"github.com/kcraybould/behavior/models"
	_ "github.com/lib/pq"
	"github.com/xo/dburl"
)

var flagVerbose = flag.Bool("v", false, "verbose")
var flagURL = flag.String("url", "postgres://Kevin:pass@localhost/behaviour?sslmode=disable", "url")

func main() {
	var err error

	//connect to the db
	db, err := dburl.Open(*flagURL)
	if err != nil {
		log.Fatal(err)
	}

	med, err := models.MedByMedID(db, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(med)

}
