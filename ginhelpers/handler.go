package ginhelpers

import (
	"github.com/gin-gonic/gin"
	"github.com/princeparmar/go-helpers/clienthelper"
	"github.com/princeparmar/go-helpers/context"
)

func ConvertContextToGinContext(g *gin.Context) context.IContext {
	ctx := context.NewContextWithParent(g.Request.Context())
	for k, v := range g.Keys {
		ctx.SetValue(context.ContextKey(k), v)
	}
	return ctx
}

func GetGinHandler(factory func() clienthelper.APIExecutor) gin.HandlerFunc {
	return func(g *gin.Context) {
		ctx := ConvertContextToGinContext(g)
		h := clienthelper.GetHandler(factory)
		h(g.Writer, g.Request.WithContext(ctx))
	}
}
