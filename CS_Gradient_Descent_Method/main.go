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
	dk01  float64
	dk02  float64
	tgrad gradient
	thess hessian_matrix
	rhess hessian_matrix
	coorM = [2]coordinate{}
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

type hessian_matrix struct {
	h11 float64
	h12 float64
	h21 float64
	h22 float64
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

func matrixH() hessian_matrix {
	var h11 = 2 * a
	var h12 = b
	var h21 = b
	var h22 = 2 * c

	var h hessian_matrix = hessian_matrix{h11, h12, h21, h22}

	return h

}

func reverse_matrixH() hessian_matrix {
	var h11 = (2 * c) / (4*a*c - b*b)
	var h12 = -b / (4*a*c - b*b)
	var h21 = -b / (4*a*c - b*b)
	var h22 = (2 * a) / (4*a*c - b*b)

	var h hessian_matrix = hessian_matrix{h11, h12, h21, h22}

	return h

}

func CS_gradient_descent_method() {

	thess = matrixH()
	rhess = reverse_matrixH()

	fmt.Printf("Матрица Гессе H(x) = [%3.3f, %3.3f]\n", thess.h11, thess.h12)
	fmt.Printf("                     [%3.3f, %3.3f]\n", thess.h21, thess.h22)
	fmt.Printf("Обратная матрица Гессе H^(-1)(x) = [%3.3f, %3.3f]\n", rhess.h11, rhess.h12)
	fmt.Printf("                                   [%3.3f, %3.3f]\n", rhess.h21, rhess.h22)
	for i := 0; i >= 0; i++ {
		tgrad = grad(x01, x02)

		fmt.Printf("Шаг равен [%d]\n", i+1)
		fmt.Printf("Градиент равен ▽f(x) = [%f, %f]\n", tgrad.x1, tgrad.x2)
		//fmt.Printf("x(%.6f , %.6f)", x01, x02)
		//fmt.Printf("Интервал [a0, b0] = [%.3f , %.3f]\n", a0, b0)

		if math.Sqrt(math.Pow(tgrad.x1, 2)+math.Pow(tgrad.x2, 2)) <= E_One {
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

				if rhess.h11 > 0 && rhess.h11*rhess.h22-rhess.h21*rhess.h12 > 0 {
					dk01 = -(rhess.h11*tgrad.x1 + rhess.h12*tgrad.x2)
					dk02 = -(rhess.h21*tgrad.x1 + rhess.h22*tgrad.x2)
				} else {
					dk01 = -tgrad.x1
					dk02 = -tgrad.x2
				}

				if dk01 == -(rhess.h11*tgrad.x1+rhess.h12*tgrad.x2) && dk02 == -(rhess.h21*tgrad.x1+rhess.h22*tgrad.x2) {
					tk = 1
					coorM[0] = coordinate{x01, x02}
					x01 = x01 + tk*dk01
					x02 = x02 + tk*dk02
				} else if dk01 == -tgrad.x1 && dk02 == -tgrad.x2 {
					for calcAnswer(x01, x02) >= calcAnswer(coorM[0].x01, coorM[0].x02) {
						tk = tk / 2
					}
					coorM[1] = coordinate{x01, x02}
					x01 = x01 + tk*dk01
					x02 = x02 + tk*dk02
				}

				var xk21 = math.Sqrt(math.Pow(x01-coorM[1].x01, 2) + math.Pow(x02-coorM[1].x02, 2))
				var xk10 = math.Sqrt(math.Pow(coorM[1].x01-coorM[0].x01, 2) + math.Pow(coorM[1].x02-coorM[0].x02, 2))
				var leng21 = math.Abs(calcAnswer(x01, x02) - calcAnswer(coorM[1].x01, coorM[1].x02))
				var leng10 = math.Abs(calcAnswer(coorM[1].x01, coorM[1].x02) - calcAnswer(coorM[0].x01, coorM[0].x02))

				if xk21 < E_TWo && xk10 < E_TWo && leng21 < E_TWo && leng10 < E_TWo {

					fmt.Printf("Величина шага 'tk' равен [%.3f]\n", tk)
					fmt.Printf("x* = x(%0.6f , %0.6f)", x01, x02)
					fmt.Printf(" f(x*) = %f\n\n", calcAnswer(x01, x02))

					break
				}

			}
		}

		fmt.Printf("Величина шага 'tk' равен [%.3f]\n", tk)
		fmt.Printf("x(%.6f , %.6f) ", x01, x02)
		fmt.Printf("dk(%.6f , %.6f)", dk01, dk02)
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
	// fmt.Fprintln(os.Stdout, "Введите величину шага tk")
	// fmt.Fscan(os.Stdin, &tk)

	CS_gradient_descent_method()

}
