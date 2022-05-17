package cdwarehouse


type CdReviews struct {
	cdID string
	reviews []CdReview
}

type CdReview struct {
	Rating       int
	Comment      string
	CustomerName string
}

func NewCdReviews(cdID string) *CdReviews {
	return &CdReviews{
		cdID: cdID,
		reviews: []CdReview{},
	}
}

func (cr *CdReviews) AddReview(customer *Customer, rating int, comment string) bool {
	if rating < 1 || rating > 10 {
		return false
	}

	if !customer.HasPurchased(cr.cdID) {
		return false
	}

	review := CdReview{
		Rating:       rating,
		Comment:      comment,
		CustomerName: customer.Name,
	}

	cr.reviews = append(cr.reviews, review)

	return true
}
