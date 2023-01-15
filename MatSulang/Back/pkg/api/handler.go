package api

import (
	"Back/pkg/model"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Res string `json:"res"`
}

func (apis *APIs) Login(c *gin.Context) {
	req := &model.Login{}

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}

	res, err := apis.db.Login(*req)
	fmt.Println(res)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}
	c.JSON(http.StatusOK, res)
}
