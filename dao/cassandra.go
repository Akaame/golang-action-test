package dao

import (
	"fmt"

	"github.com/gocql/gocql"
)

// Client is a wrapper for gocql cassandra session
type Client struct {
	*gocql.Session
	config *gocql.ClusterConfig
}

// NewClient Creates a gocql cassandra session wrapper
func NewClient(hosts ...string) (*Client, error) {
	clusterConfig := gocql.NewCluster(hosts...)
	clusterConfig.Keyspace = "messaging"
	session, err := clusterConfig.CreateSession()
	if err == nil {
		return nil, fmt.Errorf("Error during cassandra client initialiation: %w", err)
	}

	return &Client{Session: session, config: clusterConfig}, nil
}
