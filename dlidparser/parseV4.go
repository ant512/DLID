package dlidparser

import (
	"strings"
)

func parseV4(data string, issuer string) (*DLIDLicense, error) {

	start, end, err := dataRangeV2(data)

	/*
		if end > len(data) {
			//lots of states don't count correct - VA i'm looking at you
			end = len(data) - 1
		}
	*/

	//license files don't really contain extra data
	end = len(data) - 1

	if err != nil {
		return nil, err
	}

	payload := data[start:end]
	return parseDataV4(payload, issuer)
}

func parseDataV4(licenceData string, issuer string) (*DLIDLicense, error) {

	// Version 4 of the DLID card spec was published in 2009.

	/*
		This would be nice but encdoing errors - SC 2014 issued - fail this

		if !strings.HasPrefix(licenceData, "DL") && !strings.HasPrefix(licenceData, "ID") {
			return nil, errors.New("Missing header in licence data chunk")
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

		case "DCU":
			license.NameSuffix = data

		case "DAC":
			license.FirstName = data

		case "DAD":
			names := strings.Split(data, ",")
			license.MiddleNames = names

		case "DCG":
			license.Country = data

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

	// At this point we should know the country and the postal code (both are
	// mandatory fields) so we can undo the desperate mess the standards body
	// made of the postal code field.

	if license.Country == "USA" && len(strings.TrimSpace(license.Postal)) == 9 {

		// Another change to the postal code field!  Surprise!  This time the
		// standards guys trimmed the field down to 9 characters, which makes
		// sense because US zip codes are only 9 digits long.  Canadian post
		// codes are only 6 characters.  Why was the original spec 11 digits?
		// Because the standards guys are *nuts*.
		//
		// We will extract the 5-digit zip and the +4 section.  If the +4 is all
		// zeros we can discard it.

		zip := license.Postal[:5]
		plus4 := license.Postal[5:9]

		if plus4 == "0000" {
			license.Postal = zip
		} else {
			license.Postal = zip + "+" + plus4
		}
	}

	if license.IssuerName == "Wyoming" || (license.IssuerName == "West Virginia") {
		license.DateOfBirth = parseDateV3(dateOfBirth, "CANADA")
		license.ExpiryDate = parseDateV3(expiryDate, "CANADA")
		license.IssueDate = parseDateV3(issueDate, "CANADA")
	} else // Now we can parse the dates, too.
	if len(license.Country) > 0 {
		license.DateOfBirth = parseDateV3(dateOfBirth, license.Country)
		license.ExpiryDate = parseDateV3(expiryDate, license.Country)
		license.IssueDate = parseDateV3(issueDate, license.Country)
	}

	return license, nil
}
