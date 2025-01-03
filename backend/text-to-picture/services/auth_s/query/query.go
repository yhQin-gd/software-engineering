package query

import (
	"errors"
	"log"
	"net/http"
	d "text-to-picture/models/init"         // 数据库模型初始化
	"text-to-picture/models/repository/user_r" // 用户数据访问层
	u "text-to-picture/models/user"          // 用户模型

	"github.com/gin-gonic/gin"               // Gin 框架
	"gorm.io/gorm"                           // GORM ORM
)



// @Summary 获取用户信息
// @Description 根据用户名、邮箱或用户ID获取用户信息
// @Tags user
// @Accept json
// @Produce json
// @Param username query string false "用户名"
// @Param email query string false "邮箱"
// @Param id query int false "用户ID"
// @Success 200 {object} map[string]interface{} "获取用户信息成功"
// @Failure 400 {object} map[string]interface{} "无效的请求数据"
// @Failure 404 {object} map[string]interface{} "用户未找到"
// @Failure 500 {object} map[string]interface{} "查询失败"
// @Router /auth/user/info [get]
func GetUserInfo(c *gin.Context) {

	// 从上下文中获取用户名
	username, exists := c.Get("username")
	if !exists {
		log.Printf("未找到用户名")
		c.JSON(401, gin.H{
			"success": false,
			"message": "未找到用户信息",
		})
		return
	}

	var user *u.UserInformation
	var err error
	// 根据用户名获取用户信息
	user, err = user_r.GetUserByName(d.DB, username.(string))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "用户未找到"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "查询失败", "error": err})
		return
	}

	// 返回查到的用户信息
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// 获取所有用户信息
// @Summary 获取所有用户信息
// @Description 获取系统中所有用户的列表
// @Tags users
// @Produce json
// @Success 200 {object} map[string]interface{} "获取用户列表成功"
// @Failure 500 {object} map[string]interface{} "获取用户列表失败"
// @Router /user/all [get]
func GetAllUsersInfo(c *gin.Context) {
	// 获取所有用户信息
	users, err := user_r.GetAllUsers(d.DB) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取用户列表失败", "error": err.Error()})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "获取用户列表成功",
		"users":   users,
	})
}
