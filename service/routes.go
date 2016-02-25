package service

// taken from http://thenewstack.io/make-a-restful-json-api-go/
import (
    "net/http"
    "github.com/gorilla/mux"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
        Methods(route.Method).
        Path(route.Pattern).
        Name(route.Name).
        Handler(route.HandlerFunc)
    }

    // add static route for files
    router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))
    return router
}

var routes = Routes{
    Route{
        "ListDevices",
        "POST",
        "/",
        ListDevices,
    },
    Route{
        "BeepDevice",
        "POST",
        "/beep",
        BeepDevice,
    },
}
