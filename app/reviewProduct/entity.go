package reviewproduct

import "codingTest/app/members"

type ReviewProduct struct {
	ID_Review  int              `gorm:"column:id_review; primary_key; AUTO_INCREMENT;"`
	ID_Product int              `gorm:"column:id_product"`
	Member_ID  int              `gorm:"column:id_member"`
	DescReview string           `gorm:"column:desc_review"`
	Members    []members.Member `gorm:"many2many:like_reviews;ForeignKey:id_member;joinForeignKey:id_member;References:id_member;joinReferences:id_member;"`
}
