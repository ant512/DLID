package dlidparser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//Parse the data string from a pdf417 driver's license barcode
func Parse(data string) (*DLIDLicense, error) {

	// This parser is based on standards from here:
	//
	// http://www.aamva.org/DL-ID-Card-Design-Standard/
	//
	// There are currently 7 standards, and all versions since v1 have used a
	// slightly different header definition.

	// The standard says that the 3rd byte in the header should be 0x1e (record
	// separator) but South Carolina and Pennsylvania use 0x1c (file separator)
	// because they're special.  We don't even bother checking that byte.

	// PA and CT appear to have used old versions of the spec because they use
	// "AAMVA" instead of "ANSI " as part of the header.

	if len(data) < 17 {
		return nil, errors.New("Data does not contain expected header - ")
	}

	if strings.HasPrefix(data, "@\n\u001e\rANSI6") {
		data = strings.Replace(data, "ANSI6", "ANSI 6", 1)
	}
	//OREGON / AZ
	if data[0:8] == "@\r\nANSI " || data[0:8] == "@\n\rANSI " /* for AZ circa 2009 licenses */ {
		data = "@\n\u001e\rANSI " + data[8:]
	}

	if data[0:2] != "@\n" ||
		data[3] != '\r' ||
		(data[4:9] != "ANSI " && data[4:9] != "AAMVA") {
		return nil, fmt.Errorf("Data does not contain expected header %v", string(data[4:9]))
	}

	issuer := data[9:15]
	version, err := strconv.Atoi(data[15:17])

	if err != nil {
		return nil, errors.New("Data does not contain a version number")
	}

	var license *DLIDLicense
	switch version {
	case 1:
		license, err = parseV1(data, issuer)
	case 2:
		license, err = parseV2(data, issuer)
	case 3:
		license, err = parseV3(data, issuer)
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		fallthrough
	case 7:
		license, err = parseV4(data, issuer)
	case 8:
		license, err = parseV4(data, issuer)
	case 9:
		//aamva 09 is 2016 the latest spec - compat w/ the v4
		license, err = parseV4(data, issuer)
	default:
		err = errors.New("Unsupported DLID version number")
	}
	return license, err
}
