package times

import (
	"time"
)

const (
	ASSIGN_TRIP_TIME           = time.Duration(5) * time.Second
	CHANGE_TRIP_STATE_TIME     = time.Duration(5) * time.Second
	DELIVERY_TIME              = time.Duration(25) * time.Second
	ESTIMATE_NEW_DELIVERY_TIME = time.Duration(15) * time.Second
)
