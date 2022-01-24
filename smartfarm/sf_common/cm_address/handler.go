package cm_address

import (
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/errs"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_common/cm_auth"
	"net/http"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetProvinceList(c *gin.Context) {
	reqModel := (&cm_auth.Service{}).PrepareDataGet(c, c.Request.Header.Get("Bearer"))
	if reqModel == nil {
		return
	}
	respModel,err := h.service.GetProvinceList(config.GetStatus().Active)
	if err != nil {
		if errx, ok := err.(*errs.ErrContext); ok {
			if httpCode, ok := mapErrorCode[errx.Code]; ok {
				c.JSON(httpCode, err)
				return
			}
		}
		(&errs.Service{}).ErrMsgInternal(c, err)
		return
	}
	c.JSON(http.StatusOK, respModel)
}