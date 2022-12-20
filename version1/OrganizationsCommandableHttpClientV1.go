package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type OrganizationsCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewOrganizationsCommandableHttpClientV1() *OrganizationsCommandableHttpClientV1 {
	return NewOrganizationsCommandableHttpClientV1WithConfig(nil)
}

func NewOrganizationsCommandableHttpClientV1WithConfig(config *cconf.ConfigParams) *OrganizationsCommandableHttpClientV1 {
	c := &OrganizationsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/organizations"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *OrganizationsCommandableHttpClientV1) GetOrganizations(ctx context.Context, correlationId string, filter *cdata.FilterParams,
	paging *cdata.PagingParams) (result cdata.DataPage[*OrganizationV1], err error) {
	res, err := c.CallCommand(ctx, "get_organizations", correlationId, cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	))

	if err != nil {
		return *cdata.NewEmptyDataPage[*OrganizationV1](), err
	}

	return clients.HandleHttpResponse[cdata.DataPage[*OrganizationV1]](res, correlationId)
}

func (c *OrganizationsCommandableHttpClientV1) GetOrganizationById(ctx context.Context, correlationId string, id string) (result *OrganizationV1, err error) {
	res, err := c.CallCommand(ctx, "get_organization_by_id", correlationId, cdata.NewAnyValueMapFromTuples(
		"org_id", id,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*OrganizationV1](res, correlationId)
}

func (c *OrganizationsCommandableHttpClientV1) GetOrganizationByCode(ctx context.Context, correlationId string, code string) (result *OrganizationV1, err error) {
	res, err := c.CallCommand(ctx, "get_organization_by_code", correlationId, cdata.NewAnyValueMapFromTuples(
		"code", code,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*OrganizationV1](res, correlationId)
}

func (c *OrganizationsCommandableHttpClientV1) GenerateCode(ctx context.Context, correlationId string, id string) (result string, err error) {
	res, err := c.CallCommand(ctx, "generate_code", correlationId, cdata.NewAnyValueMapFromTuples(
		"org_id", id,
	))

	if err != nil {
		return "", err
	}

	return clients.HandleHttpResponse[string](res, correlationId)
}

func (c *OrganizationsCommandableHttpClientV1) CreateOrganization(ctx context.Context, correlationId string, organization *OrganizationV1) (result *OrganizationV1, err error) {
	res, err := c.CallCommand(ctx, "create_organization", correlationId, cdata.NewAnyValueMapFromTuples(
		"organization", organization,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*OrganizationV1](res, correlationId)
}

func (c *OrganizationsCommandableHttpClientV1) UpdateOrganization(ctx context.Context, correlationId string, organization *OrganizationV1) (result *OrganizationV1, err error) {
	res, err := c.CallCommand(ctx, "update_organization", correlationId, cdata.NewAnyValueMapFromTuples(
		"organization", organization,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*OrganizationV1](res, correlationId)
}

func (c *OrganizationsCommandableHttpClientV1) DeleteOrganizationById(ctx context.Context, correlationId string, id string) (result *OrganizationV1, err error) {
	res, err := c.CallCommand(ctx, "delete_organization_by_id", correlationId, cdata.NewAnyValueMapFromTuples(
		"org_id", id,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*OrganizationV1](res, correlationId)
}
