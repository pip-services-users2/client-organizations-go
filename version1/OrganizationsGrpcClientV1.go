package version1

import (
	"context"

	"github.com/pip-services-users2/client-organizations-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type OrganizationsGrpcClientV1 struct {
	*clients.GrpcClient
}

func NewOrganizationsGrpcClientV1() *OrganizationsGrpcClientV1 {
	return &OrganizationsGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("organizations_v1.Organizations"),
	}
}

func (c *OrganizationsGrpcClientV1) GetOrganizations(ctx context.Context, correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result data.DataPage[*OrganizationV1], err error) {
	timing := c.Instrument(ctx, correlationId, "organizations_v1.get_organizations")
	defer timing.EndTiming(ctx, err)

	req := &protos.OrganizationPageRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}
	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.OrganizationPageReply)
	err = c.CallWithContext(ctx, "get_organizations", correlationId, req, reply)
	if err != nil {
		return *data.NewEmptyDataPage[*OrganizationV1](), err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return *data.NewEmptyDataPage[*OrganizationV1](), err
	}

	result = toOrganizationPage(reply.Page)

	return result, nil
}

func (c *OrganizationsGrpcClientV1) GetOrganizationById(ctx context.Context, correlationId string, id string) (result *OrganizationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "organizations_v1.get_organization_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.OrgIdRequest{
		CorrelationId: correlationId,
		OrgId:         id,
	}

	reply := new(protos.OrganizationObjectReply)
	err = c.CallWithContext(ctx, "get_organization_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toOrganization(reply.Organization)

	return result, nil
}

func (c *OrganizationsGrpcClientV1) GetOrganizationByCode(ctx context.Context, correlationId string, code string) (result *OrganizationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "organizations_v1.get_organization_by_code")
	defer timing.EndTiming(ctx, err)

	req := &protos.OrganizationCodeRequest{
		CorrelationId: correlationId,
		Code:          code,
	}

	reply := new(protos.OrganizationObjectReply)
	err = c.CallWithContext(ctx, "get_organization_by_code", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toOrganization(reply.Organization)

	return result, nil
}

func (c *OrganizationsGrpcClientV1) GenerateCode(ctx context.Context, correlationId string, id string) (result string, err error) {
	timing := c.Instrument(ctx, correlationId, "organizations_v1.generate_code")
	defer timing.EndTiming(ctx, err)

	req := &protos.OrgIdRequest{
		CorrelationId: correlationId,
		OrgId:         id,
	}

	reply := new(protos.OrganizationCodeReply)
	err = c.CallWithContext(ctx, "generate_code", correlationId, req, reply)
	if err != nil {
		return "", err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return "", err
	}

	result = reply.Code

	return result, nil
}

func (c *OrganizationsGrpcClientV1) CreateOrganization(ctx context.Context, correlationId string, organization *OrganizationV1) (result *OrganizationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "organizations_v1.create_organization")
	defer timing.EndTiming(ctx, err)

	req := &protos.OrganizationObjectRequest{
		CorrelationId: correlationId,
		Organization:  fromOrganization(organization),
	}

	reply := new(protos.OrganizationObjectReply)
	err = c.CallWithContext(ctx, "create_organization", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toOrganization(reply.Organization)

	return result, nil
}

func (c *OrganizationsGrpcClientV1) UpdateOrganization(ctx context.Context, correlationId string, organization *OrganizationV1) (result *OrganizationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "organizations_v1.update_organization")
	defer timing.EndTiming(ctx, err)

	req := &protos.OrganizationObjectRequest{
		CorrelationId: correlationId,
		Organization:  fromOrganization(organization),
	}

	reply := new(protos.OrganizationObjectReply)
	err = c.CallWithContext(ctx, "update_organization", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toOrganization(reply.Organization)

	return result, nil
}

func (c *OrganizationsGrpcClientV1) DeleteOrganizationById(ctx context.Context, correlationId string, id string) (result *OrganizationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "organizations_v1.delete_organization_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.OrgIdRequest{
		CorrelationId: correlationId,
		OrgId:         id,
	}

	reply := new(protos.OrganizationObjectReply)
	err = c.CallWithContext(ctx, "delete_organization_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toOrganization(reply.Organization)

	return result, nil
}
