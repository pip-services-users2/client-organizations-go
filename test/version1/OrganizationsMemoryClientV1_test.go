package test_version1

import (
	"testing"

	"github.com/pip-services-users2/client-organizations-go/version1"
)

type organizationsMemoryClientV1Test struct {
	client  *version1.OrganizationsMemoryClientV1
	fixture *OrganizationsClientFixtureV1
}

func newOrganizationsMemoryClientV1Test() *organizationsMemoryClientV1Test {
	return &organizationsMemoryClientV1Test{}
}

func (c *organizationsMemoryClientV1Test) setup(t *testing.T) *OrganizationsClientFixtureV1 {
	c.client = version1.NewOrganizationsMemoryClientV1()

	c.fixture = NewOrganizationsClientFixtureV1(c.client)

	return c.fixture
}

func (c *organizationsMemoryClientV1Test) teardown(t *testing.T) {
	c.fixture = nil
}

func TestMemoryCrudOperations(t *testing.T) {
	c := newOrganizationsMemoryClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}
