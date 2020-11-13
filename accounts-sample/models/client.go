package models

import "time"

// Client struct for client info.
type Client struct {
	ID        string    `json:"id,omitempty"`
	FullName  string    `json:"full_name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Accounts  []Account `json:"accounts,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
