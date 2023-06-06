package example

import (
	"fmt"
	"net/http"
	"sample/axum"
)

func sample_inner(r *axum.Request, srv axum.Service) axum.ResponsePacker {
	// fmt.Println("[BEFORE] Hi, Septem. This is inner function")
	rsp := srv.Handle(r.Value)
	// fmt.Println("[AFTER] Hi, Septem. This is inner function")
	return rsp
}

func root(r *http.Request) axum.ResponsePacker {
	fmt.Println("INNER 1 2 3")
	return axum.NewDataResponse(123)
}

func get(r *http.Request) axum.ResponsePacker {
	fmt.Println("GET 4 5 6")
	return axum.NewDataResponse("Hi, Septem")
}

func appDetail(q *axum.Query[map[string]interface{}]) axum.ResponsePacker {
	fmt.Println("Query", q)
	return axum.NewDataResponse("app-detail")
}

type Info struct {
	ApplicationId string `schema:"app_id"`
	Username      string `schema:"user_name"`
}

func appInfo(p *axum.Path[Info]) axum.ResponsePacker {
	fmt.Println(p.Value)
	return axum.NewDataResponse("info-path")
}

func main() {
	app := axum.NewRouter().
		Get("/info", axum.ServiceFunc(func(r *http.Request) axum.ResponsePacker {
			fmt.Println("app info")
			return axum.NewDataResponse("Yohoho")
		})).
		Get("/detail", axum.Arg1Func(appDetail))

	router := axum.NewRouter().
		Get("/", axum.ServiceFunc(root)).
		Get("/get", axum.ServiceFunc(get)).
		Get("/apps/{app_id}/users/{user_name}", axum.Arg1Func(appInfo)).
		Layer(axum.NewService(axum.Layer2Func(sample_inner))).
		SubRouter("/app", app)

	http.ListenAndServe(":8080", router)
}
