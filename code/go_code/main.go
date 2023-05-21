package main

import (            // you'd better use these imports or else...
	"fmt"       // to fmt.Println on stdout (stderr: print()
	"strconv"   // to convert integer and floats to string for printing
	"math"    // to do fancy math stuff
	//"unused"  // do not import unused libraries
)

const (                        // vars that do/can not change
	pi float32 = 3.1415    // define and assign float of 32 bits  ( danger! also: iota -> 0 )
	first_iota = iota      // const with special value, everytime you make a const it autoincrements, so 1
	second_iota = iota     // +1 -> 2
	third_iota = iota + 3  // +1 -> 3, +3 so here: 6
	fourth_iota            // go magic: increment by 1 again and +3 again so 4 + 3 = 7
	fifth_iota             // go magic: increment by 1 again and +3 again so 5 + 3 = 8
	sixth_iota = iota      // increment by 1 again so 6
	unused_const = 13      // and iota -> 7
	seventh_iota = iota    // 8, dangerous! better make a second const block for iota range
)

const (                        // we made this especially for iota range, do not define other const in here!
	first_iota2  = iota    // 0
	second_iota2 = iota    // 1
	third_iota2  = iota    // 2
)                              // so yes you can have more than 1 const block

var (                             // vars you can change, note: no code in this block!
	msg        string = "hoi" // no single quotes allowed!
	i          int    = 39    // integer
	unused_var int    = 53    // in this block: no error if not used anywhere
	a_bool     bool   = true  // booleans have lowercase true/false
        no_fill    int            // in this block: no need to fill
	res        int    = 13    // res is also used/declared in multiplyTwoNumbers() below, but scope differs
)

func say(m string) {   // function, input only, still needs type
	fmt.Println(m) // now you can say() stuff in stead of fmt.Println() stuff
}

func unused_func(blah int) {   // unused functions: no error
	//non_existant_var++   // nope! still error for non-existant var in unused function
}

func multiplyTwoNumbers(x, y int) (res int) { // in and output both need to be declared
	res = x * y                           // res var is declared in func definition on prev line.
	return res                            // this res is different from var res from var block above
}

func blab() {
	fmt.Println("-------------------------")
	fmt.Println(res) // yes it is accessible, because declared in var block
	fmt.Println("-------------------------")
}

