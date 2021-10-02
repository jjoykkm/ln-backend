package Helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/models/model_other"
	"net/http"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{service: service}
}

func GetModelFromBody(c *gin.Context) model_other.PostBody {
	var bodyModel model_other.PostBody

	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	json.Unmarshal([]byte(jsonData), &bodyModel)
	//fmt.Printf("%+v/n", bodyModel)
	return bodyModel
}
