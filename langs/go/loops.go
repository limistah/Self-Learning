package main

import "fmt"

func main () {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Skip the increment part
	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	// Skips the initialization and increment
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	// Infinite loop
	sum := 0
	// for {
	// 	sum++
	// }
	fmt.Println(sum)

	// Continous statement
	for i := 0; i < 10; i++ {
		if  i%2 != 0 {
			continue
		}
		fmt.Printf("%d\n", i)
	}

	// Break statement
	count := 0
	for i := 0; true; i++ {
		if i% 13 == 0 {
			fmt.Printf("%d is divisible by 13 \n", i)
			count++
		}
		if count == 10 {
			break
		}
	}

	// Labels
	people := [10]string{"Helen, Mark", "Brenda", "Antonio", "Micheal"}
	friends := [2]string{"Mark", "Mar"}
outer:
	   for index, name := range people {
		   for _, friend := range friends {
			   if  name == friend {
				   fmt.Printf("Found a friend %q at index %d\n", friend, index)
				   break outer
			   }
		   }
	   }
	fmt.Println("Next instruction after the break")

	// Goto statement
	i := 0
loop:
   if i < 5 {
	   fmt.Println(i)
	   i++
	   goto loop
   }

   // Would throw no new variable is allowed when jumping to a goto
//    goto todo
//    x :=5
// todo:
//    fmt.Println("Something here")


	
}