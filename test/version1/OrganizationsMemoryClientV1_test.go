package test_version1

import (
	"testing"

	"github.com/pip-services-users2/client-organizations-go/version1"
)

type organizationsMockClientV1Test struct {
	client  *version1.OrganizationsMockClientV1
	fixture *OrganizationsClientFixtureV1
}

func newOrganizationsMockClientV1Test() *organizationsMockClientV1Test {
	return &organizationsMockClientV1Test{}
}

func (c *organizationsMockClientV1Test) setup(t *testing.T) *OrganizationsClientFixtureV1 {
	c.client = version1.NewOrganizationsMockClientV1()

	c.fixture = NewOrganizationsClientFixtureV1(c.client)

	return c.fixture
}

func (c *organizationsMockClientV1Test) teardown(t *testing.T) {
	c.fixture = nil
}

func TestMockCrudOperations(t *testing.T) {
	c := newOrganizationsMockClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}
