package cm_plant

import (
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"github.com/jjoykkm/ln-backend/errs"
	"net/http"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetFertAndCatList(c *gin.Context) {
	var reqModel model_other.ReqModel
	reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
		})
		return
	}
	respModel,err := h.service.GetFertAndCatList(config.GetStatus().Active)
	if err != nil {
		if errx, ok := err.(*errs.ErrContext); ok {
			if httpCode, ok := mapErrorCode[errx.Code]; ok {
				c.JSON(httpCode, err)
				return
			}
		}
		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
			Code: "80000",
			Err:  err,
		})
		return
	}
	c.JSON(http.StatusOK, respModel)
}