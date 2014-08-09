package main

import (
	"github.com/vandosant/euler"
	"fmt"
)

func main() {
	fmt.Println("Problem 1")
	fmt.Println("Multiples of 3 and 5")
	fmt.Println(`If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.`)
	fmt.Println("Answer:")
	fmt.Println(euler.SumMultiples(1000))
	fmt.Println("---------------")
	fmt.Println("Problem 2")
	fmt.Println("Even Fibonacci numbers")
	fmt.Println(`Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
	By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.`)
	fmt.Println("Answer:")
	fmt.Println(euler.SumEvenValues(euler.Fibonacci(4000000)))
	fmt.Println("---------------")
	fmt.Println("Problem 3")
	fmt.Println("Largest prime factor")
	fmt.Println("The prime factors of 13195 are 5, 7, 13 and 29.")
	fmt.Println("What is the largest prime factor of the number 600851475143 ?")
	fmt.Println("Answer:")
	fmt.Println(euler.PrimeFactorization(600851475143))
	fmt.Println("---------------")
	fmt.Println("Problem 4")
	fmt.Println("Largest palindrome product")
	fmt.Println("A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.")
	fmt.Println("Find the largest palindrome made from the product of two 3-digit numbers.")
	fmt.Println("Answer:")
	fmt.Println(euler.LargestPalindrome(3))
	fmt.Println("---------------")
	fmt.Println("Problem 5")
	fmt.Println("Smallest multiple")
	fmt.Println("2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.")
	fmt.Println("What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?")
	fmt.Println("Answer:")
	fmt.Println(euler.SmallestMultiple(20))
}
