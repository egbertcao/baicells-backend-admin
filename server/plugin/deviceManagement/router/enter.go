package router

type RouterGroup struct {
	DeviceRouter
	SecurityRouter
}

var RouterGroupApp = new(RouterGroup)
