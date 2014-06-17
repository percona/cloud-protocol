package proto

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (n *NullString) MarshalJSON() (b []byte, err error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.String)
}

func (n *NullString) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		n.String = ""
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &n.String)
	if err != nil {
		return err
	}
	n.Valid = true
	return nil
}

type NullInt64 struct {
	sql.NullInt64
}

func (n *NullInt64) MarshalJSON() (b []byte, err error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Int64)
}

func (n *NullInt64) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		n.Int64 = 0
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &n.Int64)
	if err != nil {
		return err
	}
	n.Valid = true
	return nil
}
