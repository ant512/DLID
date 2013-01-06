DLID
====

This is a simple library that can extract some basic information from data
retrieved from the PDF417 barcode on the back of most US driving licenses.  Most
licenses that feature these barcodes follow the DL/ID standard as laid out here:

 * [http://www.aamva.org/DL-ID-Card-Design-Standard][1]

 [1]: http://www.aamva.org/DL-ID-Card-Design-Standard

The barcode encodes data such as the driver's name, address, zip code, etc.

Although the parser is by no means complete, it will currently extract:

 - First name
 - Middle name
 - Last name
 - Street
 - City
 - State
 - Country
 - Postal code
 - Sex
 - Social security number
 - Date of birth

The DL/ID standard has proven difficult for both the standards body to define
and for implementors to follow.  Issues encountered so far include:

 - The standard is vague, and early versions omit crucial information such as
   the delimiter to use in fields that contain multiple pieces of information
   (each state therefore chooses their own at random).
 - The standard makes senseless changes between versions, such as changing from
   ISO format dates (good) to US-style dates (bad) to using both formats
   simultaneously and choosing the format based on whether the license is issued
   in the US or Canada (ugly).
 - The standard mixes fixed-width and variable-width fields arbitrarily, but
   also includes field separators, making the fixed-width fields pointless.
 - The fixed-width fields are frequently too large for the data they contain.
 - Different versions of the standard switch from one bad field encoding scheme
   to another and rarely arrive at a reasonable approach.
 - South Carolinans can't read and used the wrong ASCII code as a delimiter in
   the data header.
 - South Carolinans can't count and their implementation of the data header has
   an off-by-one bug.
 - Coloradoans can't follow simple instructions and mis-ordered the data in the
   name field.


Usage
-----

Fetch the package in the usual way:

    go get org.bitbucket/ant512/dlid

Then import it into your source:

    import "org.bitbucket/ant512/dlid/dlidparser"

Call it like so:

    s, err := dlidparser.Parse("barcodedata")


Links
-----

 - [Development blog][2]
 - [BitBucket page][3]

  [2]: http://simianzombie.com
  [3]: http://bitbucket.org/ant512/dlid


Email
-----

  Contact me at <ant@simianzombie.com>.
