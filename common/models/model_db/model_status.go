package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table status
//-------------------------------------------------------------------------------//
//model status
type Status struct {
	DBCommonGet
	StatusId        string	 `json:"status_id"`
	StatusName      string	 `json:"status_name"`
}
// New instance status
func (u *Status) New() *Status {
	return &Status{
		DBCommonGet:      	u.DBCommonGet ,
		StatusId:			u.StatusId ,
		StatusName:			u.StatusName ,
	}
}

// Custom table name for GORM
func (Status) TableName() string {
	return config.DB_STATUS
}
