package v2

type BaseController struct {
}

type Group struct {
	BaseController
	KlineController
	NoticeController
	ConfigController
}
