package client

const chunkSize = 64

func UploadImageInChunks(imgBytes []byte, sendCallback func(chunk []byte) error) error {
	// Send image
	for currentByte := 0; currentByte < len(imgBytes); currentByte += chunkSize {
		var chunk []byte

		if currentByte+chunkSize > len(imgBytes) {
			chunk = imgBytes[currentByte:] // All
		} else {
			chunk = imgBytes[currentByte : currentByte+chunkSize]
		}

		err := sendCallback(chunk)
		if err != nil {
			return err
		}
	}

	return nil
}
