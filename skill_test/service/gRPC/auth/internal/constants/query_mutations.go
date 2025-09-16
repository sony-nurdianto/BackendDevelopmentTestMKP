package constants

const (
	QueryInsertUser = `
		insert into users
			(card_id,full_name, email,phone,balance)
		values
			($1,$2,$3,$4,$5)
		returning id, card_id, full_name, email, phone, balance, status, current_trip_terminal_id, registered_at, updated_at
	`
)
