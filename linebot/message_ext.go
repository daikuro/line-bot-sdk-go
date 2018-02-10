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

// ============
// event.go
// ============
// この関数の中
//    func (e *Event) MarshalJSON() ([]byte, error) {
//case *FileMessage:
//raw.Message = &rawEventMessage{
//Type:     MessageTypeFile,
//ID:       m.ID,
//FileName: m.FileName,
//FileSize: m.FileSize,
//}
//

//type rawEventMessage struct {
//	ID        string      `json:"id"`
//	Type      MessageType `json:"type"`
//	Text      string      `json:"text,omitempty"`
//	Duration  int         `json:"duration,omitempty"`
//	Title     string      `json:"title,omitempty"`
//	Address   string      `json:"address,omitempty"`
//	Latitude  float64     `json:"latitude,omitempty"`
//	Longitude float64     `json:"longitude,omitempty"`
//	PackageID string      `json:"packageId,omitempty"`
//	StickerID string      `json:"stickerId,omitempty"`
//	FileName  string      `json:"fileName,omitempty"`
//	FileSize  int         `json:"fileSize,omitempty"`
//}

//この関数の中
// func (e *Event) UnmarshalJSON(body []byte) (err error) {
//
//case MessageTypeFile:
//e.Message = &FileMessage{
//ID:       rawEvent.Message.ID,
//FileName: rawEvent.Message.FileName,
//FileSize: rawEvent.Message.FileSize,
//}
