package postgresql

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"sync"

	"github.com/jasonhancock/cobraflags/flags"
	"github.com/jasonhancock/go-logger"
	"github.com/nsqio/go-nsq"
	"github.com/spf13/pflag"
)

type Config struct {
	Addr string

	SSLCACert string
	SSLKey    string
	SSLCert   string

	flags.FlagSet
}

func NewConfig(flagSet *pflag.FlagSet) *Config {
	var c Config

	c.Add(
		flagSet,

		flags.New(
			&c.Addr,
			"nsq-addr",
			"The address:port of the nsq server.",
			flags.Env("NSQ_ADDR"),
			flags.Default("127.0.0.1:4150"),
			flags.Required(),
		),

		flags.New(
			&c.SSLCACert,
			"nsq-ssl-ca-cert",
			"The path to the CA certificate file.",
			flags.Env("NSQ_SSL_CA_CERT"),
		),

		flags.New(
			&c.SSLCert,
			"nsq-ssl-cert",
			"The path to the certificate file.",
			flags.Env("NSQ_SSL_CERT"),
		),

		flags.New(
			&c.SSLKey,
			"nsq-ssl-key",
			"The path to the ssl private key file.",
			flags.Env("NSQ_SSL_KEY"),
		),
	)

	return &c
}

type options struct {
	tlsConfig *tls.Config
	userAgent string
}

// Option is used to customize the configuration.
type Option func(*options)

// WithTLSConfig specifies the *tls.Config to use.
func WithTLSConfig(c *tls.Config) Option {
	return func(o *options) {
		o.tlsConfig = c
	}
}

// WithUserAgent specifies the user agent string in the client.
func WithUserAgent(ua string) Option {
	return func(o *options) {
		o.userAgent = ua
	}
}

func (cfg *Config) Consumer(
	ctx context.Context,
	l *logger.L,
	wg *sync.WaitGroup,
	topic, channel string,
	h nsq.Handler,
	opts ...Option,
) (*nsq.Consumer, error) {
	conf, err := cfg.baseConfig(opts...)
	if err != nil {
		return nil, err
	}

	consumer, err := nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		return nil, fmt.Errorf("getting nsq consumer: %w", err)
	}
	setupLoggersConsumer(l, consumer)
	consumer.AddHandler(h)

	if err := consumer.ConnectToNSQDs([]string{cfg.Addr}); err != nil {
		return nil, fmt.Errorf("connecting to nsq: %w", err)
	}

	consumerWaitStop(ctx, wg, consumer)

	return consumer, nil
}

func consumerWaitStop(ctx context.Context, wg *sync.WaitGroup, consumer *nsq.Consumer) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		consumer.Stop()
		<-consumer.StopChan
	}()
}

func (cfg *Config) baseConfig(opts ...Option) (*nsq.Config, error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	c := nsq.NewConfig()

	if o.userAgent != "" {
		c.Set("user_agent", o.userAgent)
	}

	if o.tlsConfig != nil {
		c.TlsConfig = o.tlsConfig
	} else if cfg.SSLCert != "" && cfg.SSLKey != "" {
		ca, err := os.ReadFile(cfg.SSLCACert)
		if err != nil {
			return nil, fmt.Errorf("reading CA cert: %w", err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM([]byte(ca))

		cert, err := tls.LoadX509KeyPair(cfg.SSLCert, cfg.SSLKey)
		if err != nil {
			return nil, fmt.Errorf("loading TLS keypair: %w", err)
		}

		o.tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
			MinVersion:   tls.VersionTLS12,
		}
	}

	if o.tlsConfig != nil {
		if err := c.Set("tls_v1", true); err != nil {
			return nil, fmt.Errorf("enabling TLS: %w", err)
		}
		if err := c.Set("tls_config", o.tlsConfig); err != nil {
			return nil, fmt.Errorf("setting tls config: %w", err)
		}
	}

	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("config didn't validate: %w", err)
	}

	return c, nil
}

func (cfg *Config) Producer(l *logger.L, opts ...Option) (*nsq.Producer, error) {
	conf, err := cfg.baseConfig(opts...)
	if err != nil {
		return nil, err
	}

	producer, err := nsq.NewProducer(cfg.Addr, conf)
	if err != nil {
		return nil, fmt.Errorf("intializing nsq producer: %w", err)
	}

	setupLoggersProducer(l, producer)

	return producer, nil
}

type logLevelFunc func(msg any, keyvals ...any)

func (f logLevelFunc) Output(calldepth int, s string) error {
	f(s)
	return nil
}

func setupLoggersProducer(l *logger.L, producer *nsq.Producer) {
	producer.SetLoggerForLevel(logLevelFunc(l.Debug), nsq.LogLevelDebug)
	producer.SetLoggerForLevel(logLevelFunc(l.Info), nsq.LogLevelInfo)
	producer.SetLoggerForLevel(logLevelFunc(l.Warn), nsq.LogLevelWarning)
	producer.SetLoggerForLevel(logLevelFunc(l.Err), nsq.LogLevelError)
	producer.SetLoggerForLevel(logLevelFunc(l.Err), nsq.LogLevelMax)
}

func setupLoggersConsumer(l *logger.L, consumer *nsq.Consumer) {
	consumer.SetLoggerForLevel(logLevelFunc(l.Debug), nsq.LogLevelDebug)
	consumer.SetLoggerForLevel(logLevelFunc(l.Info), nsq.LogLevelInfo)
	consumer.SetLoggerForLevel(logLevelFunc(l.Warn), nsq.LogLevelWarning)
	consumer.SetLoggerForLevel(logLevelFunc(l.Err), nsq.LogLevelError)
	consumer.SetLoggerForLevel(logLevelFunc(l.Err), nsq.LogLevelMax)
}
