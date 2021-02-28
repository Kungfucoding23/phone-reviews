package models

import (
	"errors"
	"time"
)

const (
	maxLenInComment = 400
)

//Review represent an anon review from some website
type Review struct {
	ID       int64
	Stars    int       // 1 - 5
	Comment  string    // max 400 chars
	Date     time.Time // created at
	GadgetID int64
}

//CreateReviewCMD command to create a new review
type CreateReviewCMD struct {
	Stars    int    `json:"stars"`
	Comment  string `json:"comment"`
	GadgetID int64  `json:"gadget_id"`
}

//Validate validates the stars and comments len
func (cmd *CreateReviewCMD) Validate() error {
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("stars must be between 1 - 5")
	}

	if len(cmd.Comment) > maxLenInComment {
		return errors.New("comment must be less than 400 chars")
	}
	return nil
}
