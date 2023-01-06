package todo

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

	Todos struct {
		ID              int       `json:"id"`
		ActivityGroupID int       `json:"activity_group_id"`
		Title           string    `json:"title"`
		IsActive        bool      `json:"is_active"`
		Priority        string    `json:"priority"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		DeletedAt       *time.Time `json:"-"`
	}
)

func (a Todos) Bind(r *http.Request) error {

	switch r.Method {
	case "POST":
		if a.Title == "" {
			return errors.New("title cannot be null")
		}
		if a.ActivityGroupID == 0 {
			return errors.New("activity_group_id cannot be null")
		}
	}

	return nil
}