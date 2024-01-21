package db

import (
	"os"
	"strings"
	"github.com/gocql/gocql"
)

var Session *gocql.Session

func InitScyllaDB() (*gocql.Session, error) {
	keyspace := os.Getenv("SCYLLADB_KEYSPACE")
	publicIPs := os.Getenv("SCYLLADB_PUBLIC_IPS")

	publicIPList := strings.Split(publicIPs, ",")

	cluster := gocql.NewCluster(publicIPList...)
	cluster.Keyspace = keyspace

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	Session = session
	
	return session, nil
}
