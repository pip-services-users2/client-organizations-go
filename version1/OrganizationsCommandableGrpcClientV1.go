package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type OrganizationsCommandableGrpcClientV1 struct {
	*clients.CommandableGrpcClient
}

func NewOrganizationsCommandableGrpcClientV1() *OrganizationsCommandableGrpcClientV1 {
	return NewOrganizationsCommandableGrpcClientV1WithConfig(nil)
}

func NewOrganizationsCommandableGrpcClientV1WithConfig(config *cconf.ConfigParams) *OrganizationsCommandableGrpcClientV1 {
	c := &OrganizationsCommandableGrpcClientV1{
		CommandableGrpcClient: clients.NewCommandableGrpcClient("v1/organizations"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *OrganizationsCommandableGrpcClientV1) GetOrganizations(ctx context.Context, correlationId string, filter *cdata.FilterParams,
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

func (c *OrganizationsCommandableGrpcClientV1) GetOrganizationById(ctx context.Context, correlationId string, id string) (result *OrganizationV1, err error) {
	res, err := c.CallCommand(ctx, "get_organization_by_id", correlationId, cdata.NewAnyValueMapFromTuples(
		"org_id", id,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*OrganizationV1](res, correlationId)
}

func (c *OrganizationsCommandableGrpcClientV1) GetOrganizationByCode(ctx context.Context, correlationId string, code string) (result *OrganizationV1, err error) {
	res, err := c.CallCommand(ctx, "get_organization_by_code", correlationId, cdata.NewAnyValueMapFromTuples(
		"code", code,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*OrganizationV1](res, correlationId)
}

func (c *OrganizationsCommandableGrpcClientV1) GenerateCode(ctx context.Context, correlationId string, id string) (result string, err error) {
	res, err := c.CallCommand(ctx, "generate_code", correlationId, cdata.NewAnyValueMapFromTuples(
		"org_id", id,
	))

	if err != nil {
		return "", err
	}

	return clients.HandleHttpResponse[string](res, correlationId)
}

func (c *OrganizationsCommandableGrpcClientV1) CreateOrganization(ctx context.Context, correlationId string, organization *OrganizationV1) (result *OrganizationV1, err error) {
	res, err := c.CallCommand(ctx, "create_organization", correlationId, cdata.NewAnyValueMapFromTuples(
		"organization", organization,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*OrganizationV1](res, correlationId)
}

func (c *OrganizationsCommandableGrpcClientV1) UpdateOrganization(ctx context.Context, correlationId string, organization *OrganizationV1) (result *OrganizationV1, err error) {
	res, err := c.CallCommand(ctx, "update_organization", correlationId, cdata.NewAnyValueMapFromTuples(
		"organization", organization,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*OrganizationV1](res, correlationId)
}

func (c *OrganizationsCommandableGrpcClientV1) DeleteOrganizationById(ctx context.Context, correlationId string, id string) (result *OrganizationV1, err error) {
	res, err := c.CallCommand(ctx, "delete_organization_by_id", correlationId, cdata.NewAnyValueMapFromTuples(
		"org_id", id,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*OrganizationV1](res, correlationId)
}
