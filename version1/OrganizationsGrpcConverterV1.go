package version1

import (
	"encoding/json"

	"github.com/pip-services-users2/client-organizations-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]any) map[string]string {
	r := map[string]string{}

	for k, v := range val {
		r[k] = convert.StringConverter.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]any {
	var r map[string]any

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value any) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) any {
	if value == "" {
		return nil
	}

	var m any
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromOrganization(organization *OrganizationV1) *protos.Organization {
	if organization == nil {
		return nil
	}

	obj := &protos.Organization{
		Id:          organization.Id,
		Code:        organization.Code,
		CreateTime:  convert.StringConverter.ToString(organization.CreateTime),
		CreatorId:   organization.CreatorId,
		Deleted:     organization.Deleted,
		Active:      organization.Active,
		Version:     organization.Version,
		Name:        organization.Name,
		Description: organization.Description,
		Address:     organization.Address,
		Center:      toJson(organization.Center),
		Radius:      organization.Radius,
		Geometry:    toJson(organization.Geometry),
		Boundaries:  toJson(organization.Boundaries),
		Language:    organization.Language,
		Timezone:    organization.Timezone,
		Industry:    organization.Industry,
		OrgSize:     organization.OrgSize,
		Purpose:     organization.Purpose,
		Params:      toJson(organization.Params),
	}

	return obj
}

func toOrganization(obj *protos.Organization) *OrganizationV1 {
	if obj == nil {
		return nil
	}

	organization := &OrganizationV1{
		Id:          obj.Id,
		Code:        obj.Code,
		CreateTime:  convert.DateTimeConverter.ToDateTime(obj.CreateTime),
		CreatorId:   obj.CreatorId,
		Deleted:     obj.Deleted,
		Active:      obj.Active,
		Version:     obj.Version,
		Name:        obj.Name,
		Description: obj.Description,
		Address:     obj.Address,
		Center:      fromJson(obj.Center),
		Radius:      obj.Radius,
		Geometry:    fromJson(obj.Geometry),
		Boundaries:  fromJson(obj.Boundaries),
		Language:    obj.Language,
		Timezone:    obj.Timezone,
		Industry:    obj.Industry,
		OrgSize:     obj.OrgSize,
		Purpose:     obj.Purpose,
		Params:      fromJson(obj.Params),
	}

	return organization
}

func fromOrganizationPage(page data.DataPage[*OrganizationV1]) *protos.OrganizationPage {
	obj := &protos.OrganizationPage{
		Total: int64(page.Total),
		Data:  make([]*protos.Organization, len(page.Data)),
	}

	for i, v := range page.Data {
		organization := v
		obj.Data[i] = fromOrganization(organization)
	}

	return obj
}

func toOrganizationPage(obj *protos.OrganizationPage) data.DataPage[*OrganizationV1] {
	if obj == nil {
		return *data.NewEmptyDataPage[*OrganizationV1]()
	}

	organizations := make([]*OrganizationV1, len(obj.Data))

	for i, v := range obj.Data {
		organizations[i] = toOrganization(v)
	}

	page := *data.NewDataPage(organizations, int(obj.Total))

	return page
}
