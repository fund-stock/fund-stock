package v1

type BaseController struct {
}

type Group struct {
	BaseController
	KlineController
	ConfigController
}
