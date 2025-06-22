package dynamo

type GpsItem struct {
	Latitude  *float64 `dynamodbav:"Latitude"`
	Longitude *float64 `dynamodbav:"Longitude"`
	DeviceID  string   `dynamodbav:"DeviceID"`
	Timestamp string   `dynamodbav:"Timestamp"`
}

type GyroscopeItem struct {
	X         *float64 `dynamodbav:"X"`
	Y         *float64 `dynamodbav:"Y"`
	Z         *float64 `dynamodbav:"Z"`
	Timestamp string   `dynamodbav:"Timestamp"`
	DeviceID  string   `dynamodbav:"DeviceID"`
}

type PhotoItem struct {
	ImageBase64 string `dynamodbav:"ImageBase64"`
	Timestamp   string `dynamodbav:"Timestamp"`
	DeviceID    string `dynamodbav:"DeviceID"`
}
