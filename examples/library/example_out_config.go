// Code generated by options-gen. DO NOT EDIT.

package main

type OptConfigSetter func(o *Config)

func NewConfig(
	options ...OptConfigSetter,
) Config {
	o := Config{}

	// Setting defaults from field tag (if present)

	for _, opt := range options {
		opt(&o)
	}
	return o
}

func WithSomeName(opt string) OptConfigSetter {
	return func(o *Config) {
		o.name = opt

	}
}

func (o *Config) Validate() error {
	return nil
}
