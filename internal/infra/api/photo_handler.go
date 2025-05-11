package api

// type PhotoHandlers struct {
// 	CreatePhotoUseCase *usecase.CreatePhotoUseCase
// }

// func NewPhotoHandlers(createPhotoUseCase *usecase.CreatePhotoUseCase) *PhotoHandlers {
// 	return &PhotoHandlers{
// 		CreatePhotoUseCase: createPhotoUseCase,
// 	}
// }

// func (h *PhotoHandlers) CreatePhotoHandler(c *gin.Context) {
// 	var input domain.PhotoDto
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrMissingGPSInvalidFields})
// 		return
// 	}

// 	photo, err := h.CreatePhotoUseCase.Execute(input)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"photo":      photo,
// 		"recognized": photo.Recognized,
// 	})
// }
