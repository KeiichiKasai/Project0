package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/dao"
)

func recharge(c *gin.Context) {
	username := c.PostForm("/username")
	amount := c.PostForm("/amount")
	_, err := dao.DB.Exec("update user set money=money+? where username=?", amount, username)
	if err != nil {
		fmt.Printf("update failed,err:%v", err)
	}
	_, err = dao.DB.Exec("insert into moneydata(username,money)values (?,?)", username, amount)
	if err != nil {
		fmt.Printf("insert failed,err:%v", err)
	}
	c.JSON(200, gin.H{
		"message": "充值成功",
	})
}
