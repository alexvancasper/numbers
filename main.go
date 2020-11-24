package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "math/rand"
    "strconv"
)

func hello(w http.ResponseWriter, req *http.Request) {
    var help string
    help = `
    <p><b>GET /number/10</b><br>
    The request will return 10 random integers</p>


    <p><b>GET /number/10/5</b><br>
    The request will return 5 random integers with maximum is 10.</p>

    <p><b>GET /number/10/20/5</b><br>
    The request will return 5 random integers in range 10 - 20</p>`
    fmt.Fprintf(w, help)
}

// func headers(w http.ResponseWriter, req *http.Request) {

//     for name, headers := range req.Header {
//         for _, h := range headers {
//             fmt.Fprintf(w, "%v: %v\n", name, h)
//         }
//     }
// }

func number(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    number, _ := strconv.Atoi(vars["number"])

    for i:=0;i<number;i++{
        fmt.Fprintf(w, "%d\n", rand.Int())
    }
}

func number_max(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    number, _ := strconv.Atoi(vars["number"])
    max, _ := strconv.Atoi(vars["max"])

    for i:=0;i<number;i++{
        fmt.Fprintf(w, "%d\n", rand.Intn(max))
    }
}

func number_min_max(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    number, _ := strconv.Atoi(vars["number"])
    min, _ := strconv.Atoi(vars["min"])
    max, _ := strconv.Atoi(vars["max"])

    for i:=0;i<number;i++{
        fmt.Fprintf(w, "%d\n", rand.Intn(max-min+1)+min)
    }
}


func main() {

    r := mux.NewRouter()
    r.HandleFunc("/", hello).Methods("GET")
    r.HandleFunc("/number/{number:[0-9]+}", number).Methods("GET")
    r.HandleFunc("/number/{max:[0-9]+}/{number:[0-9]+}", number_max).Methods("GET")
    r.HandleFunc("/number/{min:[0-9]+}/{max:[0-9]+}/{number:[0-9]+}", number_min_max).Methods("GET")

    // r.HandleFunc("/prime/{max:[0-9]+}/{number:[0-9]+}", numnber).Methods("GET")

    fmt.Println("Server is listening...")
    http.Handle("/", r)
    http.ListenAndServe(":80", nil)
}