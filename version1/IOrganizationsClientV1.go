package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IOrganizationsClientV1 interface {
	GetOrganizations(ctx context.Context, correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result data.DataPage[*OrganizationV1], err error)

	GetOrganizationById(ctx context.Context, correlationId string, id string) (result *OrganizationV1, err error)

	GetOrganizationByCode(ctx context.Context, correlationId string, code string) (result *OrganizationV1, err error)

	GenerateCode(ctx context.Context, correlationId string, orgId string) (result string, err error)

	CreateOrganization(ctx context.Context, correlationId string, organization *OrganizationV1) (result *OrganizationV1, err error)

	UpdateOrganization(ctx context.Context, correlationId string, organization *OrganizationV1) (result *OrganizationV1, err error)

	DeleteOrganizationById(ctx context.Context, correlationId string, id string) (result *OrganizationV1, err error)
}
