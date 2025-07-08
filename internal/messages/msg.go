package messages

import "encoding/json"

type DomainMessageHeader struct {
	MobileNumber       string `json:"user_id"`
	Timestamp          int64  `json:"timestamp"`
	PacketName         string `json:"packet_name"`
	SourceService      string `json:"source_service"`
	DestinationService string `json:"destination_service"`
}

// DomainMessageMessage represents a WebSocket message with routing information
type DomainMessage struct {
	Header  DomainMessageHeader `json:"header"`
	Payload json.RawMessage     `json:"payload"`
}

func NewDomainMessage(
	MobileNumber string,
	PacketName string,
	SourceService string,
	DestinationService string,
	data interface{},
) (*DomainMessage, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return &DomainMessage{
		Header: DomainMessageHeader{
			MobileNumber:       MobileNumber,
			PacketName:         PacketName,
			SourceService:      SourceService,
			DestinationService: DestinationService,
		},
		Payload: dataBytes,
	}, nil
}
