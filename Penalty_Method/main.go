package main

import (
	"fmt"
	"math"
	"os"
)

var (
	a float64 = 10
	b float64 = 3
	c float64 = -8

	ag float64 = 3
	bg float64 = 1
	cg float64 = -10

	x01 float64
	x02 float64

	E float64 = 0.05

	E_One float64 = 0.15
	E_TWo float64 = 0.20

	r0 float64
	C  float64 = 5

	dk01 float64
	dk02 float64

	tgrad gradient

	coorM         = [2]coordinate{}
	tk    float64 = 0.5

	thess hessian_matrix
	rhess hessian_matrix
	// tgrad gradient
	// thess hessian_matrix
	// rhess hessian_matrix
	// coorM = [2]coordinate{}

)

type coordinate struct {
	x01 float64
	x02 float64
}

type gradient struct {
	x1 float64
	x2 float64
}

func grad(x1 float64, x2 float64, r float64) gradient {

	var g1 = 20*x1 + r*(9*x1+3*x2-33)
	var g2 = 6*x2 + r*(x2+3*x1-11)
	var g gradient = gradient{g1, g2}

	return g

}

type hessian_matrix struct {
	h11 float64
	h12 float64
	h21 float64
	h22 float64
}

func calcAnswer(x1 float64, x2 float64, r float64) float64 {

	return a*math.Pow(x1, 2) + b*math.Pow(x2, 2) + c + (r/2)*math.Pow((ag*x1+bg*x2+cg), 2)

}

func calcPanswer(x1 float64, x2 float64, r float64) float64 {

	return r / 2 * math.Pow((ag*x1+bg*x2+cg), 2)

}

func matrixH() hessian_matrix {
	var h11 = 9*r0 + 20
	var h12 = 3 * r0
	var h21 = 3 * r0
	var h22 = r0 + 6

	var h hessian_matrix = hessian_matrix{h11, h12, h21, h22}

	return h

}

func reverse_matrixH() hessian_matrix {
	var h11 = (9*r0*r0 + 74*r0 + 120) / ((9*r0 + 20) * (74*r0 + 120))
	var h12 = -3 * r0 / (74*r0 + 120)
	var h21 = -3 * r0 / (74*r0 + 120)
	var h22 = (9*r0 + 20) / (74*r0 + 120)

	var h hessian_matrix = hessian_matrix{h11, h12, h21, h22}

	return h

}

func CS_gradient_descent_method(x1 float64, x2 float64) coordinate {

	thess = matrixH()
	rhess = reverse_matrixH()

	// fmt.Printf("Матрица Гессе H(x) = [%3.3f, %3.3f]\n", thess.h11, thess.h12)
	// fmt.Printf("                     [%3.3f, %3.3f]\n", thess.h21, thess.h22)
	// fmt.Printf("Обратная матрица Гессе H^(-1)(x) = [%3.3f, %3.3f]\n", rhess.h11, rhess.h12)
	// fmt.Printf("                                   [%3.3f, %3.3f]\n", rhess.h21, rhess.h22)
	for i := 0; i >= 0; i++ {
		tgrad = grad(x1, x2, r0)

		// fmt.Printf("Шаг равен [%d]\n", i+1)
		// fmt.Printf("Градиент равен ▽f(x) = [%f, %f]\n", tgrad.x1, tgrad.x2)
		//fmt.Printf("x(%.6f , %.6f)", x01, x02)
		//fmt.Printf("Интервал [a0, b0] = [%.3f , %.3f]\n", a0, b0)

		if math.Sqrt(math.Pow(tgrad.x1, 2)+math.Pow(tgrad.x2, 2)) <= E_One {
			//log.Println("x* =" + "x(" + fmt.Sprintf("%.3f", x01) + "," + fmt.Sprintf("%.3f", x02) + ")" + "   k = " + fmt.Sprint(i+1))
			//fmt.Printf("Шаг равен [%d]\n", i+1)
			fmt.Printf("x* = x(%.6f , %.6f)\n", x1, x2)
			return coordinate{x1, x2}

		} else {
			if i >= 20 {
				//fmt.Printf("Шаг равен [%d]\n", i+1)
				fmt.Printf("x* = x(%.6f , %.6f)\n", x1, x2)
				return coordinate{x1, x2}

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
					tk = 0.5
					coorM[0] = coordinate{x1, x2}
					x1 = x1 + tk*dk01
					x2 = x2 + tk*dk02
				} else if dk01 == -tgrad.x1 && dk02 == -tgrad.x2 {
					for calcAnswer(x1, x2, r0) >= calcAnswer(coorM[0].x01, coorM[0].x02, r0) {
						tk = tk / 2
					}
					coorM[1] = coordinate{x1, x2}
					x1 = x1 + tk*dk01
					x2 = x2 + tk*dk02
				}

				var xk21 = math.Sqrt(math.Pow(x1-coorM[1].x01, 2) + math.Pow(x2-coorM[1].x02, 2))
				var xk10 = math.Sqrt(math.Pow(coorM[1].x01-coorM[0].x01, 2) + math.Pow(coorM[1].x02-coorM[0].x02, 2))
				var leng21 = math.Abs(calcAnswer(x1, x2, r0) - calcAnswer(coorM[1].x01, coorM[1].x02, r0))
				var leng10 = math.Abs(calcAnswer(coorM[1].x01, coorM[1].x02, r0) - calcAnswer(coorM[0].x01, coorM[0].x02, r0))

				if xk21 < E_TWo && xk10 < E_TWo && leng21 < E_TWo && leng10 < E_TWo {

					// fmt.Printf("Величина шага 'tk' равен [%.3f]\n", tk)
					// fmt.Printf("x* = x(%0.6f , %0.6f)\n", x1, x2)
					// fmt.Printf(" f(x*) = %f\n\n", calcAnswer(x1, x2, r0))
					return coordinate{x1, x2}

				}

			}
		}

		// fmt.Printf("Величина шага 'tk' равен [%.3f]\n", tk)
		// fmt.Printf("x(%.6f , %.6f) ", x1, x2)
		// fmt.Printf("dk(%.6f , %.6f)", dk01, dk02)
		// fmt.Printf(" f(x) = %.6f\n\n", calcAnswer(x1, x2, r0))
	}
	return coordinate{x1, x2}
}

