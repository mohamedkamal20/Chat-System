package utils

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func InitCassandra() {
	cluster := gocql.NewCluster("cassandra") // replace with your Cassandra host
	cluster.Port = 9042
	cluster.Keyspace = "chat_system" // replace with your keyspace name
	cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout = 20 * time.Second

	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("Error connecting to Cassandra:", err)
	}
	log.Println("Cassandra connection established")
}
