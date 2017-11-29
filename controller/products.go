package controllers

import (
	"github.com/frzifus/dbwt/model"
	"github.com/gernest/utron/controller"
	// "github.com/gorilla/schema"
	"net/http"
	"strconv"
)

type products struct {
	controller.BaseController
	Routes []string
}

func NewProducts() controller.Controller {
	return &products{
		Routes: []string{
			"get;/Products;Products",
			"get;/Detail;Detail",
		},
	}
}

func (p *products) Products() {
	// r := p.Ctx.Request()
	// user := p.URL.Query().Get("user")
	p.Ctx.Data["title"] = "Produkte"
	p.Ctx.Template = "products/products"
	p.HTML(http.StatusOK)
}

func (p *products) Detail() {
	r := p.Ctx.Request()
	productID := r.URL.Query().Get("product_id")

	id, err := strconv.Atoi(productID)

	if err != nil {
		// p.Ctx.Data["Message"] = err.Error()
		// p.Ctx.Template = "error"
		// p.HTML(http.StatusInternalServerError)
		// return
		p.Ctx.Redirect("/Products", http.StatusFound)
	}
	product := model.Product{}
	p.Ctx.DB.First(&product, id)
	p.Ctx.Data["title"] = "Detail"
	p.Ctx.Data["product"] = product
	p.Ctx.Template = "products/detail"
	p.HTML(http.StatusOK)
}