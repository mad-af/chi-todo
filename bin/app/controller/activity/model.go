package activity

import (
	"errors"
	"net/http"
	"time"
)

// TABLE
type (
	Activity struct {
		ID        int       `json:"id"`
		Email     *string    `json:"email"`
		Title     string    `json:"title"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt *time.Time `json:"-"`
	}
)

func (a Activity) Bind(r *http.Request) error {

	switch r.Method {
	case "POST":
		if a.Title == "" {
			return errors.New("title cannot be null")
		}
	case "PATCH":
		if a.Title == "" {
			return errors.New("title cannot be null")
		}
	}

	return nil
}