package util

import (
	"database/sql"
	"time"
)

func NullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func NullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  *t,
		Valid: true,
	}
}
