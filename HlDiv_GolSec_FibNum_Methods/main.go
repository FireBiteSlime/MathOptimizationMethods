package main

import (
	"fmt"
	"math"
	"os"
)

var a, b, c, d float64

func Half_Division_Method() {
	var a0, b0, E float64
	var x, y, z, L, Fx, Fy, Fz float64

	fmt.Fprintln(os.Stdout, "Начальный интервал неопределённости L0 = [a0, b0], введите a0")
	fmt.Fscan(os.Stdin, &a0)
	fmt.Fprintln(os.Stdout, "Начальный интервал неопределённости L0 = [a0, b0], введите b0")
	fmt.Fscan(os.Stdin, &b0)
	fmt.Fprintln(os.Stdout, "Введите требуемую  точность E")
	fmt.Fscan(os.Stdin, &E)

	x = (a0 + b0) / 2
	var i float64
	for i = 0; i >= 0; i++ {

		fmt.Printf("Шаг равен [%.0f]\n", i+1)
		fmt.Printf("Интервал [a0, b0] = [%.3f , %.3f]\n", a0, b0)

		L = b0 - a0
		Fx = a*math.Pow(x, 3) - b*math.Pow(x, 2) - c*x + 1
		y = a0 + math.Abs(L)/4
		z = b0 - math.Abs(L)/4
		Fy = a*math.Pow(y, 3) - b*math.Pow(y, 2) - c*y + 1
		Fz = a*math.Pow(z, 3) - b*math.Pow(z, 2) - c*z + 1

		if Fy < Fx {
			b0 = x
			//a0 = a0
			x = y
		} else if Fy > Fx {
			if Fz < Fx {
				a0 = x
				x = z
			} else if Fz >= Fx {
				a0 = y
				b0 = z
				//x = x
			}
		}
		L = math.Abs(b0 - a0)
		if math.Abs(L) <= E {
			fmt.Printf("x* принадлежит отрезку [%.3f, %.3f]\n", a0, b0)
			fmt.Printf("Середина данного интервала x* = %.3f\n", x)
			fmt.Printf("Характиристика относительного уменьшения начального интервала неопределённости R(%.0f) = %.3f\n\n", i+1, 1/(math.Pow(2, i/2)))
			break
		}

		fmt.Printf("x равен [%.3f]\n", x)
		fmt.Printf("z равен [%.3f]\n", z)
		fmt.Printf("y равен [%.3f]\n", y)
		fmt.Printf("Fy равна [%.3f]\n", Fx)
		fmt.Printf("Fz равна [%.3f]\n", Fy)
		fmt.Printf("Fz равна [%.3f]\n", Fz)
		fmt.Printf("L2k равна [%.3f]\n\n", L)
	}
}

func Golden_Section_Method() {
	var a0, b0, E float64
	var y, z, Fy, Fz, delta float64

	fmt.Fprintln(os.Stdout, "Начальный интервал неопределённости L0 = [a0, b0], введите a0")
	fmt.Fscan(os.Stdin, &a0)
	fmt.Fprintln(os.Stdout, "Начальный интервал неопределённости L0 = [a0, b0], введите b0")
	fmt.Fscan(os.Stdin, &b0)
	fmt.Fprintln(os.Stdout, "Введите требуемую  точность E")
	fmt.Fscan(os.Stdin, &E)

	var i float64
	y = a0 + ((3-math.Sqrt(5))/2)*(b0-a0)
	z = a0 + b0 - y
	for i = 0; i >= 0; i++ {

		fmt.Printf("Шаг равен [%.0f]\n", i+1)
		fmt.Printf("Интервал [a0, b0] = [%.3f , %.3f]\n", a0, b0)

		Fy = a*math.Pow(y, 3) - b*math.Pow(y, 2) - c*y + 1
		Fz = a*math.Pow(z, 3) - b*math.Pow(z, 2) - c*z + 1

		if Fy <= Fz {
			//a0 = a0
			b0 = z
			z = y
			y = a0 + b0 - y
		} else {
			a0 = y
			//b0 = b0
			y = z
			z = a0 + b0 - z

		}
		delta = math.Abs(a0 - b0)
		if delta <= E {
			fmt.Printf("x* принадлежит отрезку [%.3f, %.3f]\n", a0, b0)
			fmt.Printf("В качестве приближения можно взять середину этого интервала x* = %.3f\n", (a0+b0)/2)
			fmt.Printf("Характиристика относительного уменьшения начального интервала неопределённости R(%.0f) = %.3f\n\n", i+1, math.Pow((math.Sqrt(5)-1)/2, i-1))
			break
		}

		fmt.Printf("z равен [%.3f]\n", z)
		fmt.Printf("y равен [%.3f]\n", y)
		fmt.Printf("Fy равна [%.3f]\n", Fy)
		fmt.Printf("Fz равна [%.3f]\n", Fz)
		fmt.Printf("Delta равна [%.3f]\n\n", delta)

	}
}

