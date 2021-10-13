package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"image"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table users
//-------------------------------------------------------------------------------//
//model users
type Users struct {
	Uid      		uuid.UUID	 	`json:"uid"`
	Username     	string	 	 	`json:"username"`
	Password      	string	 	 	`json:"password"`
	FullName      	string	 	 	`json:"full_name"`
	SurName      	string	 	 	`json:"sur_name"`
	NickName      	string	 	 	`json:"nick_name"`
	Gender			uuid.UUID	 	`json:"gender"`
	BirthDate		time.Time	 	`json:"birth_date"`
	MobilePhone     string	 	 	`json:"mobile_phone"`
	Telephone      	string	 	 	`json:"telephone"`
	Mail      		string	 	 	`json:"mail"`
	Image	      	image.Image	 	`json:"image"`
	CreateDate		time.Time	 	`json:"create_date"`
	ChangeDate	    time.Time	 	`json:"change_date"`
	StatusId		uuid.UUID	 	`json:"status_id"`
	UserNo			string		 	`json:"user_no"`
}
// New instance users
func (u *Users) New() *Users {
	return &Users{
		Uid:       		u.Uid ,
		Username:       u.Username ,
		Password:       u.Password ,
		FullName:       u.FullName ,
		SurName:       	u.SurName ,
		NickName:       u.NickName ,
		Gender:       	u.Gender ,
		BirthDate:      u.BirthDate ,
		MobilePhone:    u.MobilePhone ,
		Telephone:      u.Telephone ,
		Mail:       	u.Mail ,
		Image:       	u.Image ,
		CreateDate:     u.CreateDate ,
		ChangeDate:     u.ChangeDate ,
		StatusId:       u.StatusId ,
		UserNo:       	u.UserNo ,

	}
}

// Custom table name for GORM
func (Users) TableName() string {
	return config.DB_USERS
}


//-------------------------------------------------------------------------------//
//							Table users only data
//-------------------------------------------------------------------------------//
//model users
type UsersShort struct {
	Uid      		uuid.UUID	 	`json:"uid"`
	Username     	string	 	 	`json:"username"`
	NickName      	string	 	 	`json:"nick_name"`
	Image	      	string	 		`json:"image"`
	CreateDate		time.Time	 	`json:"create_date"`
	ChangeDate	    time.Time	 	`json:"change_date"`
	StatusId		uuid.UUID	 	`json:"status_id"`
	UserNo			string		 	`json:"user_no"`
}
// New instance users
func (u *UsersShort) New() *UsersShort {
	return &UsersShort{
		Uid:       		u.Uid ,
		Username:       u.Username ,
		NickName:       u.NickName ,
		Image:       	u.Image ,
		CreateDate:     u.CreateDate ,
		ChangeDate:     u.ChangeDate ,
		StatusId:       u.StatusId ,
		UserNo:       	u.UserNo ,
	}
}

// Custom table name for GORM
func (UsersShort) TableName() string {
	return config.DB_USERS
}

