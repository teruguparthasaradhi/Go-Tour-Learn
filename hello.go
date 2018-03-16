/*
Written by:Parthasaradhi Terugu
Practice of go from #tourofgo
<<<<<<< HEAD

pushing this into git 13th march 2018 5:19 PM
*/


package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

//func add(x int, y int )int {
/*we can initialize the variales
  single or multiple
*/
func add(x, y int) int {
	return x + y
}

/*
   return types are declared in second bracket
*/
func swap(s1, s2 string) (string, string) {
	return s2, s1
}

/*
   even we can return the variable without mentioning after return
   by mentioning in second bracket
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/*
   ----- var c, python,java bool------- all the variables are intialized to false as boolean type
   -----var c, python ,jave=true,false,"no!"---- if there is an intializer then the type can be ommited, the variable will
   take the type of the initializer
*/
var c, python, java = true, false, "no!"

/*
Go basic types are like
bool, string,
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte //alias for uint8
rune //alias for int 32
	represents a Unicode code point
float32 float64
complex64 complex128
*/
var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
	/*Here we are shifting 1 to 64 times left
	  a<<b = a*(2^b)
	  like take 64 bits
	  0000000000000..-64-..0000000001 after 1<<64
	  1000000000000..-64-..0000000000
	*/
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

/*
In func zero of different data types as intialized to nothing
they have default like 0 for numberic types
false for boolean type
"" for empty strings
*/
func zero() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

/*
var i int =42
var f float64= float64(i)
var u uint=uint(f)
this is written simply in function converion
unlilke in c, in Go assignments between items of different
types requires an explicit conversion
var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	even after removing the float64 or unit we get same results
Unlike in C, in Go assignment between items of
 different type requires an explicit conversion
*/
func conversion() {
	i := 42
	f := float64(i)
	u := uint(f)
	u = u + 1
	var x, y int = 3, 4
	var f1 = math.Sqrt(float64(x*x + y*y))
	var z = uint(f1)
	fmt.Println(x, y, z)
}
func typeinference() {
	var i int
	j := i           //j is also an int
	j = j + 1        //to omit the declaration error so that i dont get j is not used
	f := 3.141       //this will be float
	g := 0.86 + 0.5i //this will we be complex type
	fmt.Printf("%T %T %T %T\n", i, j, f, g)
}

/*
Constants are declared like variables, but with const key word
constants can be string,character,boolean and numeric values
constants cannot be declared using the syntax ":="
*/
const Pi = 3.141

func constant() {
	const name = "sherlock"
	fmt.Println("Happy birthday", name)
	fmt.Println(Pi, "which is declared as constant")
}

const (
	Big = 1 << 100
	//1 is shifted to left by 100 times,
	// 1 followed by 100 zeros
	Small = Big >> 99
	//1 is shifter to right by 99 times,
	//10 in binary 2 in decimal
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
func numericconst() {
	fmt.Println("Small, Int,float= ", Small, needInt(Small), needFloat(Small))
	fmt.Println("Big,float=", float64(Big), needFloat(Big))
}
func forloop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
}

/*
for continued loop
The init and post statement are optional.
sum := 1
	for ; sum < 1000; {
		sum += sum
	}
*/
func forwhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}

}

//infinite loop condition
func forever() {
	for {

	}
}

