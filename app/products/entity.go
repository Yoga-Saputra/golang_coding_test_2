package products

type Product struct {
	IDProduct int `gorm:"primarykey; AUTO_INCREMENT; column:id_product;"`
	Name      string
	Product   string
	Price     float64
}
