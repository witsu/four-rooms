package datetime

import "time"

type Date time.Time

func Parse(s string) (Date, error) {
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return Date{}, err
	}
	return Date(date), nil
}

func (d *Date) UnmarshalJSON(bytes []byte) error {
	date, err := time.Parse(`"2006-01-02"`, string(bytes))
	if err != nil {
		return err
	}
	*d = Date(date)

	return nil
}

func (d *Date) String() string {
	return time.Time(*d).Format(time.DateOnly)
}