//this function always return a positive value
func ifloop(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func ifshort(x float64) float64 {
	if v := math.Pow(x, 2); v > x*2 {
		return v + 1
	}
	return x * 2
}
func ifelse(x float64) float64 {
	if v := math.Pow(x, 2); v > x*2 {
		return v + 1

	} else {
		fmt.Printf("%g < %g ", v, x*2)
	}
	return x
}

/*
Go's switch is like the one in C, C++, Java, JavaScript, and PHP,
 except that Go only runs the selected case,
 not all the cases that follow. In effect,
 the break statement that is needed at the end
of each case in those languages is provided
 automatically in Go.
 Another important difference is that
 Go's switch cases need not be constants,
 and the values involved need not be integers.
*/
func switchcase() {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {

	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println(" linux")
	default:
		fmt.Printf(" %s-default", os)
	}

}
func weekday() {
	today := time.Now().Weekday()
	switch time.Sunday {

	case today + 0:
		fmt.Println("weekend")
	case today - 1:
		fmt.Println("weekend")
	default:
		fmt.Printf("Its working day")
	}
}
func datetime() {
	timeNow := time.Now()
	switch {
	case timeNow.Hour() < 12:
		fmt.Println("Good Morining!")
	case timeNow.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

/*
A defer statement defers the execution
 of a function until
the surrounding function returns.

The deferred call's arguments are evaluated
immediately, but the function call is not
 executed until the surrounding function returns.

 Deferred function calls are pushed onto
  a stack. When a function returns, its deferred
  calls are executed in last-in-first-out order.
*/

func defergo() {
	defer fmt.Println("this should be printed last")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //this should print numbers in reverse order
	}
}
func pointers() {
	i, j := 30, 99
	p := &i
	fmt.Println(*p) //read i through pointer
	*p = 33         //set i through the pointer
	fmt.Println(i)
	q := &j
	*q = *q / 33   //divide through pointer
	fmt.Println(j) //chek new vauew of j
}

// A truct is colletions of fields
type Vertex struct {
	X int
	Y int
}

func Structfileds() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4 //struct fields are accesses using a dot
	fmt.Println(v.X)
	p := &v
	p.X = 1e9
	fmt.Println(v)
	/*
			Struct Literals
		A struct literal denotes a newly allocated struct
		value by listing the values of its fields.
		You can list just a subset of fields by using the
		Name: syntax. (And the order of named fields is
			 irrelevant.)
		The special prefix & returns a pointer to the
		 struct value.*/
	var (
		v1 = Vertex{1, 2}  //has type vertex
		v2 = Vertex{X: 1}  //y:0 is implicit
		v3 = Vertex{}      //X:0 and Y:0
		p1 = &Vertex{1, 2} // has type * Vertex

	)
	fmt.Println(v1, v2, v3, p1)
}

/*
The type [n]T is an array of n values of type T.

The expression

var a [10]int
declares a variable a as an array of ten integers.
*/
func Arrays() {
	var arr [10]string
	arr[0] = "Sherlock"
	arr[1] = "Holmes"
	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	/*The type []T is a slice with elements of type T.
	A slice is formed by specifying two indices,
	 a low and high bound, separated by a colon:

	a[low : high]
	This selects a half-open range which
	 includes the first element,
	 but excludes the last one.
	*/
	var s []int = primes[1:4]
	fmt.Println(s)
	/*
	   Slices are like references to arrays
	   A slice does not store any data, it just describes
	    a section of an underlying array.

	   Changing the elements of a slice modifies the
	   corresponding elements of its underlying array.

	   Other slices that share the same underlying
	    array will see those changes.
	*/
	names := [4]string{
		"sherlock",
		"lock",
		"holmes",
		"watson",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "parthu" //this change applies to all
	// since array shares the reference
	fmt.Println(names)
	fmt.Println(a, b)
	fmt.Println(a)
	fmt.Println(b)
}

/*
Slice literals
A slice literal is like an array literal without
the length.
This is an array literal:

[3]bool{true, true, false}

And this creates the same array as above,
 then builds a slice that references it:

[]bool{true, true, false}
*/

func SliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
	}
	fmt.Println(s)
}

/*
		For the array

	var a [10]int
	these slice expressions are equivalent:

	a[0:10]
	a[:10]
	a[0:]
	a[:]
*/
func SliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)
	s = s[1:5]
	fmt.Println(s)
	s = s[:3]
	fmt.Println(s)
	s = s[0:2]
	fmt.Println(s)

}

/*
A slice has both a length and a capacity.

The length of a slice is the number of
elements it contains.

The capacity of a slice is the number of
elements in the underlying array, counting
from the first element in the slice.

The length and capacity of a slice s can
 be obtained using the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it,
provided it has sufficient capacity.
Try changing one of the slice operations in the
example program to extend it beyond its capacity
 and see what happens.


*/
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*check description above printslice...*/
func SliceLengthCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]
	s = s[:0]
	printSlice(s) //len=0 cap=6 []
	s = s[:4]
	printSlice(s) //len=4 cap=6 [2 3 5 7]
	s = s[2:]
	printSlice(s) //len=2 cap=4 [5 7]
	s = s[2:]
	printSlice(s) //len=0 cap=2 []
	//	s = s[2:]
	//	printSlice(s)
	//panic: runtime error: slice bounds out of range

	/*Nill Slices*/
	var s1 []int
	printSlice(s1) //len=0 cap=0 []
	if s1 == nil {
		fmt.Println("nil!")
	}

}

/*
Creating a slice with make
Slices can be created with the built-in make function;
 this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and
returns a slice that refers to that array:

a := make([]int, 5)  // len(a)=5
To specify a capacity, pass a third argument to make:

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
*/
func SlicesWithMake() {
	a := make([]int, 5)
	printSlice(a)
	b := make([]int, 0, 5)
	printSlice(b)
	c := b[:2]
	printSlice(c)
	d := c[2:5]
	printSlice(d)

}
func SlicesOfSlices() {
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		//fmt.Printf("%s\n",Strings.Join(board[i]," ")
		fmt.Printf("%s\n", strings.Join(board[i], " "))


	}
}

/*
func append(s []T, vs ...T) []T
The first parameter s of append is a slice of type T,
 and the rest are T values to append to the slice.

The resulting value of append is a slice containing
all the elements of the original slice plus the
provided values.

If the backing array of s is too small to fit all
the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
*/
func SliceAppend() {
	var s []int
	printSlice(s)
	s = append(s, 0) //append works on nil
	printSlice(s)
	s = append(s, 1) // slice grows as needed
	printSlice(s)
	s = append(s, 2, 3, 4, 5) //we can add more than one element at a time
	printSlice(s)

}

