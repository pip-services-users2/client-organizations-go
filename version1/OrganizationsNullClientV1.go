package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type OrganizationsNullClientV1 struct {
}

func NewOrganizationsNullClientV1() *OrganizationsNullClientV1 {
	return &OrganizationsNullClientV1{}
}

func (c *OrganizationsNullClientV1) GetOrganizations(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*OrganizationV1], err error) {
	return *data.NewEmptyDataPage[*OrganizationV1](), nil
}

func (c *OrganizationsNullClientV1) GetOrganizationById(ctx context.Context, correlationId string, id string) (result *OrganizationV1, err error) {
	return nil, nil
}

func (c *OrganizationsNullClientV1) GetOrganizationByCode(ctx context.Context, correlationId string, code string) (result *OrganizationV1, err error) {
	return nil, nil
}

func (c *OrganizationsNullClientV1) GenerateCode(ctx context.Context, correlationId string, id string) (result string, err error) {
	return "", nil
}

func (c *OrganizationsNullClientV1) CreateOrganization(ctx context.Context, correlationId string, organization *OrganizationV1) (result *OrganizationV1, err error) {
	return organization, nil
}

func (c *OrganizationsNullClientV1) UpdateOrganization(ctx context.Context, correlationId string, organization *OrganizationV1) (result *OrganizationV1, err error) {
	return organization, nil
}

func (c *OrganizationsNullClientV1) DeleteOrganizationById(ctx context.Context, correlationId string, id string) (result *OrganizationV1, err error) {
	return nil, nil
}
