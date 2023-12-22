package router

import (
	"io/ioutil"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/yongchengchen/gf-ireader/app/api"
)

func init() {
	s := g.Server()

	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.GET("/news", api.NewsApi.GetRecords)
		group.GET("/news/:id", api.NewsApi.GetRecord)
		group.GET("/at/:day/news", api.NewsApi.GetDayRecords)
		group.DELETE("/news/:id", api.NewsApi.UpdateRecord) //news readed
	})

	// s.BindHandler("/ws/:token", api.WsSsh)

	path := gfile.MainPkgPath() + "/dist"

	s.BindStatusHandler(404, func(r *ghttp.Request) {
		// r.Response.w
		file := path + "/index.html"
		c, err := ioutil.ReadFile(file)
		if err != nil {
			r.Response.WriteStatus(404, "Not Found")
		}
		r.Response.WriteStatus(200, c)
	})

	// logrus.Println(path)
	s.SetServerRoot(path)
	s.SetPort(8199)
}
