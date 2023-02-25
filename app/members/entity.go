package members

type Member struct {
	IDMember  int    `gorm:"primarykey; AUTO_INCREMENT; column:id_member;"`
	Gender    string `gorm:"column:gender; type:enum('L','P')"`
	Username  string `gorm:"column:username;"`
	Skintype  string `gorm:"column:skintype;"`
	Skincolor string `gorm:"column:skincolor;"`
}
