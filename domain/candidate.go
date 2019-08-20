package domain

import "database/sql"

// Candidate domain
type Candidate struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Desc      sql.NullString `json:"description"`
	Picture   sql.NullString `json:"pic"`
	VoteCount sql.NullInt64  `json:"vote_count"`
}
