package controller

import (
	db "AirBnbProject/db/sqlc"
	"AirBnbProject/models"
	"AirBnbProject/services"
	"database/sql"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductController struct {
	storedb services.Store
}

func NewProductController(store services.Store) *ProductController {
	return &ProductController{
		storedb: store,
	}
}

type ProductCreateDto struct {
	//ProductID       int16    `form:"product_id"`
	ProductName     string   `form:"product_name" binding:"required"`
	SupplierID      *int16   `form:"supplier_id"`
	QuantityPerUnit *string  `form:"quantity_per_unit"`
	UnitPrice       *float32 `form:"unit_price"`
	UnitsInStock    *int16   `form:"units_in_stock"`
	UnitsOnOrder    *int16   `form:"units_on_order"`
	ReorderLevel    *int16   `form:"reorder_level"`
	Discontinued    int32    `form:"discontinued"`
	//Filename        *multipart.FileHeader `form:"filename" binding:"required"`
	Filename *SingleFileUpload
}

type SingleFileUpload struct {
	Filename *multipart.FileHeader `form:"filename" binding:"required"`
}

type MultipleFileUpload struct {
	Filename []*multipart.FileHeader `form:"filename" binding:"required"`
}

func (handler *ProductController) UploadMultipleProductImage(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	files := form.File["filename"]

	for _, v := range files {
		extension := filepath.Ext(v.Filename)
		// Generate random file name for the new uploaded file so it doesn't override the old file with same name
		newFileName := uuid.New().String() + extension

		// The file is received, so let's save it
		if err := c.SaveUploadedFile(v, "./public/"+newFileName); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"status": "ok", "message": "multiple product has been uploaded."})
}

func (handler *ProductController) CreateProduct(c *gin.Context) {
	var payload *ProductCreateDto

	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	fileUpload, err := c.FormFile("filename")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(fileUpload.Filename)
	// Generate random file name for the new uploaded file
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(fileUpload, "./public/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	args := &db.CreateProductParams{
		ProductName:     payload.ProductName,
		SupplierID:      payload.SupplierID,
		QuantityPerUnit: payload.QuantityPerUnit,
		UnitPrice:       payload.UnitPrice,
		UnitsInStock:    payload.UnitsInStock,
		UnitsOnOrder:    payload.UnitsOnOrder,
		Discontinued:    payload.Discontinued,
		ProductImage:    &newFileName,
	}

	product, err := handler.storedb.CreateProduct(c, *args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (handler *ProductController) DeleteProduct(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("id"))

	_, err := handler.storedb.FindProductById(c, int16(productId))

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	err = handler.storedb.DeleteProduct(c, int16(productId))
	if err != nil {
		if err != nil {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "data has been deleted"})
}

func (handler *ProductController) FindAllProduct(c *gin.Context) {
	products, err := handler.storedb.FindAllProduct(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, products)
}

func (handler *ProductController) FindAllProductPaging(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	args := &db.FindAllProductPagingParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	products, err := handler.storedb.FindAllProductPaging(c, *args)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, products)

}

func (handler *ProductController) FindProductById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := handler.storedb.FindProductById(c, int16(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, product)
}

type UpdateCreateDto struct {
	ProductID       string   `form:"product_id"`
	ProductName     string   `form:"product_name" binding:"required"`
	SupplierID      *int16   `form:"supplier_id"`
	CategoryID      *int16   `form:"category_id"`
	QuantityPerUnit *string  `form:"quantity_per_unit"`
	UnitPrice       *float32 `form:"unit_price"`
	UnitsInStock    *int16   `form:"units_in_stock"`
	UnitsOnOrder    *int16   `form:"units_on_order"`
	ReorderLevel    *int16   `form:"reorder_level"`
	Discontinued    int32    `form:"discontinued"`
	ProductImage    *string  `form:"product_image"`
	Filename        *SingleFileUpload
}

func (handler *ProductController) UpdateProduct(c *gin.Context) {
	var payload *UpdateCreateDto
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.UpdateProductParams{
		ProductName:     payload.ProductName,
		SupplierID:      payload.SupplierID,
		CategoryID:      payload.CategoryID,
		QuantityPerUnit: payload.QuantityPerUnit,
		UnitsInStock:    payload.UnitsInStock,
		UnitPrice:       payload.UnitPrice,
		UnitsOnOrder:    payload.UnitsOnOrder,
		ReorderLevel:    payload.ReorderLevel,
		Discontinued:    payload.Discontinued,
		ProductImage:    &payload.ProductName,
		ProductID:       int16(id),
	}

	category, err := handler.storedb.UpdateProduct(c, *args)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, category)
}
