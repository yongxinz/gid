package homemodel

import "gid/app/libs"

func (s *Service) NewAllocSnowFlakeId() (a *libs.Worker, err error) {
	return libs.NewWorker(1)
}

func (s *Service) SnowFlakeGetId() int64 {
	return s.snowFlake.GetId()
}
