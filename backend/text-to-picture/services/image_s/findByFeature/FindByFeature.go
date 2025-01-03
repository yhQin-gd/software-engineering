package findByFeature

import (
	"log"
	"net/http"

	d "text-to-picture/models/init"
	"text-to-picture/models/repository/image_r"

	"github.com/gin-gonic/gin"
)

// @Summary 根据特征查找图片
// @Description 根据提供的特征列表查找图片
// @Tags 图片管理
// @Accept json
// @Produce json
// @Param feature query []string true "特征列表"
// @Param isOwn query string false "是否只查找自己的图片"
// @Success 200 {object} map[string]interface{} "成功查找图片"
// @Failure 400 {object} map[string]interface{} "请求错误"
// @Failure 401 {object} map[string]interface{} "未授权"
// @Failure 500 {object} map[string]interface{} "内部服务器错误"
// @Router /auth/image/feature [get]
func FindByFeature(c *gin.Context) {
	// 从查询参数中获取特征列表
	features := c.QueryArray("feature") // 获取特征列表数组
	isOwn := c.Query("isOwn")           // 获取是否只查找自己的图片的标志

	var username string
	username = ""
	if isOwn == "true" || isOwn == "True" || isOwn == "TRUE" {
		// 如果 isOwn 标志为 true，尝试从上下文中获取用户名
		userName, exists := c.Get("username")
		if !exists {
			// 未找到用户名，返回未授权错误
			log.Printf("未找到用户名")
			c.JSON(401, gin.H{
				"success": false,
				"message": "未找到用户信息",
			})
			return
		}
		username = userName.(string) // 提取用户名
	}

	// 调用业务逻辑层函数查找图片
	images, err := image_r.FindByFeature(d.DB, username, features)
	if err != nil {
		// 查询失败，返回内部服务器错误
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "根据关键字查询图片失败",
			"error":   err.Error(),
		})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"images": images, // 查询到的图片列表
	})
}
