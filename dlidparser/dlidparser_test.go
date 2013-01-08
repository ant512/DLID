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

func TestV1ColoradoParser(t *testing.T) {
	s, err := Parse("@\n\x1e\rANSI 6360200102DL00390187ZV02260031DLDAQ0123456789ABC\nDAAJOHN,Q,PUBLIC\nDAG123 MAIN STREET\nDAIANYTOWN\nDAJVA\nDAK123459999  \nDARDM  \nDAS          \nDAT     \nDAU509\nDAW175\nDAYBL \nDAZBR \nDBA20011201\nDBB19761123\nDBCM\nDBD19961201\rZVZVAJURISDICTIONDEFINEDELEMENT\r")

	if err != nil {
		t.Error("V1 Colorado parser failed")
	}

	if s.FirstName() != "JOHN" {
		t.Error("V1 Colorado parser extracted wrong first name")
	}

	if s.MiddleNames()[0] != "Q" {
		t.Error("V1 Colorado parser extracted wrong middle name")
	}

	if s.LastName() != "PUBLIC" {
		t.Error("V1 Colorado parser extracted wrong last name")
	}
}

func TestV1Parser(t *testing.T) {
	s, err := Parse("@\n\x1e\rANSI 6360000102DL00390187ZV02260031DLDAQ0123456789ABC\nDAAPUBLIC,JOHN,Q\nDAG123 MAIN STREET\nDAIANYTOWN\nDAJVA\nDAK123459999  \nDARDM  \nDAS          \nDAT     \nDAU509\nDAW175\nDAYBL \nDAZBR \nDBA20011201\nDBB19761123\nDBCM\nDBD19961201\rZVZVAJURISDICTIONDEFINEDELEMENT\r")

	if err != nil {
		t.Error("V1 parser failed")
	}

	if s.IssuerName() != "Virginia" {
		t.Error("V1 parser extracted wrong issuer")
	}

	if s.FirstName() != "JOHN" {
		t.Error("V1 parser extracted wrong first name")
	}

	if s.MiddleNames()[0] != "Q" {
		t.Error("V1 parser extracted wrong middle name")
	}

	if s.LastName() != "PUBLIC" {
		t.Error("V1 parser extracted wrong last name")
	}

	if s.Country() != "USA" {
		t.Error("V1 parser got wrong country")
	}

	if s.Street() != "123 MAIN STREET" {
		t.Error("V1 parser got wrong street")
	}

	if s.City() != "ANYTOWN" {
		t.Error("V1 parser got wrong city")
	}

	if s.State() != "VA" {
		t.Error("V1 parser got wrong state")
	}

	if s.Postal() != "123459999" {
		t.Error("V1 parser got wrong postal code")
	}

	if s.DateOfBirth().Day() != 23 {
		t.Error("V1 parser got wrong date of birth day")
	}

	if s.DateOfBirth().Month() != 11 {
		t.Error("V1 parser got wrong date of birth month")
	}

	if s.DateOfBirth().Year() != 1976 {
		t.Error("V1 parser got wrong date of birth year")
	}

	if s.CustomerId() != "0123456789ABC" {
		t.Error("V1 parser got wrong customer id")
	}

	if s.EndorsementCodes() != "" {
		t.Error("V1 parser got wrong endorsement codes")
	}

	if s.VehicleClass() != "DM" {
		t.Error("V1 parser got wrong vehicle class")
	}

	if s.RestrictionCodes() != "" {
		t.Error("V1 parser got wrong restriction codes")
	}

	if s.Sex() != DriverSexMale {
		t.Error("V1 parser got wrong sex")
	}

	if s.ExpiryDate().Day() != 1 {
		t.Error("V1 parser got wrong expiry day")
	}

	if s.ExpiryDate().Month() != 12 {
		t.Error("V1 parser got wrong expiry month")
	}

	if s.ExpiryDate().Year() != 2001 {
		t.Error("V1 parser got wrong expiry year")
	}

	if s.IssueDate().Day() != 1 {
		t.Error("V1 parser got wrong issue day")
	}

	if s.IssueDate().Month() != 12 {
		t.Error("V1 parser got wrong issue month")
	}

	if s.IssueDate().Year() != 1996 {
		t.Error("V1 parser got wrong issue year")
	}
}

func TestV2Parser(t *testing.T) {
	
}

