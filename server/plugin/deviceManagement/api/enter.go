package api

type ApiGroup struct {
	DeviceApi
	AuthenticationApi
	AuthorizationApi
}

var ApiGroupApp = new(ApiGroup)
