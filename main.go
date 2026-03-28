package mygopackage

import (
    "context"
    _ "embed"
    "errors"
    "fmt"
    "log"

    runlicense "github.com/runlicense/sdk-go"
)

//go:embed keys/runlicense.key
var publicKey string

var ErrUnlicensed = errors.New("my-go-package: license verification failed")

var licenseErr error
var license *runlicense.LicensePayload

func init() {
    l, err := runlicense.Activate(context.Background(), "ok200paul/my-go-package", publicKey)
    if err != nil {
        var licErr *runlicense.LicenseError
        if errors.As(err, &licErr) {
            licenseErr = fmt.Errorf("%w: %s", ErrUnlicensed, licErr.Message)
        } else {
            licenseErr = ErrUnlicensed
        }
    }
    license = l
    if license != nil {
        log.Printf("Licensed to: %s", license.CustomerID)
    }
}

type Client struct{}

func New() (*Client, error) {
    if licenseErr != nil {
        return nil, licenseErr
    }
    return &Client{}, nil
}

func (c *Client) DoSomething() string {
    return "something done"
}