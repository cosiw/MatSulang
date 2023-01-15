package db

import (
	"Back/pkg/model"
	"encoding/json"
	"fmt"
)

func (h *DBHandler) Login(user model.Login) (model.TMEMBER, error) {

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
		fmt.Println("marshal Error!")
	}
	return aMember, err
}
