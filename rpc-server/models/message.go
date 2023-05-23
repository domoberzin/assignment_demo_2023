package models

type Message struct {
    ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
    Sender     string    `gorm:"size:200" json:"sender"`
    Recipient      string    `gorm:"size:3000" json:"recipient" `
	Text      string    `gorm:"size:3000" json:"text"`
    // CreatedAt time.Time `json:"created_at,omitempty"`
    // UpdatedAt time.Time `json:"updated_at,omitempty"`
}


// TableName method sets table name for Post model
func (message *Message) TableName() string {
    return "message"
}

//ResponseMap -> response map method of Message
func (message *Message) ResponseMap() map[string]interface{} {
    resp := make(map[string]interface{})
    resp["id"] = message.ID
    resp["sender"] = message.Sender
	resp["recipient"] = message.Recipient
    resp["body"] = message.Text
    // resp["created_at"] = message.CreatedAt
    // resp["updated_at"] = message.UpdatedAt
    return resp
}
