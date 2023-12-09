package bb

import "time"

func ParseUnixTimestamp(ms int64) *time.Time {
	ts := time.Unix(0, ms*int64(time.Millisecond))
	return &ts
}

func ParseRFC3339(timestamp string) (*time.Time, error) {
	timestamp = RemoveNonPrintableCharactersFromString(timestamp)
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
