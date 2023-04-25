package context

import "context"

type ContextKey string

type IContext interface {
	context.Context
	SetValue(key ContextKey, value any)
	AllKeys() []ContextKey
}

type contextWrapper struct {
	context.Context
	m map[ContextKey]struct{}
}

func NewContext() IContext {
	return &contextWrapper{Context: context.Background()}
}

func NewContextWithParent(ctx context.Context) IContext {
	return &contextWrapper{Context: ctx}
}

func (c *contextWrapper) SetValue(key ContextKey, value any) {
	c.Context = context.WithValue(c.Context, key, value)
	c.m[key] = struct{}{}
}

func (c *contextWrapper) AllKeys() []ContextKey {
	out := []ContextKey{}
	for k := range c.m {
		out = append(out, k)
	}
	return out
}
