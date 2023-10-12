package server

import (
	"context"
	"fmt"
	"github.com/flyteorg/flyte/flyteartifacts/pkg/configuration"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/artifact"
	"github.com/flyteorg/flyte/flytestdlib/promutils"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	_ "net/http/pprof" // Required to serve application.
)

type ArtifactService struct {
	artifact.UnimplementedArtifactRegistryServer
	ArtifactProvider ArtifactHandler
	TriggerProvider  TriggerHandler
	TagProvider      TagHandler
	LineageProvider  LineageHandler
	Metrics          ServiceMetrics
}

func NewArtifactService(scope promutils.Scope) *ArtifactService {
	cfg := configuration.ApplicationConfig.GetConfig().(*configuration.ApplicationConfiguration)
	fmt.Println(cfg)

	return &ArtifactService{
		Metrics: InitMetrics(scope),
	}
}

func HttpRegistrationHook(ctx context.Context, gwmux *runtime.ServeMux, grpcAddress string, grpcConnectionOpts []grpc.DialOption, _ promutils.Scope) error {
	err := artifact.RegisterArtifactRegistryHandlerFromEndpoint(ctx, gwmux, grpcAddress, grpcConnectionOpts)
	if err != nil {
		return errors.Wrap(err, "error registering execution service")
	}
	return nil
}

func GrpcRegistrationHook(ctx context.Context, server *grpc.Server, scope promutils.Scope) error {
	serviceImpl := NewArtifactService(scope)
	artifact.RegisterArtifactRegistryServer(server, serviceImpl)

	return nil
}

func (a *ArtifactService) CreateArtifact(ctx context.Context, request *artifact.CreateArtifactRequest) (*artifact.CreateArtifactResponse, error) {
	return &artifact.CreateArtifactResponse{}, nil
}

func (a *ArtifactService) GetArtifact(ctx context.Context, request *artifact.GetArtifactRequest) (*artifact.GetArtifactResponse, error) {
	return &artifact.GetArtifactResponse{}, nil
}

func (a *ArtifactService) CreateTrigger(ctx context.Context, request *artifact.CreateTriggerRequest) (*artifact.CreateTriggerResponse, error) {
	return &artifact.CreateTriggerResponse{}, nil
}

func (a *ArtifactService) DeleteTrigger(ctx context.Context, request *artifact.DeleteTriggerRequest) (*artifact.DeleteTriggerResponse, error) {
	return &artifact.DeleteTriggerResponse{}, nil
}

func (a *ArtifactService) AddTag(ctx context.Context, request *artifact.AddTagRequest) (*artifact.AddTagResponse, error) {
	return &artifact.AddTagResponse{}, nil
}

func (a *ArtifactService) RegisterProducer(ctx context.Context, request *artifact.RegisterProducerRequest) (*artifact.RegisterResponse, error) {
	return &artifact.RegisterResponse{}, nil
}

func (a *ArtifactService) RegisterConsumer(ctx context.Context, request *artifact.RegisterConsumerRequest) (*artifact.RegisterResponse, error) {
	return &artifact.RegisterResponse{}, nil
}

func (a *ArtifactService) SearchArtifacts(ctx context.Context, request *artifact.SearchArtifactsRequest) (*artifact.SearchArtifactsResponse, error) {
	return &artifact.SearchArtifactsResponse{}, nil
}
