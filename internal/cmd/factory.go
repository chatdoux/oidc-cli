package cmd

import (
	"github.com/chatdoux/oidc-cli/internal/browser"
	"github.com/chatdoux/oidc-cli/internal/clock"
	"github.com/chatdoux/oidc-cli/internal/ioreader"
	"github.com/chatdoux/oidc-cli/internal/logger"
	"github.com/chatdoux/oidc-cli/internal/tokencache"
)

type Factory struct {
	Browser        browser.Opener
	Reader         ioreader.Reader
	TokenCacheRepo tokencache.Repo
	Logger         logger.LogWriter
	Clock          clock.Clock
}

func NewFactory() *Factory {
	return &Factory{
		Browser:        browser.New(),
		Reader:         ioreader.New(),
		TokenCacheRepo: tokencache.NewRepository(),
		Logger:         logger.New(),
		Clock:          clock.New(),
	}
}
