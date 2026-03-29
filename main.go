package mygopackage

import (
    "context"
    _ "embed"
    "errors"
    "fmt"
    "log/slog"

    runlicense "github.com/runlicense/sdk-go"
)

//go:embed keys/runlicense.key
var publicKey string

var ErrUnlicensed = errors.New("my-go-package: license verification failed")

type Option func(*options)

type options struct {
    logger *slog.Logger
}

// WithLogger enables verbose logging of the license verification pipeline.
func WithLogger(l *slog.Logger) Option {
    return func(o *options) {
        o.logger = l
    }
}

type Client struct{}

func New(opts ...Option) (*Client, error) {
    var o options
    for _, fn := range opts {
        fn(&o)
    }

    var rlOpts []runlicense.Option
    if o.logger != nil {
        rlOpts = append(rlOpts, runlicense.WithLogger(o.logger))
    }

    result, err := runlicense.Activate(context.Background(), "ok200paul/my-go-package", publicKey, rlOpts...)
    if err != nil {
        var licErr *runlicense.LicenseError
        if errors.As(err, &licErr) {
            return nil, fmt.Errorf("%w: %s", ErrUnlicensed, licErr.Error())
        }
        return nil, ErrUnlicensed
    }

    if o.logger != nil {
        o.logger.Info("licensed", "customer_id", result.License.CustomerID)
    }

    return &Client{}, nil
}

func (c *Client) DoSomething() string {
    return "something done"
}
