package services

import (
	"math"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"telemetry-api/internal/dtos/requests"
	"telemetry-api/internal/dtos/response"
	"telemetry-api/internal/models"
	"telemetry-api/internal/utils"
)

type TelemetryService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewTelemetryService(db *gorm.DB, logger *zap.Logger) *TelemetryService {
	return &TelemetryService{db: db, logger: logger}
}

func (s *TelemetryService) CreateGyroscope(gyroscopeDTO *requests.CreateGyroscopeRequest) error {
	gyroscope := models.Gyroscope{
		DeviceID:  gyroscopeDTO.DeviceID,
		X:         gyroscopeDTO.X,
		Y:         gyroscopeDTO.Y,
		Z:         gyroscopeDTO.Z,
		Timestamp: gyroscopeDTO.Timestamp,
	}

	if err := s.db.Create(&gyroscope).Error; err != nil {
		s.logger.Error("failed to create gyroscope entry in service", zap.Error(err))
		return err
	}
	return nil
}

func (s *TelemetryService) CreateGPS(gpsDTO *requests.CreateGPSRequest) error {
	gps := models.GPS{
		DeviceID:  gpsDTO.DeviceID,
		Latitude:  gpsDTO.Latitude,
		Longitude: gpsDTO.Longitude,
		Timestamp: gpsDTO.Timestamp,
	}

	if err := s.db.Create(&gps).Error; err != nil {
		s.logger.Error("failed to create gps entry in service", zap.Error(err))
		return err
	}
	return nil
}

func (s *TelemetryService) CreateTelemetryPhoto(telemetryPhotoDTO *requests.CreateTelemetryPhotoRequest) error {
	if !utils.IsValidImageBase64(telemetryPhotoDTO.Photo) {
		return utils.ErrInvalidPhotoFormat
	}

	telemetryPhoto := models.TelemetryPhoto{
		DeviceID:  telemetryPhotoDTO.DeviceID,
		Photo:     telemetryPhotoDTO.Photo,
		Timestamp: telemetryPhotoDTO.Timestamp,
	}

	if err := s.db.Create(&telemetryPhoto).Error; err != nil {
		s.logger.Error("failed to create telemetry photo entry in service", zap.Error(err))
		return err
	}

	return nil
}

