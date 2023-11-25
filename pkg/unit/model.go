package unit

type Unit struct {
	Number    string      `json:"number"`
	Mqtt      interface{} `json:"mqtt"`
	Invid     string      `json:"invid"`
	UnitGuid  string      `json:"unitguid"`
	MessageID string      `json:"messageid"`
	Text      string      `json:"text"`
	Context   interface{} `json:"context"`
	Class     string      `json:"class"`
	Level     string      `json:"level"`
	Area      string      `json:"area"`
	Addr      string      `json:"addr"`
	Block     interface{} `json:"block"`
	Type_     interface{} `json:"type_"`
	Bit       interface{} `json:"bit"`
	InvertBit interface{} `json:"invertbit"`
}

type ProcessedFile struct {
	Name string `json:"filepath" bson:"filepath"`
}
