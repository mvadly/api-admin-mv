package product

import (
	"mime/multipart"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Product struct {
	ProductName string                  `form:"product_name" json:"product_name"`
	Qty         int16                   `form:"qty" json:"qty"`
	QtyUnit     string                  `form:"qty_unit" json:"qty_unit"`
	Price       float32                 `form:"price" json:"price"`
	QtyPrice    string                  `form:"qty_price" json:"qty_price"`
	Status      int                     `form:"status" json:"status"`
	ComodityId  string                  `form:"comodity_id" json:"comodity_id"`
	ImgProduct  []*multipart.FileHeader `form:"img_product" json:"img_product"`
	DateProduct string                  `form:"date_product" json:"date_product"`
	ExpProduct  string                  `form:"exp_product" json:"exp_product"`
}

func (p Product) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.ProductName, validation.Required, validation.Length(2, 100)),
		validation.Field(&p.Qty, validation.Required),
		validation.Field(&p.QtyUnit, validation.Required, validation.Length(1, 20)),
		validation.Field(&p.Price, validation.Required),
		validation.Field(&p.QtyPrice, validation.Required, validation.Length(1, 20)),
		validation.Field(&p.ComodityId, validation.Required),
		validation.Field(&p.ImgProduct, validation.Required),
		validation.Field(&p.DateProduct, validation.Required),
		validation.Field(&p.ExpProduct, validation.Required),
	)
}
