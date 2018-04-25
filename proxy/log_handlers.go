package proxy

import (
	"net/http"

	"github.com/elazarl/goproxy"
	log "github.com/sirupsen/logrus"

	"github.com/9seconds/crawlera-headless-proxy/config"
)

func handlerLogReqInitial(proxy *goproxy.ProxyHttpServer, conf *config.Config) handlerTypeReq {
	return func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		log.WithFields(log.Fields{
			"reqid":          getState(ctx).id,
			"method":         req.Method,
			"url":            req.URL,
			"proto":          req.Proto,
			"content-length": req.ContentLength,
			"remote-addr":    req.RemoteAddr,
			"headers":        req.Header,
		}).Debug("Incoming HTTP request.")
		return req, nil
	}
}

func handlerLogReqSent(proxy *goproxy.ProxyHttpServer, conf *config.Config) handlerTypeReq {
	return func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		log.WithFields(log.Fields{
			"reqid":          getState(ctx).id,
			"method":         req.Method,
			"url":            req.URL,
			"proto":          req.Proto,
			"content-length": req.ContentLength,
			"remote-addr":    req.RemoteAddr,
			"headers":        req.Header,
		}).Debug("HTTP request sent.")
		return req, nil
	}
}

func handlerLogRespInitial(proxy *goproxy.ProxyHttpServer, conf *config.Config) handlerTypeResp {
	return func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		log.WithFields(log.Fields{
			"reqid":           getState(ctx).id,
			"method":          resp.Request.Method,
			"url":             resp.Request.URL,
			"proto":           resp.Proto,
			"content-length":  resp.ContentLength,
			"headers":         resp.Header,
			"status":          resp.Status,
			"uncompressed":    resp.Uncompressed,
			"request-headers": resp.Request.Header,
		}).Debug("HTTP response")
		return resp
	}
}