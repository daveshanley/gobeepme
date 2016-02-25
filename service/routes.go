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


    /*
    router.
        Methods("GET").
        Path("/").
        Name("Static").
        Handler(http.FileServer(http.Dir("./static")))



    /*
    router.Handle("/", IndexHandler).Methods("GET")

  http.Handle("/", router)
  http.ListenAndServe(":8000", nil)
     */


//myRouter.Handle('/images/{rest}', http.StripPrefix("/images/", http.FileServer(http.Dir(HomeFolder + "images/"))))
    //2. myRouter.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir(HomeFolder + "images/"))))


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

/*
fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
 */
