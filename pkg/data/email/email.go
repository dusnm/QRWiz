package email

import (
	"errors"
	"fmt"
	"net/mail"
	"net/url"
	"strings"
)

type (
	Email struct {
		To      *mail.Address
		CC      []*mail.Address
		BCC     []*mail.Address
		Subject string
		Body    string
	}
)

func splitAndParseAddresses(addresses string) ([]*mail.Address, error) {
	parsedAddresses := make([]*mail.Address, 0, 10)
	if addresses == "" {
		return parsedAddresses, nil
	}

	for _, address := range strings.Split(addresses, ",") {
		parsedAddress, err := mail.ParseAddress(address)
		if err != nil {
			return parsedAddresses, err
		}

		parsedAddresses = append(parsedAddresses, parsedAddress)
	}

	return parsedAddresses, nil
}

func joinAddresses(addresses []*mail.Address) string {
	stringAddresses := make([]string, 0, len(addresses))
	for _, address := range addresses {
		stringAddresses = append(stringAddresses, address.Address)
	}

	return strings.Join(stringAddresses, ",")
}

func New(
	to string,
	cc string,
	bcc string,
	subject string,
	body string,
) (Email, error) {
	if to == "" {
		return Email{}, errors.New("to cannot be empty")
	}

	toAddress, err := mail.ParseAddress(to)
	if err != nil {
		return Email{}, err
	}

	ccAddresses, err := splitAndParseAddresses(cc)
	if err != nil {
		return Email{}, err
	}

	bccAddresses, err := splitAndParseAddresses(bcc)
	if err != nil {
		return Email{}, err
	}

	return Email{
		To:      toAddress,
		CC:      ccAddresses,
		BCC:     bccAddresses,
		Subject: subject,
		Body:    body,
	}, nil
}

func (e Email) String() string {
	urlValues := url.Values{}
	template := fmt.Sprintf("mailto:%s", e.To.Address)

	if len(e.CC) > 0 {
		cc := joinAddresses(e.CC)
		urlValues.Add("cc", cc)
	}

	if len(e.BCC) > 0 {
		bcc := joinAddresses(e.BCC)
		urlValues.Add("bcc", bcc)
	}

	if e.Subject != "" {
		urlValues.Add("subject", e.Subject)
	}

	if e.Body != "" {
		urlValues.Add("body", e.Body)
	}

	if query := urlValues.Encode(); query != "" {
		template = template + "?" + query
	}

	return template
}
