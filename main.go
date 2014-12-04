package servant

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
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

	m.Get("/api/search", GetCardList)
	m.Get("/api/model", CreateModels)
	m.Get("/api/product", GetProductList)
	m.Get("/api/illustrator", GetIllustratorList)
	m.Get("/api/constraint", GetConstraintList)
	m.Get("/api/type", GetTypeList)
	m.Get("/api/import/:from/:to", func(r render.Render, params martini.Params, req *http.Request) {
		for i := ToInt(params["from"]); i <= ToInt(params["to"]); i++ {
			CreateCard(i, req)
		}
		r.JSON(200, "finish!")
	})
	m.Get("/api/twitter/login", LoginTwitter)
	m.Get("/api/twitter/callback", CallbackTwitter)
	m.Get("/api/loginUser", LoginUser)
	m.Get("/api/deck", GetPublicDeckList)
	m.Post("/api/deck", binding.Json(FormDeck{}), CreateDeck)
	m.Delete("/api/deck/:owner/:id", DeleteDeck)
	m.Get("/api/deck/:id", GetDeck)
	m.Get("/api/card/:expansion", GetCardByExpansion)
	m.Get("/api/card/:expansion/:no", GetCard)
	m.Get("/api/:userId/deck", GetUserDeckList)
	m.Post("/api/setTestSession", SetTestSettion)
	m.Get("/api/card/hoge/:name/:id", func(r render.Render, params martini.Params, req *http.Request) {
		r.JSON(200, GetCardByNameAndId(params["name"], ToInt(params["id"]), req))
	})
	http.ListenAndServe(":8080", m)
	http.Handle("/", m)
}
