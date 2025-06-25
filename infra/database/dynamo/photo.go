package dynamo

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/google/uuid"
	"github.com/yanvic/challenge/core/entity"
)

func SavePhotoAnalysis(photo entity.Photo, result entity.PhotoAnalysisResult, rekognitionResult *rekognition.DetectFacesOutput, matchedID string, similarity float64, duration int64) error {
	record := PhotoAnalysisItem{
		UUID:                uuid.New().String(),
		DeviceID:            photo.DeviceID,
		PhotoID:             generatePhotoID(photo),
		Timestamp:           photo.Timestamp,
		RekognitionResponse: convertRekognitionResponse(rekognitionResult),
		IsRecognized:        result.Recognized,
		SimilarityScore:     similarity,
		MatchedPhotoID:      matchedID,
		AnalysisDurationMs:  duration,
		CreatedAt:           time.Now().UTC().Format(time.RFC3339),
	}

	item, err := attributevalue.MarshalMap(record)
	if err != nil {
		return err
	}

	tableName := "photo_analysis"
	_, err = Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})

	return err
}

func generatePhotoID(photo entity.Photo) string {
	return photo.DeviceID + "_" + photo.Timestamp
}

func convertRekognitionResponse(result *rekognition.DetectFacesOutput) map[string]interface{} {
	return map[string]interface{}{
		"faces_count": len(result.FaceDetails),
		"faces":       result.FaceDetails,
	}
}
