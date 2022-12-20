package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-organizations-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type organizationsCommandableGrpcClientV1Test struct {
	client  *version1.OrganizationsCommandableGrpcClientV1
	fixture *OrganizationsClientFixtureV1
}

func newOrganizationsCommandableGrpcClientV1Test() *organizationsCommandableGrpcClientV1Test {
	return &organizationsCommandableGrpcClientV1Test{}
}

func (c *organizationsCommandableGrpcClientV1Test) setup(t *testing.T) *OrganizationsClientFixtureV1 {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewOrganizationsCommandableGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewOrganizationsClientFixtureV1(c.client)

	return c.fixture
}

func (c *organizationsCommandableGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableGrpcCrudOperations(t *testing.T) {
	c := newOrganizationsCommandableGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}
