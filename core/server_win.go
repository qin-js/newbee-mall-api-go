//go:build windows
// +build windows

package core

import (
	"gee"
	"net/http"
	"time"
)

func initServer(address string, router *gee.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
