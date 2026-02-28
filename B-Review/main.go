package main

// import "fmt"

// import (
// 	"runtime"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// )

// // "fmt"

// // func main() {

// //  arr:= []int{1,2,3,4}

// //  reverseArray(arr)

// // }

// // func reverseArray(arr []int){
// //    left := 0
// //    right := len(arr)-1

// //   for left = 0 ; left < right ; left++{
// //    arr[left],arr[right] = arr[right],arr[left]
// //    right --
// // }

// //  fmt.Println(arr)

// // }

// // type Node struct{
// //   Data int
// //   Next *Node
// // }

// // type LinkedList struct{
// //   Top *Node
// // }

// // func FindLength(pos1,pos2 int)int{

// //  var arr LinkedList

// //   curr := arr.Top
// //   count := 0

// //  for curr.Data != pos1{
// //    curr = curr.Next
// // }

// //  for curr.Next.Data != pos2{
// //    count ++
// // }

// //  return count

// // }

// // func ReverseLinkList(n LinkedList){

// //   curr := n.Top
// //   fir := n.Top
// //  for curr.Next == nil{
// //  curr = curr.Next
// // }

// // }

// type Claims struct {
// 	UserID uint
// 	email  string
// 	jwt.Claims
// }

// var jwtsecret = []byte("Secret-Key")

// func JwtMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 		auth := ctx.GetHeader("Authorization")

// 		if auth == "" {
// 			ctx.JSON(401, gin.H{"error": "Please Login"})
// 			ctx.Abort()
// 			return
// 		}

// 		tokenstr := strings.TrimPrefix(auth, "Bearer ")
// 		claims := Claims{}

// 		token, err := jwt.ParseWithClaims(tokenstr, claims, func(t *jwt.Token) (interface{}, error) {
// 			return jwtsecret, nil
// 		})

// 		if err != nil || !token.Valid {
// 			ctx.JSON(401, gin.H{"error": "Invalid Token"})
// 			ctx.Abort()
// 			return
// 		}

// 		ctx.Set("email", claims.email)
// 		ctx.Set("user_id", claims.UserID)

// 	}
// }

// func main() {
// 	runtime.GOMAXPROCS(1)
// }

// func Print[T any](v T) {
// 	fmt.Println(v)
// }

// func main() {
// 	Print(1)

// 	Print("Hello World")
//     Print(true)

// }

type Node struct {
	Data int
	Next *Node
}

func Reverse(head *Node) *Node {
	var prev *Node
	curr := head

	//  1->2->3->4->nil

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev

}

func Rever(head *Node) *Node {

	var prev *Node
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev

}

func ReverseFirstK(head *Node, k int) *Node {

	var prev *Node

	curr := head
	count := 0

	for curr != nil && count < k {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		count++
	}

	head.Next = curr

	return prev

}

func ReverseBetween(head *Node, x, y int) *Node {
	if head == nil || x == y {
		return head
	}

	dummy := &Node{Next: head}
	prev := dummy

	for i := 1; i < x; i++ {
		prev = prev.Next
	}

	start := prev.Next
	curr := start.Next

	for i := 0; i < x-y; i++ {
		start.Next = curr.Next
		curr.Next = prev.Next
		prev.Next = curr
		curr = start.Next
	}

	return dummy.Next
}
