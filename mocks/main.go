package main

import (
	"log"
	"os"

	"github.com/checkr/openmock"
	"github.com/checkr/openmock/swagger_gen/restapi"
	"github.com/checkr/openmock/swagger_gen/restapi/operations"
	"github.com/dsrhub/dsrhub/idl_dsrhub"
	loads "github.com/go-openapi/loads"
	"github.com/golang/protobuf/proto"
	flags "github.com/jessevdk/go-flags"
)

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewOpenMockAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "OpenMock"
	parser.LongDescription = "OpenMock is a Go service that can mock services in integration tests, staging environment, or anywhere.  The goal is to simplify the process of writing mocks in various channels.  Currently it supports three channels: HTTP Kafka AMQP (e.g. RabbitMQ) The admin API allows you to manipulate the mock behaviour provided by openmock, live.  The base path for the admin API is \"/api/v1\".\n"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	// add our custom openmock functionality
	om := &openmock.OpenMock{}
	om.ParseEnv()
	om.GRPCServiceMap = map[string]openmock.GRPCService{
		"idl.dsrhub.DSRHubService": {
			"CreateDSR": openmock.GRPCRequestResponsePair{
				Request:  proto.MessageV2(&idl_dsrhub.CreateDSRRequest{}),
				Response: proto.MessageV2(&idl_dsrhub.CreateDSRResponse{}),
			},
		},
	}

	server.ConfigureAPI(om)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
