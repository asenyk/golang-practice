package models

type AddressPostals struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	GUID       string `json:"guid"`
	StreetName string `json:"street_name"`
}
