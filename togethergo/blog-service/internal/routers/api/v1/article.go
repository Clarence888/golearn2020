package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (A Article) Get(c *gin.Context) {
	fmt.Println(c.Param("id"))   //获取传值的参数
	fmt.Println(c.Query("name")) //获取拼接的query参数
	//app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

func (A Article) List(c *gin.Context) {

}

func (A Article) Create(c *gin.Context) {

}

func (A Article) Update(c *gin.Context) {

}

func (A Article) Delete(c *gin.Context) {

}
