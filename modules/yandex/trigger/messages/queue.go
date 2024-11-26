package messages

// QueueAttrs
type QueueAttrs map[string]string

// QueueMessageAttrs
type QueueMessageAttrs map[string]map[string]string

// QueueMessage
type QueueMessage struct {
	MessageID    string            `json:"message_id"`
	Body         string            `json:"body"`
	Attrs        QueueAttrs        `json:"attributes"`
	MessageAttrs QueueMessageAttrs `json:"message_attributes"`
	MD5OfBody    string            `json:"md5_of_body"`
	MD5OfAttrs   string            `json:"md5_of_message_attributes"`
}

// Queue
type Queue struct {
	QueueID string       `json:"queue_id"`
	Message QueueMessage `json:"message"`
}
