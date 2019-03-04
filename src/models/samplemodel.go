package models

// SampleModel for storing name and description
type SampleModel struct {
	ID          int    `json:"ID,sting"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}
