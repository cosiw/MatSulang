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

	res, err := apis.db.Login(req)
	fmt.Println("res : ", *res)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (apis *APIs) InsertUser(c *gin.Context) {
	req := &model.TMEMBER{}

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}
	fmt.Println(req)
	res, err := apis.db.InsertUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}
	if res == 0 {
		c.JSON(http.StatusNoContent, &Response{Res: "No contents"})
	}
	c.JSON(http.StatusOK, res)
}

func (apis *APIs) DeleteUser(c *gin.Context) {
	req := c.Param("id")

	res, err := apis.db.DeleteUser(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad Request"})
		return
	}

	if res == 0 {
		c.JSON(http.StatusNoContent, &Response{Res: "No Contents"})
		return
	}

	c.JSON(http.StatusOK, &Response{Res: fmt.Sprintln(res)})

}
