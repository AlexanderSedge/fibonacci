package main

import "fmt" 


func main() {
	
	fmt.Println("How many integers would you like from the Fibonacci sequence?")
	var count int
	fmt.Scanln(&count)
	arr := make([]int, count, count)
	if(count>0){
		arr[0] = 1
	} 
	if(count>1){
		arr[1] = 1
	} 
	fmt.Println("Press 1 to use array-based calculation (fast), press 2 to use a recursive function (inefficient, not recommended for sequences > 40ish).")
	var choice int
	fmt.Scanln(&choice)
	for (choice<1 || choice>2){  //if selection is out of bounds
		fmt.Println("Invalid selection. Please enter 1 or 2.")
		fmt.Scanln(&choice)
	}
	if(choice==1){ //array selected
		fmt.Println("Using simple array. Calculating: ")
		arrayfib(count, arr)
	}
	if(choice==2){ //recursive selected
		fmt.Println("Using recursive function. Calculating: ")
		recursefib(count, arr)
	}
	fmt.Println("Sequence:")
	for i:=0; i<count; i++ { //print by iterating over array
		fmt.Printf("%d ", arr[i])
	}
}

//RECURSEFIB//
//Generates a fibonacci sequence using recursive function calls.
//input: number of integers desired, plus a pre-initialized slice which functions as a pointer to our array
//output: none, but modifies the input array
//VERY inefficient but this is a technical skills test, and fibonacci is kinda the classic "recursive" problem, so it's included even if it's suboptimal
func recursefib(count int, arr []int) int{
	if(count==0){
		return 0
	} 
	if(count<3){
		return 1
	} 
	var sum int = recursefib(count-1, arr)+recursefib(count-2, arr)
	arr[count-1] = sum
	return sum
	
}

//ARRAYFIB//
//Generates a fibonacci sequence using a pointer to a pre-existing array.
//input: number of integers desired, plus a pre-initialized slice which functions as a pointer to our array
//output: none, but modifies the input array
func arrayfib(count int, arr []int) {
	//first two integers are always 1
	if(count==0){
		return
	} 
	if(count>2) {
		for j:=2; j<count; j++ {
			arr[j] = arr[j-1]+arr[j-2]
		}
		return
	}

}