package repositories

import (
	"github.com/bed72/oohferta/src/domain/constants"
)

var canHaveSuccessfulBody = [4]int{
	constants.StatusOK,
	constants.StatusCreated,
	constants.StatusAccepted,
	constants.StatusNoContent,
}

var canHaveErrorBody = [7]int{
	constants.StatusBadRequest,
	constants.StatusUnauthorized,
	constants.StatusForbidden,
	constants.StatusNotFound,
	constants.StatusNotAllowed,
	constants.StatusConflict,
	constants.StatusInternalServerError,
}

func HasErrorBody(status int) bool {
	for _, s := range canHaveErrorBody {
		if s == status {
			return true
		}
	}

	return false
}

func HasSuccessfulBody(status int) bool {
	for _, s := range canHaveSuccessfulBody {
		if s == status {
			return true
		}
	}

	return false
}
