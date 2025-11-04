package utils

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	// Assuming you have a utility function to convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

type NullString struct {
	sql.NullString
}

// Implement sql.Scanner (so you can scan DB rows into it)
func (ns *NullString) Scan(value interface{}) error {
	return ns.NullString.Scan(value)
}

// Implement driver.Valuer (optional, for inserts/updates)
func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.String, nil
}

// Implement JSON marshalling (so output is "value" or null)
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.String)
}
