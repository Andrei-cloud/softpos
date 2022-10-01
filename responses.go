package softpos

type CreateResponse struct {
	Reference string `json:"reference"`
}

type conflict struct {
	Reason string `json:"reason,omitempty"`
	Field  string `json:"field,omitempty"`
	Value  string `json:"value,omitempty"`
	Type   int    `json:"type,omitempty"`
}
