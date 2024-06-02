package utils

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func InitCassandra() {
	cluster := gocql.NewCluster("cassandra", "cassandra2") // replace with your Cassandra host
	cluster.Port = 9042
	cluster.Keyspace = "chat_system" // replace with your keyspace name
	cluster.Consistency = gocql.Quorum

	cluster.ProtoVersion = 4 // Ensure the protocol version is set

	// Setting up reconnection policies
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 5}
	cluster.ConnectTimeout = time.Second * 10
	cluster.SocketKeepalive = time.Second * 30

	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("Error connecting to Cassandra:", err)
	}
	log.Println("Cassandra connection established")
}
