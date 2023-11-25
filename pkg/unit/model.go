package unit

// Info written in file to process
// @Description Unit info
// @Description with number, mqtt, invid, unitguid,
// @Description with messageid, text, context, class,
// @Description with level, area, addr, block,
// @Description with type_, bit, invertbit
type Unit struct {
	Number    string      `json:"number"`    // Unit number
	Mqtt      interface{} `json:"mqtt"`      // MQTT
	Invid     string      `json:"invid"`     // Invid
	UnitGuid  string      `json:"unitguid"`  // Unit GUID
	MessageID string      `json:"messageid"` // Message ID
	Text      string      `json:"text"`      // Text
	Context   interface{} `json:"context"`   // Context
	Class     string      `json:"class"`     // Class
	Level     string      `json:"level"`     // Level
	Area      string      `json:"area"`      // Area
	Addr      string      `json:"addr"`      // Addr
	Block     interface{} `json:"block"`     // Block
	Type_     interface{} `json:"type_"`     // Type
	Bit       interface{} `json:"bit"`       // Bit
	InvertBit interface{} `json:"invertbit"` // Invert bit
}

// Processed file info
// @Description Processed file info
type ProcessedFile struct {
	Name string `json:"filepath" bson:"filepath"` // File path
}
