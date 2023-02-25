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

// mapping  slice review product
func FormatRProductSlice(rProducts []ReviewProduct) []ReviewProductFormatter {
	rProductFormatter := []ReviewProductFormatter{}

	for _, rProduct := range rProducts {
		reviewProductFormatter := FormatRProduct(rProduct)
		rProductFormatter = append(rProductFormatter, reviewProductFormatter)
	}

	return rProductFormatter
}
