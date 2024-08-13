package middleware

import "github.com/go-telegram/bot"

type Middlewares interface {
	DefaultMiddleware(next bot.HandlerFunc) bot.HandlerFunc
}
