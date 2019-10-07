package calendar

import "errors"

// Date struct containing Year, Month and Day as ints
type Date struct {
	year  int
	month int
	day   int
}

// Year get date year
func (d *Date) Year() int {
	return d.year
}

// Month get date month
func (d *Date) Month() int {
	return d.month
}

// Day get date day
func (d *Date) Day() int {
	return d.day
}

// SetYear Accepts 4 digit int
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invaid year")
	}
	d.year = year
	return nil
}

// SetMonth Accepts 2 digit int
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invaid month")
	}
	d.month = month
	return nil
}

// SetDay Accepts 2 digit int
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invaid day")
	}
	d.day = day
	return nil
}
