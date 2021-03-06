package monitoring

// Option type for passing additional options to NewRegistry.
type Option func(options) options

type options struct {
	publishExpvar bool
}

var defaultOptions = options{
	publishExpvar: false,
}

// PublishExpvar enables publishing all registered variables via expvar interface.
// Note: expvar does not allow removal of any stats.
func PublishExpvar(o options) options {
	o.publishExpvar = true
	return o
}

// IgnorePublishExpvar disables publishing expvar variables in a sub-registry.
func IgnorePublishExpvar(o options) options {
	o.publishExpvar = false
	return o
}

func applyOpts(in *options, opts []Option) *options {
	if len(opts) == 0 {
		return ensureOptions(in)
	}

	tmp := *ensureOptions(in)
	for _, opt := range opts {
		tmp = opt(tmp)
	}
	return &tmp
}

func ensureOptions(in *options) *options {
	if in != nil {
		return in
	}

	tmp := defaultOptions
	return &tmp
}
