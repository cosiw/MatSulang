package model

type TMEMBER struct {
	MemberID string `json:"MEMBERID"`
	UserID   string `json:"USERID"`
	Password string `json:"PASSWORD"`
	Nickname string `json:"NICKNAME"`
	Report   string `json:"REPORT"`
	Image    string `json:"IMAGE"`
}
type TALCOHOL struct {
	ALCOHOLID string `json:"AlcoholID"`
	ALNAME    string `json:"AlName"`
	CATEGORY  string `json:"Category"`
}

type TBOARD struct {
	BOARDID  string `json:"BoardID"`
	TITLE    string `json:"Title"`
	CONTENTS string `json:"Contents"`
	MEMBERID string `json:"MemberID"`
	PAIRID   string `json:"PairID"`
	LIKE     string `json:"Like"`
	CRTIME   string `json:"Crtime"`
}

type TFOOD struct {
	FOODID   string `json:"FoodID"`
	FOODNAME string `json:"FoodName"`
	CATEGORY string `json:"Category"`
}

type TPAIRING struct {
	PAIRID    string `json:"PairID"`
	PAIRNAME  string `json:"PairName"`
	FOODID    string `json:"FoodID"`
	ALCOHOLID string `json:"AlcoholID"`
}

type TREPLY struct {
	ReplyID   string `json:"REPLYID"`
	BoardID   string `json:"BOARDID"`
	Title     string `json:"TITLE"`
	CONTENTS  string `json:"CONTENTS"`
	MemeberID string `json:"MEMBERID"`
}

type TIMAGE struct {
	IMAGEID string `json:"ImageID"`
	TYPE    string `json:"Type"`
	PATH    string `json:"Path"`
}

type Login struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}
