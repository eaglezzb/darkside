package model

type UserInfoModel struct {
	Uid  		string 	`json:"uid" form:"uid"`
	Name	 	string 	`json:"name" form:"name"`
	Password	string  `json:"password" form:"password"`
}