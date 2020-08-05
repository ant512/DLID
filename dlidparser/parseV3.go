package dlidparser

import (
	"strings"
	"time"
)

func parseV3(data string, issuer string) (*DLIDLicense, error) {

	start, end, err := dataRangeV2(data)
	if err != nil {
		return nil, err
	}

	if end >= len(data) {
		//lots of states don't count correct - VA i'm looking at you
		end = len(data) - 1
	}

	payload := data[start:end]

	license, err := parseDataV3(payload, issuer)

	if err != nil {
		return nil, err
	}

	return license, nil
}

func parseDataV3(licenceData string, issuer string) (*DLIDLicense, error) {

	// Version 3 of the DLID card spec was published in 2005.  It is currently
	// (as of 2012) used in Wisconsin.

	/*
		if !strings.HasPrefix(licenceData, "DL") {
			err := errors.New("Missing header in licence data chunk")
			return nil, err
		}
	*/

	licenceData = licenceData[2:]

	components := strings.Split(licenceData, "\n")

	license := &DLIDLicense{}

	license.IssuerID = issuer
	license.IssuerName = issuers[issuer]

	var dateOfBirth string
	var expiryDate string
	var issueDate string

	for component := range components {

		if len(components[component]) < 3 {
			continue
		}

		identifier := components[component][0:3]
		data := components[component][3:]

		data = strings.Trim(data, " ")

		switch identifier {
		case "DCA":
			license.VehicleClass = data
		case "DCB":
			license.RestrictionCodes = data
		case "DCD":
			license.EndorsementCodes = data
		case "DCS":
			license.LastName = data
		case "DCG":
			license.Country = data
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

			license.FirstName = names[0]

			if len(names) > 1 {
				license.MiddleNames = names[1:]
			}

		case "DAG":
			license.Street = data

		case "DAI":
			license.City = data

		case "DAJ":
			license.State = data

		case "DAK":
			license.Postal = data

		case "DAQ":
			license.CustomerID = data
		case "DBA":
			expiryDate = data
		case "DBB":
			dateOfBirth = data
		case "DBC":
			switch data {
			case "1":
				license.Sex = DriverSexMale
			case "2":
				license.Sex = DriverSexFemale
			default:
				license.Sex = DriverSexNone
			}

		case "DBD":
			issueDate = data
		}
	}

	//if empty default to USA - Michigan - doesn't set DCG - based on license issue 1.20.2017
	//without country - doesn't parse dates
	if license.Country == "" {
		license.Country = "USA"
	}

	// At this point we should know the country and the postal code (both are
	// mandatory fields) so we can undo the desperate mess the standards body
	// made of the postal code field.

	if strings.Contains(license.Country, "USA") && len(strings.TrimSpace(license.Postal)) == 9 {

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

		// Naturally, some Texas licences ignore the spec and just use 5
		// characters if they don't have a +4 section.

		zip := license.Postal[:5]
		plus4 := license.Postal[5:9]

		if plus4 == "0000" {
			license.Postal = zip
		} else {
			license.Postal = zip + "+" + plus4
		}
	}

	// Now we can parse the birth date, too.
	if len(license.Country) > 0 {
		license.DateOfBirth = parseDateV3(dateOfBirth, license.Country)
		license.ExpiryDate = parseDateV3(expiryDate, license.Country)
		license.IssueDate = parseDateV3(issueDate, license.Country)
	}
	return license, nil
}

func parseDateV3(data string, country string) time.Time {

	// And now we get the payoff for the previous awful decision to switch to
	// Lumpy Date Format: we're now supporting the international big-endian
	// date format used in Canada and V1 of the spec (yyyyMMdd) and the US
	// lumpy date format *in the same field*.  I can understand that different
	// versions of a standard don't agree with each other, but now we've got two
	// implementations of a standard within a single field in a single version
	// of the standard.  Breathtakingly stupid.

	if len(data) != 8 {
		return time.Unix(0, 0)
	}
	order := []string{"20060102", "01022006"}
	if strings.Contains(country, "USA") {
		order = []string{"01022006", "20060102"}
	}
	for _, format := range order {
		t, err := time.Parse(format, data)
		if err == nil {
			return t
		}
	}
	return time.Unix(0, 0)
}