func TestV3Parser(t *testing.T) {
	s, err := Parse("@\n\x1e\rANSI 636015030002DL00410217ZT02020022DLDCAB\nDCBLP\nDCDP\nDBA04052018\nDCSJONES\nDCTJAMES ROBERT R\nDBD07082012\nDBB10111978\nDBC1\nDAYBRO\nDAU 70 IN\nDAG123 SOME STREET\nDAICITY 12\nDAJTX\nDAK902100000  \nDAQ22334455\nDCF11111111111111111111\nDCGUSA\nDCHB   \nDAZBRO\nDCU\rZTZTA220\nZTBW\n")

	if err != nil {
		t.Error("V3 parser failed")
	}

	if s.IssuerName() != "Texas" {
		t.Error("V3 parser extracted wrong issuer")
	}

	if s.FirstName() != "JAMES" {
		t.Error("V3 parser extracted wrong first name")
	}

	if len(s.MiddleNames()) != 2 {
		t.Error("V3 parser failed to extract middle names")
	}

	if s.MiddleNames()[0] != "ROBERT" || s.MiddleNames()[1] != "R" {
		t.Error("V3 parser extracted wrong middle names")
	}

	if s.LastName() != "JONES" {
		t.Error("V3 parser extracted wrong last name")
	}

	if s.DateOfBirth().Day() != 11 {
		t.Error("V3 parser got wrong date of birth day")
	}

	if s.DateOfBirth().Month() != 10 {
		t.Error("V3 parser got wrong date of birth month")
	}

	if s.DateOfBirth().Year() != 1978 {
		t.Error("V3 parser got wrong date of birth year")
	}

	if s.CustomerId() != "22334455" {
		t.Error("V3 parser got wrong customer id")
	}

	if s.EndorsementCodes() != "P" {
		t.Error("V3 parser got wrong endorsement codes")
	}

	if s.VehicleClass() != "B" {
		t.Error("V3 parser got wrong vehicle class")
	}

	if s.RestrictionCodes() != "LP" {
		t.Error("V3 parser got wrong restriction codes")
	}

	if s.Country() != "USA" {
		t.Error("V3 parser got wrong country")
	}

	if s.Street() != "123 SOME STREET" {
		t.Error("V3 parser got wrong street")
	}

	if s.City() != "CITY 12" {
		t.Error("V3 parser got wrong city")
	}

	if s.State() != "TX" {
		t.Error("V3 parser got wrong state")
	}

	if s.Postal() != "90210" {
		t.Error("V3 parser got wrong postal code")
	}

	if s.Sex() != DriverSexMale {
		t.Error("V3 parser got wrong sex")
	}

	if s.ExpiryDate().Day() != 5 {
		t.Error("V4 parser got wrong expiry day")
	}

	if s.ExpiryDate().Month() != 4 {
		t.Error("V3 parser got wrong expiry month")
	}

	if s.ExpiryDate().Year() != 2018 {
		t.Error("V3 parser got wrong expiry year")
	}

	if s.IssueDate().Day() != 8 {
		t.Error("V3 parser got wrong issue day")
	}

	if s.IssueDate().Month() != 7 {
		t.Error("V3 parser got wrong issue month")
	}

	if s.IssueDate().Year() != 2012 {
		t.Error("V3 parser got wrong issue year")
	}
}

