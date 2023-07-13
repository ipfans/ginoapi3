package ginoapi3

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*RouterGroup
	engine *gin.Engine
	schema *openapi3.T

	schemaPath       string
	schemaUIPath     string
	schemaMiddleware []gin.HandlerFunc
	disableSchema    bool
}

// InitSchemaStruct init schema struct
func initSchemaStruct() *openapi3.T {
	// The default info is the name of the executable file and the version is "1.0.0"
	_, fn := filepath.Split(os.Args[0])

	return &openapi3.T{
		OpenAPI: DefaultSchemaVersion,
		Info: &openapi3.Info{
			Title:   fn,
			Version: "1.0.0",
		},
		Paths: openapi3.Paths{},
	}
}

// New is a alias of `gin.New()`
func New() *Engine {
	e := gin.New()
	return &Engine{
		RouterGroup: &RouterGroup{
			group: &e.RouterGroup,
		},
		engine:       e,
		schema:       initSchemaStruct(),
		schemaPath:   "/openapi.json",
		schemaUIPath: "/openapi",
	}
}

// Default is a alias of `gin.Default()`
func Default() *Engine {
	e := gin.Default()
	return &Engine{
		RouterGroup: &RouterGroup{
			group: &e.RouterGroup,
		},
		engine:       e,
		schema:       initSchemaStruct(),
		schemaPath:   "/openapi.json",
		schemaUIPath: "/openapi",
	}
}

// Info sets the `Info` field of the OpenAPI schema.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#info-object
func (e *Engine) Info(info *openapi3.Info) {
	e.schema.Info = info
}

// Engine returns the underlying `*gin.Engine` instance.
// Use this method only if you want to extend the functionality of Gin. All `RouteGroup` methods will be skipped in openapi if you set it in the underlying gin engine.
func (e *Engine) Engine() *gin.Engine {
	return e.engine
}

// Schema returns the underlying `*openapi3.T` instance.
func (e *Engine) Schema() *openapi3.T {
	return e.schema
}

// DisableSchemaHandler disables the schema handler. The schema handler is enabled by default.
func (e *Engine) DisableSchemaHandler() {
	e.disableSchema = true
}

// SchemaPath sets the path of the OpenAPI schema file. The default path is `/openapi.json`.
func (e *Engine) SchemaPath(path string) {
	e.schemaPath = path
}

// SchemaUIPath sets the path of the OpenAPI schema UI. The default path is `/openapi`.
func (e *Engine) SchemaUIPath(path string) {
	e.schemaUIPath = path
}

// setupSchema sets up the schema handler and schema UI handler.
func (e *Engine) setupSchema() {
	if e.disableSchema {
		return
	}

	var schemaHandlers, schemaUIHandlers []gin.HandlerFunc
	if len(e.schemaMiddleware) > 0 {
		schemaHandlers = append(schemaHandlers, e.schemaMiddleware...)
		schemaUIHandlers = append(schemaUIHandlers, e.schemaMiddleware...)
	}
	schemaHandlers = append(schemaHandlers, e.SchemaHandler())
	schemaUIHandlers = append(schemaUIHandlers, e.SchemaUIHandler(nil))

	e.engine.GET(e.schemaPath, schemaHandlers...)
	e.engine.GET(e.schemaUIPath, schemaUIHandlers...)
}

// SchemaHandler returns the schema handler. You can use it as your wish.
func (e *Engine) SchemaHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, e.schema)
	}
}

// SchemaUIHandler returns the schema UI handler with extensions ability. You can use it as your wish.
func (e *Engine) SchemaUIHandler(opt *RedocUIOption) gin.HandlerFunc {
	if opt == nil {
		opt = &RedocUIOption{
			HideDownloadButton: true,
		}
	}
	optJSON, _ := json.Marshal(opt)
	return func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Writer.WriteHeader(http.StatusOK)
		ui := fmt.Sprintf(redocUI, e.schemaPath, string(optJSON))
		_, _ = c.Writer.Write([]byte(ui))
	}
}

// SchemaMiddleware sets the middleware for the schema handler. It will be called before the schema handler and schema UI handler is called.
func (e *Engine) SchemaMiddleware(middleware ...gin.HandlerFunc) {
	e.schemaMiddleware = middleware
}

// Use is a alias of `gin.Engine.Use()`
func (e *Engine) Use(middleware ...gin.HandlerFunc) {
	e.engine.Use(middleware...)
}

// HandleContext is a alias of `gin.Engine.HandleContext()`
func (e *Engine) HandleContext(c *gin.Context) {
	e.engine.HandleContext(c)
}

// LoadHTMLFiles is a alias of `gin.Engine.LoadHTMLFiles()`
func (e *Engine) LoadHTMLFiles(files ...string) {
	e.engine.LoadHTMLFiles(files...)
}

// LoadHTMLGlob is a alias of `gin.Engine.LoadHTMLGlob()`
func (e *Engine) LoadHTMLGlob(pattern string) {
	e.engine.LoadHTMLGlob(pattern)
}

// NoMethod is a alias of `gin.Engine.NoMethod()`
func (e *Engine) NoMethod(handlers ...gin.HandlerFunc) {
	e.engine.NoMethod(handlers...)
}

// NoRoute is a alias of `gin.Engine.NoRoute()`
func (e *Engine) NoRoute(handlers ...gin.HandlerFunc) {
	e.engine.NoRoute(handlers...)
}

// Routes is a alias of `gin.Engine.Routes()`
func (e *Engine) Routes() (routes gin.RoutesInfo) {
	return e.engine.Routes()
}

// Run is a alias of `gin.Engine.Run()`. It will also automatically generate the OpenAPI schema file and serve it.
func (e *Engine) Run(addr ...string) (err error) {
	e.setupSchema()
	return e.engine.Run(addr...)
}

// RunFd is a alias of `gin.Engine.RunFd()`. It will also automatically generate the OpenAPI schema file and serve it.
func (e *Engine) RunFd(fd int) (err error) {
	e.setupSchema()
	return e.engine.RunFd(fd)
}

// RunListener is a alias of `gin.Engine.RunListener()`. It will also automatically generate the OpenAPI schema file and serve it.
func (e *Engine) RunListener(listener net.Listener) (err error) {
	e.setupSchema()
	return e.engine.RunListener(listener)
}

// RunUnix is a alias of `gin.Engine.RunUnix()`. It will also automatically generate the OpenAPI schema file and serve it.
func (e *Engine) RunTLS(addr, certFile, keyFile string) (err error) {
	e.setupSchema()
	return e.engine.RunTLS(addr, certFile, keyFile)
}

// RunUnix is a alias of `gin.Engine.RunUnix()`. It will also automatically generate the OpenAPI schema file and serve it.
func (e *Engine) RunUnix(file string) (err error) {
	e.setupSchema()
	return e.engine.RunUnix(file)
}

// ServeHTTP is a alias of `gin.Engine.ServeHTTP()`. It will also automatically generate the OpenAPI schema file and serve it.
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.setupSchema()
	e.engine.ServeHTTP(w, req)
}

// SetFuncMap is a alias of `gin.Engine.SetFuncMap()`
func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.engine.SetFuncMap(funcMap)
}

// SetHTMLTemplate is a alias of `gin.Engine.SetHTMLTemplate()`
func (e *Engine) SetHTMLTemplate(templ *template.Template) {
	e.engine.SetHTMLTemplate(templ)
}

// SetTrustedProxies is a alias of `gin.Engine.SetTrustedProxies()`
func (e *Engine) SetTrustedProxies(trustedProxies []string) error {
	return e.engine.SetTrustedProxies(trustedProxies)
}
