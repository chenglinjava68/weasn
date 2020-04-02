package open_platform

import (
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/open-platform/auth"
	"github.com/prodbox/weasn/open-platform/base"
	"github.com/prodbox/weasn/open-platform/server"
)

type Application struct {
	*base.Client
	opts context.Options
	pool *context.Pool
}

func New(opts ...context.Option) *Application {
	options := context.NewOptions(opts...).InitOptions(defaultOptions())

	app := &Application{
		Client: base.New(options),
		opts:   options,
	}

	app.pool = context.New(func() context.Context {
		return app.allocateContext()
	})
	return app
}

// defaultOptions 默认服务
func defaultOptions() []context.Option {
	return []context.Option{
		defaultHttpClient(),
		defaultCache(),
		defaultEncrypter(),
		defaultAccessToken(),
	}
}

// 初始化上下文
func (this *Application) allocateContext() context.Context {
	return context.NewContext(this.opts)
}

// Server 服务端
func (this *Application) Server(opts ...server.Option) server.Server {
	return server.New(this.pool, auth.NewVerifyTicket(this.opts)).InitOptions(opts...)
}

func (this *Application) AccessToken() context.AccessToken {
	return this.opts.AccessToken
}

// OfficialAccount 代公众号实现业务
func (this *Application) OfficialAccount(appId, refreshToken string) {
}
