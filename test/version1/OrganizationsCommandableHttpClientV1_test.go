package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-organizations-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type organizationsCommandableHttpClientV1Test struct {
	client  *version1.OrganizationsCommandableHttpClientV1
	fixture *OrganizationsClientFixtureV1
}

func newOrganizationsCommandableHttpClientV1Test() *organizationsCommandableHttpClientV1Test {
	return &organizationsCommandableHttpClientV1Test{}
}

func (c *organizationsCommandableHttpClientV1Test) setup(t *testing.T) *OrganizationsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewOrganizationsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewOrganizationsClientFixtureV1(c.client)

	return c.fixture
}

func (c *organizationsCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newOrganizationsCommandableHttpClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}
