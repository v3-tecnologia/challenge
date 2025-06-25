package dynamo

type GpsItem struct {
	UUID      string   `dynamodbav:"uuid"`
	Latitude  *float64 `dynamodbav:"Latitude"`
	Longitude *float64 `dynamodbav:"Longitude"`
	DeviceID  string   `dynamodbav:"DeviceID"`
	Timestamp string   `dynamodbav:"Timestamp"`
}

type GyroscopeItem struct {
	UUID      string   `dynamodbav:"uuid"`
	X         *float64 `dynamodbav:"X"`
	Y         *float64 `dynamodbav:"Y"`
	Z         *float64 `dynamodbav:"Z"`
	Timestamp string   `dynamodbav:"Timestamp"`
	DeviceID  string   `dynamodbav:"DeviceID"`
}

type PhotoItem struct {
	UUID        string `dynamodbav:"uuid"`
	ImageBase64 string `dynamodbav:"ImageBase64"`
	Timestamp   string `dynamodbav:"Timestamp"`
	DeviceID    string `dynamodbav:"DeviceID"`
}

type PhotoAnalysisItem struct {
	UUID                string                 `dynamodbav:"uuid"`
	DeviceID            string                 `dynamodbav:"device_id"`
	PhotoID             string                 `dynamodbav:"photo_id"`
	Timestamp           string                 `dynamodbav:"timestamp"`
	RekognitionResponse map[string]interface{} `dynamodbav:"rekognition_response"`
	IsRecognized        bool                   `dynamodbav:"is_recognized"`
	SimilarityScore     float64                `dynamodbav:"similarity_score"`
	MatchedPhotoID      string                 `dynamodbav:"matched_photo_id"`
	AnalysisDurationMs  int64                  `dynamodbav:"analysis_duration_ms"`
	CreatedAt           string                 `dynamodbav:"created_at"`
}
