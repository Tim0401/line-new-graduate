package line

import (
	"log"
	"net"

	"github.com/Tim0401/line-new-graduate/pkg/line/controller"
	"github.com/Tim0401/line-new-graduate/pkg/line/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Interface インターフェイス定義
type ApiInterface interface {
	Serve()
}

type Api struct {
}

func NewApi() ApiInterface {
	return &Api{}
}

func (a *Api) Serve() {
	listenPort, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	catService := &controller.CatController{}
	// 実行したい実処理をseverに登録する
	pb.RegisterCatServer(server, catService)
	reflection.Register(server)
	server.Serve(listenPort)
}
