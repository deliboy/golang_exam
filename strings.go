package main

import "fmt"

var (
	start= []rune("가")[0]
	end = []rune("힣")[0]
)

//받침이 있는 문자인지 확인
func HasConsonantSuffix(s string) bool{
	numEnds :=28
	result :=false

	for _, r:= range s{
		if start <=r && r <end{
			index :=int(r-start)
			result = index%numEnds !=0
		}
	}

	return result;
}
func main() {

	

	var f = HasConsonantSuffix("간")
	fmt.Println(f);


	// for i, r:= range "나낙낚"{
	// 	fmt.Println(i, r)
	// }
	 for i:=start; i<=end; i++{
	 	fmt.Println(i,string(i))
	 }
//
	fmt.Println("가나다")

}
