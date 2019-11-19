package geometry_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/guruprasad0110/go-rpc-server-client/geometry"
	"github.com/guruprasad0110/go-rpc-server-client/pb"
)

const (
	HOST_IP = "127.0.0.1"
	PORT1   = "2019"
	PORT2   = "2020"
)

var ServerAddr1 = HOST_IP + ":" + PORT1
var ServerAddr2 = HOST_IP + ":" + PORT2

func TestSquare(t *testing.T) {
	s := geometry.NewServer(ServerAddr1, nil)

	var dimensions pb.Dimensions = pb.Dimensions{Shape: "Square", Unit: "Cm", Dimension: []int32{12, 12}}
	s.GetArea(context.TODO(), &dimensions)

}

func TestServerStartAndStop(t *testing.T) {

	// Start Server 1
	fmt.Println("Create new server")
	s1 := geometry.NewServer(ServerAddr1, nil)
	fmt.Println("Done creating new server. Now register and start listening on port: ", PORT1)
	s1.StartServer()

	// Start Server 2
	s2 := geometry.NewServer(ServerAddr1, nil)
	fmt.Println("Done creating new server. Now register and start listening on port: ", PORT2)
	s2.StartServer()

	// Make s1 to sleep for 100 seconds during GetArea execution
	s1.Sleep(100 * time.Second)
	var dimensions pb.Dimensions = pb.Dimensions{Shape: "Square", Unit: "Cm", Dimension: []int32{13, 13}}
	// GetArea is a blocking call. It will sleep for 100 seconds (set above) and then returns result
	s1.GetArea(context.TODO(), &dimensions)

	// Add another method which is non-blocking here...

	// Be a nice citizen.. Stop all servers
	s1.StopServer()
	s2.StopServer()

}
