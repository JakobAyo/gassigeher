package models

import (
	"testing"
)

// DONE: TestReviewReactivationRequestRequest_Validate tests validation for reactivation request review
func TestReviewReactivationRequestRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     ReviewReactivationRequestRequest
		wantErr bool
	}{
		{
			name: "approved without message",
			req: ReviewReactivationRequestRequest{
				Approved: true,
				Message:  nil,
			},
			wantErr: false,
		},
		{
			name: "approved with message",
			req: ReviewReactivationRequestRequest{
				Approved: true,
				Message:  stringPtr("Account reactivated successfully"),
			},
			wantErr: false,
		},
		{
			name: "denied without message",
			req: ReviewReactivationRequestRequest{
				Approved: false,
				Message:  nil,
			},
			wantErr: false,
		},
		{
			name: "denied with message",
			req: ReviewReactivationRequestRequest{
				Approved: false,
				Message:  stringPtr("Cannot reactivate at this time"),
			},
			wantErr: false,
		},
		{
			name: "empty message pointer",
			req: ReviewReactivationRequestRequest{
				Approved: true,
				Message:  stringPtr(""),
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