func TestV4Parser(t *testing.T) {
	s, err := Parse("@\n\x1e\rANSI 636000040002DL00410282ZV03190008DLDAQT64235789\nDCSSAMPLE\nDDEN\nDACMICHAEL\nDDFN\nDADJOHN,BOB\nDDGN\nDCUJR\nDCAD\nDCBK\nDCDPH\nDBD06062008\nDBB06071986\nDBA12102012\nDBC1\nDAU068 in\nDAYBRO\nDAG2300 WEST BROAD STREET\nDAIRICHMOND\nDAJVA\nDAK232690000 \nDCF2424244747474786102204\nDCGUSA\nDCK123456789\nDDAM\nDDB06062008\nDDC06062009\nDDD1\rZVZVA01\r")

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

	if s.NameSuffix() != "JR" {
		t.Error("V4 parser extracted wrong name suffix")
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

	if s.ExpiryDate().Day() != 10 {
		t.Error("V4 parser got wrong expiry day")
	}

	if s.ExpiryDate().Month() != 12 {
		t.Error("V4 parser got wrong expiry month")
	}

	if s.ExpiryDate().Year() != 2012 {
		t.Error("V4 parser got wrong expiry year")
	}

	if s.IssueDate().Day() != 6 {
		t.Error("V4 parser got wrong issue day")
	}

	if s.IssueDate().Month() != 6 {
		t.Error("V4 parser got wrong issue month")
	}

	if s.IssueDate().Year() != 2008 {
		t.Error("V4 parser got wrong issue year")
	}
}

func TestV5Parser(t *testing.T) {
	s, err := Parse("@\n\x1e\rANSI 636000050002DL00410282ZV03190008DLDAQT64235789\nDCSSAMPLE\nDDEN\nDACMICHAEL\nDDFN\nDADJOHN,BOB\nDDGN\nDCUJR\nDCAD\nDCBK\nDCDPH\nDBD06062008\nDBB06071986\nDBA12102012\nDBC1\nDAU068 in\nDAYBRO\nDAG2300 WEST BROAD STREET\nDAIRICHMOND\nDAJVA\nDAK232690000 \nDCF2424244747474786102204\nDCGUSA\nDCK123456789\nDDAM\nDDB06062008\nDDC06062009\nDDD1\rZVZVA01\r")

	if err != nil {
		t.Error("V5 parser failed")
	}

	if s.IssuerName() != "Virginia" {
		t.Error("V5 parser extracted wrong issuer")
	}

	if s.FirstName() != "MICHAEL" {
		t.Error("V5 parser extracted wrong first name")
	}

	if len(s.MiddleNames()) != 2 {
		t.Error("V5 parser failed to extract middle names")
	}

	if s.MiddleNames()[0] != "JOHN" || s.MiddleNames()[1] != "BOB" {
		t.Error("V5 parser extracted wrong middle names")
	}

	if s.LastName() != "SAMPLE" {
		t.Error("V5 parser extracted wrong last name")
	}

	if s.NameSuffix() != "JR" {
		t.Error("V4 parser extracted wrong name suffix")
	}

	if s.DateOfBirth().Day() != 7 {
		t.Error("V5 parser got wrong date of birth day")
	}

	if s.DateOfBirth().Month() != 6 {
		t.Error("V5 parser got wrong date of birth month")
	}

	if s.DateOfBirth().Year() != 1986 {
		t.Error("V5 parser got wrong date of birth year")
	}

	if s.CustomerId() != "T64235789" {
		t.Error("V5 parser got wrong customer id")
	}

	if s.EndorsementCodes() != "PH" {
		t.Error("V5 parser got wrong endorsement codes")
	}

	if s.VehicleClass() != "D" {
		t.Error("V5 parser got wrong vehicle class")
	}

	if s.RestrictionCodes() != "K" {
		t.Error("V5 parser got wrong restriction codes")
	}

	if s.Country() != "USA" {
		t.Error("V5 parser got wrong country")
	}

	if s.Street() != "2300 WEST BROAD STREET" {
		t.Error("V5 parser got wrong street")
	}

	if s.City() != "RICHMOND" {
		t.Error("V5 parser got wrong city")
	}

	if s.State() != "VA" {
		t.Error("V5 parser got wrong state")
	}

	if s.Postal() != "23269" {
		t.Error("V5 parser got wrong postal code")
	}

	if s.Sex() != DriverSexMale {
		t.Error("V5 parser got wrong sex")
	}

	if s.ExpiryDate().Day() != 10 {
		t.Error("5V5 parser got wrong expiry day")
	}

	if s.ExpiryDate().Month() != 12 {
		t.Error("V5 parser got wrong expiry month")
	}

	if s.ExpiryDate().Year() != 2012 {
		t.Error("V5 parser got wrong expiry year")
	}

	if s.IssueDate().Day() != 6 {
		t.Error("V5 parser got wrong issue day")
	}

	if s.IssueDate().Month() != 6 {
		t.Error("V5 parser got wrong issue month")
	}

	if s.IssueDate().Year() != 2008 {
		t.Error("V5 parser got wrong issue year")
	}
}

func TestV6Parser(t *testing.T) {
	s, err := Parse("@\n\x1e\rANSI 636000060002DL00410282ZV03190008DLDAQT64235789\nDCSSAMPLE\nDDEN\nDACMICHAEL\nDDFN\nDADJOHN,BOB\nDDGN\nDCUJR\nDCAD\nDCBK\nDCDPH\nDBD06062008\nDBB06071986\nDBA12102012\nDBC1\nDAU068 in\nDAYBRO\nDAG2300 WEST BROAD STREET\nDAIRICHMOND\nDAJVA\nDAK232690000 \nDCF2424244747474786102204\nDCGUSA\nDCK123456789\nDDAM\nDDB06062008\nDDC06062009\nDDD1\rZVZVA01\r")

	if err != nil {
		t.Error("V6 parser failed")
	}

	if s.IssuerName() != "Virginia" {
		t.Error("V6 parser extracted wrong issuer")
	}

	if s.FirstName() != "MICHAEL" {
		t.Error("V6 parser extracted wrong first name")
	}

	if len(s.MiddleNames()) != 2 {
		t.Error("V6 parser failed to extract middle names")
	}

	if s.MiddleNames()[0] != "JOHN" || s.MiddleNames()[1] != "BOB" {
		t.Error("V6 parser extracted wrong middle names")
	}

	if s.LastName() != "SAMPLE" {
		t.Error("V6 parser extracted wrong last name")
	}

	if s.NameSuffix() != "JR" {
		t.Error("V4 parser extracted wrong name suffix")
	}

	if s.DateOfBirth().Day() != 7 {
		t.Error("V6 parser got wrong date of birth day")
	}

	if s.DateOfBirth().Month() != 6 {
		t.Error("V6 parser got wrong date of birth month")
	}

	if s.DateOfBirth().Year() != 1986 {
		t.Error("V6 parser got wrong date of birth year")
	}

	if s.CustomerId() != "T64235789" {
		t.Error("V6 parser got wrong customer id")
	}

	if s.EndorsementCodes() != "PH" {
		t.Error("V6 parser got wrong endorsement codes")
	}

	if s.VehicleClass() != "D" {
		t.Error("V6 parser got wrong vehicle class")
	}

	if s.RestrictionCodes() != "K" {
		t.Error("V6 parser got wrong restriction codes")
	}

	if s.Country() != "USA" {
		t.Error("V6 parser got wrong country")
	}

	if s.Street() != "2300 WEST BROAD STREET" {
		t.Error("V6 parser got wrong street")
	}

	if s.City() != "RICHMOND" {
		t.Error("V6 parser got wrong city")
	}

	if s.State() != "VA" {
		t.Error("V6 parser got wrong state")
	}

	if s.Postal() != "23269" {
		t.Error("V6 parser got wrong postal code")
	}

	if s.Sex() != DriverSexMale {
		t.Error("V6 parser got wrong sex")
	}

	if s.ExpiryDate().Day() != 10 {
		t.Error("V6 parser got wrong expiry day")
	}

	if s.ExpiryDate().Month() != 12 {
		t.Error("V6 parser got wrong expiry month")
	}

	if s.ExpiryDate().Year() != 2012 {
		t.Error("V6 parser got wrong expiry year")
	}

	if s.IssueDate().Day() != 6 {
		t.Error("V6 parser got wrong issue day")
	}

	if s.IssueDate().Month() != 6 {
		t.Error("V6 parser got wrong issue month")
	}

	if s.IssueDate().Year() != 2008 {
		t.Error("V6 parser got wrong issue year")
	}
}

func TestV7Parser(t *testing.T) {
	s, err := Parse("@\n\x1e\rANSI 636000070002DL00410282ZV03190008DLDAQT64235789\nDCSSAMPLE\nDDEN\nDACMICHAEL\nDDFN\nDADJOHN,BOB\nDDGN\nDCUJR\nDCAD\nDCBK\nDCDPH\nDBD06062008\nDBB06071986\nDBA12102012\nDBC1\nDAU068 in\nDAYBRO\nDAG2300 WEST BROAD STREET\nDAIRICHMOND\nDAJVA\nDAK232690000 \nDCF2424244747474786102204\nDCGUSA\nDCK123456789\nDDAM\nDDB06062008\nDDC06062009\nDDD1\rZVZVA01\r")

	if err != nil {
		t.Error("V7 parser failed")
	}

	if s.IssuerName() != "Virginia" {
		t.Error("V7 parser extracted wrong issuer")
	}

	if s.FirstName() != "MICHAEL" {
		t.Error("V7 parser extracted wrong first name")
	}

	if len(s.MiddleNames()) != 2 {
		t.Error("V7 parser failed to extract middle names")
	}

	if s.MiddleNames()[0] != "JOHN" || s.MiddleNames()[1] != "BOB" {
		t.Error("V7 parser extracted wrong middle names")
	}

	if s.LastName() != "SAMPLE" {
		t.Error("V7 parser extracted wrong last name")
	}

	if s.NameSuffix() != "JR" {
		t.Error("V4 parser extracted wrong name suffix")
	}

	if s.DateOfBirth().Day() != 7 {
		t.Error("V7 parser got wrong date of birth day")
	}

	if s.DateOfBirth().Month() != 6 {
		t.Error("V7 parser got wrong date of birth month")
	}

	if s.DateOfBirth().Year() != 1986 {
		t.Error("V7 parser got wrong date of birth year")
	}

	if s.CustomerId() != "T64235789" {
		t.Error("V7 parser got wrong customer id")
	}

	if s.EndorsementCodes() != "PH" {
		t.Error("V7 parser got wrong endorsement codes")
	}

	if s.VehicleClass() != "D" {
		t.Error("V7 parser got wrong vehicle class")
	}

	if s.RestrictionCodes() != "K" {
		t.Error("V7 parser got wrong restriction codes")
	}

	if s.Country() != "USA" {
		t.Error("V7 parser got wrong country")
	}

	if s.Street() != "2300 WEST BROAD STREET" {
		t.Error("V7 parser got wrong street")
	}

	if s.City() != "RICHMOND" {
		t.Error("V7 parser got wrong city")
	}

	if s.State() != "VA" {
		t.Error("V7 parser got wrong state")
	}

	if s.Postal() != "23269" {
		t.Error("V7 parser got wrong postal code")
	}

	if s.Sex() != DriverSexMale {
		t.Error("V7 parser got wrong sex")
	}

	if s.ExpiryDate().Day() != 10 {
		t.Error("V7 parser got wrong expiry day")
	}

	if s.ExpiryDate().Month() != 12 {
		t.Error("V7 parser got wrong expiry month")
	}

	if s.ExpiryDate().Year() != 2012 {
		t.Error("V7 parser got wrong expiry year")
	}

	if s.IssueDate().Day() != 6 {
		t.Error("V7 parser got wrong issue day")
	}

	if s.IssueDate().Month() != 6 {
		t.Error("V7 parser got wrong issue month")
	}

	if s.IssueDate().Year() != 2008 {
		t.Error("V7 parser got wrong issue year")
	}
}

func TestV7CanadaParser(t *testing.T) {
	s, err := Parse("@\n\x1e\rANSI 636000070002DL00410282ZV03190008DLDAQT64235789\nDCSSAMPLE\nDDEN\nDACMICHAEL\nDDFN\nDADJOHN,BOB\nDDGN\nDCUJR\nDCAD\nDCBK\nDCDPH\nDBD20080606\nDBB19860607\nDBA20121210\nDBC1\nDAU068 in\nDAYBRO\nDAG2300 WEST BROAD STREET\nDAIRICHMOND\nDAJVA\nDAK232690000 \nDCF2424244747474786102204\nDCGCAN\nDCK123456789\nDDAM\nDDB20080606\nDDC20090606\nDDD1\rZVZVA01\r")

	if err != nil {
		t.Error("V7 Canada parser failed")
	}

	if s.DateOfBirth().Day() != 7 {
		t.Error("V7 Canada parser got wrong date of birth day")
	}

	if s.DateOfBirth().Month() != 6 {
		t.Error("V7 Canada parser got wrong date of birth month")
	}

	if s.DateOfBirth().Year() != 1986 {
		t.Error("V7 Canada parser got wrong date of birth year")
	}

	if s.ExpiryDate().Day() != 10 {
		t.Error("V7 Canada parser got wrong expiry day")
	}

	if s.ExpiryDate().Month() != 12 {
		t.Error("V7 Canada parser got wrong expiry month")
	}

	if s.ExpiryDate().Year() != 2012 {
		t.Error("V7 Canada parser got wrong expiry year")
	}

	if s.IssueDate().Day() != 6 {
		t.Error("V7 Canada parser got wrong issue day")
	}

	if s.IssueDate().Month() != 6 {
		t.Error("V7 Canada parser got wrong issue month")
	}

	if s.IssueDate().Year() != 2008 {
		t.Error("V7 Canada parser got wrong issue year")
	}
}
