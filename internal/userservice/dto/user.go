package dto

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Mobile    string    `json:"mobile" db:"mobile"`
	Name      string    `json:"name" db:"name"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	LastSeen  time.Time `json:"last_seen" db:"last_seen"`
}

type UserPresence struct {
	UserID   int       `json:"user_id" db:"user_id"`
	IsOnline bool      `json:"is_online" db:"is_online"`
	LastSeen time.Time `json:"last_seen" db:"last_seen"`
}

type UserPresenceVisibility struct {
	UserID        int  `json:"user_id" db:"user_id"`     // The owner of the presence info
	ViewerID      int  `json:"viewer_id" db:"viewer_id"` // The user who wants to see it
	CanViewStatus bool `json:"can_view_status" db:"can_view_status"`
}

type UserChatContact struct {
	UserID      int       `json:"user_id" db:"user_id"`
	ContactID   int       `json:"contact_id" db:"contact_id"`
	LastMessage time.Time `json:"last_message" db:"last_message"`
}
