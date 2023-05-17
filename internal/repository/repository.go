package repository

import (
	"time"

	"github.com/ZeCarlos88/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InserReservation(res models.Reservation) (int, error)

	InsertRoomRestriction(r *models.RoomRestriction) error

	SearchAvailabiltyByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabiltyForAllRooms(start, end time.Time) ([]models.Room, error)
}
