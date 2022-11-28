package router

type RouterGroup struct {
	DeviceRouter
	SecurityRouter
	TemplateRouter
	InstanceRouter
}

var RouterGroupApp = new(RouterGroup)
