package repository

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tranm/gassigeher/internal/models"
)

// setupTestDB creates a test database
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	// Create tables
	schema := `
	CREATE TABLE bookings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		dog_id INTEGER NOT NULL,
		date TEXT NOT NULL,
		walk_type TEXT CHECK(walk_type IN ('morning', 'evening')),
		scheduled_time TEXT NOT NULL,
		status TEXT DEFAULT 'scheduled' CHECK(status IN ('scheduled', 'completed', 'cancelled')),
		completed_at TIMESTAMP,
		user_notes TEXT,
		admin_cancellation_reason TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(dog_id, date, walk_type)
	);
	`

	if _, err := db.Exec(schema); err != nil {
		t.Fatalf("Failed to create schema: %v", err)
	}

	return db
}

func TestBookingRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewBookingRepository(db)

	booking := &models.Booking{
		UserID:        1,
		DogID:         1,
		Date:          "2025-12-01",
		WalkType:      "morning",
		ScheduledTime: "09:00",
	}

	err := repo.Create(booking)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if booking.ID == 0 {
		t.Error("Expected booking ID to be set")
	}

	if booking.Status != "scheduled" {
		t.Errorf("Expected status to be 'scheduled', got %s", booking.Status)
	}
}

func TestBookingRepository_CheckDoubleBooking(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewBookingRepository(db)

	// Create first booking
	booking := &models.Booking{
		UserID:        1,
		DogID:         1,
		Date:          "2025-12-01",
		WalkType:      "morning",
		ScheduledTime: "09:00",
	}
	repo.Create(booking)

	// Check for double booking
	isBooked, err := repo.CheckDoubleBooking(1, "2025-12-01", "morning")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !isBooked {
		t.Error("Expected dog to be marked as booked")
	}

	// Check different walk type
	isBooked, err = repo.CheckDoubleBooking(1, "2025-12-01", "evening")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if isBooked {
		t.Error("Expected evening slot to be available")
	}
}

func TestBookingRepository_AutoComplete(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewBookingRepository(db)

	// Create past booking
	yesterday := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
	booking := &models.Booking{
		UserID:        1,
		DogID:         1,
		Date:          yesterday,
		WalkType:      "morning",
		ScheduledTime: "09:00",
	}
	repo.Create(booking)

	// Run auto-complete
	count, err := repo.AutoComplete()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if count != 1 {
		t.Errorf("Expected 1 booking to be completed, got %d", count)
	}

	// Verify booking is completed
	completed, _ := repo.FindByID(booking.ID)
	if completed.Status != "completed" {
		t.Errorf("Expected status 'completed', got %s", completed.Status)
	}
}

func TestBookingRepository_Cancel(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewBookingRepository(db)

	booking := &models.Booking{
		UserID:        1,
		DogID:         1,
		Date:          "2025-12-01",
		WalkType:      "morning",
		ScheduledTime: "09:00",
	}
	repo.Create(booking)

	reason := "Dog is sick"
	err := repo.Cancel(booking.ID, &reason)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify cancellation
	cancelled, _ := repo.FindByID(booking.ID)
	if cancelled.Status != "cancelled" {
		t.Errorf("Expected status 'cancelled', got %s", cancelled.Status)
	}

	if cancelled.AdminCancellationReason == nil || *cancelled.AdminCancellationReason != reason {
		t.Error("Expected cancellation reason to be set")
	}
}
