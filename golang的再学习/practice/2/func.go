package main

import "fmt"

func a() (int) {
	fmt.Println("Hello, world")
	return 1
}

func b(a, b int) (sum int){
	sum = a + b
	return
}

func add(a, b int) (sum int){
    anonymous := func(x, y int) (int){
        return x + y
    }
    return anonymous(a, b)
}

func swap(a, b int) (int,ret int){
    return a + b, ret
}

func sumList(arr ...int) (sum int){
	for _, v := range arr{
		sum += v
	}

	return sum
}

func sumLista(arr []int) (sum int){
	for _, v := range arr{
		sum += v
	}

	return sum
}

func main() (){
	ret := a()
	fmt.Println(ret)

	ret = b(1, 2)
	fmt.Println(ret)

	ret = add(2, 3)
	fmt.Println(ret)

	ret, _ = swap(3, 4)
	fmt.Println(ret)
	fmt.Println(swap(3, 4))

	ret = sumList(1, 2, 3, 4)
	fmt.Println(ret)

//	arr := [...]{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4}
	ret = sumList(slice...)
	fmt.Println(ret)
//	ret = sumList(arr...)
//	fmt.Println(ret)

	fmt.Printf("%T\n", sumList);
	fmt.Printf("%T\n", sumLista);

}
