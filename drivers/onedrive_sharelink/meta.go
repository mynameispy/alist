package onedrive_sharelink

import (
	"net/http"

	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/op"
)

type Addition struct {
	// Usually one of two
	driver.RootPath
	// define other
	ShareLinkURL       string `json:"url" required:"true"`
	ShareLinkPassword  string `json:"password"`
	IsSharepoint       bool
	downloadLinkPrefix string
	Headers            http.Header
	HeaderTime         int64
	RodAddress         string `json:"rod_address" required:"true" default:"ws://127.0.0.1:7317"`
}

var config = driver.Config{
	Name:        "Onedrive Sharelink",
	OnlyProxy:   true,
	NoUpload:    true,
	DefaultRoot: "/",
	CheckStatus: false,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &OnedriveSharelink{}
	})
}
