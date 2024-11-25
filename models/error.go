package models

import (
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
)

var (
	ErrAccessForbidden     = errors.New("access forbidden")
	ErrUserAlreadyExist    = errors.New("Username or phone already exist")
	ErrDataNotFound        = errors.New("data not found")
	ErrInvalidUserPassword = errors.New("Invalid username or password")
	ErrUserPasswordEmpty   = errors.New("Username or password must not be empty")
	ErrLoginFailed         = errors.New("Login failed")
	ErrInvalidToken        = errors.New("Invalid token")
	ErrUnauthorizedToken   = errors.New("UnAuthorized token")
	ErrFailedToAuthorized  = errors.New("UnAuthorized token")
	ErrFailedGenerateToken = errors.New("Failed to generate token")
	ErrUpdateToken         = errors.New("Update token failed")
	ErrNoRows              = errors.New("no rows in result set")
	/*category*/
	ErrCategoryNameDuplicate = errors.New("Duplicate category name.")
	ErrCategoryNotFound      = errors.New("category not found")
)

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewValidationError(err error) *Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = err.Error()
	return &e
}

func NewError(err error) *Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["message"] = err.Error()
	return &e
}

func Nullable[T any](row *T, err error) (*T, error) {
	if err == nil {
		return row, nil
	}

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return nil, err
}

func NullableList[T any](rows []*T, err error) ([]*T, error) {
	if err == nil {
		return rows, nil
	}

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return nil, err
}

func NullableID(row string, err error) (string, error) {
	if err == nil {
		return row, nil
	}

	if err == pgx.ErrNoRows {
		return "", nil
	}

	return "", err
}

func ConvertToApiErr(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.ConstraintName {
		case "category_name_uq":
			return ErrCategoryNameDuplicate
		case "user_name_uq":
			return ErrUserAlreadyExist
		case "no rows in result set":
			return ErrNoRows
		}
	}
	return nil
}
