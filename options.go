package tcp_server

type Options struct {
}

type Option func(*Options)

func WhenNewMessage() Option {
	return func(o *Options) {

	}
}
