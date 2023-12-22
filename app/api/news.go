package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/yongchengchen/gf-ireader/app/model"
	"github.com/yongchengchen/gf-ireader/library/response"
)

// 用户API管理对象
var NewsApi = new(newsApi)

type newsApi struct{}

// redis Get command
func (a *newsApi) GetRecords(r *ghttp.Request) {
	var news []model.News
	err := g.DB("sqlite").Model("news").Where("readed", 0).Scan(&news)

	if err != nil {
		response.JsonExit(r, 500, err.Error())
	}
	response.JsonExit(r, 200, "success", news)
}

func (a *newsApi) GetDayRecords(r *ghttp.Request) {
	var news []model.News

	var day = r.Get("day")

	err := g.DB("sqlite").Model("news").Where("readed", 0).Where("created_at", day).Scan(&news)

	if err != nil {
		response.JsonExit(r, 500, err.Error())
	}
	response.JsonExit(r, 200, "success", news)
}

func (a *newsApi) GetRecord(r *ghttp.Request) {
	var (
		item *model.News
	)
	var id = r.Get("id")

	if err := g.DB("sqlite").Model("news").Where("id", id).Scan(&item); err != nil {
		response.JsonExit(r, 400, err.Error())
	}

	response.JsonExit(r, 200, "success", item)
}

func (a *newsApi) InsertRecord(r *ghttp.Request) {
	var (
		data *model.News
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 400, err.Error())
	}

	v, err := g.DB("sqlite").GetValue(gctx.New(), "select id from news order by id desc limit 1")
	if err != nil {
		response.JsonExit(r, 400, err.Error())
	}
	data.Id = 1 + v.Uint()
	ret, err := g.DB("sqlite").Model("news").Insert(data)
	if err != nil {
		response.JsonExit(r, 400, err.Error())
	}

	response.JsonExit(r, 200, "success", ret)
}

func (a *newsApi) UpdateRecord(r *ghttp.Request) {
	var (
		data *model.News
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 400, err.Error())
	}

	ret, err := g.DB("sqlite").Model("news").Update(data, "id", data.Id)
	if err != nil {
		response.JsonExit(r, 400, err.Error())
	}

	response.JsonExit(r, 200, "success", ret)
}

func (a *newsApi) DeleteRecord(r *ghttp.Request) {
	var (
		data *model.News
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 400, err.Error())
	}

	ret, err := g.DB("sqlite").Model("news").Delete("id", data.Id)
	if err != nil {
		response.JsonExit(r, 400, err.Error())
	}

	response.JsonExit(r, 200, "success", ret)
}
