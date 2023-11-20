package contact

import (
	"errors"
	"fmt"
	"net/mail"
	"net/url"
	"strconv"

	"github.com/dusnm/QRWiz/pkg/escape"
	"github.com/nyaruka/phonenumbers"
)

type (
	Address struct {
		PostalCode      uint64
		StreetAndNumber string
		City            string
		State           string
		Country         string
	}

	Contact struct {
		Name    string
		Surname string
		Address Address
		Phone   *phonenumbers.PhoneNumber
		Email   *mail.Address
		Website *url.URL
	}
)

func New(
	name string,
	surname string,
	postalCode string,
	streetAndNumber string,
	city string,
	state string,
	countryAlpha2 string,
	countryAlpha3 string,
	phone string,
	email string,
	website string,
) (Contact, error) {
	if name == "" {
		return Contact{}, errors.New("name required")
	}

	if surname == "" {
		return Contact{}, errors.New("surname required")
	}

	name = escape.SpecialCharacters(escape.HexString(name))
	surname = escape.SpecialCharacters(escape.HexString(surname))

	if streetAndNumber != "" {
		streetAndNumber = escape.SpecialCharacters(escape.HexString(streetAndNumber))
	}

	if city != "" {
		city = escape.SpecialCharacters(escape.HexString(city))
	}

	if state != "" {
		state = escape.SpecialCharacters(escape.HexString(state))
	}

	var pCode uint64
	if postalCode != "" {
		v, err := strconv.ParseUint(postalCode, 10, 64)
		if err != nil {
			return Contact{}, err
		}

		pCode = v
	}

	var pNumber *phonenumbers.PhoneNumber
	if phone != "" {
		v, err := phonenumbers.Parse(phone, countryAlpha2)
		if err != nil {
			return Contact{}, err
		}

		if !phonenumbers.IsValidNumber(v) {
			return Contact{}, errors.New("invalid phone number for given country")
		}

		pNumber = v
	}

	var eMail *mail.Address
	if email != "" {
		v, err := mail.ParseAddress(email)
		if err != nil {
			return Contact{}, err
		}

		eMail = v
	}

	var wSite *url.URL
	if website != "" {
		v, err := url.Parse(website)
		if err != nil {
			return Contact{}, err
		}

		wSite = v
	}

	return Contact{
		Name:    name,
		Surname: surname,
		Address: Address{
			PostalCode:      pCode,
			StreetAndNumber: streetAndNumber,
			City:            city,
			State:           state,
			Country:         countryAlpha3,
		},
		Phone:   pNumber,
		Email:   eMail,
		Website: wSite,
	}, nil
}

func (c Contact) formatAddress() string {
	address := "ADR:,,"

	if street := c.Address.StreetAndNumber; street != "" {
		address += street
	}

	address += ","

	if city := c.Address.City; city != "" {
		address += city
	}

	address += ","

	if state := c.Address.State; state != "" {
		address += state
	}

	address += ","

	if postalCode := c.Address.PostalCode; postalCode != 0 {
		address += strconv.FormatUint(postalCode, 10)
	}

	return address + "," + c.Address.Country + ";"
}

func (c Contact) formatEmail() string {
	if email := c.Email; email != nil {
		return "EMAIL:" + email.Address + ";"
	}

	return ""
}

func (c Contact) formatPhone() string {
	if phone := c.Phone; phone != nil {
		phoneS := phonenumbers.Format(
			phone,
			phonenumbers.INTERNATIONAL,
		)

		return "TEL:" + phoneS + ";"
	}

	return ""
}

func (c Contact) formatWebsite() string {
	if website := c.Website; website != nil {
		return "URL:" + website.String() + ";"
	}

	return ""
}

func (c Contact) String() string {
	template := "MECARD:"
	address := c.formatAddress()
	email := c.formatEmail()
	phone := c.formatPhone()
	website := c.formatWebsite()
	name := fmt.Sprintf(
		"N:%s,%s;",
		c.Surname,
		c.Name,
	)

	for _, item := range []string{address, email, name, phone, website} {
		if item != "" {
			template += item
		}
	}

	return template + ";"
}
