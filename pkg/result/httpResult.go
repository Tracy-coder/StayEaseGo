package result

import (
	"StayEaseGo/pkg/xerr"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

func HttpResult(c *app.RequestContext, code int, obj interface{}, err error) {
	if err == nil {
		c.JSON(code, obj)
	} else {
		errCode := xerr.SERVER_COMMON_ERROR
		errMsg := err.Error()
		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*xerr.CodeError); ok {
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) {
					errCode = grpcCode
					errMsg = gstatus.Message()
				}
			}
		}

		log.Errorf("【API-ERR】 : %+v ", err)

		c.JSON(http.StatusInternalServerError, utils.H{
			"errCode": errCode,
			"errMsg":  errMsg,
		})
	}
}
