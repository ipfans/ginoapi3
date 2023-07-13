package ginoapi3

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	group *gin.RouterGroup
}

func (r *RouterGroup) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{
		group: r.group.Group(relativePath, handlers...),
	}
}

func (r *RouterGroup) GET(relativePath string, opt []OperationOption, handlers ...gin.HandlerFunc) *RouterGroup {
	r.group.GET(relativePath, handlers...)
	return r
}

func (r *RouterGroup) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	r.group.Handle(httpMethod, relativePath, handlers...)
	return r
}
