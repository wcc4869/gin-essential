package repository

import (
	"github.com/wcc4869/ginessential/common"
	"github.com/wcc4869/ginessential/model"
)

// 创建分类
func CreateCategory(category *model.Category) (err error) {
	err = common.DB.Create(category).Error
	if err != nil {
		return err
	}
	return nil
}

// 跟新
func UpdateCategory(category *model.Category) (err error) {
	err = common.DB.Save(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCategoryById(id int) (category *model.Category, err error) {
	category = &model.Category{}
	err = common.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func DeleteCategory(category *model.Category, id int) bool {
	err := common.DB.Where("id=?", id).Delete(category).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

// 分类列表
func Show(categories *[]model.Category, option map[string]interface{}) (err error) {
	//err = common.DB.Find(categories).Error
	con := common.DB.Debug().Offset(option["offset"]).Limit(option["limit"])
	name, ok := option["name"].(string)
	ln := "%" + name + "%" // interface to string
	if ok {
		con = con.Where("name LIKE ?", ln)
	}
	order := option["order"].(string)
	orderType := option["order_type"].(string)
	con = con.Order(order + " " + orderType)
	//fmt.Println(name)
	con.Find(&categories)
	//err = con.Find(&categories).Error
	//if err != nil {
	//	return err
	//}
	return nil
}
