package server

import (
	"fmt"
	"github.com/apache/dubbo-kubernetes/pkg/config/intercp"
	"github.com/apache/dubbo-kubernetes/pkg/core"
	"github.com/apache/dubbo-kubernetes/pkg/core/runtime/component"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"net/http"
	"time"
)

var log = core.Log.WithName("intercp-server")

const (
	grpcMaxConcurrentStreams = 1000000
	grpcKeepAliveTime        = 15 * time.Second
)

type InterCpServer struct {
	config     intercp.InterCpServerConfig
	grpcServer *grpc.Server
	instanceId string
}

var _ component.Component = &InterCpServer{}

func New(
	config intercp.InterCpServerConfig,
	instanceId string,
) (*InterCpServer, error) {
	grpcOptions := []grpc.ServerOption{
		grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    grpcKeepAliveTime,
			Timeout: grpcKeepAliveTime,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             grpcKeepAliveTime,
			PermitWithoutStream: true,
		}),
	}

	grpcOptions = append(grpcOptions)
	grpcServer := grpc.NewServer(grpcOptions...)

	return &InterCpServer{
		config:     config,
		grpcServer: grpcServer,
		instanceId: instanceId,
	}, nil
}

func (d *InterCpServer) Start(stop <-chan struct{}) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", d.config.Port))
	if err != nil {
		return err
	}
	log := log.WithValues(
		"instanceId",
		d.instanceId,
	)

	errChan := make(chan error)
	go func() {
		defer close(errChan)
		if err := d.grpcServer.Serve(lis); err != nil {
			if err != http.ErrServerClosed {
				log.Error(err, "terminated with an error")
				errChan <- err
				return
			}
		}
		log.Info("terminated normally")
	}()
	log.Info("starting", "interface", "0.0.0.0", "port", d.config.Port, "tls", true)

	select {
	case <-stop:
		log.Info("stopping gracefully")
		d.grpcServer.GracefulStop()
		log.Info("stopped")
		return nil
	case err := <-errChan:
		return err
	}
}

func (d *InterCpServer) NeedLeaderElection() bool {
	return false
}

func (d *InterCpServer) GrpcServer() *grpc.Server {
	return d.grpcServer
}
