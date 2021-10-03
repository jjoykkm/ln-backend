package helper

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"net/http"
)

func GetModelFromBody(c *gin.Context) model_other.BodyReq {
	var bodyReq model_other.BodyReq

	if err := c.Bind(&bodyReq); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	//fmt.Printf("%+v/n", bodyReq)
	return bodyReq
}

func ConvertToJson(data interface{}) {
	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println(string(result))
}

func ConcatJoin(typeJoin, leftTable, rightTable, joinKey string) string {
	if typeJoin == "" || leftTable == "" || rightTable == "" || joinKey == "" {
		return ""
	}
	sql := fmt.Sprintf("%s join %s on %s.%s = %s.%s",
		typeJoin, rightTable, leftTable, joinKey, rightTable, joinKey)
	fmt.Println(sql)
	return sql
}