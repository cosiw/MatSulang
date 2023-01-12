package model

type TMEMBER struct {
	MEMBERID string `json:"MemberID"`
	USERID   string `json:"UserID"`
	PASSWORD string `json:"Password"`
	NICKNAME string `json:"Nickname"`
	REPORT   string `json:"report"`
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
	REPLYID  string `json:"ReplyID"`
	BOARDID  string `json:"BoardID"`
	TITLE    string `json:"Title"`
	CONTENTS string `json:"Contents"`
	MEMBERID string `json:"MemeberID"`
}

type TIMAGE struct {
	IMAGEID string `json:"ImageID"`
	TYPE    string `json:"Type"`
	PATH    string `json:"Path"`
}
