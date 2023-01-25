package db

import (
	"Back/pkg/model"
	"encoding/json"
	"fmt"
)

func (h *DBHandler) Login(user *model.Login) (*model.TMEMBER, error) {

	condition := fmt.Sprintf("USERID = '%s' AND PASSWORD = '%s'", user.UserID, user.Password)

	member, err := h.SelectData("TMEMBER", "*", condition)

	if err != nil {
		fmt.Println("user Select Error!")
		return nil, err
	}
	fmt.Println(member)
	member = removeArrChars(member)

	var aMember model.TMEMBER
	err = json.Unmarshal([]byte(member), &aMember)
	if err != nil {
		fmt.Println("Unmarshal Error!")
		return nil, err
	}

	return &aMember, err
}

func (h *DBHandler) InsertUser(user *model.TMEMBER) (int64, error) {
	values := fmt.Sprintf("'%s','%s','%s',%s", user.UserID, user.Password, user.Nickname, user.Report)
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

func (h *DBHandler) UpdateUser(id string, user *model.TMEMBER) (int64, error) {
	value := fmt.Sprintf("USERID = '%s', NICKNAME = '%s', IMAGE = '%s'", user.UserID, user.Nickname, user.Image)
	cond := fmt.Sprintf("MEMBERID = %s", id)
	return h.UpdateData("TMEMBER", value, cond)

}
