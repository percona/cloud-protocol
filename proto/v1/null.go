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

func NewNullInt64(v int64) NullInt64 {
    return NullInt64{sql.NullInt64{v, true}}
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

func NewNullFloat64(v float64) NullFloat64 {
    return NullFloat64{sql.NullFloat64{v, true}}
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (n *NullFloat64) MarshalJSON() (b []byte, err error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Float64)
}

func (n *NullFloat64) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		n.Float64 = 0
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &n.Float64)
	if err != nil {
		return err
	}
	n.Valid = true
	return nil
}

// sql package doesn't allow me to create proper float32
// so conversion is used to achieve the same
type NullFloat32 struct {
    sql.NullFloat64
}

func (n *NullFloat32) MarshalJSON() (b []byte, err error) {
    if !n.Valid {
        return []byte("null"), nil
    }
    return json.Marshal(float32(n.Float64))
}

func (n *NullFloat32) UnmarshalJSON(b []byte) error {
    if bytes.Equal(b, []byte("null")) {
        n.Float64 = 0
        n.Valid = false
        return nil
    }

    var v float32
    err := json.Unmarshal(b, &v)
    if err != nil {
        return err
    }

    n.Valid = true
    n.Float64 = float64(v)
    return nil
}

