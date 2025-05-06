package providers

type recognitionProviderInMemory struct{}

func NewRecognitionProviderInMemory() RecognitionProvider {
	return &recognitionProviderInMemory{}
}

func (r *recognitionProviderInMemory) CompareFaces(sourceImage, targetImage []byte) (bool, error) {
	return string(sourceImage) == string(targetImage), nil
}
