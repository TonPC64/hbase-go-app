package models

import (
	"encoding/json"
)

const (
	MessageResultTypeCreate  = "CreateMessage"
	MessageResultTypeProcess = "ProcessMessage"
)

// MessageResult is produce to notification
type MessageResult struct {
	Type string            `json:"type"`
	Data MessageResultData `json:"data"`
}

// MessageResultData is Data of result
type MessageResultData struct {
	BucketID string `json:"bucketId"`
	Message
}

// Message is Shema mysql
type Message struct {
	BucketID   string           `json:"bucketId" gorm:"type:varchar(255)"`
	ChannelID  string           `json:"channelId" gorm:"type:varchar(255)" binding:"required"`
	ID         string           `json:"id" gorm:"PRIMARY_KEY;unique;type:varchar(255)" binding:"required"`
	Message    *json.RawMessage `json:"message" gorm:"type:text" binding:"required"`
	Recipient  string           `json:"recipient" gorm:"type:varchar(255)" binding:"required"`
	Sender     string           `json:"sender" gorm:"type:varchar(255)" binding:"required"`
	SenderType string           `json:"senderType" gorm:"type:ENUM('user', 'system', 'admin')"`
	ThreadID   string           `json:"threadId" gorm:"column:thread_id;type:varchar(255)" binding:"required"`
	Timestamp  int64            `json:"timestamp" gorm:"type:bigint(11) unsigned"`
}

// SenderType struct
type SenderType string

// SenderType value for push message tracking
const (
	SenderTypeUser   SenderType = "user"
	SenderTypeSystem SenderType = "system"
	SenderTypeAdmin  SenderType = "admin"
)

func (s SenderType) String() string {
	return string(s)
}
