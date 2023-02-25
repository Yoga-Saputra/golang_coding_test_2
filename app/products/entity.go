package products

type Product struct {
	IDProduct int `gorm:"AUTO_INCREMENT"`
	Name      string
	Product   string
	Price     float64
}
