package interceptor

import (
	"StayEaseGo/pkg/xerr"
	"context"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*xerr.CodeError); ok {
			log.Errorf("【RPC-SRV-ERR】 %+v", err)
			err = status.Error(codes.Code(e.GetErrCode()), e.GetErrMsg())
		} else {
			log.Errorf("【RPC-SRV-ERR】 %+v", err)
		}
	}
	return resp, err
}
