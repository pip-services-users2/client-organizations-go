package version1

import (
	"context"
	"strings"
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type OrganizationsMemoryClientV1 struct {
	organizations []*OrganizationV1
}

func NewOrganizationsMemoryClientV1() *OrganizationsMemoryClientV1 {
	return &OrganizationsMemoryClientV1{
		organizations: make([]*OrganizationV1, 0),
	}
}

func (c *OrganizationsMemoryClientV1) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}

	if value == "" || search == "" {
		return false
	}

	return strings.Index(strings.ToLower(value), search) >= 0
}

func (c *OrganizationsMemoryClientV1) matchSearch(item *OrganizationV1, search string) bool {
	search = strings.ToLower(search)
	if c.matchString(item.Id, search) {
		return true
	}
	if c.matchString(item.Name, search) {
		return true
	}

	return false
}

func (c *OrganizationsMemoryClientV1) composeFilter(filter *data.FilterParams) func(*OrganizationV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	search, searchOk := filter.GetAsNullableString("search")
	id, idOk := filter.GetAsNullableString("id")
	code, codeOk := filter.GetAsNullableString("code")
	active, activeOk := filter.GetAsNullableBoolean("active")
	deleted := filter.GetAsBooleanWithDefault("deleted", false)

	return func(item *OrganizationV1) bool {
		if idOk && item.Id != id {
			return false
		}
		if codeOk && item.Code != code {
			return false
		}
		if activeOk && item.Active != active {
			return false
		}
		if !deleted && item.Deleted {
			return false
		}
		if searchOk && !c.matchSearch(item, search) {
			return false
		}

		return true
	}
}

func (c *OrganizationsMemoryClientV1) GetOrganizations(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*OrganizationV1], err error) {
	filterFunc := c.composeFilter(filter)

	organizations := make([]*OrganizationV1, 0)

	for _, organization := range c.organizations {
		if filterFunc(organization) {
			organizations = append(organizations, organization)
		}
	}

	return *data.NewDataPage[*OrganizationV1](organizations, len(organizations)), nil
}

func (c *OrganizationsMemoryClientV1) GetOrganizationById(ctx context.Context, correlationId string, id string) (result *OrganizationV1, err error) {
	var organization *OrganizationV1

	for _, org := range c.organizations {
		if org.Id == id {
			organization = org
			break
		}
	}

	return organization, nil
}

func (c *OrganizationsMemoryClientV1) GetOrganizationByCode(ctx context.Context, correlationId string, code string) (result *OrganizationV1, err error) {
	var organization *OrganizationV1

	for _, org := range c.organizations {
		if org.Code == code {
			organization = org
			break
		}
	}

	return organization, nil
}

func (c *OrganizationsMemoryClientV1) GenerateCode(ctx context.Context, correlationId string, id string) (result string, err error) {
	return id, nil
}

func (c *OrganizationsMemoryClientV1) CreateOrganization(ctx context.Context, correlationId string, organization *OrganizationV1) (result *OrganizationV1, err error) {
	if organization.Id == "" {
		organization.Id = data.IdGenerator.NextLong()
	}

	if organization.Code == "" {
		organization.Code = "org_id:" + data.IdGenerator.NextShort()
	}

	organization.CreateTime = time.Now()

	c.organizations = append(c.organizations, organization)

	return organization, nil
}

func (c *OrganizationsMemoryClientV1) UpdateOrganization(ctx context.Context, correlationId string, organization *OrganizationV1) (result *OrganizationV1, err error) {
	index := -1

	for i, org := range c.organizations {
		if org.Id == organization.Id {
			index = i
			break
		}
	}

	if index > -1 {
		c.organizations[index] = organization
		return organization, nil
	} else {
		return nil, nil
	}

}

func (c *OrganizationsMemoryClientV1) DeleteOrganizationById(ctx context.Context, correlationId string, id string) (result *OrganizationV1, err error) {
	var organization *OrganizationV1

	for _, org := range c.organizations {
		if org.Id == id {
			organization = org
			break
		}
	}

	if organization != nil {
		organization.Deleted = true
	}

	return organization, nil
}
