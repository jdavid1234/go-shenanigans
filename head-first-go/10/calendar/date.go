// B"H

package calendar

import "errors"

// -- -----------------------------
// Note the lowercase fields to keep them unexported.
// I.e. only accessible via the getters/setters.
type Date struct {
	year  int
	month int
	day   int
}

// -- -----------------------------
func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

// -- -----------------------------
// year setter method (note how we need pointer receivers):
func (d *Date) SetYear(year int) error {

	if year < 1 {
		return errors.New("invalid year")
	}

	d.year = year

	return nil
}

// -- -----------------------------
// month setter method (note how we need pointer receivers):
func (d *Date) SetMonth(month int) error {

	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}

	d.month = month

	return nil
}

// -- -----------------------------
// day setter method (note how we need pointer receivers):
func (d *Date) SetDay(day int) error {

	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}

	d.day = day

	return nil
}
