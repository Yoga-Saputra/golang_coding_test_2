package likereview

type LikeReview struct {
	ID_Review int `gorm:"column:id_review" json:"id_review"`
	ID_Member int `gorm:"column:id_member" json:"id_member"`
}
