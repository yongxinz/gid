package preinit

import "gid/app/modules/home/homemodel"

var (
	Srv *homemodel.Service
)

func init() {
	Srv = homemodel.NewService()
}
