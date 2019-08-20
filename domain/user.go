package domain

// User domain
type User struct {
	ID             int64  `json:"id"`
	Username       string `json:"username"`
	Fullname       string `json:"fullname"`
	Passhash       string `json:"passhash"`
	IdentityNumber int64  `json:"identity_number"`
	IsVoted        bool   `json:"is_voted"`
	Role           int    `json:"role"`
	Loginable      bool   `json:"loginable"`
}
