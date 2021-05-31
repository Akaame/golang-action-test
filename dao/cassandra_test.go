package dao

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type CassandraTestSuite struct {
	suite.Suite
	testContext       context.Context
	cassandraInstance testcontainers.Container
	ip                string
	port              nat.Port
}

func TestCassandraTestSuite(t *testing.T) {
	suite.Run(t, new(CassandraTestSuite))
}

func (c *CassandraTestSuite) SetupSuite() {
	c.testContext = context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "cassandra",
		ExposedPorts: []string{"9042/tcp"},
		WaitingFor:   wait.ForListeningPort("9042"),
	}
	cassandraContainer, err := testcontainers.GenericContainer(c.testContext, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal(err)
	}
	ip, err := cassandraContainer.Host(c.testContext)
	if err != nil {
		log.Fatal(err)
	}
	port, err := cassandraContainer.MappedPort(c.testContext, "80")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v, %v\n", ip, port)
	c.cassandraInstance = cassandraContainer
	c.ip = ip
	c.port = port
}

func (c *CassandraTestSuite) TearDownSuite() {
	c.cassandraInstance.Terminate(c.testContext)
}

func TestNewClient(t *testing.T) {
	type args struct {
		hosts []string
	}
	tests := []struct {
		name           string
		args           args
		expectedIPHost string
		expectedPort   int
		wantErr        bool
	}{
		{
			name:           "happy_path",
			args:           args{hosts: []string{"localhost"}},
			expectedIPHost: "localhost",
			expectedPort:   9042,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.hosts...)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedIPHost, got.config.Hosts[0])
			assert.Equal(t, tt.expectedPort, got.config.Port)
		})
	}
}
