package dlidparser

import (
	"time"
)

type DriverSex int

const (
	DriverSexNone DriverSex = iota
	DriverSexMale
	DriverSexFemale
)

type DLIDLicense struct {
	firstName            string
	middleName           string
	lastName             string
	street               string
	city                 string
	state                string
	country              string
	postal               string
	sex                  DriverSex
	socialSecurityNumber string
	dateOfBirth          time.Time
}

func (d *DLIDLicense) SetFirstName(s string) {
	d.firstName = s
}

func (d *DLIDLicense) FirstName() string {
	return d.firstName
}

func (d *DLIDLicense) SetMiddleName(s string) {
	d.middleName = s
}

func (d *DLIDLicense) MiddleName() string {
	return d.middleName
}

func (d *DLIDLicense) SetLastName(s string) {
	d.lastName = s
}

func (d *DLIDLicense) LastName() string {
	return d.lastName
}

func (d *DLIDLicense) SetStreet(s string) {
	d.street = s
}

func (d *DLIDLicense) Street() string {
	return d.street
}

func (d *DLIDLicense) SetCity(s string) {
	d.city = s
}

func (d *DLIDLicense) City() string {
	return d.city
}

func (d *DLIDLicense) SetState(s string) {
	d.state = s
}

func (d *DLIDLicense) State() string {
	return d.state
}

func (d *DLIDLicense) SetCountry(s string) {
	d.country = s
}

func (d *DLIDLicense) Country() string {
	return d.country
}

func (d *DLIDLicense) SetPostal(s string) {
	d.postal = s
}

func (d *DLIDLicense) Postal() string {
	return d.postal
}

func (d *DLIDLicense) SetSex(s DriverSex) {
	d.sex = s
}

func (d *DLIDLicense) Sex() DriverSex {
	return d.sex
}

func (d *DLIDLicense) SetSocialSecurityNumber(s string) {
	d.socialSecurityNumber = s
}

func (d *DLIDLicense) SocialSecurityNumber() string {
	return d.socialSecurityNumber
}

func (d *DLIDLicense) SetDateOfBirth(t time.Time) {
	d.dateOfBirth = t
}

func (d *DLIDLicense) DateOfBirth() time.Time {
	return d.dateOfBirth
}
