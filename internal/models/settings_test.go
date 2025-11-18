package models

import (
	"testing"
)

// DONE: TestUpdateSettingRequest_Validate tests UpdateSettingRequest validation
func TestUpdateSettingRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     UpdateSettingRequest
		wantErr bool
	}{
		{
			name: "Valid request with numeric value",
			req: UpdateSettingRequest{
				Value: "14",
			},
			wantErr: false,
		},
		{
			name: "Valid request with text value",
			req: UpdateSettingRequest{
				Value: "Some configuration value",
			},
			wantErr: false,
		},
		{
			name: "Valid request with boolean-like value",
			req: UpdateSettingRequest{
				Value: "true",
			},
			wantErr: false,
		},
		{
			name: "Empty value",
			req: UpdateSettingRequest{
				Value: "",
			},
			wantErr: true,
		},
		{
			name: "Missing value",
			req:  UpdateSettingRequest{},
			wantErr: true,
		},
		{
			name: "Whitespace only value",
			req: UpdateSettingRequest{
				Value: "   ",
			},
			wantErr: false, // Current implementation only checks for empty string, not whitespace
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Additional check: if error expected, verify error type
			if tt.wantErr && err != nil {
				if _, ok := err.(*ValidationError); !ok {
					t.Errorf("Expected ValidationError, got %T", err)
				}
			}
		})
	}
}
