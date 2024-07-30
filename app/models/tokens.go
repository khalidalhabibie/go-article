package models

type Tokens struct {
	Access  string `json:"access" groups:"public"`
	Refresh string `json:"refresh" groups:"public"`
}
