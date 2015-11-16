package oauth

import (
	"database/sql"
)

// Returns properly confiigured sql.NullInt64 for Client foreign keys
func clientIDOrNull(client *Client) sql.NullInt64 {
	if client == nil {
		return sql.NullInt64{Int64: 0, Valid: false}
	}

	return sql.NullInt64{Int64: int64(client.ID), Valid: true}
}

// Returns properly confiigured sql.NullInt64 for User foreign keys
func userIDOrNull(user *User) sql.NullInt64 {
	if user == nil {
		return sql.NullInt64{Int64: 0, Valid: false}
	}

	return sql.NullInt64{Int64: int64(user.ID), Valid: true}
}

// Returns properly confiigured sql.NullString
func stringOrNull(str string) sql.NullString {
	if str == "" {
		return sql.NullString{String: "", Valid: false}
	}

	return sql.NullString{String: str, Valid: true}
}
