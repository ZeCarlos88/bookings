package models

import "time"

type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatesAt   time.Time
	UpdatedAT   time.Time
}

type Room struct {
	ID        int
	RoomName  string
	CreatesAt time.Time
	UpdatedAT time.Time
}

type Restriction struct {
	ID              int
	RestrictionName string
	CreatesAt       time.Time
	UpdatedAT       time.Time
}

type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatesAt time.Time
	UpdatedAT time.Time
	Room      Room
}

type RoomRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	ReservationID int
	RestrictionID int
	CreatesAt     time.Time
	UpdatedAT     time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Restriction
}
