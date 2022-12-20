package test_version1

import (
	"context"
	"testing"
	"time"

	"github.com/pip-services-users2/client-organizations-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
)

type OrganizationsClientFixtureV1 struct {
	Client        version1.IOrganizationsClientV1
	ORGANIZATION1 *version1.OrganizationV1
	ORGANIZATION2 *version1.OrganizationV1
}

func NewOrganizationsClientFixtureV1(client version1.IOrganizationsClientV1) *OrganizationsClientFixtureV1 {
	return &OrganizationsClientFixtureV1{
		Client: client,
		ORGANIZATION1: &version1.OrganizationV1{
			Id:          data.IdGenerator.NextLong(),
			Code:        "",
			Name:        "Organization #1",
			Description: "Test organization #1",
			CreateTime:  time.Now(),
			CreatorId:   "123",
			Active:      true,
		},
		ORGANIZATION2: &version1.OrganizationV1{
			Id:          data.IdGenerator.NextLong(),
			Code:        "",
			Name:        "Organization #2",
			Description: "Test organization #2",
			CreateTime:  time.Now(),
			CreatorId:   "123",
			Active:      true,
		},
	}
}

func (c *OrganizationsClientFixtureV1) clear() {
	page, _ := c.Client.GetOrganizations(context.Background(), "", nil, nil)

	for _, v := range page.Data {
		organization := v
		c.Client.DeleteOrganizationById(context.Background(), "", organization.Id)
	}
}

func (c *OrganizationsClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create one organization
	organization, err := c.Client.CreateOrganization(context.Background(), "", c.ORGANIZATION1)
	assert.Nil(t, err)

	assert.NotNil(t, organization)
	assert.Equal(t, organization.Name, c.ORGANIZATION1.Name)
	assert.Equal(t, organization.Description, c.ORGANIZATION1.Description)
	assert.NotEmpty(t, organization.Code)

	organization1 := organization

	// Create another organization
	organization, err = c.Client.CreateOrganization(context.Background(), "", c.ORGANIZATION2)
	assert.Nil(t, err)

	assert.NotNil(t, organization)
	assert.Equal(t, organization.Name, c.ORGANIZATION2.Name)
	assert.Equal(t, organization.Description, c.ORGANIZATION2.Description)
	assert.NotEmpty(t, organization.Code)

	//organization2 := organization

	// Get all organizations
	page, err1 := c.Client.GetOrganizations(context.Background(), "", nil, nil)
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) >= 2)

	// Get organization by code
	organization, err = c.Client.GetOrganizationByCode(context.Background(), "", organization1.Code)
	assert.Nil(t, err)

	assert.NotNil(t, organization)
	assert.Equal(t, organization.Id, organization1.Id)

	// Generate code
	code, err2 := c.Client.GenerateCode(context.Background(), "", organization1.Id)
	assert.Nil(t, err2)

	assert.NotEmpty(t, code)

	// Update the organization
	organization1.Description = "Updated Content 1"
	organization, err = c.Client.UpdateOrganization(context.Background(), "", organization1)
	assert.Nil(t, err)

	assert.NotNil(t, organization)
	assert.Equal(t, organization.Description, "Updated Content 1")
	assert.Equal(t, organization.Name, organization1.Name)

	organization1 = organization

	// Delete organization
	organization, err = c.Client.DeleteOrganizationById(context.Background(), "", organization1.Id)
	assert.Nil(t, err)

	// Try to get deleted organization
	organization, err = c.Client.GetOrganizationById(context.Background(), "", organization1.Id)
	assert.Nil(t, err)

	assert.NotNil(t, organization)
	assert.True(t, organization.Deleted)
	//assert.Nil(t, organization)
}
