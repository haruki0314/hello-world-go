package main


import (
	// "time"
	"fmt"

	// "os"
)

// type T struct{
// 	User User
// }

// type User struct {
// 	//構造体定義 
// 	Name string
// 	Age int
// 	// X,Y int
// }
// type Users []*User
// //構造体のポインタ変数を返すコンストラクタ関数をgoではよく扱う
// func NewUser(/*引数*/name string,age int)/*ポインタ型の引数:*/ *User{
// 	return &User{Name:name,Age:age}
// }

// type Person struct{
// 	Name string
// 	Age int
// }

// type Car struct{
// 	Number string
// 	Model string
// }

// func (p *Person) ToString() string{
// 	 return fmt.Sprintf("Name=%v,Age=%v",p.Name,p.Age)
// }
// func(c *Car) ToString() string{
// 	return fmt.Sprintf("Number=%v,Model=%v",c.Number,c.Model)
// }

// type Stringfy interface{
// 	ToString() string
// }

// func updateUser2(user *User){
// 	user.Name = "A"
// 	user.Age = 1000
// }

// func (u User)SayName(){
// 	fmt.Println(u.Name)
// }


// func (u User)SetName1(name string){
// 	u.Name = name
// }

//基本的にメソッドに定義するレシーバはポインタ型にするべき
// func (u *User)SetName2(name string){
// 	u.Name = name
// }

// func Plus(x int,y int) int{
// 	return x+y
// }
// func Div(x,y int)(int ,int){
// 	q := x/y;
// 	a := x%y;
// 	return q,a;
// }
// func Double(price int)(result int){
// 	result = price * 2
// 	return 
// }
// func CallFuntion(f func()){
// 	f()
// }
// func Later() func(string) string{
// 	var store string 
// 	return func(next string) string{
// 		s := store
// 		store = next
// 		return s
// 	}
// }

// func integers() func() int{
// 	i := 0
// 	return func() int{
// 		i++
// 		return i
// 	}
// }


// func TestDefer(){
// 	defer fmt.Println("end")
// 	fmt.Println("start")
// }

// func sub(){
// 	for{
// 		fmt.Println("Sub Loop")
// 		time.Sleep(100*time.Millisecond)
// 	}
// }

// func Sum(s ...int) int{
// 	a:=0
// 	for _,v := range s{
// 		fmt.Println(a)
// 		a += v
// 	}
// 	return a
// }

// func Sum(s ...int) int{
// 	n:=0
// 	for _,v:=range s{
// 		n+= v
// 	}
// 	return n
// }

// func receiever(c chan int){
// 	for{
// 		i :=<-c
// 		fmt.Println(i)
// 	}
// }

// func (u *User) SetName(){
// 	u.Name = "A"
// }

// type MyInt int

type error interface{
	Error()string
}
type MyError struct{
	Message string
	ErrCode int
}

func(e *MyError) Error() string{
	return e.Message
}

func RaiseError() error{
	return &MyError{Message:"カスタムエラーが発生しました。",ErrCode:1}
}


