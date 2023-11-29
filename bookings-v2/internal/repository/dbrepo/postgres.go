package dbrepo

import (
	"context"
	"github.com/tsawler/bookings-app/internal/models"
	"time"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	var insertedAt int

	stmt := `insert into reservations (first_name, last_name, email, phone, start_date,
                          end_date, room_id, created_at, updated_at)
                          values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&insertedAt)

	if err != nil {
		return insertedAt, err
	}

	return insertedAt, nil
}

func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	stmt := `insert into room_restrictions(start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at)
			values
			    ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.RestrictionID,
		r.RestrictionID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}
	return nil
}

// SearchAvailabilityByDatesByRoomID searches for the availability of a room between the given start and end time.
// It returns a boolean value indicating whether the room is available or not, and an error if any.
// The start and end time parameters specify the range for availability search.
// The roomID parameter specifies the ID of the room for which availability is being checked.
func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	query := `select
    			count(id) 
			  from 
 				room_restrictions rr 
			  where 
			    room_id = $1
 				and $2 < end_date and $3 > start_date;`

	row := m.DB.QueryRowContext(ctx, query, roomID, start, end)
	var numRows int
	err := row.Scan(&numRows)

	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}
	return false, nil
}

// SearchRoomAvailabilityForAllRooms searches for room availability for all rooms between the given start and end time.
// It returns a slice of models.Room and an error.
// The start and end time parameters specify the range for availability search.
func (m *postgresDBRepo) SearchRoomAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	var rooms []models.Room

	query := `
			select 
			    r.id, r.room_name
			from 
			    rooms r 
			where r.id not in 
			(select rr.room_id from room_restrictions rr where $1 < rr.end_date and $2 > rr.start_date);
`
	rows, err := m.DB.QueryContext(ctx, query, start, end)

	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.ID, &room.RoomName)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil

}
