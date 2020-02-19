package models

// Error structure
type Error struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Message    string `json:"message"`
	StatusCode int64  `json:"status_code"`
}
