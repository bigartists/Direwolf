package model

import (
	"fmt"
)

var (
	ErrDuplicateModelBaseURL = fmt.Errorf("combination of model and base_url already exists")
	InValidUserId            = "invalid user id"
)
