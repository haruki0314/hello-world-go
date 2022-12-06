package main


import "fmt"

func anythig(a interface{})

func Plus(x int,y int) int{
	return x+y
}
func Div(x,y int)(int ,int){
	q := x/y;
	a := x%y;
	return q,a;
}
func Double(price int)(result int){
	result = price * 2
	return 
}
func CallFuntion(f func()){
	f()
}
func Later() func(string) string{
	var store string 
	return func(next string) string{
		s := store
		store = next
		return s
	}
}

func integers() func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}

func main(){
	var x12 interface{} = 3
	i1c2 := x12.(int)
	fmt.Println(i1c2)


	i:=Plus(1,2)
	fmt.Println(i)
	fmt.Println(Div(1,2));
	f := func(x int,y int) int{
		return x+y;
	}(1,2)
	fmt.Println(f)
	
	//関数を引数にとる関数
	CallFuntion(func(){
		fmt.Println("Im a funciton")
		})
	
	funcArg := Later()
	fmt.Println(funcArg("1a"))
	fmt.Println(funcArg("2a"))
	fmt.Println(funcArg("3a"))
	fmt.Println(funcArg("4a"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	a := 0
	if a==2{
		fmt.Println("two")
	}else{
		fmt.Println("one")
	}

	x:=0
	if x:=2; x==2{
		fmt.Println(x)
	}
	fmt.Println(x)

	i=0
	for {
		i++;
		if i == 3{
			break
		}
		fmt.Println(i)
	}
	/*while文*/
	point := 0
	for point <10 {
		fmt.Println(point)
		point++
	}

	for forArgs := 0; forArgs< 10; forArgs++ {
		if forArgs == 3 {
			continue
		}
		if forArgs== 6 {
			break
		}
		fmt.Println(forArgs)
	}

	arrArgsOutput := [3]int{1,2,3}
	for i1a := 0; i1a < len(arrArgsOutput); i1a ++ {
		fmt.Println(arrArgsOutput[i1a])
	}

	ArrArgsOurtput2 := [3]int{1,2,3}
	for i1b,v1b := range ArrArgsOurtput2 {
		fmt.Println(i1b,v1b)
	}

	sl := []string{"Python","PHP","GO"}
	for ic1,vc3:= range sl {
		fmt.Println(ic1,vc3)
	}

	var m1 map[string]int =map[string]int{"apple":100,"banana":200}
	for k,v := range m1{
		fmt.Println(k,v)
	}

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

	x
}