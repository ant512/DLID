package dlidparser

import (
	"testing"
)

func TestBadHeader(t *testing.T) {
	_, err := Parse("@\n\x1d\rANSI 636")

	if err == nil {
		t.Error("Bad header did not cause an error")
	} else if err.Error() != "Data does not contain expected header" {
		t.Error("Wrong error message returned")
	}
}

func TestIllegalVersion(t *testing.T) {
	_, err := Parse("@\n\x1e\rANSI 636000080002")

	if err == nil {
		t.Error("Illegal version not detected")
	} else if err.Error() != "Unsupported DLID version number" {
		t.Error("Wrong error message returned")
	}
}

func TestV4Parser(t *testing.T) {
	s, err := Parse("@\n\x1e\rANSI 636000040002DL00410281ZV03190008DLDAQT64235789\nDCSSAMPLE\nDDEN\nDACMICHAEL\nDDFN\nDADJOHN,BOB\nDDGN\nDCUJR\nDCAD\nDCBK\nDCDPH\nDBD06062008\nDBB06071986\nDBA12102012\nDBC1\nDAU068 in\nDAYBRO\nDAG2300 WEST BROAD STREET\nDAIRICHMOND\nDAJVA\nDAK232690000 \nDCF2424244747474786102204\nDCGUSA\nDCK123456789\nDDAM\nDDB06062008\nDDC06062009\nDDD1\rZVZVA01\r")

	if err != nil {
		t.Error("V4 parser failed")
	}

	if s.IssuerName() != "Virginia" {
		t.Error("V4 parser extracted wrong issuer")
	}

	if s.FirstName() != "MICHAEL" {
		t.Error("V4 parser extracted wrong first name")
	}

	if len(s.MiddleNames()) != 2 {
		t.Error("V4 parser failed to extract middle names")
	}

	if s.MiddleNames()[0] != "JOHN" || s.MiddleNames()[1] != "BOB" {
		t.Error("V4 parser extracted wrong middle names")
	}

	if s.LastName() != "SAMPLE" {
		t.Error("V4 parser extracted wrong last name")
	}

	if s.DateOfBirth().Day() != 7 {
		t.Error("V4 parser got wrong date of birth day")
	}

	if s.DateOfBirth().Month() != 6 {
		t.Error("V4 parser got wrong date of birth month")
	}

	if s.DateOfBirth().Year() != 1986 {
		t.Error("V4 parser got wrong date of birth year")
	}

	if s.CustomerId() != "T64235789" {
		t.Error("V4 parser got wrong customer id")
	}

	if s.EndorsementCodes() != "PH" {
		t.Error("V4 parser got wrong endorsement codes")
	}

	if s.VehicleClass() != "D" {
		t.Error("V4 parser got wrong vehicle class")
	}

	if s.RestrictionCodes() != "K" {
		t.Error("V4 parser got wrong restriction codes")
	}

	if s.Country() != "USA" {
		t.Error("V4 parser got wrong country")
	}

	if s.Street() != "2300 WEST BROAD STREET" {
		t.Error("V4 parser got wrong street")
	}

	if s.City() != "RICHMOND" {
		t.Error("V4 parser got wrong city")
	}

	if s.State() != "VA" {
		t.Error("V4 parser got wrong state")
	}

	if s.Postal() != "23269" {
		t.Error("V4 parser got wrong postal code")
	}

	if s.Sex() != DriverSexMale {
		t.Error("V4 parser got wrong sex")
	}
}

func TestV5Parser(t *testing.T) {
	_, err := Parse("@\n\x1e\rANSI 636000050002DL00410278ZV03190008DLDAQT64235789\nDCSSAMPLE\nDDEN\nDACMICHAEL\nDDFN\nDADJOHN\nDDGN\nDCUJR\nDCAD\nDCBK\nDCDPH\nDBD06062008\nDBB06061986\nDBA12102012\nDBC1\nDAU068 in\nDAYBRO\nDAG2300 WEST BROAD STREET\nDAIRICHMOND\nDAJVA\nDAK232690000 \nDCF2424244747474786102204\nDCGUSA\nDCK123456789\nDDAM\nDDB06062008\nDDC06062009\nDDD1\rZVZVA01\r")

	if err != nil {
		t.Error("V5 parser failed")
	}
}

func TestV6Parser(t *testing.T) {
	_, err := Parse("@\n\x1e\rANSI 636000060002DL00410278ZV03190008DLDAQT64235789\nDCSSAMPLE\nDDEN\nDACMICHAEL\nDDFN\nDADJOHN\nDDGN\nDCUJR\nDCAD\nDCBK\nDCDPH\nDBD06062008\nDBB06061986\nDBA12102012\nDBC1\nDAU068 in\nDAYBRO\nDAG2300 WEST BROAD STREET\nDAIRICHMOND\nDAJVA\nDAK232690000 \nDCF2424244747474786102204\nDCGUSA\nDCK123456789\nDDAM\nDDB06062008\nDDC06062009\nDDD1\rZVZVA01\r")

	if err != nil {
		t.Error("V6 parser failed")
	}
}

func TestV7Parser(t *testing.T) {
	_, err := Parse("@\n\x1e\rANSI 636000070002DL00410278ZV03190008DLDAQT64235789\nDCSSAMPLE\nDDEN\nDACMICHAEL\nDDFN\nDADJOHN\nDDGN\nDCUJR\nDCAD\nDCBK\nDCDPH\nDBD06062008\nDBB06061986\nDBA12102012\nDBC1\nDAU068 in\nDAYBRO\nDAG2300 WEST BROAD STREET\nDAIRICHMOND\nDAJVA\nDAK232690000 \nDCF2424244747474786102204\nDCGUSA\nDCK123456789\nDDAM\nDDB06062008\nDDC06062009\nDDD1\rZVZVA01\r")

	if err != nil {
		t.Error("V7 parser failed")
	}
}
