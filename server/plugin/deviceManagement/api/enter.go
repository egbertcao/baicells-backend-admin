package api

type ApiGroup struct {
	DeviceApi
	AuthenticationApi
	AuthorizationApi
	TemplateApi
	InstanceApi
}

var ApiGroupApp = new(ApiGroup)
