package main

import "fmt"

func main(){
	var bit8 int8 = 255
	var bit16 int16 = 65535
	var bit32 int32 = 42924967295
	var bit64 int64  // banyak

	// Contoh Expresi Aritmatika Aritmatika
	var op1,op2 int = 3,5
	fmt.Printf(" %d + %d = %d \n", op1, op2,op1 + op2) // 3 + 5 = 8
	fmt.Printf(" %d - %d = %d \n", op1, op2,op1 – op2) // 3 – 5 = -2
	fmt.Printf(" %d * %d = %d \n", op1, op2,op1 * op2) // 3 * 5 = 15
	fmt.Printf(" %d / %d = %f \n", op1, op2,op1 / op2) // 3 / 5 = 0,66667
	fmt.Printf(" %d % %d = %d \n", op1, op2,op1 % op2) // 3 % 5 = 2

	// Contoh Expressi Relasional

	var int1,int2 int = 21,40
    var str1,str2 string = “Fransiska”,”fransisko”
    var bool1, bool2 bool = true,false

    //Pencocokan
    fmt.Println(bool1 == bool1) // true
    fmt.Println(str1 == str2) // false
    fmt.Println(int1 == int2) // false

    // Pentidakcocokan
    fmt.Println(int1 != int2) // true
    fmt.Println(str1 != str2) // true
    fmt.Println(bool1 != bool2) // true
    
    fmt.Println(int1 > int2) // 21 > 40 , false
    fmt.Println(int1 < int2) // 21 < 40 , true

    fmt.Println(int1 <= int2) // 21 <= 40 , true
    fmt.Println(int1 >= int2) // 21 >= 40 , true 

}