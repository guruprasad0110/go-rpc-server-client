package geometry_test

import (
	"context"
	"testing"

	"github.com/guruprasad0110/go-rpc-server-client/geometry"
	"github.com/guruprasad0110/go-rpc-server-client/pb"
)

const (
	HOST_IP = "127.0.0.1"
	PORT    = "2019"
)

var ServerAddr = HOST_IP + ":" + PORT

func TestSquare(t *testing.T) {
	s := geometry.NewServer(ServerAddr, nil)

	var dimensions pb.Dimensions = pb.Dimensions{Shape: "Square", Unit: "Cm", Dimension: []int32{12, 12}}
	s.GetArea(context.TODO(), &dimensions)

}
