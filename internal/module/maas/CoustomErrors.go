package maas

import (
	"fmt"
)

var (
	ErrDuplicateMaasBaseURL = fmt.Errorf("combination of maas and base_url already exists")
)
