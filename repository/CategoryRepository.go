package repository

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
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

func Ecxel() {

	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet1")
	categories := map[string]string{"A1": "id", "B1": "name", "C1": "age"}
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}

	values := map[string]string{"A2": "1", "B2": "WCC", "C2": "12", "A3": "2", "B3": "LKK", "C3": "24"}

	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}

	// 设置单元格的值
	//f.SetCellValue("Sheet1", "A2", "Hello world.")
	//f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("t.xlsx"); err != nil {
		println(err.Error())
	}

}
