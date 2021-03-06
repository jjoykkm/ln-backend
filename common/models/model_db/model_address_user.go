package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table address_user
//-------------------------------------------------------------------------------//
//model address_user
type AddressUser struct {
	DBCommonGet
	AddressUserId	 	 string	 	 `json:"address_user_id"`
	HouseNo				 string  	 `json:"house_no"`
	Alley				 string  	 `json:"alley"`
	Road				 string  	 `json:"road"`
	LocationX			 float64	 `json:"location_x"`
	LocationY			 float64	 `json:"location_y"`
	SubDistrictId		 string	 	 `json:"sub_district_id"`
	DistrictId			 string	 	 `json:"district_id"`
	ProvinceId			 string	 	 `json:"province_id"`
	CountryId			 string	 	 `json:"country_id"`
	Uid					 string	 	 `json:"uid"`
	Moo					 string  	 `json:"moo"`
}
// New instance address_user
func (u *AddressUser) New() *AddressUser {
	return &AddressUser{
		DBCommonGet:      	u.DBCommonGet ,
		AddressUserId:      u.AddressUserId ,
		HouseNo:            u.HouseNo ,
		Alley:              u.Alley ,
		Road:               u.Road ,
		LocationX:          u.LocationX ,
		LocationY:          u.LocationY ,
		SubDistrictId:      u.SubDistrictId ,
		DistrictId:         u.DistrictId ,
		ProvinceId:         u.ProvinceId ,
		CountryId:          u.CountryId ,
		Uid:             	u.Uid ,
		Moo:                u.Moo ,
	}
}

// Custom table name for GORM
func (AddressUser) TableName() string {
	return config.DB_ADDRESS_USER
}

