package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-listing-api/database"
	"product-listing-api/model"
	"strconv"
)

func CreateProduct(context *gin.Context) {
	var product model.Product

	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.Database.Create(&product).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": product})
}

func GetAllProducts(context *gin.Context) {
	var products []model.Product
	if err := database.Database.Find(&products).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProduct(context *gin.Context) {
	var product model.Product
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.Database.Where("ID=?", id).Find(&product).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": product})
}

func UpdateProduct(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input map[string]interface{}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var product model.Product

	err = database.Database.First(&product, id).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.Database.Model(&product).Updates(input).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": product})
}

func DeleteProduct(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.Database.Delete(&model.Product{}, id).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
