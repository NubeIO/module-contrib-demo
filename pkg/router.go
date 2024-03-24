package pkg

import (
	"encoding/json"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/lib-module-go/nmodule"
	"github.com/NubeIO/lib-module-go/router"
	"net/http"
	"time"
)

type health struct {
	FormattedDateTime string    `json:"formatted_date_time"`
	TimeDate          time.Time `json:"time_date"`
}

var route *router.Router

func (m *Module) CallModule(method nhttp.Method, urlString string, headers http.Header, body []byte) ([]byte, error) {
	mod := (nmodule.Module)(m)
	return route.CallHandler(&mod, method, urlString, headers, body)
}

func InitRouter() {
	route = router.NewRouter()

	route.Handle(nhttp.GET, "/api/ping", GetHealthStatus) // http://0.0.0.0:1660/api/modules/module-contrib-demo/api/ping
	route.Handle(nhttp.GET, "/api/weather/:town", GetWeather)
}

func GetHealthStatus(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return json.Marshal(health{
		FormattedDateTime: time.Now().Format(time.Stamp),
		TimeDate:          time.Now().UTC(),
	})
}

func GetWeather(m *nmodule.Module, r *router.Request) ([]byte, error) {
	weather, _, err := (*m).(*Module).getWeather(r.PathParams["town"])
	if err != nil {
		return nil, err
	}
	return weather, nil
}
