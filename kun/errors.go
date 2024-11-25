package kun

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

type Error error

var (
	// ErrNotFound error
	ErrNotFound Error = errors.New("not found")

	// ErrAlreadyExists error
	ErrAlreadyExists Error = errors.New("already exists")
)

// Handle error
func HandleError(err error) error {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.Join(ErrNotFound, err)
		}

		if errors.Is(err, sql.ErrNoRows) {
			return errors.Join(ErrNotFound, err)
		}

		if err.Error() == "record not found" {
			return errors.Join(ErrNotFound, err)
		}
	}
	return err
}
