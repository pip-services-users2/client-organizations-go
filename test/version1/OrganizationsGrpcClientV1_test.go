package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-organizations-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type organizationsGrpcClientV1Test struct {
	client  *version1.OrganizationsGrpcClientV1
	fixture *OrganizationsClientFixtureV1
}

func newOrganizationsGrpcClientV1Test() *organizationsGrpcClientV1Test {
	return &organizationsGrpcClientV1Test{}
}

func (c *organizationsGrpcClientV1Test) setup(t *testing.T) *OrganizationsClientFixtureV1 {
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

	c.client = version1.NewOrganizationsGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewOrganizationsClientFixtureV1(c.client)

	return c.fixture
}

func (c *organizationsGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestGrpcCrudOperations(t *testing.T) {
	c := newOrganizationsGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}