// func CS_gradient_descent_method(x1 float64, x2 float64, r float64) coordinate {

// 	sx1 := x1
// 	sx2 := x2

// 	for i := 0; i >= 0; i++ {
// 		tgrad = grad(sx1, sx2, r)
// 		fmt.Printf("	Шаг вычисления минимума методом градиентного спуска равен [%d]\n", i+1)
// 		fmt.Printf("	Градиент равен ▽f(x) = [%f, %f]\n", tgrad.x1, tgrad.x2)
// 		fmt.Printf("	Величина шага 'tk' равен [%.3f]\n", tk)
// 		fmt.Printf("	x(%.6f , %.6f)", sx1, sx2)
// 		fmt.Printf("	f(x) = %.6f\n\n", calcAnswer(sx1, sx2, r))
// 		if math.Sqrt(math.Pow(tgrad.x1, 2)+math.Pow(tgrad.x2, 2)) < E_One {
// 			//log.Println("x* =" + "x(" + fmt.Sprintf("%.3f", x01) + "," + fmt.Sprintf("%.3f", x02) + ")" + "   k = " + fmt.Sprint(i+1))
// 			//fmt.Printf("Шаг равен [%d]\n", i+1)
// 			fmt.Printf("	x* = x(%.6f , %.6f)", sx1, sx2)
// 			return coordinate{sx1, sx2}
// 		} else {
// 			// if i == 0 {
// 			// 	fmt.Fprintln(os.Stdout, "Введите величину шага tk")
// 			// 	fmt.Fscan(os.Stdin, &tk)
// 			// }
// 			for i > -1 {

// 				if i > 0 {

// 					coorM[0] = coorM[1]
// 				}

// 				coorM[1] = coordinate{sx1, sx2}

// 				sx1 = sx1 - tk*tgrad.x1
// 				sx2 = sx2 - tk*tgrad.x2
// 				fmt.Printf("проверка(%.6f , %.6f)", sx1, sx2)
// 				fmt.Printf("	Величина шага 'tk' равен [%.3f]\n", tk)
// 				if calcAnswer(sx1, sx2, r)-calcAnswer(coorM[1].x01, coorM[1].x02, r) >= 0 {
// 					tk = tk / 2
// 				} else {
// 					break
// 				}
// 			}

// 			var xk21 = math.Sqrt(math.Pow(sx1-coorM[1].x01, 2) + math.Pow(sx2-coorM[1].x02, 2))
// 			var xk10 = math.Sqrt(math.Pow(coorM[1].x01-coorM[0].x01, 2) + math.Pow(coorM[1].x02-coorM[0].x02, 2))
// 			var leng21 = math.Abs(calcAnswer(sx1, sx2, r) - calcAnswer(coorM[1].x01, coorM[1].x02, r))
// 			var leng10 = math.Abs(calcAnswer(coorM[1].x01, coorM[1].x02, r) - calcAnswer(coorM[0].x01, coorM[0].x02, r))

// 			if xk21 < E_TWo && xk10 < E_TWo && leng21 < E_TWo && leng10 < E_TWo {

