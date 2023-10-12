package lib

import "time"

func ConvertToEpoch(timestamp time.Time) int64 {
	var epochTimestamp int64
	if !timestamp.IsZero() {
		epochTimestamp = timestamp.Unix()
	}

	return epochTimestamp
}
