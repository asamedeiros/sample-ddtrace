package plugin

import (
	"context"
	"fmt"
	"strings"

	"github.com/asamedeiros/kong-go-sample-ddtrace/pkg/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/Kong/go-pdk"
)

const (
	Version  = "0.2"
	Priority = 13
)

type Config interface {
	Access(kong *pdk.PDK)
}

type pluginConfig struct {
	log    log.Log
	tracer trace.Tracer
	//m   *entities.AutoGenerated
}

// NewPlugin returns a new plugin configuration.
func NewPlugin(log log.Log, tracer trace.Tracer) Config {

	/* nm := &entities.AutoGenerated{
		WorkspaceName: "workspace_name",
		TraceID: entities.TraceID{
			W3C: "unknown",
		},
	} */

	return &pluginConfig{
		log:    log,
		tracer: tracer,
		//m:   nm,
	}
}

// Access is executed for every request from a client
// and, before it is being proxied to the upstream service.
func (c *pluginConfig) Access(kong *pdk.PDK) {

	traceparent, err := kong.Request.GetHeader("traceparent")
	//kong.Log.Err(fmt.Sprintf("by_header_kong_5_traceparent, traceparent: %s", traceparent))

	ctx := context.Background()

	if err == nil {
		// prepare carrier to set traceparent into context
		carrier := propagation.MapCarrier{}
		carrier.Set("traceparent", traceparent)
		//c.log.Info(fmt.Sprint(carrier))
		// reads tracecontext from the carrier into a returned Context.
		ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
	}

	_, span := c.tracer.Start(ctx, "access")
	defer span.End()

	//c.log.With("trace_id", traceid).Info("logando o trace_id")

	/* n1, err := kong.Ctx.GetSharedString("traceparent")
	if err == nil {
		kong.Log.Err(fmt.Sprintf("by_shared_kong_5, trace_id: {\"w3c\":\"%s\"}", n1))
	} */

	/* n2, err := kong.Nginx.GetCtxAny("tracecontext")
	if err == nil {
		kong.Log.Err(fmt.Sprintf("by_ctx_kong_5, trace_id: {\"w3c\":\"%s\"}", n2))
	} */

	//kong.Log.Err("error_kong_3, a: b, f: d")

	// You can now use your logger in your code.
	c.log.Info("something really cool")

	c.log.Error("something really cool")

	c.log.Warn("something really cool")

	c.log.Debug("something really cool")

	// You can set context for trace correlation using zap.Any or zap.Reflect
	//c.log.Info("setting context", zap.Any("context", ctx))

	traceid := "unknown"
	//spanid := "unknown"
	if traceparent != "" {
		traceid = strings.Split(traceparent, "-")[1]
		//spanid = strings.Split(traceparent, "-")[2]
	}

	kong.Log.Err(fmt.Sprintf("by_header_kong_5, namespace: %s, trace_id: %s", "sample-ddtrace", traceid))

	/*
		tracer := otel.Tracer("example/main")
		ctx, span := tracer.Start(ctx, "example")
		defer span.End()

		log.With("a", "b")(&log.JSONFormatter{})

		standardFields := log.Fields{
			"dd.trace_id": convertTraceID(span.SpanContext().TraceID().String()),
			"dd.span_id":  convertTraceID(span.SpanContext().SpanID().String()),
			"dd.service":  "dataplane",
			"dd.env":      "production",
			"dd.version":  "otel2023052301",
		}

		log.WithFields(standardFields).WithContext(ctx).Info("hello world") */

	//_log.Printf("log_print_1")

	//c.log.Error(fmt.Sprintf("error_2 - %s", "opa"))

	//c.log.Info(fmt.Sprintf("info_2 - %s", "opa"))

	//kong.Log.Debug("debug_kong_2")

	//kong.Log.Err("error_kong_2")

	//kong.Log.Info("info_kong_2")

	//kong.Log.Err(fmt.Sprintf("m_error_kong_2, trace_id: %s", c.m.TraceID.W3C)) //veio outro... justamente pq pegou o request antigo

	/* tracedata, _ := kong.Nginx.GetCtxString("tracedata")
	kong.Log.Err(fmt.Sprintf("m_error_kong_2, %s", tracedata)) */

	/* h, _ := kong.Request.GetHeaders(-1)
	rHeader := make(map[string]string)
	for k := range h {
		rHeader[strings.ToLower(k)] = h[k][0]
	} */

	//c.accessError(kong, 401)

	/* h, _ := kong.Request.GetHeaders(-1)
	rHeader := make(map[string]string)
	for k := range h {
		rHeader[strings.ToLower(k)] = h[k][0]
	}

	rPath, _ := kong.Request.GetPath()
	rMethod, _ := kong.Request.GetMethod()
	rHost, _ := kong.Request.GetHost()
	rRawQuery, _ := kong.Request.GetRawQuery()
	rRemoteAddr, _ := kong.Client.GetIp()
	rBody, _ := kong.Request.GetRawBody()

	req := &entities.PermissionRequest{
		Header:  rHeader,
		Method:  rMethod,
		RawBody: rBody,
		URL: &url.URL{
			Host:     rHost,
			Path:     rPath,
			RawQuery: rRawQuery,
		},
		RemoteAddr: rRemoteAddr,
	}

	reqLog := c.log.With("plugin", "sample-ddtrace").
		With("x-request-id", req.GetHeader("x-request-id")).
		With("method", req.Method).
		With("path", req.URL.Path).
		With("host", req.GetHeader("host")).
		With("user-agent", req.GetHeader("user-agent")).
		With("remote-addr", req.RemoteAddr).
		With("cf-ray", req.GetHeader("cf-ray")).
		With("aws-xray", req.GetHeader("x-amzn-trace-id"))

	reqLog.Error(fmt.Sprintf("error_3 - %s", "opa")) */

	//c.accessError(kong, rsl.Status)
}

/* func (c *pluginConfig) Log(kong *pdk.PDK) {

	str, err := kong.Log.Serialize()
	if err == nil {
		kong.Log.Info("serialize_kong: ", str)
	}
} */

/* func convertTraceID(id string) string {
	if len(id) < 16 {
		return ""
	}
	if len(id) > 16 {
		id = id[16:]
	}
	intValue, err := strconv.ParseUint(id, 16, 64)
	if err != nil {
		return ""
	}
	return strconv.FormatUint(intValue, 10)
} */

/* func (c *pluginConfig) accessError(kong *pdk.PDK, code int) {
	headers := make(map[string][]string)
	if code == http.StatusUnauthorized {
		kong.Response.AddHeader("X-sample-ddtrace", "Unauthorized")
		kong.Response.Exit(code, []byte("Unauthorized"), headers)
	} else {
		kong.Response.AddHeader("X-sample-ddtrace", "Forbidden")
		kong.Response.Exit(code, []byte("Forbidden"), headers)
	}
} */
