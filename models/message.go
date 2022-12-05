package models

type MessageAttributes struct {
	Src   string
	Prod  string
	Type  string
	Stype string
	Op    string
}

type MessageModel struct {
	Attributes      MessageAttributes
	Payload         string
	SpecificPayload string
}
