package servant

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

var m *martini.Martini
var SESSION_KEY = "servant_session"

func init() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(sessions.Sessions(SESSION_KEY, sessions.NewCookieStore([]byte(SESSION_KEY))))

	m.Get("/api/search", func(r render.Render, req *http.Request) {
		res := GetCardList(r, req)
		r.JSON(200, res)
	})

	m.Get("/api/model", func(r render.Render, req *http.Request) {
		res := CreateModels(r, req)
		r.JSON(200, res)
	})

	m.Get("/api/product", func(r render.Render, req *http.Request) {
		res := GetProductList(r, req)
		r.JSON(200, res)
	})

	m.Get("/api/illustrator", func(r render.Render, req *http.Request) {
		res := GetIllustratorList(r, req)
		r.JSON(200, res)
	})

	m.Get("/api/constraint", func(r render.Render, req *http.Request) {
		res := GetConstraintList(r, req)
		r.JSON(200, res)
	})

	m.Get("/api/type", func(r render.Render, req *http.Request) {
		res := GetTypeList(r, req)
		r.JSON(200, res)
	})

	m.Get("/api/import/:from/:to", func(r render.Render, params martini.Params, req *http.Request) {
		for i := ToInt(params["from"]); i <= ToInt(params["to"]); i++ {
			CreateCard(i, req)
		}
		r.JSON(200, "finish!")
	})

	m.Get("/api/twitter/login", LoginTwitter)
	m.Get("/api/twitter/callback", CallbackTwitter)
	m.Get("/api/loginUser", LoginUser)
	m.Post("/api/deck", CreateDeck)
	m.Get("/api/deck/:id", GetDeck)
	m.Get("/api/card/:expansion", GetCardByExpansion)
	m.Get("/api/card/:expansion/:no", GetCard)

	m.Post("/api/setTestSession", SetTestSettion)

	http.ListenAndServe(":8080", m)
	http.Handle("/", m)
}
