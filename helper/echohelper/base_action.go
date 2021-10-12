package echohelper

import (
	"github.com/labstack/echo"
	"net/http"
)

type Bind func(c echo.Context, v interface{}) error
type Render func(c echo.Context, code int, data interface{}) error

func DefaultBind(c echo.Context, v interface{}) error {
	return c.Bind(v)
}

func DefaultRender(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, data)
}

type BaseAction struct {
	bind   Bind
	render Render
}

func (a BaseAction) SetBind(bind Bind) {
	a.bind = bind
}

func (a BaseAction) SetRender(render Render) {
	a.render = render
}

func (a BaseAction) GetBind() Bind {
	if a.bind == nil {
		return DefaultBind
	} else {
		return a.bind
	}
}

func (a BaseAction) GetRender() Render {
	if a.render == nil {
		return DefaultRender
	} else {
		return a.render
	}
}

func (a BaseAction) Route(r *echo.Group) {}

func (a BaseAction) Name() string {
	return ""
}

func (a BaseAction) Path() string {
	return ""
}

func (a BaseAction) Bind(c echo.Context, v interface{}) error {
	return a.GetBind()(c, v)
}

type Error struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func (a BaseAction) ErrorRender(c echo.Context, code int, eCode int, eMsg error) error {
	e := Error{
		ErrorCode: eCode,
		ErrorMsg:  eMsg.Error(),
	}
	return a.GetRender()(c, code, e)
}

func (a BaseAction) DataRender(c echo.Context, v interface{}) error {
	return a.GetRender()(c, http.StatusOK, v)
}

func (a BaseAction) RenderRequestError(c echo.Context, err error) error {
	return a.ErrorRender(c, http.StatusBadRequest, 1000, err)
}

func (a BaseAction) RenderWrapper(c echo.Context, v interface{}, err error) error {
	if err != nil {
		return a.ErrorRender(c, http.StatusBadGateway, 1001, err)
	}
	return a.DataRender(c, v)
}
