package homemodel

import (
	"gid/app/libs"

	"github.com/hulklab/yago"
)

func (s *Service) NewAllocSnowFlakeId() (a *libs.Worker, err error) {
	conf := yago.Config.GetStringMap("app")
	id := conf["snow_flake_id"].(int64)
	return libs.NewWorker(id)
}

func (s *Service) SnowFlakeGetId() int64 {
	return s.snowFlake.GetId()
}
