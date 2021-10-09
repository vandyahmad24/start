package students

import "time"

type Students struct {
	ID        int    `json:"id"`
	Name      string `json:"name" `
	Age       int    `json:"age"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
