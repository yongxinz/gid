package homerpc

import (
	"context"
	"net/http"

	"github.com/hulklab/yago"
	"github.com/hulklab/yago/coms/logger"

	"gid/app/g/preinit"
	pb "gid/app/modules/home/homerpc/homepb"
)

type SegmentRpc struct {
}

func init() {
	SegmentRpc := new(SegmentRpc)
	pb.RegisterGidServer(yago.RpcServer, SegmentRpc)
}

func (r *SegmentRpc) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	return &pb.PingReply{Data: "Hello "}, nil
}

func (r *SegmentRpc) GetId(ctx context.Context, in *pb.IdRequest) (*pb.IdReply, error) {
	var id int64
	var err error

	tag := in.GetTag()
	logger.Ins().Infof("grpc get id, receive tag:  %s", tag)

	out := &pb.IdReply{
		Status: &pb.Status{
			Code: http.StatusOK,
		},
	}

	if id, err = preinit.Srv.GetId(tag); err != nil {
		out.Status.Code = http.StatusInternalServerError
		out.Status.Msg = err.Error()
		return out, nil
	}
	out.Id = id
	return out, nil
}

func (r *SegmentRpc) GetSnowId(ctx context.Context, in *pb.SnowIdRequest) (*pb.SnowIdReply, error) {
	logger.Ins().Infof("grpc get snow id:  %s")
	out := &pb.SnowIdReply{
		Status: &pb.Status{
			Code: http.StatusOK,
		},
		Id: preinit.Srv.SnowFlakeGetId(),
	}

	return out, nil
}
