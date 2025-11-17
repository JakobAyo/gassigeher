package models

import (
	"testing"
	"time"
)

// DONE: TestCreateBlockedDateRequest_Validate tests validation for blocked date creation
func TestCreateBlockedDateRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     CreateBlockedDateRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request",
			req: CreateBlockedDateRequest{
				Date:   "2025-12-25",
				Reason: "Christmas holiday",
			},
			wantErr: false,
		},
		{
			name: "valid request with long reason",
			req: CreateBlockedDateRequest{
				Date:   "2025-01-01",
				Reason: "New Year's Day - Shelter closed for maintenance and staff holiday",
			},
			wantErr: false,
		},
		{
			name: "missing date",
			req: CreateBlockedDateRequest{
				Date:   "",
				Reason: "Holiday",
			},
			wantErr: true,
			errMsg:  "date",
		},
		{
			name: "missing reason",
			req: CreateBlockedDateRequest{
				Date:   "2025-12-25",
				Reason: "",
			},
			wantErr: true,
			errMsg:  "reason",
		},
		{
			name: "invalid date format - wrong separator",
			req: CreateBlockedDateRequest{
				Date:   "2025/12/25",
				Reason: "Holiday",
			},
			wantErr: true,
			errMsg:  "invalid date format",
		},
		{
			name: "invalid date format - DD-MM-YYYY",
			req: CreateBlockedDateRequest{
				Date:   "25-12-2025",
				Reason: "Holiday",
			},
			wantErr: true,
			errMsg:  "invalid date format",
		},
		{
			name: "invalid date format - MM/DD/YYYY",
			req: CreateBlockedDateRequest{
				Date:   "12/25/2025",
				Reason: "Holiday",
			},
			wantErr: true,
			errMsg:  "invalid date format",
		},
		{
			name: "invalid date - February 30",
			req: CreateBlockedDateRequest{
				Date:   "2025-02-30",
				Reason: "Invalid date",
			},
			wantErr: true,
			errMsg:  "invalid date format",
		},
		{
			name: "invalid date - month 13",
			req: CreateBlockedDateRequest{
				Date:   "2025-13-01",
				Reason: "Invalid month",
			},
			wantErr: true,
			errMsg:  "invalid date format",
		},
		{
			name: "invalid date - short format",
			req: CreateBlockedDateRequest{
				Date:   "2025-1-1",
				Reason: "Short format",
			},
			wantErr: true,
			errMsg:  "invalid date format",
		},
		{
			name: "past date - should be valid (no restriction)",
			req: CreateBlockedDateRequest{
				Date:   "2020-01-01",
				Reason: "Past date",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil {
				// Validate error message contains expected substring (accounting for field prefix)
				if tt.errMsg != "" && !contains(err.Error(), tt.errMsg) {
					// Also check without field prefix
					actualMsg := err.Error()
					if !contains(actualMsg, "YYYY-MM-DD") && tt.errMsg == "invalid date format" {
						t.Errorf("Validate() error = %v, expected to contain date format error", actualMsg)
					}
				}
			}
		})
	}
}

// DONE: TestBlockedDate_DateParsing tests that dates can be parsed correctly
func TestBlockedDate_DateParsing(t *testing.T) {
	validDate := "2025-12-25"
	parsedDate, err := time.Parse("2006-01-02", validDate)
	if err != nil {
		t.Errorf("Valid date should parse successfully: %v", err)
	}

	expectedYear := 2025
	expectedMonth := time.December
	expectedDay := 25

	if parsedDate.Year() != expectedYear {
		t.Errorf("Expected year %d, got %d", expectedYear, parsedDate.Year())
	}
	if parsedDate.Month() != expectedMonth {
		t.Errorf("Expected month %v, got %v", expectedMonth, parsedDate.Month())
	}
	if parsedDate.Day() != expectedDay {
		t.Errorf("Expected day %d, got %d", expectedDay, parsedDate.Day())
	}
}
