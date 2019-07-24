package chapter7

import (
	"fmt"
	"math"
	"strings"
)

type Var string

type literal float64

type Expr interface {
	Eval(env Env) float64 //返回每个表达式在Env这个上下文中对应的具体的值

	Check(vars map[Var]bool) error
}

//一元操作符[-,+]
type unary struct {
	op rune
	x  Expr
}

//二元操作符x[+|-|*|/]y
type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string //pow|sin|sqrt
	args []Expr
}

//上下文
type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}
func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (l literal) Eval(env Env) float64 {
	return float64(l)
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env) //递归调用，当env是一个复杂类型的时候，比如说是call的时候，先调用call自身的Eval方法，将返回的结果再执行这里的方法
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unSupport unary operation:%q\n", u.op))
}
func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q\n", u.op)
	}
	return u.x.Check(vars) //递归检查，可能u的属性x是一个复合属性，比如(x+y) --> binary
}

func (bn binary) Eval(env Env) float64 {
	switch bn.op {
	case '+':
		return bn.x.Eval(env) + bn.y.Eval(env)
	case '-':
		return bn.x.Eval(env) - bn.y.Eval(env)
	case '*':
		return bn.x.Eval(env) * bn.y.Eval(env)
	case '/':
		return bn.x.Eval(env) / bn.y.Eval(env)
	}
	panic(fmt.Sprintf("unSupport binary operation:%q\n", bn.op))
}
func (bn binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op:%q\n", b.op)
	}
	if err := bn.x.Check(vars); err != nil {
		return err
	}

	return bn.y.Check(vars)
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
	panic(fmt.Sprintf("unSupport call:%s", c.fn))
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok{
		return fmt.Errorf("unknown function %q\n",c.fn)
	}
	if len(c.args) != arity{
		return fmt.Errorf("call to %s has %d args,want %d",c.fn,len(c.args),arity)
	}

	for _,arg := range c.args{
		//检查每一个变量
		if err := arg.Check(vars);err != nil{
			return err
		}
	}
	return nil
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}
//测试
func TestEval() {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x,3) + pow(y,3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x,3) + pow(y,3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}

	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}

		//expr,err := Parse
	}
}
