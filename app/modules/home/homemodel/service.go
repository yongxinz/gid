package homemodel

import (
	"gid/app/libs"
)

type Service struct {
	alloc     *Alloc
	snowFlake *libs.Worker
}

func NewService() (s *Service) {
	var err error
	s = &Service{}

	if s.alloc, err = s.NewAllocId(); err != nil {
		panic(err)
	}
	if s.snowFlake, err = s.NewAllocSnowFlakeId(); err != nil {
		panic(err)
	}

	return
}
