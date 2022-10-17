package main

import (
	"fmt"
	"math"
	"os"
)

var (
	a     float64
	b     float64
	c     float64
	E     float64
	E_One float64
	E_TWo float64
	M     int
	x01   float64
	x02   float64
	tgrad gradient
	coorM = [1]coordinate{}
	tk    float64
)

type coordinate struct {
	x01 float64
	x02 float64
}

type gradient struct {
	x1 float64
	x2 float64
}

func grad(x1 float64, x2 float64) gradient {

	var g1 = a*2*x1 + b*x2
	var g2 = b*x1 + c*2*x2
	var g gradient = gradient{g1, g2}

	return g

}

func calcAnswer(x1 float64, x2 float64) float64 {

	var answer = a*math.Pow(x1, 2) + b*x1*x2 + c*math.Pow(x2, 2)
	return answer
}

func CS_gradient_descent_method() {

	for i := 0; i >= 0; i++ {
		tgrad = grad(x01, x02)
		fmt.Printf("Шаг равен [%d]\n", i+1)
		fmt.Printf("Градиент равен ▽f(x) = [%f, %f]\n", tgrad.x1, tgrad.x2)
		if math.Sqrt(math.Pow(tgrad.x1, 2)+math.Pow(tgrad.x2, 2)) < E_One {
			//log.Println("x* =" + "x(" + fmt.Sprintf("%.3f", x01) + "," + fmt.Sprintf("%.3f", x02) + ")" + "   k = " + fmt.Sprint(i+1))
			//fmt.Printf("Шаг равен [%d]\n", i+1)
			fmt.Printf("x* = x(%.6f , %.6f)", x01, x02)
			break
		} else {
			if i >= M {
				//fmt.Printf("Шаг равен [%d]\n", i+1)
				fmt.Printf("x* = x(%.6f , %.6f)", x01, x02)
				break
			} else {
				// if i == 0 {
				// 	fmt.Fprintln(os.Stdout, "Введите величину шага tk")
				// 	fmt.Fscan(os.Stdin, &tk)
				// }

				for i > -1 {
					coorM[0] = coordinate{x01, x02}

					x01 = x01 - tk*tgrad.x1
					x02 = x02 - tk*tgrad.x2

					if calcAnswer(x01, x02)-calcAnswer(coorM[0].x01, coorM[0].x02) >= 0 {
						tk = tk / 2
					} else {
						break
					}
				}

				if math.Sqrt(math.Pow(x01-coorM[0].x01, 2)+math.Pow(x02-coorM[0].x02, 2)) < E_TWo && math.Abs(calcAnswer(x01, x02)-calcAnswer(coorM[0].x01, coorM[0].x02)) < E_TWo {
					//fmt.Printf("Шаг равен [%d]\n", i+1)
					//fmt.Printf("Градиент равен ▽f(x) = [%f, %f]\n", tgrad.x1, tgrad.x2)
					fmt.Printf("Величина шага 'tk' равен [%.3f]\n", tk)
					fmt.Printf("x* = x(%.6f , %.6f)", x01, x02)
					fmt.Printf(" f(x*) = %.6f\n\n", calcAnswer(x01, x02))

					break
				}

			}
		}

		fmt.Printf("Величина шага 'tk' равен [%.3f]\n", tk)
		fmt.Printf("x(%.6f , %.6f)", x01, x02)
		fmt.Printf(" f(x) = %.6f\n\n", calcAnswer(x01, x02))
	}
}

func main() {

	fmt.Fprintln(os.Stdout, "Введите коэффициент (a) перед (x1^2)")
	fmt.Fscan(os.Stdin, &a)
	fmt.Fprintln(os.Stdout, "Введите коэффициент (b) перед (x1x2)")
	fmt.Fscan(os.Stdin, &b)
	fmt.Fprintln(os.Stdout, "Введите коэффициент (c) перед (x2^2)")
	fmt.Fscan(os.Stdin, &c)

	fmt.Fprintln(os.Stdout, "Введите x01")
	fmt.Fscan(os.Stdin, &x01)
	fmt.Fprintln(os.Stdout, "Введите x02")
	fmt.Fscan(os.Stdin, &x02)
	// fmt.Fprintln(os.Stdout, "Введите E")
	// fmt.Fscan(os.Stdin, &E)
	fmt.Fprintln(os.Stdout, "Введите E1")
	fmt.Fscan(os.Stdin, &E_One)
	fmt.Fprintln(os.Stdout, "Введите E2")
	fmt.Fscan(os.Stdin, &E_TWo)
	fmt.Fprintln(os.Stdout, "Введите предельное число интераций M")
	fmt.Fscan(os.Stdin, &M)
	fmt.Fprintln(os.Stdout, "Введите величину шага tk")
	fmt.Fscan(os.Stdin, &tk)

	CS_gradient_descent_method()

}
