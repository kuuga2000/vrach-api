package models

type Object struct {
	ObjectId     int    `json:"id"`
	ObjectName   string `json:"name"`
	ObjectGender int    `json:"gender"`
}

type Objects struct {
	Objects []Object `json:"objects"`
}
