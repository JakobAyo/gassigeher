package models

import (
	"testing"
)

func TestCreateBookingRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     CreateBookingRequest
		wantErr bool
	}{
		{
			name: "Valid request",
			req: CreateBookingRequest{
				DogID:         1,
				Date:          "2025-12-01",
				WalkType:      "morning",
				ScheduledTime: "09:00",
			},
			wantErr: false,
		},
		{
			name: "Missing dog ID",
			req: CreateBookingRequest{
				Date:          "2025-12-01",
				WalkType:      "morning",
				ScheduledTime: "09:00",
			},
			wantErr: true,
		},
		{
			name: "Invalid date format",
			req: CreateBookingRequest{
				DogID:         1,
				Date:          "01-12-2025",
				WalkType:      "morning",
				ScheduledTime: "09:00",
			},
			wantErr: true,
		},
		{
			name: "Invalid walk type",
			req: CreateBookingRequest{
				DogID:         1,
				Date:          "2025-12-01",
				WalkType:      "afternoon",
				ScheduledTime: "09:00",
			},
			wantErr: true,
		},
		{
			name: "Invalid time format",
			req: CreateBookingRequest{
				DogID:         1,
				Date:          "2025-12-01",
				WalkType:      "morning",
				ScheduledTime: "25:00",
			},
			wantErr: true,
		},
		{
			name: "Empty date",
			req: CreateBookingRequest{
				DogID:         1,
				Date:          "",
				WalkType:      "morning",
				ScheduledTime: "09:00",
			},
			wantErr: true,
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

func TestMoveBookingRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     MoveBookingRequest
		wantErr bool
	}{
		{
			name: "Valid request",
			req: MoveBookingRequest{
				Date:          "2025-12-01",
				WalkType:      "evening",
				ScheduledTime: "17:00",
				Reason:        "Dog not feeling well",
			},
			wantErr: false,
		},
		{
			name: "Missing reason",
			req: MoveBookingRequest{
				Date:          "2025-12-01",
				WalkType:      "evening",
				ScheduledTime: "17:00",
			},
			wantErr: true,
		},
		{
			name: "Invalid date",
			req: MoveBookingRequest{
				Date:          "invalid",
				WalkType:      "evening",
				ScheduledTime: "17:00",
				Reason:        "Test",
			},
			wantErr: true,
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
