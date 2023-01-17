package product

import (
	"bookstore/models/db"
	"bookstore/models/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (pc *productController) ProductList(c *gin.Context) {
	result, err := pc.productRepo.ProductList()

	user, _ := c.Get("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "failed",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    result,
		})
		return
	}
}

func (pc *productController) ProductCreate(c *gin.Context) {
	var (
		requestProduct = new(request.Product)
		err            error
	)

	err = c.Bind(requestProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed",
		})
	}

	user, _ := c.Get("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "failed",
		})
		return
	}

	dataProductDetails := []db.ProductDetail{}
	for _, v := range requestProduct.Details {
		dataProductDetail := db.ProductDetail{}
		dataProductDetail.Name = v.Name
		dataProductDetail.Stock = int32(v.Stock)
		dataProductDetail.Gambar = v.Gambar
		dataProductDetail.CreatedBy = int8(user.(db.User).ID)
		dataProductDetails = append(dataProductDetails, dataProductDetail)
	}

	dataProduct := db.Product{
		Name:           requestProduct.Name,
		Gambar:         requestProduct.Gambar,
		ProductDetails: dataProductDetails,
		Base: db.Base{
			CreatedBy: int8(user.(db.User).ID),
		},
	}

	result, err := pc.productRepo.ProductCreate(dataProduct)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed",
		})
		return
	} else {
		c.JSON(201, gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    result,
		})
		return
	}
}

func (pc *productController) ProductDetail(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed",
		})
	}

	user, _ := c.Get("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "failed",
		})
		return
	}

	result, err := pc.productRepo.ProductDetail(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "failed",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    result,
		})
		return
	}
}

func (pc *productController) ProductUpdate(c *gin.Context) {
	var (
		requestProduct = new(request.Product)
		err            error
	)

	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed",
		})
		return
	}

	err = c.Bind(requestProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed",
		})
		return
	}

	user, _ := c.Get("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "failed",
		})
		return
	}

	product, err := pc.productRepo.ProductDetail(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "failed",
		})
		return
	} else {

		_, err := pc.productRepo.ProductUpdate(product, *requestProduct)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "failed",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "success",
			})
			return
		}
	}

}

func (pc *productController) ProductDelete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed",
		})
		return
	}

	user, _ := c.Get("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "failed",
		})
		return
	}

	_, err := pc.productRepo.ProductDelete(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "failed",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
		})
		return
	}
}
