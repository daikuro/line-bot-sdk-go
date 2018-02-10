package linebot

import (
	"encoding/json"
)

const (
	MessageTypeFile MessageType = "file"
	EventTypeFile   EventType   = "file"
)

// FileMessage type
type FileMessage struct {
	ID       string
	FileName string
	FileSize int
}

func (*FileMessage) message() {}

// MarshalJSON method of FileMessage
func (m *FileMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type     MessageType `json:"type"`
		FileName string      `json:"fileName"`
		FileSize int         `json:"fileSize"`
	}{
		Type:     MessageTypeFile,
		FileName: m.FileName,
		FileSize: m.FileSize,
	})
}
