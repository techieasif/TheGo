package repository

import (
	"github.com/tsawler/bookings-app/internal/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchRoomAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
}
