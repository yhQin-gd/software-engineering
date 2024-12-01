package query

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	d "text-to-picture/models/init"
	"text-to-picture/models/repository/image_r"

	//u "text-to-picture/models/user"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
	获取用户生成的图像			GetUserImages
	获取用户的收藏图像			GetUserFavoritedImages

	查询指定的某张图像			GetImage
	查询指定时间段内的所有图像	 GetImagesWithinTimeRange
	获取所有图像信息			GetAllImages
*/

// 获取用户生成的图像
// @Summary 获取用户生成的图像
// @Description 根据用户名或用户ID获取用户生成的图像
// @Tags user
// @Accept json
// @Produce json
// @Param username query string false "用户名"
// @Param id query int false "用户ID"
// @Success 200 {object} map[string]interface{}  "获取用户的图像成功"
// @Failure 400 {object} map[string]interface{} "无效的用户名或用户id"
// @Failure 500 {object} map[string]interface{} "查询用户图片失败"
// @Router /getuserimages [get]
func GetUserImages(c *gin.Context) {
	// username := c.Query("username") // 从请求中获取用户名
	// userIdStr := c.Query("id") // 从请求中获取用户ID（字符串）
	// userId, err := strconv.Atoi(userIdStr) // 将字符串转换为整数
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

	// if username != ""{
	images, err := image_r.GetUserImagesByUsername(d.DB, username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "查询用户图片失败", "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "获取用户的图像成功", "images": images})
	return

	// }else if err == nil{// id转username
	// 	var user u.UserInformation
	// 	err := d.DB.Table("userinformation").Where("id = ?", userId).First(&user).Error // 使用 Find 而不是 First
	// 		if err != nil {
	// 			c.JSON(http.StatusBadRequest, gin.H{"message": "无效的用户id"})
	// 		}

	// 		images, err := image_r.GetUserImagesByUsername(d.DB, user.UserName)
	// 		if err != nil {
	// 			c.JSON(http.StatusInternalServerError, gin.H{"message": "查询用户图片失败","error":err})
	// 			return
	// 		}
	// 		c.JSON(http.StatusOK, gin.H{"message":"获取用户的图像成功","images":images})
	// 		return

	// }else {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "无效的用户名或用户id"})
	// 	return
	// }
}

// 获取用户的收藏图像
// @Summary 获取用户收藏的图片
// @Description 根据用户名或用户ID获取用户收藏的图片
// @Tags images
// @Accept json
// @Produce json
// @Param username query string false "用户名"
// @Param id query string false "用户ID"
// @Success 200 {array} image.ImageInformation "获取用户收藏的图片成功"
// @Failure 400  {object} map[string]interface{} "无效的用户ID或用户名"
// @Failure 500  {object} map[string]interface{} "查询用户收藏的图片失败"
// @Router /getuserfavoritedimages [get]
func GetUserFavoritedImages(c *gin.Context) {
	// username := c.Query("username") // 从请求中获取用户名
	// userIdStr := c.Query("id") // 从请求中获取用户ID（字符串）
	// userId, err := strconv.Atoi(userIdStr) // 将字符串转换为整数

	// if username != ""{
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

	images, err := image_r.GetUserFavoritedImagesByUsername(d.DB, username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "查询用户收藏的图片失败", "err": err})
		return
	}
	c.JSON(http.StatusOK, images)
	//	return

	// }else if err == nil{// id转username
	// 	var user u.UserInformation
	// 	err := d.DB.Table("userinformation").Where("id = ?", userId).First(&user).Error // 使用 Find 而不是 First
	// 		if err != nil {
	// 			c.JSON(http.StatusBadRequest, gin.H{"message": "无效的用户id","error":err})
	// 		}

	// 		images, err := image_r.GetUserFavoritedImagesByUsername(d.DB, user.UserName)
	// 		if err != nil {
	// 			c.JSON(http.StatusInternalServerError, gin.H{"message": "查询用户收藏的图片失败","err":err})
	// 			return
	// 		}
	// 		c.JSON(http.StatusOK, images)
	// 		return

	// }else {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "无效的用户ID或用户名"})
	// 	return
	// }

}

// 查询指定的某张图像
func GetImage(c *gin.Context) {
	url := c.Query("url")
	username := c.Query("username")          // 从请求中获取用户名
	imageIdStr := c.Query("id")              // 从请求中获取图片ID（字符串）
	imageId, err := strconv.Atoi(imageIdStr) // 将字符串转换为整数

	if url != "" {
		image, err := image_r.GetImageByUrl(d.DB, url)
		if err != nil {
			// 检查错误类型
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"message": "未找到相关图片"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "查询用户的图片失败", "error": err})
			}
			return
		}
		c.JSON(http.StatusOK, image)
	} else if username != "" {
		image, err := image_r.GetImageByUsername(d.DB, username)
		if err != nil {
			// 检查错误类型
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"message": "未找到相关图片"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "查询用户的图片失败", "error": err})
			}
			return
		}
		c.JSON(http.StatusOK, image)

	} else if err == nil {
		image, err := image_r.GetImageById(d.DB, imageId)
		if err != nil {
			// 检查错误类型
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"message": "未找到相关图片"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "查询用户的图片失败", "error": err})
			}
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "查询图像成功", "image": image})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "无效的图像ID或用户名"})
	}
}

// 查询指定时间段内的所有图像
func GetImagesWithinTimeRange(c *gin.Context) {
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

	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")

	//定义正则表达式来检测时间字符串是否包含时间部分
	timeRegex := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2})(\.\d+)?Z$`)
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

	var startTime, endTime time.Time
	var err error

	//检查 start_time 是否包含时间部分
	if timeRegex.MatchString(startTimeStr) { //含时间部分
		startTime, err = time.Parse("2006-01-02T15:04:05Z", startTimeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "无效的开始时间格式", "error": err.Error()})
			return
		}
	} else if dateRegex.MatchString(startTimeStr) { //不含时间部分
		startTime, err = time.Parse("2006-01-02", startTimeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "无效的开始时间格式", "error": err.Error()})
			return
		}
		startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, time.UTC)
	} else { //都不符合
		c.JSON(http.StatusBadRequest, gin.H{"message": "无效的开始时间格式"})
		return
	}

	//检查 end_time 是否包含时间部分
	if timeRegex.MatchString(endTimeStr) { //含时间部分
		endTime, err = time.Parse("2006-01-02T15:04:05Z", endTimeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "无效的结束时间格式", "error": err.Error()})
			return
		}
	} else if dateRegex.MatchString(endTimeStr) { //不含时间部分
		endTime, err = time.Parse("2006-01-02", endTimeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "无效的结束时间格式", "error": err.Error()})
			return
		}
		endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 23, 59, 59, 999999999, time.UTC)
	} else { //都不符合
		c.JSON(http.StatusBadRequest, gin.H{"message": "无效的结束时间格式"})
		return
	}

	images, err := image_r.GetImagesInfoWithinTimeRange(d.DB, username.(string), startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "查询图像列表失败", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "查询图像列表成功",
		"images":  images,
	})
}

// 获取所有图像信息
// @Summary 获取所有图像信息
// @Description 获取系统中所有图像的信息
// @Tags images
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "获取图像列表成功"
// @Failure 500 {object} map[string]interface{} "获取图像列表失败"
// @Router /getallimages [get]
func GetAllImages(c *gin.Context) {
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

	images, err := image_r.GetAllImagesInfo(d.DB, username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取图像列表失败", "error": err.Error()})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "获取图像列表成功",
		"images":  images,
	})
}
