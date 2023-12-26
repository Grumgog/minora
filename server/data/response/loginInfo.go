package response

import "time"

type LoginInfo struct {
	Id           uint      `json:"id" validate:"required"`
	Username     string    `json:"username" validate:"required"`
	Email        string    `json:"email" validate:"required"`
	CreatedAt    time.Time `json:"createdAt" validate:"required"`
	LastLogin    time.Time `json:"lastLogin"`
	LastActionAt time.Time `json:"lastActionAt"`
	LastAction   Action    `json:"lastAction"`
}

type Action string

const (
	READ            Action = "READ"
	WRITE           Action = "WRITE"
	UPDATE          Action = "UPDATE"
	CREATE_RESOURCE Action = "CREATE_RESOURCE"
)
