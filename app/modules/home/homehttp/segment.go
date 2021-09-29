package homehttp

import (
	"net/http"

	"gid/app/g"
	"gid/app/g/preinit"
	"gid/app/modules/home/homedao"
	"gid/app/modules/home/homemodel"

	"github.com/hulklab/yago"
	"github.com/hulklab/yago/base/basehttp"
)

type SegmentHttp struct {
	basehttp.BaseHttp
}

func init() {
	SegmentHttp := new(SegmentHttp)

	userGroup := yago.NewHttpGroupRouter("/")
	userGroup.Get("/ping", SegmentHttp.PongAction)
	userGroup.Post("/tag", SegmentHttp.AddAction)
	userGroup.Get("/id/:tag", SegmentHttp.GetIdAction)
	userGroup.Get("/snowid", SegmentHttp.GetSnowIdAction)

	yago.SetHttpNoRouter(SegmentHttp.NoRouterAction)
}

func (h *SegmentHttp) NoRouterAction(c *yago.Ctx) {
	c.JSON(http.StatusNotFound, g.Hash{
		"error": "404, page not exists",
	})
}

func (h *SegmentHttp) PongAction(c *yago.Ctx) {
	c.SetData("pong")
}

func (h *SegmentHttp) GetIdAction(c *yago.Ctx) {
	var (
		tag string
		id  int64
		err error
	)

	// tag 不能空
	tag = c.Param("tag")
	if tag == "" {
		c.SetData("param error")
		return
	}

	// 查询 ID
	if id, err = preinit.Srv.GetId(tag); err != nil {
		c.SetData("get id error")
	}

	c.SetData(id)
}

func (h *SegmentHttp) AddAction(c *yago.Ctx) {
	var err error
	var s homedao.Segments
	if err := c.ShouldBind(&s); err != nil {
		c.SetError(err)
		return
	}

	if err = homemodel.NewSegmentsModel().Add(&s); err != nil {
		c.SetError(err)
		return
	}

	c.SetDataOrErr("OK", err)
}

func (h *SegmentHttp) GetSnowIdAction(c *yago.Ctx) {
	c.SetData(preinit.Srv.SnowFlakeGetId())
}
