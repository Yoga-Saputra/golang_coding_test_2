package reviewproduct

type InputReviewProduct struct {
	ID_Product int    `json:"id_product"`
	ID_Memebr  int    `json:"id_member"`
	DescReview string `json:"desc_review"`
}
