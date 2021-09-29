package homemodel

import (
	"gid/app/modules/home/homedao"
	"time"

	"github.com/hulklab/yago"
	"github.com/hulklab/yago/base/basemodel"
	"github.com/hulklab/yago/coms/logger"
	"github.com/hulklab/yago/coms/orm"
	"xorm.io/xorm"
)

type SegmentsModel struct {
	basemodel.BaseModel
}

func NewSegmentsModel(opts ...basemodel.Option) *SegmentsModel {
	model := &SegmentsModel{}

	opts = append(
		opts,
		basemodel.WithDefaultPageSize(1, 20),
		basemodel.WithDefaultOrder(
			basemodel.Order{"create_time": -1},
		),
	)

	model.Init(opts...)

	return model
}

func (m *SegmentsModel) Add(s *homedao.Segments) error {
	// 判断 tag 是否已存在
	segment := &homedao.Segments{BizTag: s.BizTag}

	exists, err := m.GetSession().Exist(segment)
	if exists == true {
		return yago.NewErr("tag " + s.BizTag + " exists")
	}

	// 添加 tag
	_, err = m.GetSession().Insert(s)
	if err != nil {
		logger.Ins().Errorf("add tag error")
		return yago.WrapErr(yago.ErrSystem, err)
	}

	return nil
}

func (m *SegmentsModel) GetList() (s []homedao.Segments, err error) {
	if err = m.GetSession().Find(&s); err != nil {
		logger.Ins().Errorf("get all tag error")
	}
	return
}

func (m *SegmentsModel) GetNextId(tag string) (s *homedao.Segments, err error) {
	err = orm.Ins().Transactional(func(session *xorm.Session) error {
		if _, err := session.Exec(
			"update segments set max_id=max_id+step, update_time=? where biz_tag = ?", time.Now(), tag); err != nil {
			logger.Ins().Errorf("update segments error")
			return err
		}

		s = &homedao.Segments{
			BizTag: tag,
		}
		_, _ = orm.Ins().Get(s)

		return nil
	})

	return
}
