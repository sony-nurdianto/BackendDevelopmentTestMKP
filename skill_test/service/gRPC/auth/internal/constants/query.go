package constants

const (
	QueryGetUserByID = `
		select 
			id, 
			card_id, 
			full_name, 
			email, 
			phone, 
			balance, 
			status, 
			current_trip_terminal_id, 
			registered_at, 
			updated_at
		from users
		where 
			card_id = $1
	`

	QueryGetUserByEmail = `
		select 
			id, 
			card_id, 
			full_name, 
			email, 
			phone, 
			balance, 
			status, 
			current_trip_terminal_id, 
			registered_at, 
			updated_at
		from users
		where 
			email = $1
	`
)
