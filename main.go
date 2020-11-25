package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "math/rand"
    "strconv"
    // "encoding/json"
    "html/template"
)


type User struct {
    Username string
    Password string 
}

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

func user(w http.ResponseWriter, req *http.Request) {
    // w.Header().Set("Content-Type","application/json")

    for key, value := range req.Form{
        fmt.Printf("%s = %s\n", key, value)
    }

    p := User{Username: req.FormValue("username"), Password: req.FormValue("password")}
    // js, _ := json.Marshal(p)
    // w.Write(js)

    tmplt, err := template.ParseFiles("static/user.html")
    if err != nil {
        fmt.Println(err)
        return
    }    
    tmplt.Execute(w, p)
}

func signup(w http.ResponseWriter, req *http.Request) {
    req.ParseForm()

    fmt.Printf("USERNAME => %s\n", req.FormValue("username"))
    fmt.Printf("PASSWORD => %s\n", req.FormValue("password"))
    fmt.Printf("HIDDEN => %s\n", req.FormValue("hidden"))
    // fmt.Println("---------------------")
    // for key, value := range req.Form{
    //     fmt.Printf("%s = %s\n", key, value)
    // }
    http.Redirect(w, req, "/user", http.StatusSeeOther)
}

func indexHTMLTemplateVariableHandler(response http.ResponseWriter, request *http.Request) {
    // tmplt := template.New("IndexTemplated.html")       //create a new template with some name
    tmplt, err := template.ParseFiles("static/index.html") //parse some content and generate a template, which is an internal representation
    if err != nil {
        fmt.Println(err)
        return
    }
    // p := Student{Id: 1, Name: "Aisha"} //define an instance with required field
    tmplt.Execute(response, nil) //merge template ‘t’ with content of ‘p’
}

// func TemplatedHandler(response http.ResponseWriter, request *http.Request) {
//     tmplt := template.New("hello template")
//     tmplt, _ = tmplt.Parse("Top Student: {{.Id}} - {{.Name}}!")

//     p := Student{Id: 1, Name: "Aisha"} //define an instance with required field

//     tmplt.Execute(response, p) //merge template ‘t’ with content of ‘p’
// }

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
    r.HandleFunc("/hello", hello).Methods("GET")
    r.HandleFunc("/signup", signup).Methods("POST")
    r.HandleFunc("/user", user).Methods("GET")
    r.HandleFunc("/", indexHTMLTemplateVariableHandler).Methods("GET")
    r.HandleFunc("/number/{number:[0-9]+}", number).Methods("GET")
    r.HandleFunc("/number/{max:[0-9]+}/{number:[0-9]+}", number_max).Methods("GET")
    r.HandleFunc("/number/{min:[0-9]+}/{max:[0-9]+}/{number:[0-9]+}", number_min_max).Methods("GET")

    // r.HandleFunc("/prime/{max:[0-9]+}/{number:[0-9]+}", numnber).Methods("GET")

    fmt.Println("Server is listening...")
    http.Handle("/", r)
    http.ListenAndServe(":80", nil)
}