package models

import "testing"

func NewReview(stars int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Stars:   stars,
		Comment: comment,
	}
}

func TestCreateReviewValidateNoError(t *testing.T) {
	review := NewReview(4, "The iphone C looks good")

	err := review.Validate()

	if err != nil {
		t.Error("the validation did not pass")
		t.Fail()
	}
}

func TestCreateReviewValidateWithError(t *testing.T) {
	review := NewReview(8, "The iphone looks REALLY good")
	err := review.Validate()
	if err == nil {
		t.Error("should fail with more than 5 stars")
		t.Fail()
	}
}
