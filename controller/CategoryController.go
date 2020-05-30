package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wcc4869/ginessential/model"
	"github.com/wcc4869/ginessential/repository"
	"github.com/wcc4869/ginessential/response"
	"strconv"
)

// 创建分类
func CreateCategory(ctx *gin.Context) {
	name := ctx.DefaultPostForm("name", "default")

	category := model.Category{Name: name}
	err := repository.CreateCategory(&category)
	if err != nil {
		response.Error(ctx, nil, "创建失败")
	}
	response.Success(ctx, gin.H{"category": category}, "创建成功")
}

// 更新分类
func UpdateCategory(ctx *gin.Context) {
	//category := model.Category{}
	categoryId, _ := strconv.Atoi(ctx.Request.FormValue("id"))
	category, err := repository.GetCategoryById(categoryId)

	if err != nil {
		response.Error(ctx, gin.H{"category_id": categoryId}, "没有该分类 ")
		return
	}
	name := ctx.Request.FormValue("name")
	if name == "" {
		response.Error(ctx, nil, "name 不能为空")
		return
	}

	category.Name = name
	//ctx.ShouldBind(&category)
	err = repository.UpdateCategory(category)

	if err != nil {
		response.Error(ctx, nil, "更新失败")
	} else {
		response.Success(ctx, gin.H{"category": category}, "更新成功")
	}
}

// 删除分类
func DeleteCategory(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Request.FormValue("id"))
	category, err := repository.GetCategoryById(categoryId)

	if err != nil {
		response.Error(ctx, gin.H{"category_id": categoryId}, "没有该分类 ")
		return
	}
	re := repository.DeleteCategory(category, categoryId)
	if re {
		response.Success(ctx, nil, "删除成功")
	} else {
		response.Error(ctx, nil, "删除失败")
	}
}

// 获取分类
func GetCategory(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Request.FormValue("id"))
	category, err := repository.GetCategoryById(categoryId)

	if err != nil {
		response.Error(ctx, gin.H{"category_id": categoryId}, "没有该分类 ")
		return
	}

	response.Success(ctx, gin.H{"data": category}, "获取成功")
}

// 分类列表
func GetCategories(ctx *gin.Context) {
	var categories []model.Category
	name := ctx.Request.FormValue("name")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))        // 页数
	per_num, _ := strconv.Atoi(ctx.DefaultQuery("per_num", "2")) // 每页个数

	where := make(map[string]interface{})
	if name != "" {
		where["name"] = name
	}
	where["offset"] = (page - 1) * per_num
	where["limit"] = per_num

	err := repository.Show(&categories, where)
	if err != nil {
		response.Error(ctx, nil, "获取失败")
		return
	}
	response.Success(ctx, gin.H{"data": categories}, "获取成功")

}
