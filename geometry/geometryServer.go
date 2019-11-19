package geometry

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"github.com/guruprasad0110/go-rpc-server-client/pb"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// A Server interface
type Server interface {
	// Start the gRPC server, non-blocking.
	StartServer() error

	// Stop the gRPC server gracefully.
	StopServer() error

	// Methods defined by the server.
	pb.GeometryServer
}

// Internal implementation of pb.GeometryServer interface.
type server struct {
	laddr      string       // host:port listen address
	tlsConfig  *tls.Config  // if secure
	grpcServer *grpc.Server // gRPC server instance of this agent
}

// NewServer makes a new Server that listens on laddr
// If tlsConfig is nil, the sever is insecure.
func NewServer(laddr string, tlsConfig *tls.Config) Server {
	// Set log flags here so other pkgs can't override in their init().
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile | log.LUTC)

	s := &server{
		laddr:     laddr,
		tlsConfig: tlsConfig,
	}

	// Create a gRPC server and register this agent a implementing the
	// RCEAgentServer interface and protocol
	var grpcServer *grpc.Server
	if tlsConfig != nil {
		opt := grpc.Creds(credentials.NewTLS(tlsConfig))
		grpcServer = grpc.NewServer(opt)
	} else {
		grpcServer = grpc.NewServer()
	}
	s.grpcServer = grpcServer

	return s
}

// Start the server for remote connection
func (s *server) StartServer() error {
	// Register the Geometry service with the gRPC server.
	pb.RegisterGeometryServer(s.grpcServer, s)

	lis, err := net.Listen("tcp", s.laddr)
	if err != nil {
		return err
	}
	go s.grpcServer.Serve(lis)
	if s.tlsConfig != nil {
		log.Printf("secure server listening on %s", s.laddr)
	} else {
		log.Printf("insecure server listening on %s", s.laddr)
	}
	return nil
}

func (s *server) StopServer() error {
	s.grpcServer.GracefulStop()
	log.Printf("server stopped on %s", s.laddr)
	return nil
}

// Start implementing service methods.
func (s *server) GetArea(ctx context.Context, req *pb.Dimensions) (*pb.Area, error) {
	fmt.Println("Start calculating the area")
	fmt.Println(req)
	var area int32 = 1
	for _, dim := range req.Dimension {
		area = area * dim
	}
	fmt.Println("ARea == ", area)
	return nil, nil
}
