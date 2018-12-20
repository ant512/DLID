package dlidparser

import (
	"encoding/json"
	"time"
)

//DriverSex - states record a sex
type DriverSex int

//Possibilities for different sex
const (
	DriverSexNone DriverSex = iota
	DriverSexMale
	DriverSexFemale
)

//DLIDLicense holds extracted driver's license info
type DLIDLicense struct {
	FirstName             string    `json:"first_name,omitempty"`
	MiddleNames           []string  `json:"middle_names,omitempty"`
	LastName              string    `json:"last_name,omitempty"`
	NameSuffix            string    `json:"name_suffix,omitempty"`
	Street                string    `json:"street,omitempty"`
	City                  string    `json:"city,omitempty"`
	State                 string    `json:"state,omitempty"`
	Country               string    `json:"country,omitempty"`
	Postal                string    `json:"postal,omitempty"`
	Sex                   DriverSex `json:"sex,omitempty"`
	SocialSecurityNumber  string    `json:"social_security_number,omitempty"`
	DateOfBirth           time.Time `json:"date_of_birth,omitempty"`
	IssuerID              string    `json:"issuer_id,omitempty"`
	IssuerName            string    `json:"issuer_name,omitempty"`
	ExpiryDate            time.Time `json:"expiry_date,omitempty"`
	IssueDate             time.Time `json:"issue_date,omitempty"`
	VehicleClass          string    `json:"vehicle_class,omitempty"`
	RestrictionCodes      string    `json:"restriction_codes,omitempty"`
	EndorsementCodes      string    `json:"endorsement_codes,omitempty"`
	CustomerID            string    `json:"customer_id,omitempty"`
	DocumentDiscriminator string    `json:"document_discriminator,omitempty"`
}

func (d DLIDLicense) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}
