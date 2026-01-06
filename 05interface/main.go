package main

import 	"fmt"


type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width,height float64
}


func (c Circle) area() float64{
    return 3.14 * c.radius*c.radius
}

func (c Circle) perimeter() float64{

     return 2*3.14*c.radius 
 }

func (r Rectangle) area() float64{

  return  r.height*r.width
}
func (r Rectangle) perimeter() float64{

   return 2 * (r.height+r.width)
}

func printInfo(s Shape){
  
   fmt.Println("Area : ",s.area())
   fmt.Println("perimetrer : ",s.perimeter())
}

func main() {
    c := Circle{radius: 5}
    r := Rectangle{width: 4, height: 6}

    printInfo(c)
    printInfo(r)
}

