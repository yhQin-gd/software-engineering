package like

import (
	"errors"
	"fmt"
	"net/http"
	"text-to-picture/models/image"
	db "text-to-picture/models/init"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReqBody struct {
	URL string `json:"url"`
}

// @Summary 点赞图片
// @Description 点赞图片接口
// @Tags image
// @Accept json
// @Produce json
// @Param requestBody body like.ReqBody true "图片 URL"
// @Success 200 {object} map[string]interface{} "点赞成功"
// @Failure 400 {object} map[string]interface{} "缺少图片 URL"
// @Failure 401 {object} map[string]interface{} "名字解析出错"
// @Failure 409 {object} map[string]interface{} "用户已点赞该图片"
// @Failure 500 {object} map[string]interface{} "数据库操作错误"
// @Router /auth/like [put]
func LikeImage(c *gin.Context) {
	// 解析请求中的图片 URL 和 token

	var reqBody ReqBody
	err := c.BindJSON(&reqBody)
	if err != nil {
		return
	}
	imageURL := reqBody.URL

	if imageURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": "Missing image URL"})
		return
	}

	usernames, exist := c.Get("username")
	if !exist || usernames == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "名字解析出错"})
		return
	}
	username, _ := usernames.(string)
	// 开始事务
	tx := db.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "点赞数据库开始出错"})
		return
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var imageLike image.ImageLike

	// 查询用户是否有点赞记录
	if err := tx.Where("username = ? AND picture = ?", username, imageURL).First(&imageLike).Error; err == nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果找到了记录，则返回冲突状态码
		//if imageLike.UserName != "root" { // 这一行可以移除，除非有特殊原因保留
		c.JSON(http.StatusConflict, gin.H{
			"code":  409,
			"error": "用户已经点赞过该图片",
		})
		return
		//}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 只有当发生非“未找到记录”的错误时才回滚事务并返回 500
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "Database query error",
		})
		return
	}

	// 获取当前点赞数
	var currentLikeCount int
	fmt.Printf("%v  %v", username, imageURL)
	if err := tx.Model(&image.ImageInformation{}).Where("picture = ?", imageURL).Select("likecount").Row().Scan(&currentLikeCount); err != nil {
		fmt.Printf("%v  %v %v", username, imageURL, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": err.Error()}) //"sql: no rows in result set"
		return
	}

	newImageLike := image.ImageLike{
		Picture:  imageURL,
		UserName: username,
		Num:      1,
	}

	if err := tx.Create(&newImageLike).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": err.Error()})
		return
	}
	if err := tx.Model(&image.ImageInformation{}).Where("picture = ?", imageURL).Update("likecount", currentLikeCount+1).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":          200,
		"current_likes": currentLikeCount + 1,
		"message":       "Image liked successfully"})
}
