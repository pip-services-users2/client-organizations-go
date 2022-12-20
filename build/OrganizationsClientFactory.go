package build

import (
	clients1 "github.com/pip-services-users2/client-organizations-go/version1"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type OrganizationsClientFactory struct {
	*cbuild.Factory
}

func NewOrganizationsClientFactory() *OrganizationsClientFactory {
	c := &OrganizationsClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-organizations", "client", "null", "*", "1.0")
	memoryClientDescriptor := cref.NewDescriptor("service-organizations", "client", "memory", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-organizations", "client", "commandable-http", "*", "1.0")
	cmdGrpcClientDescriptor := cref.NewDescriptor("service-organizations", "client", "commandable-grpc", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-organizations", "client", "grpc", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewOrganizationsNullClientV1)
	c.RegisterType(memoryClientDescriptor, clients1.NewOrganizationsMemoryClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewOrganizationsCommandableHttpClientV1)
	c.RegisterType(cmdGrpcClientDescriptor, clients1.NewOrganizationsCommandableGrpcClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewOrganizationsGrpcClientV1)

	return c
}
