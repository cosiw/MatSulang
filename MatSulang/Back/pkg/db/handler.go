package db

import (
	"Back/pkg/model"
	"encoding/json"
	"fmt"
)

func (h *DBHandler) Login(user *model.Login) (*model.TMEMBER, error) {

	condition := fmt.Sprintf("UserID = '%s' AND PASSWORD = '%s'", user.UserID, user.Password)

	member, err := h.SelectData("TMEMBER", "*", condition)
	fmt.Println(member)
	if err != nil {
		fmt.Println("user Select Error!")
	}

	member = removeArrChars(member)
	fmt.Println("remove : ", member)
	var aMember model.TMEMBER
	err = json.Unmarshal([]byte(member), &aMember)
	if err != nil {
		fmt.Println("Unmarshal Error!")
	}
	return &aMember, err
}

func (h *DBHandler) InsertUser(user *model.TMEMBER) (int64, error) {
	values := fmt.Sprintf("'%s','%s','%s',%d", user.UserID, user.Password, user.Nickname, user.Report)
	res, err := h.InsertData("TMEMBER", "USERID, PASSWORD, NICKNAME, REPORT", values)
	if err != nil {
		fmt.Printf("InsertData Error!! : %s", err)
	}
	return res, nil
}

func (h *DBHandler) DeleteUser(id string) (int64, error) {
	cond := fmt.Sprintf("MEMBERID = %s", id)

	return h.DeleteData("TMEMBER", cond)

}
