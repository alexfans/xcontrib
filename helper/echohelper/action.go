package echohelper

type Action interface {
	GetBind() Bind
	GetRender() Render
	SetBind(bind Bind)
	SetRender(render Render)
	Name() string
	Path() string
}