// 				fmt.Printf("	Величина шага 'tk' равен [%.3f]\n", tk)
// 				fmt.Printf("	x* = x(%0.6f , %0.6f)", sx1, sx2)
// 				fmt.Printf("	f(x*) = %f\n\n", calcAnswer(sx1, sx2, r))
// 				return coordinate{sx1, sx2}
// 			}

// 			if math.Sqrt(math.Pow(x01-coorM[0].x01, 2)+math.Pow(x02-coorM[0].x02, 2)) < 0.20 && math.Abs(calcAnswer(x1, x2, r)-calcAnswer(coorM[0].x01, coorM[0].x02, r)) < 0.20 {
// 				//fmt.Printf("Шаг равен [%d]\n", i+1)
// 				//fmt.Printf("Градиент равен ▽f(x) = [%f, %f]\n", tgrad.x1, tgrad.x2)
// 				fmt.Printf("	Величина шага 'tk' равен [%.3f]\n", tk)
// 				fmt.Printf("	x* = x(%.6f , %.6f)", x1, x2)
// 				fmt.Printf("	f(x*) = %.6f\n\n", calcAnswer(x1, x2, r))
// 				return coordinate{x1, x2}
// 			}
// 		}

// 	}
// 	return coordinate{x1, x2}
// }

func Penalty_method() {
	cr := coordinate{}
	for i := 0; i >= 0; i++ {

		fmt.Printf("Шаг вычисления минимума методом штрафных функций равен [%d]\n", i+1)
		fmt.Printf("x(%.6f , %.6f)", x01, x02)
		fmt.Printf(" r0(%.6f)", r0)
		fmt.Printf(" f(x) = %.6f\n\n", calcAnswer(x01, x02, r0))

		cr = CS_gradient_descent_method(x01, x02)
		x01 = cr.x01
		x02 = cr.x02
		t := calcPanswer(x01, x01, r0)

		if t <= E {
			fmt.Printf("ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZz")

			fmt.Printf("Величина 'r' равен [%.3f]\n", r0)
			fmt.Printf("x* = x(%.6f , %.6f)", x01, x01)
			fmt.Printf("f(x*) = %.6f\n\n", calcAnswer(x01, x01, r0))
			i = -1
			break
		} else {
			r0 = C * r0
		}
		//break
	}

}

func main() {
	// fmt.Fprintln(os.Stdout, "Ввод функции вида ʄ(x) = a(x1)^2 + b(x2)^2 - c --> min")
	// fmt.Fprintln(os.Stdout, "	Введите коэффициент (a) перед (x1^2)")
	// fmt.Fscan(os.Stdin, &a)
	// fmt.Fprintln(os.Stdout, "	Введите коэффициент (b) перед (x2^2)")
	// fmt.Fscan(os.Stdin, &b)
	// fmt.Fprintln(os.Stdout, "	Введите коэффициент (c) свободный")
	// fmt.Fscan(os.Stdin, &c)

	// fmt.Fprintln(os.Stdout, "Ввод функции вида g1(x) = a(x1) + b(x2) - c ")
	// fmt.Fprintln(os.Stdout, "	Введите коэффициент (a) перед (x1^2)")
	// fmt.Fscan(os.Stdin, &ag)
	// fmt.Fprintln(os.Stdout, "	Введите коэффициент (b) перед (x2^2)")
	// fmt.Fscan(os.Stdin, &bg)
	// fmt.Fprintln(os.Stdout, "	Введите коэффициент (c) свободный")
	// fmt.Fscan(os.Stdin, &cg)

	fmt.Fprintln(os.Stdout, "Введите x01")
	fmt.Fscan(os.Stdin, &x01)
	fmt.Fprintln(os.Stdout, "Введите x02")
	fmt.Fscan(os.Stdin, &x02)

	// fmt.Fprintln(os.Stdout, "Введите E")
	// fmt.Fscan(os.Stdin, &E)
	// fmt.Fprintln(os.Stdout, "Введите r^0")
	// fmt.Fscan(os.Stdin, &r0)
	// fmt.Fprintln(os.Stdout, "Введите C")
	// fmt.Fscan(os.Stdin, &C)
	r0 = 1
	k := 0
	for k != 3 {
		fmt.Fprintln(os.Stdout, "Метод штрафов - 1")
		fmt.Fprintln(os.Stdout, "Метод барьерных функций - 2")
		fmt.Fprintln(os.Stdout, "Для выхода - 3")
		fmt.Fscan(os.Stdin, &k)
		fmt.Print("\033[H\033[2J")
		switch k {
		case 1:
			Penalty_method()

		case 2:

		case 3:
			return
		default:
			k = 0
		}
	}

}
