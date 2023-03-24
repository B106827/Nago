package controllers

import (
	"NagoBackend/models"
	"NagoBackend/server/contexts"
	"net/http"

	"github.com/labstack/echo/v4"

	productForms "NagoBackend/forms/product"
)

type ProductController struct{}

// 商品一覧
func (pc *ProductController) Index(c echo.Context) error {
	pm := models.Product{}
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

// 商品情報
func (pc *ProductController) Get(c echo.Context) error {
	productForm := new(productForms.ProductForm)
	cc := c.(*contexts.CustomContext)
	if err := cc.BindValidate(productForm); err != nil {
		return c.JSON(http.StatusOK, badRequestResponse(err))
	}
	pm := models.Product{}
	product, err := pm.GetActiveProductWithImages(productForm.Id)
	if err != nil || product == nil {
		// エラーもしくは商品が存在しない
		return c.JSON(http.StatusOK, notFoundResponse([]string{"商品が存在しません"}))
	}
	return c.JSON(http.StatusOK, successResponse(map[string]interface{}{
		"product": product,
	}))
}
