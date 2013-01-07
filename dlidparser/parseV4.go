package dlidparser

import (
	"errors"
	"strings"
)

func parseV4(data string, issuer string) (license *DLIDLicense, err error) {

	start, end, err := dataRangeV2(data)
	payload := data[start:end]

	if err != nil {
		return
	}

	license, err = parseDataV4(payload, issuer)

	if err != nil {
		return
	}

	return
}

func parseDataV4(licenceData string, issuer string) (license *DLIDLicense, err error) {

	// Version 4 of the DLID card spec was published in 2009.

	if !strings.HasPrefix(licenceData, "DL") {
		err = errors.New("Missing header in licence data chunk")
		return
	}

	licenceData = licenceData[2:]

	components := strings.Split(licenceData, "\n")

	license = new(DLIDLicense)

	license.SetIssuerId(issuer)
	license.SetIssuerName(issuers[issuer])

	var dateOfBirth string

	for component := range components {

		if len(components[component]) < 3 {
			continue
		}

		identifier := components[component][0:3]
		data := components[component][3:]

		data = strings.Trim(data, " ")

		switch identifier {
		case "DCA":
			license.SetVehicleClass(data)
			
		case "DCB":
			license.SetRestrictionCodes(data)

		case "DCD":
			license.SetEndorsementCodes(data)

		case "DCS":
			license.SetLastName(data)

		case "DAC":
			license.SetFirstName(data)

		case "DAD":

			names := strings.Split(data, ",")

			// We don't care about any other middle names.
			license.SetMiddleName(names[0])

		case "DCG":
			license.SetCountry(data)

		case "DAG":
			license.SetStreet(data)

		case "DAI":
			license.SetCity(data)

		case "DAJ":
			license.SetState(data)

		case "DAK":
			license.SetPostal(data)

		case "DBB":
			dateOfBirth = data

		case "DBC":
			switch data {
			case "1":
				license.SetSex(DriverSexMale)
			case "2":
				license.SetSex(DriverSexFemale)
			default:
				license.SetSex(DriverSexNone)
			}
		}
	}

	// At this point we should know the country and the postal code (both are
	// mandatory fields) so we can undo the desperate mess the standards body
	// made of the postal code field.

	if license.Country() == "USA" && len(license.Postal()) > 0 {

		// Another change to the postal code field!  Surprise!  This time the
		// standards guys trimmed the field down to 9 characters, which makes
		// sense because US zip codes are only 9 digits long.  Canadian post
		// codes are only 6 characters.  Why was the original spec 11 digits?
		// Because the standards guys are *nuts*.
		//
		// We will extract the 5-digit zip and the +4 section.  If the +4 is all
		// zeros we can discard it.

		zip := license.Postal()[:5]
		plus4 := license.Postal()[5:9]

		if plus4 == "0000" {
			license.SetPostal(zip)
		} else {
			license.SetPostal(zip + "+" + plus4)
		}
	}

	// Now we can parse the birth date, too.
	if len(license.Country()) > 0 {
		license.SetDateOfBirth(parseDateV3(dateOfBirth, license.Country()))
	}

	return
}
