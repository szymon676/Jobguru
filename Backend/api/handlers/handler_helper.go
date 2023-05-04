package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/szymon676/jobguru/types"
)

func WriteJSON(w http.ResponseWriter, status int, v ...any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func ParseDate(bj *types.JobReq) (time.Time, error) {
	dateStr := bj.Date
	dateLayout := "2006-01-02"

	date, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date")
	}

	return date, nil
}
