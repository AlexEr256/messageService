package dto

type MessageRequest struct {
	Creator   string `json:"creator"`
	Recipient string `json:"recipient"`
	Mail      string `json:"mail`
}

type MessageResponse struct {
	Status bool `json:"status"`
}

type Payload struct {
	Before *MessageBody `json:"before"`
	After  *MessageBody `json:"after"`
}

type MessageBody struct {
	Id        int64  `json:"id"`
	Creator   string `json:"creator"`
	Recipient string `json:"recipient"`
	Mail      string `json:"mail`
	CreatedAt int64  `json:"created"`
}

type AggregationResp struct {
	Total int `json:"total"`
}
