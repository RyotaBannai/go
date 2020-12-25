package main

/*
#include <stdio.h>
#include <math.h>
#define MSG "Hello"
#define ADD(a,b) (a + b)
void Cprint(char* str){
	printf("%s", str);
}

double MyPow(double x, double y){
	double ret= pow(x,y);
	return ret;
}

int Add(int a, int b){
	int ret = ADD(a, b);
	return ret;
}
*/
import "C"
import "fmt"

func main() {
	str := C.CString("hoge\n")
	// CString: golang の String 型を C 言語の char* に変換
	C.Cprint(str)

	ret := C.MyPow(10, 2)
	fmt.Println(ret)

	// マクロ
	res := C.MSG
	fmt.Println(res)
	res2 := C.Add(10, 10)
	fmt.Println(res2)

	callC()
}
