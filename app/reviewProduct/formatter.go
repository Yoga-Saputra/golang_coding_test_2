package reviewproduct

type ReviewProductFormatter struct {
	DescReview       string      `json:"desc_review"`
	Gender           interface{} `json:"gender"`
	Username         interface{} `json:"username"`
	Skintype         interface{} `json:"skintype"`
	Skincolor        interface{} `json:"skincolor"`
	TotalLikeReviews int         `json:"total_like_reviews"`
}

// single Review Product
func FormatRProduct(rProduct ReviewProduct) ReviewProductFormatter {
	var gender interface{}
	var username interface{}
	var skinType interface{}
	var skinColor interface{}
	// var countTo interface{}

	if len(rProduct.Members) > 0 {
		for _, member := range rProduct.Members {
			gender = member.Gender
			username = member.Username
			skinType = member.Skintype
			skinColor = member.Skincolor
		}
	}

	formater := ReviewProductFormatter{
		DescReview:       rProduct.DescReview,
		Gender:           gender,
		Username:         username,
		Skintype:         skinType,
		Skincolor:        skinColor,
		TotalLikeReviews: len(rProduct.Members),
	}

	return formater
}

type ReviewCreateFormatter struct {
	ID_Review  int    `json:"id_review"`
	ID_Product int    `son:"id_product"`
	Member_ID  int    `json:"id_member"`
	DescReview string `on:"desc_review"`
}

func FormatRProductCreate(rProduct ReviewProduct) ReviewCreateFormatter {

	formatter := ReviewCreateFormatter{
		ID_Review:  rProduct.ID_Review,
		ID_Product: rProduct.ID_Product,
		Member_ID:  rProduct.Member_ID,
		DescReview: rProduct.DescReview,
	}

	return formatter
}

// mapping  slice review product
func FormatRProductSlice(rProducts []ReviewProduct) []ReviewProductFormatter {
	rProductFormatter := []ReviewProductFormatter{}

	for _, rProduct := range rProducts {
		reviewProductFormatter := FormatRProduct(rProduct)
		rProductFormatter = append(rProductFormatter, reviewProductFormatter)
	}

	return rProductFormatter
}
