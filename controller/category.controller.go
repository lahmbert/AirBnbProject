package controller

import (
	db "AirBnbProject/db/sqlc"
	"AirBnbProject/middleware"
	"AirBnbProject/models"
	"AirBnbProject/services"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	storedb services.Store
}

func NewCategoryController(store services.Store) *CategoryController {
	return &CategoryController{
		storedb: store,
	}
}

func (cate *CategoryController) GetCategoryById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := models.Nullable(cate.storedb.FindCategoryById(c, int32(id)))

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	if category == nil {
		c.JSON(http.StatusNotFound, models.NewError(models.ErrCategoryNotFound))
		return
	}

	c.JSON(http.StatusOK, category)
}

func (cate *CategoryController) GetListCategory(c *gin.Context) {
	categories, err := cate.storedb.FindAllCategory(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (cate *CategoryController) PostCategory(c *gin.Context) {
	var payload *models.CategoryPostReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	token, _ := middleware.GenerateJWT("cool")
	log.Println(token)

	args := db.CreateCategoryParams{
		CategoryName: payload.CategoryName,
		Description:  payload.Description,
	}

	category, err := cate.storedb.CreateCategory(c, args)
	if err != nil {
		if apiErr := models.ConvertToApiErr(err); apiErr != nil {
			c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(apiErr))
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, category)

}

func (cate *CategoryController) UpdateCategory(c *gin.Context) {
	var payload *models.CategoryUpdateReq
	cateId, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.UpdateCategoryParams{
		CategoryID:   int32(cateId),
		CategoryName: payload.CategoryName,
		Description:  payload.Description,
	}

	category, err := models.Nullable(cate.storedb.UpdateCategory(c, *args))
	if err != nil {
		/* 		if apiErr := models.ConvertToApiErr(err); apiErr != nil {
			c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(apiErr))
		} */
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	/* 	if category == nil {
		c.JSON(http.StatusNotFound, models.NewError(err))
		return
	} */
	c.JSON(http.StatusOK, category)

}

func (cate *CategoryController) DeleteCategory(c *gin.Context) {
	cateId, _ := strconv.Atoi(c.Param("id"))

	_, err := cate.storedb.FindCategoryById(c, int32(cateId))

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	err = cate.storedb.DeleteCategory(c, int32(cateId))
	if err != nil {
		if err != nil {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "data has been deleted"})

}
