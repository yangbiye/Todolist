package service

import (
	"github.com/yby/todo_list/models"
)

// 增加
func AddItems(title string, context string, userID uint) (uint, error) {
	item, res := models.AddItem(title, context, userID)
	if res.Error != nil {
		return 0, res.Error
	}

	return item.ID, nil
}

// 查找
// 查所有
func SearchAll(userID string, page int) ([]models.Item, error) {
	item, res := models.SearchAll(userID, page)
	if res.Error != nil {
		return nil, res.Error
	}
	return item, nil

}

// 查已完成
func SearchFinish(userID string, page int) ([]models.Item, error) {
	item, res := models.SearchFinish(userID, page)
	if res.Error != nil {
		return nil, res.Error
	}
	return item, nil

}

// 查未完成
func SearchUnFinish(userID string, page int) ([]models.Item, error) {
	item, res := models.SearchUnFinish(userID, page)
	if res.Error != nil {
		return nil, res.Error
	}
	return item, nil

}

// 查关键字
func SearchMainField(userID string, title string, page int) ([]models.Item, error) {
	item, res := models.SearchMainField(userID, title, page)
	if res.Error != nil {
		return nil, res.Error
	}
	return item, nil

}

// 删
// 删一条
func DeleteOne(itemID string) error {
	res := models.DeleteOne(itemID)
	if res.Error != nil {
		return res.Error
	}
	return nil

}

// 删全部
func DeleteAll(userID string) error {
	res := models.DeleteAll(userID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 删除已完成
func DeleteFinish(userID string) error {
	res := models.DeleteFinish(userID)
	if res.Error != nil {
		return res.Error
	}
	return nil

}

// 删除未完成
func DeleteUnFinish(userID string) error {
	res := models.DeleteUnFinish(userID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 改
// 改一条已完成
func UpdateItem1(itemID string) error {
	res := models.UpdateItem1(itemID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 改一条未完成
func UpdateItem2(itemID string) error {
	res := models.UpdateItem2(itemID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
