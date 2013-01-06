package dlidparser

import (
	"errors"
	"strconv"
	"strings"
)

func Parse(data string) (license *DLIDLicense, err error) {

	// This parser is based on standards from here:
	//
	// http://www.aamva.org/DL-ID-Card-Design-Standard/
	//
	// There are currently 7 standards, and all versions since v1 have used a
	// slightly different header definition.

	// The standard says that the 3rd byte in the header should be 0x1e (record
	// separator) but South Carolina uses 0x1c (file separator) because they're
	// special.

	if !strings.HasPrefix(data, "@\n\x1e\rANSI ") && !(strings.HasPrefix(data, "@\n\x1c\rANSI ")) {
		return license, errors.New("Data does not contain expected header")
	}

	issuer := data[9:15]
	version, err := strconv.Atoi(data[15:17])

	if err != nil {
		return license, errors.New("Data does not contain a version number")
	}

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
	}

	return
}
