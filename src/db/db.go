package db

import (
	"os"
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"github.com/gocql/gocql"
)

var Session *gocql.Session
func init() {
	var err error
	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
	keyspace := os.Getenv("SCYLLADB_KEYSPACE")
	publicIP := os.Getenv("SCYLLADB_IP")

	cluster := gocql.NewCluster(publicIP)
	cluster.Keyspace = keyspace

	Session, err = cluster.CreateSession()

	if err != nil {
		panic(err)
	}
	
	fmt.Println("Scylla init done!")
}
