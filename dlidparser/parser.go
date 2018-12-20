package dlidparser

import (
	"errors"
	"strconv"
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

	if len(data) < 15 {
		return nil, errors.New("Data does not contain expected header")
	}

	if data[0:2] != "@\n" ||
		data[3] != '\r' ||
		(data[4:9] != "ANSI " && data[4:9] != "AAMVA") {
		return nil, errors.New("Data does not contain expected header")
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
	default:
		err = errors.New("Unsupported DLID version number")
	}
	return license, err
}
