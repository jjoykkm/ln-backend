package sf_formula_plant

import (
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/errs"
	"net/http"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetPlantCategoryList(c *gin.Context) {
	var reqModel model_other.ReqModel
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
		})
		return
	}
	// GetPlantCategoryList(status, language string) ([]model_services.ForPlantCatList, int) (*model_other.RespModel, error)
	respModel,err := h.service.GetPlantCategoryList(config.GetStatus().Active, &reqModel)
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

func (h *Handler) GetPlantCategoryItem(c *gin.Context) {
	var reqModel model_other.ReqModel
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
		})
		return
	}
	// GetPlantCategoryItem(status, plantTypeId, language string, offset int) (*model_other.RespOffsetModel, error)
	respModel,err := h.service.GetPlantCategoryItem(config.GetStatus().Active, &reqModel)
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

func (h *Handler) GetPlantOverviewByPlant(c *gin.Context) {
	var reqModel model_other.ReqModel
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
		})
		return
	}
	//  GetPlantOverviewByPlant(status, uid, plantId string, offset int) (*model_other.RespOffsetModel, error)
	respModel,err := h.service.GetPlantOverviewByPlant(config.GetStatus().Active, &reqModel)
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

func (h *Handler) GetPlantOverviewFavorite(c *gin.Context) {
	var reqModel model_other.ReqModel
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
		})
		return
	}
	//  GetPlantOverviewFavorite(status, uid, language string, offset int) (*model_other.RespOffsetModel, error)
	respModel,err := h.service.GetPlantOverviewFavorite(config.GetStatus().Active, &reqModel)
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

func (h *Handler) GetMyPlantOverview(c *gin.Context) {
	var reqModel model_other.ReqModel
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
		})
		return
	}
	//  GetMyPlantOverview(status, uid, language string, offset int) (*model_other.RespOffsetModel, error)
	respModel,err := h.service.GetMyPlantOverview(config.GetStatus().Active, &reqModel)
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

func (h *Handler) GetFormulaPlantDetail(c *gin.Context) {
	var reqModel model_other.ReqModel
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
		})
		return
	}
	//  GetFormulaPlantDetail(status, formulaPlantId, language string) (*model_other.RespModel, error)
	respModel,err := h.service.GetFormulaPlantDetail(config.GetStatus().Active, &reqModel)
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