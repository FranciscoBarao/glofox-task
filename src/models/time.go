package models

import (
	"fmt"
	"glofox-task/middleware"
	"net/http"
	"strings"
	"time"
)

// Struct so that we can use dates in the format "yyyy-mm-dd"
type CustomTime struct {
	time.Time
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	date := t.Time.Format("2006-01-02")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return middleware.NewCustomError(http.StatusBadRequest, "Malformed input date format")
	}
	t.Time = date
	return nil
}
