package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table trans_socket_area
//-------------------------------------------------------------------------------//
//model trans_socket_area
type TransSocketArea struct {
	DBCommon
	FarmAreaId      uuid.UUID	 `json:"farm_area_id"`
	SocketId     	uuid.UUID	 `json:"socket_id"`
}
// New instance trans_socket_area
func (u *TransSocketArea) New() *TransSocketArea {
	return &TransSocketArea{
		DBCommon:      	u.DBCommon ,
		FarmAreaId:		u.FarmAreaId ,
		SocketId:		u.SocketId ,
	}
}

// Custom table name for GORM
func (TransSocketArea) TableName() string {
	return config.DB_TRANS_SOCKET_AREA
}

//-------------------------------------------------------------------------------//
//							Upsert TransSocketArea
//-------------------------------------------------------------------------------//
type TransSocketAreaUS struct {
	FarmAreaId      string	 `json:"farm_area_id"`
	SocketId     	string	 `json:"socket_id"`
	StatusId		string	 `json:"status_id"`
}
func (TransSocketAreaUS) TableName() string {
	return config.DB_TRANS_SOCKET_AREA
}
