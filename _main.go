package _main

import{
	"log"
	"fmt"
	"net/http"
	"math/rand"
}

type books struct{
	Name string `json:"name"`
	SerialNo string `json:"serial_no"`
	Publication string `json:"publication"`
	Author *Details `json:"author"`
}

type Details struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age string `json:"age"`
}
var Books []Book 

func main(){

	Books=append(Books,books{Name:"the night at the train",SerialNo:"67353636",Publication:"S chand",&Author{FirstName:"lakhmir chand",LastName:"manjit chand",Age:"28"}})
	Books=append(Books,books{Name:"Helen Keler",SerialNo:"60082616",Publication:"murli mm",&Author{FirstName:"balveer ",LastName:"manjit gangadhar",Age:"28"}})

	r:=mux.NewRouter()
	r.HandleFunc("/Books",getBooks).Methods("GET")
	r.HandleFunc("/Book/{id}",getBookById).Methods("GET")
	r.HandleFunc("/Books",createBookLibrary).Methods("POST")
	r.HandleFunc("/Book/{id}", updateBookLibrary).Methods("UPDATE")
	r.HandleFunc("/Book/{id}",deleteBookFromLibrary).Methods("DELETE")

	fmt.Printf("server started at port:8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

}

func getBooks(w http.ResponseWriter , r *http.Request){
	w.Header.Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)
}

func getBookById(w http.ResponseWriter r *http.Request){
	w.Header.Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)
	params:=mux.Vars(r)
	for _,item:=range books{
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createBookLibrary(w http.ResponseWriter,r *http.Request){
	Header.Set("Content-Type","application/json")
	var book []Book
	_:=json.NewDecoder(r.body).Decode(&book)
	book.SerialNo=strconv.Itoa(rand.Intn(1000000))
	Books=append(Books,Book)
	json.NewEncoder(w).Encode(Books)
}

func deleteBookFromLibrary(w http.ResponseWriter, r *http.Request){
	w.Header.Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item:=range params{
		if item["serial_no"]==books.SerialNo{
			Books=append(Books[:index],Books[index+1:]...)
			json.NewEncoder(w).Encode(Books)
			return 
		}
	}
	
}

func updateBookLibrary(w http.ResponseWriter , r *http.Request){
	Header.Set("Content-Type","application/json")
	params:=mux.Vars(r)

	for index,item:=range params{
		if item["serial_no"]==books.SerialNo{
			Books=append(Books[:index],Books[index+1:]...)
			var book Book
			_:=json.NewDecoder(r.body).Decode(&book)
			Books=append(Books,book)
			json.NewEncoder(w).Encode(Books)
			return 
		}
	}
}