package main
import (
	"fmt"
	"time"
)
type Matrix [][]int

func newMatrix(r, c int) [][]int {
	a := make([]int, c*r)
	m := make([][]int, r)
	lo, hi := 0, c
	for i := range m {
		m[i] = a[lo:hi:hi]
		lo, hi = hi, hi+c
	}
	return m
}

func rowCount(inM Matrix) int {
	return len(inM)
}

func colCount(inM Matrix) int {
	return len(inM[0])
}
func doCannonAlgorithm (inA Matrix, inB Matrix) Matrix {
	m := rowCount(inA)     // number of rows the first matrix
	//n := colCount(inA)     // number of columns the first matrix
	//p := rowCount(inB)    // number of rows the second matrix
	q := colCount(inB)    // number of columns the second matrix

	CTotal := newMatrix(m, q)

	start := time.Now()
	time.Sleep(1111 * time.Millisecond) // just to max sure timer works delete later

	//Step ONE
	A1 := newMatrix(m, q)
	B1 := newMatrix(m, q)
	A1 = stepOneShiftA(inA, m)
	B1 = stepOneShiftB(inB, q)
	C1 := multiplyIJofEach(A1, B1)
	CTotal = addIJofEach(CTotal, C1)

	//Step TWO
	A2 := newMatrix(m, q)
	B2 := newMatrix(m, q)
	A2 = oneEachShiftA(A1)
	B2 = oneEachShiftB(B1)
	C2 := multiplyIJofEach(A2, B2)
	CTotal = addIJofEach(CTotal, C2)

	//Step THREE
	A3 := newMatrix(m, q)
	B3 := newMatrix(m, q)
	A3 = oneEachShiftA(A2)
	B3 = oneEachShiftB(B2)
	C3 := multiplyIJofEach(A3, B3)
	CTotal = addIJofEach(CTotal, C3)

	//Step FOUR
	A4 := newMatrix(m, q)
	B4 := newMatrix(m, q)
	A4 = oneEachShiftA(A3)
	B4 = oneEachShiftB(B3)
	C4 := multiplyIJofEach(A4, B4)
	CTotal = addIJofEach(CTotal, C4)

	fmt.Println("C1 = ")
	printMat(C1)
	fmt.Println("C2 = ")
	printMat(C2)
	fmt.Println("C3 = ")
	printMat(C3)
	fmt.Println("C4 = ")
	printMat(C4)

	elapsed := time.Since(start)
	fmt.Printf("Time taken to calculate %s ",elapsed)
	fmt.Println()
	fmt.Println("The result of the Matrix multiplication is C1 + C2 + C3 + C4 which =")

	return CTotal
}

func oneEachShiftA(inA Matrix) Matrix {
	for i:= 0 ; i < rowCount(inA) ; i++ {
		inA = oneShiftRow(inA, i)
	}
	return inA
}

func oneShiftRow(inA Matrix, x int) Matrix {
	for i:=0 ; i < len(inA) - 1 ; i ++ {
		inA[x][i], inA[x][i+1] = inA[x][i+1], inA[x][i]
	}
	return inA
}

func oneEachShiftB(inB Matrix) Matrix {
	for i:= 0 ; i < colCount(inB) ; i++ {
		inB = oneShiftCol(inB, i)
	}
	return inB
}

func oneShiftCol(inB Matrix, x int) Matrix {
	for i:=0 ; i < len(inB) - 1 ; i ++ {
		inB[i][x], inB[i+1][x] = inB[i+1][x], inB[i][x]
	}
	return inB
}

func multiplyIJofEach(inA Matrix, inB Matrix) Matrix {
	result := newMatrix(rowCount(inA), colCount(inB))
	for i := 0 ; i < len(inA) ; i++ {
		for j := 0 ; j < len(inA) ; j++ {
			result[i][j] = inA[i][j] * inB[i][j]
		}
	}
	return result
}

func addIJofEach (inA Matrix, inB Matrix) Matrix {
	result := newMatrix(rowCount(inA), colCount(inB))
	for i := 0 ; i < len(inA) ; i++ {
		for j := 0 ; j < len(inA) ; j++ {
			result[i][j] = inA[i][j] + inB[i][j]
		}
	}
	return result
}

func stepOneShiftA(inA Matrix, rowCount int) Matrix {			//performs 0LS on row 1, 1LS on row2, 2LS on row3, 3LS on row4
	for i := 1 ; i < rowCount; i++ {
		j := 0
		for i > j {
			inA = oneShiftRow(inA, i)
			j = j+1
		}
	}
	fmt.Println("A1 (step one is unique)= ")
	printMat(inA)
	return inA
}
func stepOneShiftB(inB Matrix, colCount int) Matrix{			//performs 0LS on row 1, 1LS on row2, 2LS on row3, 3LS on row4
	for i := 1 ; i < colCount; i++ {
		j := 0
		for i > j {
			inB = oneShiftCol(inB, i)
			j = j+1
		}
	}
	fmt.Println("B1 (step one is unique)= ")
	printMat(inB)
	return inB
}



func printMat(inM Matrix) {
	for _, i := range inM {
		for _, j := range i {
			fmt.Print(" ", j)
		}
		fmt.Println()
	}
}
func main() {
	// Use slices
	// Unlike arrays they are passed by reference,not by value
	a := Matrix{{2, 3, 4, 5}, {9, 8, 7, 6}, {5, 4, 2, 3}, {8, 7, 3, 4}}
	b := Matrix{{3, 5, 7, 6}, {2, 7, 6, 3}, {7, 5, 3, 2}, {4, 3, 2, 5}}

	fmt.Println("A = ")
	printMat(a)
	fmt.Println("B = ")
	printMat(b)

	c := doCannonAlgorithm(a, b)
	printMat(c)
}
