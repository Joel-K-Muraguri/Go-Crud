package responses

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "title") {
		return errors.New("TITLE ALREADY TAKEN")
	}
	return errors.New("INCORRECT DETAILS")
}