func (s *TelemetryService) GetGyroscopeData(deviceID string, page, limit int) (*response.GyroscopeListResponse, error) {
	var gyroscopes []models.Gyroscope
	var total int64

	query := s.db.Model(&models.Gyroscope{})
	if deviceID != "" {
		query = query.Where("device_id = ?", deviceID)
	}

	if err := query.Count(&total).Error; err != nil {
		s.logger.Error("failed to count gyroscope records", zap.Error(err))
		return nil, err
	}

	offset := (page - 1) * limit
	if err := query.Order("timestamp DESC").Limit(limit).Offset(offset).Find(&gyroscopes).Error; err != nil {
		s.logger.Error("failed to fetch gyroscope data", zap.Error(err))
		return nil, err
	}

	var gyroscopeResponses []response.GyroscopeResponse
	for _, g := range gyroscopes {
		gyroscopeResponses = append(gyroscopeResponses, response.GyroscopeResponse{
			ID:        g.ID,
			DeviceID:  g.DeviceID,
			X:         g.X,
			Y:         g.Y,
			Z:         g.Z,
			Timestamp: g.Timestamp,
			CreatedAt: g.CreatedAt,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &response.GyroscopeListResponse{
		Data: gyroscopeResponses,
		Pagination: response.PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			Total:       total,
			TotalPages:  totalPages,
		},
	}, nil
}

func (s *TelemetryService) GetGPSData(deviceID string, page, limit int) (*response.GPSListResponse, error) {
	var gpsData []models.GPS
	var total int64

	query := s.db.Model(&models.GPS{})
	if deviceID != "" {
		query = query.Where("device_id = ?", deviceID)
	}

	if err := query.Count(&total).Error; err != nil {
		s.logger.Error("failed to count GPS records", zap.Error(err))
		return nil, err
	}

	offset := (page - 1) * limit
	if err := query.Order("timestamp DESC").Limit(limit).Offset(offset).Find(&gpsData).Error; err != nil {
		s.logger.Error("failed to fetch GPS data", zap.Error(err))
		return nil, err
	}

	var gpsResponses []response.GPSResponse
	for _, g := range gpsData {
		gpsResponses = append(gpsResponses, response.GPSResponse{
			ID:        g.ID,
			DeviceID:  g.DeviceID,
			Latitude:  g.Latitude,
			Longitude: g.Longitude,
			Timestamp: g.Timestamp,
			CreatedAt: g.CreatedAt,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &response.GPSListResponse{
		Data: gpsResponses,
		Pagination: response.PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			Total:       total,
			TotalPages:  totalPages,
		},
	}, nil
}

func (s *TelemetryService) GetPhotoData(deviceID string, page, limit int) (*response.TelemetryPhotoListResponse, error) {
	var photos []models.TelemetryPhoto
	var total int64

	query := s.db.Model(&models.TelemetryPhoto{})
	if deviceID != "" {
		query = query.Where("device_id = ?", deviceID)
	}

	if err := query.Count(&total).Error; err != nil {
		s.logger.Error("failed to count photo records", zap.Error(err))
		return nil, err
	}

	offset := (page - 1) * limit
	if err := query.Order("timestamp DESC").Limit(limit).Offset(offset).Find(&photos).Error; err != nil {
		s.logger.Error("failed to fetch photo data", zap.Error(err))
		return nil, err
	}

	var photoResponses []response.TelemetryPhotoResponse
	for _, p := range photos {
		photoResponses = append(photoResponses, response.TelemetryPhotoResponse{
			ID:        p.ID,
			DeviceID:  p.DeviceID,
			Photo:     p.Photo,
			Timestamp: p.Timestamp,
			CreatedAt: p.CreatedAt,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &response.TelemetryPhotoListResponse{
		Data: photoResponses,
		Pagination: response.PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			Total:       total,
			TotalPages:  totalPages,
		},
	}, nil
}

func (s *TelemetryService) GetDevices(page, limit int) (*response.DeviceListResponse, error) {
	type DeviceStats struct {
		DeviceID        string `gorm:"column:device_id"`
		LastSeen        string `gorm:"column:last_seen"`
		GyroscopeCount  int64  `gorm:"column:gyroscope_count"`
		GPSCount        int64  `gorm:"column:gps_count"`
		PhotoCount      int64  `gorm:"column:photo_count"`
		TotalDataPoints int64  `gorm:"column:total_data_points"`
	}

	var devices []DeviceStats
	var total int64

	query := `
		WITH device_stats AS (
			SELECT 
				all_devices.device_id,
				GREATEST(
					COALESCE(MAX(g.timestamp), '1970-01-01'::timestamp),
					COALESCE(MAX(gps.timestamp), '1970-01-01'::timestamp),
					COALESCE(MAX(p.timestamp), '1970-01-01'::timestamp)
				) as last_seen,
				COALESCE(COUNT(DISTINCT g.id), 0) as gyroscope_count,
				COALESCE(COUNT(DISTINCT gps.id), 0) as gps_count,
				COALESCE(COUNT(DISTINCT p.id), 0) as photo_count,
				(COALESCE(COUNT(DISTINCT g.id), 0) + COALESCE(COUNT(DISTINCT gps.id), 0) + COALESCE(COUNT(DISTINCT p.id), 0)) as total_data_points
			FROM (
				SELECT DISTINCT device_id FROM gyroscopes
				UNION
				SELECT DISTINCT device_id FROM gps
				UNION
				SELECT DISTINCT device_id FROM telemetry_photos
			) all_devices
			LEFT JOIN gyroscopes g ON all_devices.device_id = g.device_id
			LEFT JOIN gps ON all_devices.device_id = gps.device_id
			LEFT JOIN telemetry_photos p ON all_devices.device_id = p.device_id
			GROUP BY all_devices.device_id
		)
		SELECT * FROM device_stats
		ORDER BY last_seen DESC
		LIMIT ? OFFSET ?
	`

	countQuery := `
		SELECT COUNT(DISTINCT device_id) as total FROM (
			SELECT device_id FROM gyroscopes
			UNION
			SELECT device_id FROM gps
			UNION
			SELECT device_id FROM telemetry_photos
		) all_devices
	`

	if err := s.db.Raw(countQuery).Scan(&total).Error; err != nil {
		s.logger.Error("failed to count unique devices", zap.Error(err))
		return nil, err
	}

	offset := (page - 1) * limit
	if err := s.db.Raw(query, limit, offset).Scan(&devices).Error; err != nil {
		s.logger.Error("failed to fetch device statistics", zap.Error(err))
		return nil, err
	}

	var deviceResponses []response.DeviceResponse
	for _, d := range devices {
		lastSeen, err := utils.ParseTimestamp(d.LastSeen)
		if err != nil {
			s.logger.Warn("failed to parse last_seen timestamp", zap.String("device_id", d.DeviceID), zap.Error(err))
			lastSeen = utils.GetZeroTime()
		}

		deviceResponses = append(deviceResponses, response.DeviceResponse{
			DeviceID:        d.DeviceID,
			LastSeen:        lastSeen,
			GyroscopeCount:  d.GyroscopeCount,
			GPSCount:        d.GPSCount,
			PhotoCount:      d.PhotoCount,
			TotalDataPoints: d.TotalDataPoints,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &response.DeviceListResponse{
		Data: deviceResponses,
		Pagination: response.PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			Total:       total,
			TotalPages:  totalPages,
		},
	}, nil
}
