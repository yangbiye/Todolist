package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Title   string `gorm:"type:varchar(256);not null"`
	Context string `gorm:"type:varchar(256);not null"`
	Status  bool   `gorm:"not null"`
	UserID  uint   `gorm:"not null"`
}

//增
func AddItem(title string, context string, userID uint) (Item, *gorm.DB) {
	item := Item{Title: title, Context: context, Status: false, UserID: userID}
	result := db.Create(&item)
	return item, result
}

//查
//查所有
func SearchAll(userID string, page int) ([]Item, *gorm.DB) {
	var items []Item
	result := db.Limit(10).Offset(page*10).Where("user_id = ?", userID).Find(&items)
	return items, result
}

//查已完成
func SearchFinish(userID string, page int) ([]Item, *gorm.DB) {
	var items []Item
	result := db.Limit(10).Offset(page*10).Where("status=?", true).Find(&items)
	return items, result
}

//查未完成
func SearchUnFinish(userID string, page int) ([]Item, *gorm.DB) {
	var items []Item
	result := db.Limit(10).Offset(page*10).Where("status=?", false).Find(&items)
	return items, result
}

//查关键字(指定title)
func SearchMainField(userId string, title string, page int) ([]Item, *gorm.DB) {
	var items []Item
	result := db.Limit(10).Offset(page*10).Where("title LIKE ?", "%"+title+"%").Find(&items)
	return items, result
}

//删
//删一条（按ID）
func DeleteOne(itemID string) *gorm.DB {
	return db.Delete(&Item{}, itemID)
}

//删除已完成
func DeleteFinish(userID string) *gorm.DB {
	return db.Where("status=true AND user_id=?", userID).Delete(&Item{})
}

//删除未完成
func DeleteUnFinish(userID string) *gorm.DB {
	return db.Where("status=false AND user_id=?", userID).Delete(&Item{})
}

//删除全部
func DeleteAll(userID string) *gorm.DB {
	return db.Where("user_id=?", userID).Delete(&Item{})
}

//改
//改一条已完成
func UpdateItem1(itemID string) *gorm.DB {
	return db.Model(&Item{}).Where("id = ?", itemID).Update("status", true)

}

//改一条未完成
func UpdateItem2(itemID string) *gorm.DB {
	return db.Model(&Item{}).Where("id = ?", itemID).Update("status", false)
}
