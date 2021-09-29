package homemodel

import (
	"gid/app/modules/home/homedao"

	"github.com/hulklab/yago/coms/logger"
)

func (s *Service) GetId(tag string) (id int64, err error) {
	s.alloc.Mu.Lock()
	defer s.alloc.Mu.Unlock()

	val, ok := s.alloc.BizTagMap[tag]
	if !ok {
		// tag 不存在则创建
		if err = s.CreateTag(&homedao.Segments{
			BizTag: tag,
			MaxId:  1,
			Step:   200,
		}); err != nil {
			return 0, err
		}
		val, _ = s.alloc.BizTagMap[tag]
	}
	return val.GetId(s)
}

func (s *Service) CreateTag(e *homedao.Segments) (err error) {
	if err = NewSegmentsModel().Add(e); err != nil {
		logger.Ins().Errorf("create tag error")
		return
	}
	b := &BizAlloc{
		BazTag:  e.BizTag,
		GetDb:   false,
		IdArray: make([]*IdArray, 0),
	}
	b.IdArray = append(b.IdArray, &IdArray{
		Cur:   1,
		Start: 0,
		End:   e.Step,
	})
	s.alloc.BizTagMap[e.BizTag] = b
	return
}
