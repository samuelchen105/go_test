package directory

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuhsuan105/go_test/common"
)

func HandlerRegister(rt *gin.RouterGroup) {
	rt.GET("all", ReadAll)
	rt.GET(":id", Read)
	rt.POST("", Create)
	rt.PATCH("", Update)
	rt.DELETE(":id", Delete)
}

func Create(c *gin.Context) {
	db := common.GetDatabase()
	params := Directory{Is_hide: false}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result := db.Create(&params)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func Read(c *gin.Context) {
	db := common.GetDatabase()
	param := c.Param("id")
	var res Directory

	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = db.Where(&Directory{Id: id}).First(&res).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"directory": res})
}

func ReadAll(c *gin.Context) {
	db := common.GetDatabase()
	var res []Directory

	err := db.Find(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"directories": res})
}

func Update(c *gin.Context) {
	db := common.GetDatabase()
	var params Directory
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var res Directory
	err := db.Where(&Directory{Id: params.Id}).First(&res).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = db.Model(&res).Updates(&params).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"directory": res})
}

func Delete(c *gin.Context) {
	db := common.GetDatabase()
	id := c.Param("id")

	err := db.Delete(Directory{}, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, nil)
}
