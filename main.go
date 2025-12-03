// import main
package main

import (
	"fmt"

)







// func main(){
//   fmt.Println("Hello World")
//    runtime.GOMAXPROCS(3)
// //  str := "Hello 👋  World"
// for i,x:= range "str👋 ing "{
//    fmt.Println(string(x))
//    fmt.Println(i)
// }

// x:="hello world"

// var a interface{} = nil
// fmt.Println(a == nil)

// x := new(int)
  

// defer fmt.Println("A")
// defer fmt.Println("B")
// 

// a := []int{1,2,3}
// b := a[:2]
// b = append(b, 100)
// fmt.Println(a, b)


//   go func(){
//    fmt.Println("Hello")
    
// }()

// time.Sleep( 2*time.Millisecond)


//  var mu sync.Mutex
//  var wg sync.WaitGroup
 
//  count:= 0 
//  wg.Add(100)
//   for i:=0 ; i<100 ; i++{
  
//     go func(){
//     mu.Lock()
//     count++
//     mu.Unlock()
//     wg.Done()

// } ()
// }
//   wg.Wait()
//   fmt.Println(count)

// type student struct{
//    name string
//    score int
// }

//    switch {
//    case score>80 :
//       fmt.Println("A")
//     case score > 60:
//       fmt.Println("B")
//     default :
//       fmt.Println("invalid")
// }

//  func (s student) float64,error {
// 	result := value,err

//   if err!= nil{
//    fmt.Errorf("error")
// }
//  }()

//  fmt.Scanf(student)
// }



type student struct{

  name string
  score int
}

func (s student)findError()(int,error){
   if s.score<0 ||s.score>100{
      return 0 ,fmt.Errorf("invalid nmbr")
}
   
    return s.score , nil
}


func main(){
   
    var s student 
  
    fmt.Println("Enter name")
    fmt.Scanln(&s.name)
     
    fmt.Println("Enter marks")
    fmt.Scanln(&s.score)


    switch {
    
     case s.score > 80:
      fmt.Println("A")
     case s.score > 60 :
      fmt.Println("B")
      default :
      fmt.Println("Low")    

}  


   value,err := s.findError()
 
   if err!= nil {
	fmt.Println("Error :" ,err)
   }


  fmt.Println(value)


 

 
}