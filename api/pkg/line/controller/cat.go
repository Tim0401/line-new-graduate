package controller

import (
	"context"

	"github.com/Tim0401/line-new-graduate/pkg/line/pb"
	"golang.org/x/xerrors"
)

type CatController struct {
	pb.UnimplementedCatServer
}

func (cc *CatController) GetMyCat(ctx context.Context, message *pb.GetMyCatMessage) (*pb.MyCatResponse, error) {
	switch message.TargetCat {
	case "tama":
		//たまはメインクーン
		return &pb.MyCatResponse{
			Name: "tama",
			Kind: "mainecoon",
		}, nil
	case "mike":
		//ミケはノルウェージャンフォレストキャット
		return &pb.MyCatResponse{
			Name: "mike",
			Kind: "Norwegian Forest Cat",
		}, nil
	}
	return nil, xerrors.New("Not Found YourCat")
}
