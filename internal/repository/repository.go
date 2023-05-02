package repository

import "github.com/ZeCarlos88/bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InserReservation(res models.Reservation) (int, error)

	InsertRoomRestriction(r *models.RoomRestriction) error
}
