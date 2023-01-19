package controllers

import (
	"NagoBackend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

// 商品一覧
func (pc *ProductController) Index(c echo.Context) error {
	pm := new(models.Product)
	products, err := pm.GetActiveProductsWithImages()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusOK, serverErrorResponse([]string{"エラーが発生しました"}))
	}
	if products == nil {
		return c.JSON(http.StatusOK, notFoundResponse([]string{"商品が見つかりません"}))
	}
	return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
		"products": products,
	}))
}
