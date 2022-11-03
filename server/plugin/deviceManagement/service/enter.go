package service

type ServiceGroup struct {
	DeviceService
	AuthenticationSerivice
	AuthorizationSerivice
}

var ServiceGroupApp = new(ServiceGroup)
