package products

type ProductFormatter struct {
	ProductId int     `json:"id_product"`
	Name      string  `json:"name"`
	Product   string  `json:"product"`
	Price     float64 `json:"price"`
}

func FormatProduct(product Product) ProductFormatter {

	formatter := ProductFormatter{
		ProductId: product.IDProduct,
		Name:      product.Name,
		Product:   product.Product,
		Price:     product.Price,
	}

	return formatter
}

// mapping  slice Products
func FormatProductSlice(products []Product) []ProductFormatter {
	productsFormatter := []ProductFormatter{}

	for _, product := range products {
		productFormatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}
