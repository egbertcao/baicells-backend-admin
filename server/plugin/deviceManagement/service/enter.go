package service

type ServiceGroup struct {
	DeviceService
	AuthenticationSerivice
	AuthorizationSerivice
	TemplateService
	InstanceService
}

var ServiceGroupApp = new(ServiceGroup)
