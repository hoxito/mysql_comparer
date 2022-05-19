package models

type Difference struct {
	Master      string `json:"maste,omitempty"`
	Slave       string `json:"slave,omitempty" validate:"required"`
	Differences string `json:"differences,omitempty" validate:"required"`
}
