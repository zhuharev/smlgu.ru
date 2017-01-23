package middleware

import (
	//"github.com/fatih/color"
	"github.com/go-macaron/cache"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"time"
)

type Context struct {
	*macaron.Context
	Cache   cache.Cache
	Flash   *session.Flash
	Session session.Store

	//User        *models.User
	//IsSigned    bool
	//IsBasicAuth bool
}

func Contexter() macaron.Handler {
	return func(c *macaron.Context, cache cache.Cache, sess session.Store, f *session.Flash) {
		ctx := &Context{
			Context: c,
			Cache:   cache,
			Flash:   f,
			Session: sess,
		}
		// Compute current URL for real-time change language.
		//ctx.Data["Link"] = setting.AppSubUrl + ctx.Req.URL.Path

		ctx.Data["PageStartTime"] = time.Now()

		c.Map(ctx)
	}
}
