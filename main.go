package main

import "html/template"
import "net/http"
import "log"
import "time"

type PageData struct {
	Time      time.Time
	UserInput []string
	Cars      []ACar
}

type ACar struct {
	Model string
}

var cars = []ACar{}
var userinput []string

func fetchCars() []ACar {
	var fetchedCars = []ACar{}
	var sliceofCars = []string{"Honda", "Toyota", "Porsche", "Acura"}
	for _, car := range sliceofCars {
		fetchedCars = append(fetchedCars, ACar{car})
	}
	return fetchedCars

}

func updateUI(rw http.ResponseWriter, req *http.Request) {
	start := time.Now() // used to track time to perform request

	req.ParseForm()                                                    // parse the data input by the user
	userinput = req.Form["cars"]                                       // specifically find the cars form data
	log.Printf("[updateUI] User form selection: %s", req.Form["cars"]) // print the data to the console for debuggin

	log.Printf("[updateUI] data served in %s\n", time.Since(start)) //print time to serve to console
	serveUI(rw, req)                                                // not sure if this is the right way to do this (it leaves the url in the browser bar)
}

func serveUI(rw http.ResponseWriter, req *http.Request) {
	start := time.Now() // used to track time to perform request
	cars := fetchCars() // get the data we'll use on the drop down select form

	tmpl, err := template.ParseFiles("templates/index.html") // parse the template

	if err != nil {
		log.Panic(err)
	}

	page := PageData{Time: time.Now(), UserInput: userinput, Cars: cars} // build the data to be used in the template
	err = tmpl.Execute(rw, page)                                         // apply the page data to the template

	if err != nil {
		log.Panic(err)
	}

	log.Printf("[serveUI] data served in %s\n", time.Since(start)) //print time to serve to console
}

func main() {
	serveAddress := ":8080"
	rMux := http.NewServeMux()
	rMux.HandleFunc("/", serveUI)
	rMux.HandleFunc("/action_page", updateUI)
	log.Println(http.ListenAndServe(serveAddress, rMux))
}