func main() {                     // main() is automatically started
	fmt.Println(first_iota)   // iota const: 1
	fmt.Println(second_iota)  // iota const: 2
	fmt.Println(third_iota)   // iota const: 3 + 3 = 6
	fmt.Println(fourth_iota)  // iota const: 4 + 3 = 7
	fmt.Println(fifth_iota)   // iota const: 5 + 3 = 8
	fmt.Println(sixth_iota)   // iota const: 6
	fmt.Println(seventh_iota) // iota const: 8 (because another const was made between sixth_iota and seventh_iota)
	var v1 int                // like in var block (if you print v1 now, it displays 0)
	v1++                      // avoid the "not used" error, also: v1 is now 1
	var v2 int = 4            // like in var block
	print(v2, "\n")           // print to stderr
	b := 42                   // without var keyword, var definitions need colon, also: implicitly typed here
	//b = b + 0.1             // nope! b was implicitly typed an integer above
	b = b + 10                // yup! add integer to integer
	d := int(12)              // explicitly typed
	print(d, "\n")            // need to use var, else error "declared but not used"
	//unused_var2 := 53       // nope! no unused vars outside the var block allowed
	m, n := 7, 8              // look mom, both hands!
	print(m, " ", n, "\n")    // print() needs explicit space an newline
	i = i + 3                 // simple arithmatic with previously declared var
	i += 3                    // same thing different
	say(strconv.Itoa(i))      // function call while converting int to ascii, because say() needs string
	fmt.Println(msg)          // the normal way to print
	fmt.Println(42)           // fmt.Println() can print integers
	fmt.Println(pi)           // ... and floats
	say(msg + " bla")         // strings can be added to
	say("")                   // empty allowed, just prints "\n"
        fmt.Printf("%[1]v\n", b)  // printf the value in a default format to stdout, see https://pkg.go.dev/fmt
	fmt.Printf("%[1]v\n", pi) // float gets printed as a float, int as an int etc.
	fmt.Printf("%[1]T\n", pi) // print type of the var to stdout
        fmt.Printf("%[1]b\n", b)  // printf the value as binary to stdout
        fmt.Printf("%[1]X\n", 31) // printf the value as (uppercase) hex to stdout

	s0 := strconv.FormatFloat(float64(pi), 'f', 2, 32) // convert float32 to string (FormatFloat needs float64)
	print(s0, "\n")                                    // also rounded to 2 numbers after decimal point
        
	const c0 = "pipo"                   // note: cannot change a const
	const c1 = 1313                     // not expected: no errors if unused
	fmt.Printf(c0 + " en mamaloe \n")   // works, maybe useful if no \n needed
	//fmt.Printf(c0 + 42 + "\n")        // nope! type mismatch: no mixing str and int
	const f1 = float32(math.MaxFloat32) // for some reason MaxFloat32 is a float64 if not defined a float32
        fmt.Printf("%[1]T %[1]f\n", f1)     // print type and value to stdout, "no exponent" style
	const f2 = math.MaxFloat32          // f2 is a float64 with max value for a float32
        fmt.Printf("%[1]T %[1]f\n", f2)     // same value as above, but different type
	const f3 = math.MaxFloat64          // maximum value for a float64
        fmt.Printf("%[1]T %[1]v\n", f3)     // print type and value to stdout, exponent style (which is the default)
	const f4 = math.MaxInt              // everybody likes big integers
	float_f4 := float32(f4)             // to see the magnitude of MaxInt convert to float and ...
	fmt.Printf("%[1]e\n", float_f4)     // ... print exponent style to stdout
	f4_str := strconv.Itoa(f4)          // convert to string
	fmt.Println("MaxInt = " + f4_str)   // print the actual number an int can be

	res = 13                              // fill empty var from var block, no colon needed because pre-declared
	fmt.Println(res)                      // 13 (duh!)
	blab()                                // because res is defined in var block it is accessible to the blab() function
	fmt.Println(multiplyTwoNumbers(3, 2)) // this function call has its own var called res, and even used as return
	fmt.Println("still", res)             // still 13 because multiplyTwoNumbers's var res is another scope

	// python: var1 = 1 means: var1 -> int_object_made_by_python -> 0x12345[1]
	// go:     var1 := 1 means: var1 -> 0x12345[1]
	var p0 *string = new(string)   // declare pointer to string and give address of a string as value
        //var p1 *string               // nope! p1 = <nil> and accessing *p1 gets you a seg fault
	*p0 = "value (in pointer p0)"  // insert a value in the address of the string declared with new(string)
	print(*p0, " ", p0, "\n")      // print value and adres of pointer p0
	print("address of var res is ", &res, "\n") // print address of var

	var arr0 = [3]int{1, 2, 3}         // an array of ints that cannot change
	arr1 := [3]int{2, 4, 6}            // same
	arr0_copy := arr0                  // make a copy of arr0 (not like python!)
	fmt.Println(arr0, arr0_copy, arr1) // [1 2 3] [1 2 3] [2 4 6]
	arr0[1] = 22                       // copy is still [1, 2, 3]
	fmt.Println(arr0, arr0_copy)       // [1 22 3] [1 2 3]

        slice0 := []int{2, 4, 6}           // slice is mutable array, size managed by go
	slice1 := arr1[:]                  // also made from array, slice becomes reference to the data of arr1
	fmt.Println(slice0, slice1)        // same data because slice0 and arr1 have same data 
	fmt.Printf("Address of the array: %p\n", &arr1)                   // mem address of arr1
	fmt.Printf("Address of the slice underlying array: %p\n", slice1) // mem addres of array under slice1 (same address as prev)
	arr1[1] = 44              // change is also visible in slice1
	fmt.Println(arr1, slice1) // [2 44 6] [2 44 6] because slice1 references to the data of arr1
	slice0 = append(slice0, 8) // you can append to a slice
	fmt.Println(slice0)        // [2 4 6 8]
	slice1 = append(slice1, 8) // but if there is an underlying array, it does not change along with it
	fmt.Println(arr1, slice1)  // [2 44 6]  [2 44 6 8]
	fmt.Printf("Address of the array: %p\n", &arr1)                   // mem address of arr1
	fmt.Printf("Address of the slice underlying array: %p\n", slice1) // no longer the same address
	// python: x = [1, 2, 3]; y = x; x.append(4) -> y is now [1, 2, 3, 4] because x and y both point to the same array object.
        slice0 = append(slice0, 10, 12, 14, 16) // all at once
	fmt.Println(slice0)
        slice2 := slice0[2:6] // start at element 2, up to (but not including) element 6. [2 4 6 8 10 12 14 16] -> [6 8 10 12]
	fmt.Println(slice2)

	map0 := map[string]int{"foo":42, "bar":68, "baz":55} // define hash key: string, value: integer
	fmt.Println(map0)                                    // nothing unusual here
	fmt.Println(map0["bar"])                             // nothing unusual here
	map0["baz"] = 73                                     // nothing unusual here
	fmt.Println(map0)                                    // nothing unusual here
	//map0["foo"] = [1, 2, 3]                            // nope! only int, see stuct


	type specimen struct{                                // first define the struct, can also be done at the top of the file
		coll_id int
		genus string
		species string
		location map[string]float32                 // key: string, value: float32
	}
	ant0 := specimen{}                                  // fake 0th specimen so antX = coll_id X
	ant1 := specimen{ coll_id: 1,                       // instantiate first specimen
	                  genus: "Lasius",
	                  species: "niger",
                          location: map[string]float32{"x_coord": 52.23456, "y_coord": 5.0001},
                        }
        fmt.Println(ant1)
        fmt.Println(ant1.genus)               // struct uses dot notation
        fmt.Println(ant1.location["x_coord"]) // map in struct uses normal notation
	ant2 := specimen{ coll_id: 2,                       // instantiate first specimen
	                  genus: "Lasius",
	                  species: "flavus",
                          location: map[string]float32{"x_coord": 52.234562, "y_coord": 5.0002},
                        }

	collection := []specimen{ant0, ant1, ant2} // slice of instances of specimen
	fmt.Println(collection)                    // whole slice
	fmt.Println(collection[1])                 // element 1 of slice
	fmt.Println(collection[1].genus)           // genus of element 1 of slice

	var ants []*specimen                    // slice of pointers to specimens
        ants = append(ants, &ant1, &ant2)       // put pointers to speciments in it
	fmt.Println(ants)                       // whole slice (prints addresses!)
	fmt.Println(*ants[1])                   // contents of address 1 in slice
	fmt.Println(ants[1].genus)              // genus of specimen at address 1 in slice

}
