package main

import (
	"encoding/json"
	"net/http"
)

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type Message struct {
// 	Text string `json:"text"`
// }

// func hellloHandler(w http.ResponseWriter,r *http.Request){

//    if r.Method == http.MethodGet{

//    w.Write([]byte("HEllo From GEt request"))

//    return
// }

//    if r.Method==http.MethodPost{

//     var msg Message

//      err:= json.NewDecoder(r.Body).Decode(&msg)

//   if err!=nil{

//     http.Error(w,"Invalid Json",http.StatusBadRequest)
//     return
//  }

//  w.Header().Set("Content-type","application/json")
//  json.NewEncoder(w).Encode(msg)
//  return
// }

//  http.Error(w,"Invalid Method request",http.StatusMethodNotAllowed)
// }

// func main(){

//    http.HandleFunc("/hello",hellloHandler)
//    http.ListenAndServe(":8080",nil)
// }






// type helloResponse struct{
//    Message string `json:"message"`
// }


// func hellloHandler(r http.ResponseWriter , w *http.Request){
   
//     if w.Method != http.MethodGet{
  
//      http.Error(r,"Method Not Allowed",http.StatusMethodNotAllowed)  
//      return

// }
//    response := helloResponse{
   
//     Message: "Hello World", 
// }

// w.Header.Set("Content-Type","application/json")

// json.NewEncoder(r).Encode(response)
 
// }


// func main(){
 
//   http.HandleFunc("/api/hello",hellloHandler)
//   http.ListenAndServe(":8080",nil)

// }


func helloResponse(w http.ResponseWriter,  r *http.Request){
 
  if r.Method != http.MethodGet{
    http.Error(w,"Method Not allowed",http.StatusMethodNotAllowed)
   return
}  
  
 w.Header().Set("Content-Type","application/json")

  response := map[string]string{
   "message" : "Hello",
}

 json.NewEncoder(w).Encode(response)

}


func main(){
    
 http.HandleFunc("/hello",helloResponse)
 http.ListenAndServe(":8080",nil)

}
