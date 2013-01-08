package dlidparser

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func parseV3(data string, issuer string) (license *DLIDLicense, err error) {

	start, end, err := dataRangeV2(data)
	payload := data[start:end]

	if err != nil {
		return
	}

	license, err = parseDataV3(payload, issuer)

	if err != nil {
		return
	}

	return
}

func parseDataV3(licenceData string, issuer string) (license *DLIDLicense, err error) {

	// Version 3 of the DLID card spec was published in 2005.  It is currently
	// (as of 2012) used in Wisconsin.

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

		case "DCG":
			license.SetCountry(data)

		case "DCT":

			// This field contains all of the licencee's names except last
			// name.  The V3 spec doc doesn't specify how the names are
			// separated and doesn't provide an example (unlike the 2000
			// doc).  Wisconsin use a space and Virginia use a comma.  This
			// is why standards have to be adequately documented.

			separator := " "

			if strings.Index(data, separator) == -1 {
				separator = ","
			}

			names := strings.Split(data, separator)

			license.SetFirstName(names[0])

			if len(names) > 1 {
				license.SetMiddleNames(names[1:])
			}

		case "DAG":
			license.SetStreet(data)

		case "DAI":
			license.SetCity(data)

		case "DAJ":
			license.SetState(data)

		case "DAK":
			license.SetPostal(data)

		case "DAQ":
			license.SetCustomerId(data)

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

		// For some reason known only to themselves, the standards guys took
		// the V1 and 2 postal code standards (code padded to 11 characters with
		// spaces) and replaced the spaces with zeros if a) the country is "USA"
		// and b) if the trailing "+4" portion of the postal code is unknown.
		// Quite what happens to pad Canadian postal codes (they are always 6
		// alphanumeric characters, like British postal codes) is undocumented.
		//
		// We will extract the 5-digit zip and the +4 section.  If the +4 is all
		// zeros we can discard it.  The last two digits are always useless (the
		// next version of the spec shortens the field to 9 characters) so we
		// can ignore them.

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

func parseDateV3(data string, country string) time.Time {

	// And now we get the payoff for the previous awful decision to switch to
	// Lumpy Date Format: we're now supporting the international big-endian
	// date format used in Canada and V1 of the spec (yyyyMMdd) and the US
	// lumpy date format *in the same field*.  I can understand that different
	// versions of a standard don't agree with each other, but now we've got two
	// implementations of a standard within a single field in a single version
	// of the standard.  Breathtakingly stupid.

	var day int
	var month int
	var year int
	var err error
	var location *time.Location

	if country == "USA" {
		month, err = strconv.Atoi(data[:2])

		if err != nil {
			return time.Unix(0, 0)
		}

		day, err = strconv.Atoi(data[2:4])

		if err != nil {
			return time.Unix(0, 0)
		}

		year, err = strconv.Atoi(data[4:8])

		if err != nil {
			return time.Unix(0, 0)
		}
	} else {
		year, err = strconv.Atoi(data[:4])

		if err != nil {
			return time.Unix(0, 0)
		}

		month, err = strconv.Atoi(data[4:6])

		if err != nil {
			return time.Unix(0, 0)
		}

		day, err = strconv.Atoi(data[6:8])

		if err != nil {
			return time.Unix(0, 0)
		}
	}

	location, err = time.LoadLocation("UTC")

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location)
}
