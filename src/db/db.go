package db

import (
	"log"
	"os"
	"strings"
	"github.com/gocql/gocql"
)

var session *gocql.Session

func InitScyllaDB() {
	keyspace := os.Getenv("SCYLLADB_KEYSPACE")
	publicIPs := os.Getenv("SCYLLADB_PUBLIC_IPS")

	publicIPList := strings.Split(publicIPs, ",")

	cluster := gocql.NewCluster(publicIPList...)
	cluster.Keyspace = keyspace

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	return session
}
