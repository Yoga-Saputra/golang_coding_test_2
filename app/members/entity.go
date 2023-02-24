package members

type Member struct {
	IDMember  int    `gorm:"primarykey"`
	Gender    string `gorm:"column:gender; type:enum('L','P')"`
	Username  string
	Skintype  string
	Skincolor string
}
