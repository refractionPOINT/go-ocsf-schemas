// https://schema.ocsf.io/objects/reputation
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type Reputation struct {
	Provider    string  `json:"provider"`
	Recommended string  `json:"recommended"`
	BaseScore   float64 `json:"base_score" validate:"required"`
	Score       string  `json:"score"`
	ScoreID     int     `json:"score_id" validate:"required"`
}

func ValidateReputation(rep *Reputation) (*Reputation, error) {
	err := validator.New().Struct(rep)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return rep, nil
}
