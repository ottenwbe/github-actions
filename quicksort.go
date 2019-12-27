/*
MIT License

Copyright (c) 2019 Beate Ottenwälder

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := newSlice(50)
	QuickSort(slice)
	fmt.Println(slice)
}

func newSlice(length int) []int {
	slice := make([]int, length, length)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range slice {
		slice[i] = rnd.Intn(1000)
	}

	return slice
}

// QuickSort sorts a slice of an integer array
func QuickSort(slice []int) {
	if len(slice) > 1 {
		pivot := partition(slice)
		QuickSort(slice[:pivot])
		QuickSort(slice[pivot+1:])
	}
}

func partition(slice []int) int {
	left, right := 0, len(slice)-1
	pivot := slice[right]

	for j := range slice {
		if slice[j] < pivot {
			slice[j], slice[left] = slice[left], slice[j]
			left++
		}
	}
	slice[left], slice[right] = slice[right], slice[left]
	return left
}
