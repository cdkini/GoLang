package interpret

import (
	"Okra/okra/interpreter/ast"
	"Okra/okra/okraerr"
)

type Callable interface {
	Arity() int
	Call(i *Interpreter, args []interface{}) interface{}
}

type Function struct {
	declaration ast.FuncStmt
}

func NewFunction(declaration ast.FuncStmt) *Function {
	return &Function{declaration}
}

func (f *Function) Arity() int {
	return len(f.declaration.Params)
}

func (f *Function) Call(i *Interpreter, args []interface{}) interface{} {
	env := NewEnvironment(i.globalEnv)
	for i, token := range f.declaration.Params {
		env.Define(token.Lexeme, args[i])
	}

	block := i.executeBlock(f.declaration.Body, env)
	if r, ok := block.(*ReturnValue); ok {
		return r.LiteralExpr.Val
	}
	return nil
}

type Struct struct {
	name    string
	methods map[string]*Function
}

func NewStruct(name string, methods map[string]*Function) *Struct {
	return &Struct{name, methods}
}

func (s *Struct) Arity() int {
	return 0
}

func (s *Struct) Call(i *Interpreter, args []interface{}) interface{} {
	return NewInstance(*s)

}

func (s *Struct) findMethod(method string) interface{} {
	if _, ok := s.methods[method]; ok {
		return s.methods[method]
	}
	return nil
}

type Instance struct {
	class  Struct
	fields map[string]interface{}
}

func NewInstance(class Struct) *Instance {
	return &Instance{class, make(map[string]interface{})}
}

func (i *Instance) get(property ast.Token) interface{} {
	if val, ok := i.fields[property.Lexeme]; ok {
		return val
	}
	if method := i.class.findMethod(property.Lexeme); method != nil {
		return method
	}
	okraerr.ReportErr(0, 0, "Undefined property "+property.Lexeme+".")
	return nil
}

func (i *Instance) set(property ast.Token, val interface{}) {
	i.fields[property.Lexeme] = val
}
