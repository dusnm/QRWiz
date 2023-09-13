package encoder

import (
	"bytes"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"io"
)

type (
	nopCloser struct {
		io.Writer
	}
)

func (nopCloser) Close() error {
	return nil
}

func Encode(data string) ([]byte, error) {
	qcode, err := qrcode.NewWith(
		data,
		qrcode.WithEncodingMode(qrcode.EncModeByte),
		qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionHighest),
	)

	b := bytes.NewBuffer(nil)
	w := standard.NewWithWriter(
		nopCloser{Writer: b},
		standard.WithBorderWidth(40),
		standard.WithQRWidth(12),
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
	)

	if err = qcode.Save(w); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