func main(){
	// var x12 interface{} = 3
	// i1c2 := x12.(int)
	// fmt.Println(i1c2)
	// // かたに合わせて復元した値のデータ型に応じて処理を変更できる
	// if x12 == nil {
	// 	fmt.Println("none")
	// }  else if i1a2,isInt := x12.(int); isInt{
	// 		fmt.Println(i1a2 ,"int")
	// }

	// switch x12.(type){
	// case int:
	// 	fmt.Println("int")
	// case string:
	// 	fmt.Println("string")

	// }

	// i:=Plus(1,2)
	// fmt.Println(i)
	// fmt.Println(Div(1,2));
	// f := func(x int,y int) int{
	// 	return x+y;
	// }(1,2)
	// fmt.Println(f)
	
	// //関数を引数にとる関数
	// CallFuntion(func(){
	// 	fmt.Println("Im a funciton")
	// 	})
	
	// funcArg := Later()
	// fmt.Println(funcArg("1a"))
	// fmt.Println(funcArg("2a"))
	// fmt.Println(funcArg("3a"))
	// fmt.Println(funcArg("4a"))

	// ints := integers()
	// fmt.Println(ints())
	// fmt.Println(ints())
	// fmt.Println(ints())
	// fmt.Println(ints())

	// a := 0
	// if a==2{
	// 	fmt.Println("two")
	// }else{
	// 	fmt.Println("one")
	// }

	// x:=0
	// if x:=2; x==2{
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	// i=0
	// for {
	// 	i++;
	// 	if i == 3{
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// /*while文*/
	// point := 0
	// for point <10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// for forArgs := 0; forArgs< 10; forArgs++ {
	// 	if forArgs == 3 {
	// 		continue
	// 	}
	// 	if forArgs== 6 {
	// 		break
	// 	}
	// 	fmt.Println(forArgs)
	// }

	// arrArgsOutput := [3]int{1,2,3}
	// for i1a := 0; i1a < len(arrArgsOutput); i1a ++ {
	// 	fmt.Println(arrArgsOutput[i1a])
	// }

	// ArrArgsOurtput2 := [3]int{1,2,3}
	// for i1b,v1b := range ArrArgsOurtput2 {
	// 	fmt.Println(i1b,v1b)
	// }

	// sl := []string{"Python","PHP","GO"}
	// for ic1,vc3:= range sl {
	// 	fmt.Println(ic1,vc3)
	// }

	// var m1 map[string]int =map[string]int{"apple":100,"banana":200}
	// for k,v := range m1{
	// 	fmt.Println(k,v)
	// }

	// switch n {
	// case 1,2:
	// 	fmt.Println("1,2")
	// case 3,4:
	// 	fmt.Println("3,4")
	// default:
	// 	fmt.Println("idont know")
		
	// }

	// switch n:=2; n{
	// case 1,2:
	// 	fmt.Println("1,2")
	// default:
	// 	fmt.Println("idont know")
	// }

	// n:=6
	// switch n>0 && n<4:

	// Loop:
	// for{
	// 	for{
	// 		fmt.Println("a")
	// 		break Loop
	// 	}
	// }
	// go sub()
	// for{
	// 	fmt.Println("Main Loop")
	// 	time.Sleep(200*time.Millisecond)
	// }
	// TestDefer()

	// defer func(){
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()

	// file,err := os.Create("test.txt")
	// if err != nil{
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	
	// file.Write([]byte("Hello"))

	// sl5:= []int{1,2,3,4,5}
	// fmt.Println(sl5)

	// sl2 := sl5
	// sl2 [0] = 1000
	// fmt.Println(sl2)
	// //元のアドレスにも1000がコピーされる

	// sl6 := []int{1,2,3,4,5}
	// sl3 := make([]int,5,10)
	// fmt.Println(sl6)

	// n := copy(sl3,sl6)

	// fmt.Println(n,sl3)

	// sl := []string{"A","B","C"}
	// fmt.Println(sl)

	// for _, v:= range sl{
	// 	fmt.Println(v,sl)
	// }

	// fmt.Println(Sum(1,2,3))

	// m := map[int]string{
	// 	1:"A",
	// 	2:"B",
	// }
	// fmt.Println(m)

	// m[1] = "JAPAN"
	// m[2] = "USA"

	// fmt.Println(m[2])

	// s , ok:=m[5]
	// if !ok {fmt.Println("error")}
	// fmt.Println(s,ok)
	
	// //mapから要素を削除
	// delete(m,1)
	// fmt.Println(m)

	// sl := []string{"A","B","C"}
	// fmt.Println(sl)
	
	// 	for i := range sl{
	// 		fmt.Println(sl[i])
	// 	}

	//スライス
	//可変長引数

	// fmt.Println(Sum(1,2,3))

	// var m = map[string]int{"A":100,"B":200}
	// fmt.Println(m)
	// fmt.Println(m["A"]);

	// m4:=make(map[int]string)
	// m5 , ok:= m4[1]
	// if !ok {
	// 	fmt.Println("erro")
	// }
	// 	fmt.Println(ok,m5)
	// fmt.Println(m4)

	// m:= map[string]int{
	// 	"Apple":100,
	// 	"Banan":200,
	// }
	
	// for k,v:=range m{
	// 	fmt.Println(k,v)
	// }
	
	// var ch1 chan int
	
	// ch1 = make(chan int)
	// ch2 := make(chan int)

	// fmt.Println(cap(ch1))
	// fmt.Println(cap(ch2))

	// ch3 := make(chan int,5)
	// fmt.Println(cap(ch3))

	// ch3 <- 1
	// fmt.Println(len(ch3))

	// ch3 <- 2
	// ch3 <- 3
	// fmt.Println(len(ch3))

	// i := <-ch3
	// fmt.Println(i)

	// i2 := <-ch3
	// fmt.Println(i2)
	// fmt.Println("len",len(ch3))

	// ch1:= make(chan int)
	// ch2:= make(chan int)

	// go receiever(ch1)
	// go receiever(ch2)

	// i :=0

	// for i<100{
	// 	ch1<- i
	// 	ch2 <- i

	// 	time.Sleep(50*time.Millisecond)
	// 	i++
	// }
	// fmt.Println(<-ch1)

	// user1 := User{Name:"user1"}
	// user1.SayName()

	// user2 := &User{Name:"user2"}
	// user2.serName2("8"
	// t := T{User:User{Name:"user1",Age:10}}
	// fmt.Println(t)

	// fmt.Println(t.User)
	// fmt.Println(t.User.Name)

	// t.User.SetName()
	// fmt.Println(t)
	// user1 := NewUser("user1",10)
	// fmt.Println(user1)
	// fmt.Println(*user1)

	// users := Users{}
/*user型のスライスを上で構造体として定義、スライスにuser型のオブジェクトを積んでいく*/
	// user1 := User{Name:"user1",Age:10}
	// user2 := User{Name:"user2",Age:20}
	// user3 := User{Name:"user3",Age:30}
	// user4 := User{Name:"user4",Age:40}

	// users := Users{}

	// users = append(users,&user1)
	// users = append(users,&user2)
	// users = append(users,&user3)
	// users = append(users,&user4)

/*for文で展開していく*/
	// for _,u := range users{
	// 	fmt.Println(*u)
	// }
	// fmt.Println(users)


/*make関数でuser型のスライスを作成していく*/
	// users2 := make([]*User,0)
	// users2 = append(users2,&user1,&user2)

	// for _,u := range users2{
	// 	fmt.Println(*u)
	// }

	// m:= map[int]User{
	// 	1:{Name:"user1",Age:10},
	// 	2:{Name:"user2",Age:20},
	// }
	// fmt.Println(m)

	// m2 := map[User]string{
	// 	{Name:"user1",Age:10}:"Tokyo",
	// 	{Name:"user2",Age:20}:"LA",
	// }
	// fmt.Println(m2)

	// m3 := make(map[int]User)
	// fmt.Println(m3)
	// m3[1] = User{Name: "user3"}
	// fmt.Println(m3)

	// for _,v := range m{
	// 	fmt.Println(v)
	// }

		// vs := []Stringfy{
		// 	&Person{Name:"Taro",Age:21},
		// 	&Car{Number:"123-456",Model:"AB-1234"},
		// }

		// for _,v := range vs{
		// 	fmt.Println(v.ToString())
		// }

			err := RaiseError()
			fmt.Println(err.Error())

			e,ok := err.(*MyError)
			if ok{
				fmt.Println(e.ErrCode)
			}
}