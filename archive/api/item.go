package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yby/todo_list/models"
	"github.com/yby/todo_list/service"
)

// CreateItem 创建一条代办事项
func CreateItem(c *gin.Context) {
	title := c.Query("title")
	context := c.Query("context")
	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	id, err := service.AddItems(title, context, uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":     "ok",
		"item_id": id,
	})
}

// DeleteItem 删除事项
func DeleteItem(c *gin.Context) {
	itemID := c.Query("item_id")
	action := c.Query("action")
	userID := c.Query("user_id")

	switch action {
	case "0":
		if itemID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "item_id为空",
			})
			return
		}
		if err := service.DeleteOne(itemID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
	case "1":
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "user_id为空",
			},
			)
			return
		}
		if err := service.DeleteFinish(userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
	case "2":
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "user_id为空",
			})
			return
		}
		if err := service.DeleteUnFinish(userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
	case "3":
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "user_id为空",
			})
			return
		}
		if err := service.DeleteAll(userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "action错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

// UpdateItem 更改事项状态
func UpdateItem(c *gin.Context) {
	itemID := c.Query("item_id")
	status := c.Query("status")
	if itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "item_id为空",
		})
		return
	}

	switch status {
	case "0":
		if err := service.UpdateItem2(itemID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
	case "1":
		if err := service.UpdateItem1(itemID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "未知status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func generateItemList(items []models.Item) []map[string]any {
	l := make([]map[string]any, 10)
	for i, item := range items {
		l[i] = map[string]any{
			"item_id":     item.ID,
			"context":     item.Context,
			"item_status": item.Status,
		}
	}
	return l
}

// GetItem 获取事项
func GetItem(c *gin.Context) {
	userID := c.Query("user_id")
	action := c.Query("action")
	keyWord := c.Query("key_word")
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "page不能小于1",
		})
		return
	}

	page -= 1

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "user_id为空",
		})
		return
	}

	var items []models.Item
	switch action {
	case "0":
		items, err = service.SearchFinish(userID, int(page))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
	case "1":
		items, err = service.SearchUnFinish(userID, int(page))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
	case "2":
		items, err = service.SearchAll(userID, int(page))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
	case "3":
		if keyWord == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "key_word为空",
			})
			return
		}

		items, err = service.SearchMainField(userID, keyWord, int(page))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": map[string]any{
			"item":  generateItemList(items)[:len(items)],
			"total": len(items),
		},
		"msg": "ok",
	})
}
