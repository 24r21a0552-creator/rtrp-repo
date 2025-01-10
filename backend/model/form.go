package model

type Booking struct {
	roll_no    string `json:"roll_no,omitempty"`
	email      string `json:"email,omitempty"`
	department string `json:"department,omitempty"`
	sport      string `json:"sport,omitempty"`
	date       string `json:"date,omitempty"`
	time       string `json:"time,omitempty"`
	venue      string `json:"venue,omitempty"`
}

func NewBooking(roll, mail, dept, sport, date, time, venue string) Booking {
	return Booking{
		roll_no:    roll,
		email:      mail,
		department: dept,
		sport:      sport,
		date:       date,
		venue:      venue,
	}
}
