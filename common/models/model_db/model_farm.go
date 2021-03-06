package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table farm
//-------------------------------------------------------------------------------//
//model farm
type Farm struct {
	DBCommonGet
	FarmId      	string	 	 `json:"farm_id"`
	FarmName    	string	 `json:"farm_name"`
	FarmDesc    	string	 `json:"farm_desc"`
}
// New instance farm
func (u *Farm) New() *Farm {
	return &Farm{
		DBCommonGet:      	u.DBCommonGet ,
		FarmId:			u.FarmId ,
		FarmName:		u.FarmName ,
		FarmDesc:		u.FarmDesc ,
	}
}

// Custom table name for GORM
func (Farm) TableName() string {
	return config.DB_FARM
}

//-------------------------------------------------------------------------------//
//							Upsert Farm
//-------------------------------------------------------------------------------//
type FarmUS struct {
	FarmId      	string	 `json:"farm_id" gorm:"default:uuid_generate_v4()"`
	FarmName    	string	 `json:"farm_name"`
	FarmDesc    	string	 `json:"farm_desc"`
	StatusId		string	 `json:"status_id"`
}
func (FarmUS) TableName() string {
	return config.DB_FARM
}

