package plugin

import (
	"fmt"
	_log "log"
	"strings"

	"github.com/Kong/go-pdk"
)

const (
	Version  = "0.2"
	Priority = 13
)

type Config interface {
	Access(kong *pdk.PDK)
	Log(kong *pdk.PDK)
}

type pluginConfig struct {
	//log log.Log
	//m   *entities.AutoGenerated
}

// NewPlugin returns a new plugin configuration.
func NewPlugin() Config {

	/* nm := &entities.AutoGenerated{
		WorkspaceName: "workspace_name",
		TraceID: entities.TraceID{
			W3C: "unknown",
		},
	} */

	return &pluginConfig{
		//log: log,
		//m:   nm,
	}
}

// Access is executed for every request from a client
// and, before it is being proxied to the upstream service.
func (c *pluginConfig) Access(kong *pdk.PDK) {

	_log.Printf("log_print_1")

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

	traceid := "unknown"
	traceparent, _ := kong.Request.GetHeader("traceparent")
	kong.Log.Err(fmt.Sprintf("by_header_kong_5_traceparent, traceparent: %s", traceparent))
	if traceparent != "" {
		traceid = strings.Split(traceparent, "-")[1]
	}
	kong.Log.Err(fmt.Sprintf("by_header_kong_5, trace_id: %s", traceid))

	/* n1, err := kong.Ctx.GetSharedString("traceparent")
	if err == nil {
		kong.Log.Err(fmt.Sprintf("by_shared_kong_5, trace_id: {\"w3c\":\"%s\"}", n1))
	} */

	/* n2, err := kong.Nginx.GetCtxAny("tracecontext")
	if err == nil {
		kong.Log.Err(fmt.Sprintf("by_ctx_kong_5, trace_id: {\"w3c\":\"%s\"}", n2))
	} */

	//kong.Log.Err("error_kong_3, a: b, f: d")

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

func (c *pluginConfig) Log(kong *pdk.PDK) {

	str, err := kong.Log.Serialize()
	if err == nil {
		/* m := entities.AutoGenerated{}
		json.Unmarshal([]byte(str), &m)
		//c.m = &m
		b, _ := json.Marshal(m)
		kong.Log.Info("serialize_kong: ", string(b)) */
		kong.Log.Info("serialize_kong: ", str)
	}
}

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
