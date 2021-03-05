package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/leevdh/udemy-bookstore_utilities_go/rest_errors"
)

const ErrorNoRow = "no rows in result set"

func ParseError(err error) rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRow) {
			return rest_errors.NewNotFoundError("no record matching given ID")
		}

		return rest_errors.NewInternalServiceError("database error: error parsing database response", sqlErr)
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("duplicate invalid data")
	}

	return rest_errors.NewInternalServiceError("database error: error processing request", sqlErr)
}
