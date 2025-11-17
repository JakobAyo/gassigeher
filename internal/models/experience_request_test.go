package models

import (
	"testing"
)

// DONE: TestCreateExperienceRequestRequest_Validate tests validation for experience request creation
func TestCreateExperienceRequestRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     CreateExperienceRequestRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid blue request",
			req: CreateExperienceRequestRequest{
				RequestedLevel: "blue",
			},
			wantErr: false,
		},
		{
			name: "valid orange request",
			req: CreateExperienceRequestRequest{
				RequestedLevel: "orange",
			},
			wantErr: false,
		},
		{
			name: "invalid level - green",
			req: CreateExperienceRequestRequest{
				RequestedLevel: "green",
			},
			wantErr: true,
			errMsg:  "Requested level must be 'blue' or 'orange'",
		},
		{
			name: "invalid level - empty",
			req: CreateExperienceRequestRequest{
				RequestedLevel: "",
			},
			wantErr: true,
			errMsg:  "Requested level must be 'blue' or 'orange'",
		},
		{
			name: "invalid level - uppercase",
			req: CreateExperienceRequestRequest{
				RequestedLevel: "BLUE",
			},
			wantErr: true,
			errMsg:  "Requested level must be 'blue' or 'orange'",
		},
		{
			name: "invalid level - random string",
			req: CreateExperienceRequestRequest{
				RequestedLevel: "expert",
			},
			wantErr: true,
			errMsg:  "Requested level must be 'blue' or 'orange'",
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
				// Check error message contains expected text (accounting for field prefix)
				if !contains(err.Error(), tt.errMsg) {
					t.Errorf("Validate() error = %v, expected to contain %v", err.Error(), tt.errMsg)
				}
			}
		})
	}
}

// DONE: TestReviewExperienceRequestRequest_Validate tests validation for reviewing experience requests
func TestReviewExperienceRequestRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     ReviewExperienceRequestRequest
		wantErr bool
	}{
		{
			name: "approved without message",
			req: ReviewExperienceRequestRequest{
				Approved: true,
				Message:  nil,
			},
			wantErr: false,
		},
		{
			name: "approved with message",
			req: ReviewExperienceRequestRequest{
				Approved: true,
				Message:  stringPtr("Great progress!"),
			},
			wantErr: false,
		},
		{
			name: "denied without message",
			req: ReviewExperienceRequestRequest{
				Approved: false,
				Message:  nil,
			},
			wantErr: false,
		},
		{
			name: "denied with message",
			req: ReviewExperienceRequestRequest{
				Approved: false,
				Message:  stringPtr("Need more experience"),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
