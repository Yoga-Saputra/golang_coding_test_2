package products

type InputProduct struct {
	Name    string  `json:"name" binding:"required"`
	Product string  `json:"product" binding:"required"`
	Price   float64 `json:"price" binding:"required"`
}

type GetProductDetailInput struct {
	IDProduct int `uri:"id" binding:"required"`
}
