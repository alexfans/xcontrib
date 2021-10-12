package servicehelper

import "time"

const (
	layout = "2006-01-02 15:04:05"
)

func TimeFormat(t time.Time) string {
	return t.Format(layout)
}

func TimeParse(s string) time.Time {
	t, err := time.Parse(layout, s)
	if err != nil {

	}
	return t
}
