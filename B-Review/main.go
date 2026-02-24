package main

import (

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// "fmt"

// func main() {

//  arr:= []int{1,2,3,4}

//  reverseArray(arr)

// }

// func reverseArray(arr []int){
//    left := 0
//    right := len(arr)-1

//   for left = 0 ; left < right ; left++{
//    arr[left],arr[right] = arr[right],arr[left]
//    right --
// }

//  fmt.Println(arr)

// }

// type Node struct{
//   Data int
//   Next *Node
// }

// type LinkedList struct{
//   Top *Node
// }

// func FindLength(pos1,pos2 int)int{

//  var arr LinkedList

//   curr := arr.Top
//   count := 0

//  for curr.Data != pos1{
//    curr = curr.Next
// }

//  for curr.Next.Data != pos2{
//    count ++
// }

//  return count

// }

// func ReverseLinkList(n LinkedList){

//   curr := n.Top
//   fir := n.Top
//  for curr.Next == nil{
//  curr = curr.Next
// }

// }

 type Claims struct{
 UserID uint   
 email string
 jwt.Claims
  
}

 var jwtsecret = []byte("Secret-Key")

func JwtMiddleware()gin.HandlerFunc{
 return func(ctx *gin.Context) {

  auth := ctx.GetHeader("Authorization")

 if auth == ""{
  ctx.JSON(401,gin.H{"error":"Please Login"})
  ctx.Abort()
  return 
}

 tokenstr := strings.TrimPrefix(auth,"Bearer ")
 claims := Claims{}


 token, err := jwt.ParseWithClaims(tokenstr,claims,func(t *jwt.Token) (interface{}, error) {
  return jwtsecret,nil
})
 
 if err!=nil || !token.Valid{
  ctx.JSON(401,gin.H{"error":"Invalid Token"})
  ctx.Abort()
  return 
}

 ctx.Set("email",claims.email)
 ctx.Set("user_id",claims.UserID)




}
}

