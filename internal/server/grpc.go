package server

import (
	"context"
	"fmt"
	pb "sm2/api"
)

type Sm2GrpcServer struct {
	pb.UnimplementedCmbSm2Server
}

func (s Sm2GrpcServer) GetSign(ctx context.Context, req *pb.SignReq) (rep *pb.SignRep, err error) {
	src := req.Src
	fmt.Println(src)
	return nil, nil
}

func (s Sm2GrpcServer) VerifySign(ctx context.Context, req *pb.VerifyReq) (rep *pb.VerifyRep, err error) {
	return nil, nil
}
