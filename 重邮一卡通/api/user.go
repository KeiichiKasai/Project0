package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
	"main.go/dao"
	"main.go/model"
	"net/http"
)

func register(c *gin.Context) {
	username := c.PostForm("username") //注册用户名
	password := c.PostForm("password") //注册密码
	rows, err := dao.DB.Query("select username from user where id>?", 0)
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.Username)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return
		}
		if username == u.Username {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "用户已存在",
			})
			return
		}
	}
	_, err = dao.DB.Exec("insert into user(username,password)values(?,?)", username, password)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}
func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	rows, err := dao.DB.Query("select username,password from user where id>?", 0)
	if err != nil {
		fmt.Scanf("Query failed,err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u model.User
		err = rows.Scan(&u.Username, &u.Password)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return
		}
		if username == u.Username && password == u.Password {
			tokenString, _ := middleware.GenToken(u.Username)
			c.JSON(http.StatusOK, gin.H{
				"message": "登录成功",
				"data":    gin.H{"token": tokenString},
			})
			return
		}
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "用户名或密码错误",
	})
}
