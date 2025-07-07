package messages

import "encoding/json"

type DomainMessageHeader struct {
	UserID             string `json:"user_id"`
	Timestamp          int64  `json:"timestamp"`
	PacketName         string `json:"packet_name"`
	SourceService      string `json:"source_service"`
	InternalInitiator  string `json:"internal_initiator,omitempty"`
	DestinationService string `json:"destination_service"`
}

// DomainMessageMessage represents a WebSocket message with routing information
type DomainMessage struct {
	Header  DomainMessageHeader `json:"header"`
	Payload json.RawMessage     `json:"payload"`
}