func Fibonacci_Number_Method() {
	var a0, b0, E, l float64
	var y, z, Fy, Fz, Fn, yn2, zn2 float64
	Mf := make([]float64, 0)
	var N int

	fmt.Fprintln(os.Stdout, "Начальный интервал неопределённости L0 = [a0, b0], введите a0")
	fmt.Fscan(os.Stdin, &a0)
	fmt.Fprintln(os.Stdout, "Начальный интервал неопределённости L0 = [a0, b0], введите b0")
	fmt.Fscan(os.Stdin, &b0)
	fmt.Fprintln(os.Stdout, "Введите допустимую длину конечного интервала l > 0")
	fmt.Fscan(os.Stdin, &l)
	fmt.Fprintln(os.Stdout, "Введите константу различимости E > 0")
	fmt.Fscan(os.Stdin, &E)

	Mf = append(Mf, 1)
	Mf = append(Mf, 1)
	Fn = Mf[0]
	for i := 2; Fn < math.Abs(b0-a0)/l; i++ {
		//Mf[i] = Mf[i-1] + Mf[i-2]

		Mf = append(Mf, Mf[i-1]+Mf[i-2])
		Fn = Mf[i]
		N = i + 1
	}

	fmt.Printf("Числа Фибоначчи = ")
	for index := 0; index < len(Mf); index++ {
		fmt.Printf("%.0f ", Mf[index])
	}
	fmt.Printf("N = '%d'\n\n", N)
	var k int

	y = a0 + (Mf[N-3]/Mf[N-1])*(b0-a0)
	z = a0 + (Mf[N-2]/Mf[N-1])*(b0-a0)

	for k = 0; k >= 0; k++ {

		//fmt.Printf(" N k = %d %d\n", N, k)
		fmt.Printf("Шаг равен [%d]\n", k+1)
		fmt.Printf("Интервал [a0, b0] = [%.3f , %.3f]\n", a0, b0)

		Fy = a*math.Pow(y, 3) - b*math.Pow(y, 2) - c*y + 1
		Fz = a*math.Pow(z, 3) - b*math.Pow(z, 2) - c*z + 1
		if Fy <= Fz {
			//a0 = a0
			b0 = z
			z = y
			y = a0 + (Mf[N-k-3]/Mf[N-k-1])*(b0-a0)
		} else {
			a0 = y
			//b0 = b0
			y = z
			z = a0 + (Mf[N-k-2]/Mf[N-k-1])*(b0-a0)
		}

		if k == N-3 {
			zn2 = (a0 + b0) / 2
			yn2 = zn2
			z = y + E
			y = yn2
			Fy = a*math.Pow(y, 3) - b*math.Pow(y, 2) - c*y + 1
			Fz = a*math.Pow(z, 3) - b*math.Pow(z, 2) - c*z + 1

			if Fy <= Fz {
				//a0 = a0
				b0 = z
			} else {
				a0 = y
				//b0 = b0
			}
			fmt.Printf("x* принадлежит отрезку [%.3f, %.3f]\n", a0, b0)
			fmt.Printf("В качестве приближения можно взять середину этого интервала x* = %.3f\n", (a0+b0)/2)
			fmt.Printf("Характиристика относительного уменьшения начального интервала неопределённости R(%d) = %.3f\n\n",
				k+1, 1/Fn)
			break
		}

		fmt.Printf("z равен [%.3f]\n", z)
		fmt.Printf("y равен [%.3f]\n", y)
		fmt.Printf("Fy равна [%.3f]\n", Fy)
		fmt.Printf("Fz равна [%.3f]\n\n", Fz)

	}
}

func main() {
	k := 0
	fmt.Fprintln(os.Stdout, "Введите коэффициент (a) перед (x^3)")
	fmt.Fscan(os.Stdin, &a)
	fmt.Fprintln(os.Stdout, "Введите коэффициент (b) перед (x^2)")
	fmt.Fscan(os.Stdin, &b)
	fmt.Fprintln(os.Stdout, "Введите коэффициент (c) перед (x)")
	fmt.Fscan(os.Stdin, &c)
	fmt.Fprintln(os.Stdout, "Введите коэффициент (d) свободный")
	fmt.Fscan(os.Stdin, &d)

	for k != 4 {
		fmt.Fprintln(os.Stdout, "Метод половинного деления - 1")
		fmt.Fprintln(os.Stdout, "Метод золотого сечения - 2")
		fmt.Fprintln(os.Stdout, "Метод чисел Фибоначчи - 3")
		fmt.Fprintln(os.Stdout, "Для выхода - 4")
		fmt.Fscan(os.Stdin, &k)
		fmt.Print("\033[H\033[2J")
		switch k {
		case 1:
			Half_Division_Method()

		case 2:
			Golden_Section_Method()
		case 3:
			Fibonacci_Number_Method()
		case 4:
			return
		default:
			k = 0
		}
	}

}
