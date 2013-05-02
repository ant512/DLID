package dlidparser

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func parseV2(data string, issuer string) (license *DLIDLicense, err error) {

	start, end, err := dataRangeV2(data)

	if end >= len(data) {
		err = errors.New("Payload location does not exist in data")
	}

	payload := data[start:end]

	if err != nil {
		return
	}

	license, err = parseDataV2(payload, issuer)

	if err != nil {
		return
	}

	return
}

func dataRangeV2(data string) (start int, end int, err error) {

	start, err = strconv.Atoi(data[23:27])

	if err != nil {
		err = errors.New("Data contains malformed payload location")
		return
	}

	end, err = strconv.Atoi(data[27:31])

	if err != nil {
		err = errors.New("Data contains malformed payload length")
		return
	}

	end += start

	return
}

func parseDataV2(licenceData string, issuer string) (license *DLIDLicense, err error) {

	// Version 1 of the DLID card spec was published in 2003.

	if !strings.HasPrefix(licenceData, "DL") {
		err = errors.New("Missing header in licence data chunk")
		return
	}

	licenceData = licenceData[2:]

	components := strings.Split(licenceData, "\n")

	license = new(DLIDLicense)

	license.SetIssuerId(issuer)
	license.SetIssuerName(issuers[issuer])

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

		case "DCT":

			// This field contains all of the licencee's names except last
			// name.  The V2 spec doc doesn't specify how the names are
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
			license.SetDateOfBirth(parseDateV2(data))

		case "DBC":

			// According to the spec, the standard dropped M/F and two of
			// the ANSI D-20 gender codes in this revision.  The only
			// permissible values are now "1" and "2".

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

	return
}

func parseDateV2(data string) time.Time {

	// Sooo, let me get this straight.  They switched from a reasonably-standard
	// and universal date format (yyyyMMdd) to the bizarre US lumpy format
	// (MMddyyyy)?  What were they thinking!?

	month, err := strconv.Atoi(data[:2])

	if err != nil {
		return time.Unix(0, 0)
	}

	day, err := strconv.Atoi(data[2:4])

	if err != nil {
		return time.Unix(0, 0)
	}

	year, err := strconv.Atoi(data[4:8])

	if err != nil {
		return time.Unix(0, 0)
	}

	location, err := time.LoadLocation("UTC")

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location)
}