/*
The range form of the for loop iterates over a
 slice or map.

When ranging over a slice, two values are returned for
 each iteration.
The first is the index, and the second is a copy of
the element at that index.
*/
func RangeGo() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
	for i, v := range pow {

		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/*
Range continued
You can skip the index or value by assigning to _.

If you only want the index, drop the
 ", value" entirely.
*/
func RangeCont() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // ==2^i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}

/*
imported pic in the top import section
Implement Pic. It should return a slice of length dy,
each element of which is a slice of dx 8-bit unsigned
 integers. When you run the program, it will display
 your picture, interpreting the integers as grayscale
 (well, bluescale) values.

The choice of image is up to you. Interesting
 functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8
	inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)
*/
/*
func SlicesExercies(dx, dy int) [][]uint8 {
	a := make([][]uint8, dx)
	for i := 0; i < dx; i++ {
		a[i] = make([]uint8, dx)
	}
	//var i=0,j=0 uint8
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			a[i][j] = uint8((i ^ j))
		}
	}

	return a
	//try at https://tour.golang.org/moretypes/18
}*/

/*
Maps
A map maps keys to values.

The zero value of a map is nil. A nil map has no keys, 
nor can keys be added.

The make function returns a map of the given type,
 initialized and ready for use
*/
type Vertex1 struct {
	Lat,Long float64
}
var m map[string]Vertex1
func GoMaps() {
m=make(map[string]Vertex1)
m["Bell Labs"]=Vertex1{40.68433, -74.39967}
fmt.Println(m["Bell Labs"])
}
/*
Map literals
Map literals are like struct literals,
 but the keys are required.
*/
func MapLiterals(){
	type Vertex struct {
		Lat, Long float64
	}
	/*
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	*/
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	

	fmt.Println(m)
}
/*
Mutating Maps
Insert or update an element in map m:

m[key] = elem
Retrieve an element:

elem = m[key]
Delete an element:

delete(m, key)
Test that a key is present with a two-value
 assignment:

elem, ok = m[key]
If key is in m, ok is true. If not, ok is false.

If key is not in the map, then elem is the zero 
value for the map's element type.

Note: if elem or ok have not yet been declared 
you could use a short declaration form:

elem, ok := m[key]
*/
func MUtatingMaps(){
	m1:=make(map[string]int)
	m1["Answer"]=42
	fmt.Println("The value: ",m1["Answer"])
	m1["Answer"]=48
	fmt.Println("The value: ",m1["Answer"])
	delete(m1,"Answer")
	fmt.Println("The Value: ",m1["Answer"])
	v,ok:=m1["Answer"]
	fmt.Println("The Value:",v,"Present?",ok)
}
//exercise problem
//https://tour.golang.org/moretypes/23
/*
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	a:=make(map[string]int )
	words:=strings.Fields(s)
	for _,word :=range words{
	a[word]++
	} 
	return a
	
}

func main() {
	wc.Test(WordCount)
}
*/
func MapsExercise(){
	s:= "I am learning go"
	a:=make(map[string]int)
	words:=strings.Fields(s)
	for _,word:=range words{
		a[word]++
	}
	fmt.Println(a)

	}


/*
Function values
Functions are values too.
 They can be passed around just like other values.

Function values may be used as function arguments and
 return values.
 */
 func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}


func FunctionValues(){
hypot:=func (x,y float64)float64{
	return math.Sqrt(x*x+y*y)
}
fmt.Println(hypot(5,12))
fmt.Println(compute(hypot))
fmt.Println(compute(math.Pow))

}
func main() {

	fmt.Println("Hello World!")
	fmt.Println("The time is ", time.Now())
	fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Printf("Now we have %g problem.\n", math.Sqrt(7))
	fmt.Printf("Value of pi=%g.\n", math.Pi)
	//fmt.Printf("12+13=%g\n",int(add(12,13)))
	fmt.Println(add(12, 13))
	a, b := swap("hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(30))
	var i int
	fmt.Println(i, c, python, java)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var i1, j int = 1, 2
	k := 3
	fmt.Println(i1, j, k)
	zero()
	conversion()
	typeinference()
	constant()
	numericconst()
	forloop()
	forwhile()
	//forever()
	fmt.Println(ifloop(-3), ifloop(3))
	fmt.Println(ifshort(3), ifshort(2))
	fmt.Println(ifelse(3), ifelse(2))
	switchcase()
	weekday()
	datetime()
	defergo()
	pointers()
	Structfileds()
	Arrays()
	SliceLiterals()
	SliceDefaults()
	SliceLengthCapacity()
	SlicesWithMake()
	SlicesOfSlices()
	SliceAppend()
	//pic.Show(SlicesExercies)
	RangeGo()
	RangeCont()
	GoMaps()
	MapLiterals()
	MUtatingMaps()
	//https://tour.golang.org/moretypes/23
	MapsExercise()
	FunctionValues()
}

//https://tour.golang.org/moretypes/24



//https://tour.golang.org/moretypes/6

