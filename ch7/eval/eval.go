package eval

import (
	"fmt"
	"math"
	"testing"
)

// 定义一个表达式接口
type Expr interface {
	// 返回env中的表达式值
	Eval(env Env) float64
}

// Var类型表示对一个变量的引用
type Var string

// 表示一个浮点类型的常量
type literal float64

// 表示+,-运算表达式
type unary struct {
	op rune // one of '+','-'
	x  Expr
}

// 表示+,-,*,/运算表达式
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

type call struct {
	fn   string // one of "pow","sign","sqrt"
	args []Expr
}

// 变量映射
type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return +u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) * b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", c.fn))
}

func TestEval(t *testing.T) {

	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}

	var prevExpr string

	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
			expr, err := Parse(test.expr)
			if err != nil {
				t.Error(err)
				continue
			}
			got := fmt.Sprintf("%.6g", expr.Eval(test.env))

			fmt.Printf("\t%v => %s\n", test.env, got)

			if got != test.want {
				t.Errorf("%s.Eval() in %v = %q,want %q", test.expr, test.env, got, test.want)
			}
		}
	}

}
