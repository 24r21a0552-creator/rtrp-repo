package model

type Booking struct {
	Roll_no    string `json:"roll_no,omitempty"`
	Email      string `json:"email,omitempty"`
	Department string `json:"department,omitempty"`
	Sport      string `json:"sport,omitempty"`
	Date       string `json:"date,omitempty"`
	Time       string `json:"time,omitempty"`
	Venue      string `json:"venue,omitempty"`
}

type Cancellation struct {
	Email   string `json:"email,omitempty"`
	Roll_no string `json:"roll_no,omitempty"`
	Sport   string `json:"sport,omitempty"`
	Date    string `json:"date,omitempty"`
}

func NewBooking(roll, mail, dept, sport, date, time, venue string) Booking {
	return Booking{
		Roll_no:    roll,
		Email:      mail,
		Department: dept,
		Sport:      sport,
		Date:       date,
		Venue:      venue,
	}
}
