package messaging

type GyroscopeMessage struct {
	X, Y, Z   float64
	Timestamp int64
	DeviceID  string
}

type GPSMessage struct {
	Latitude, Longitude float64
	Timestamp           int64
	DeviceID            string
}

type PhotoMessage struct {
	Photo     []byte
	Timestamp int64
	DeviceID  string
}
