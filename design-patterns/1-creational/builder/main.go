package main

func main() {
	director := Director{}

	builderA := ConcreteBuilderA{}
	director.ConstructProduct(&builderA)
	productA := builderA.GetResult()

	builderB := ConcreteBuilderB{}
	director.ConstructProduct(&builderB)
	productB := builderB.GetResult()
}

type Builder interface {
	SetOptionA(a string)
	SetOptionB(b string)
	SetOptionC(c string)
}

type ConcreteBuilderA struct {
	A string
	B string
	C string
}

func (aa *ConcreteBuilderA) SetOptionA(a string) {
	aa.A = a
}

func (aa *ConcreteBuilderA) SetOptionB(b string) {
	aa.B = b
}

func (aa *ConcreteBuilderA) SetOptionC(c string) {
	aa.C = c
}

func (aa *ConcreteBuilderA) GetResult() ProductA {
	return ProductA{
		A: aa.A,
		B: aa.B,
		C: aa.C,
	}
}

type ConcreteBuilderB struct {
	A string
	B string
	C string
}

func (bb *ConcreteBuilderB) SetOptionA(a string) {
	bb.A = a
}

func (bb *ConcreteBuilderB) SetOptionB(b string) {
	bb.B = b
}

func (bb *ConcreteBuilderB) SetOptionC(c string) {
	bb.C = c
}

func (bb *ConcreteBuilderB) GetResult() ProductB {
	return ProductB{
		A: bb.A,
		B: bb.B,
		C: bb.C,
	}
}

type ProductA struct {
	A string
	B string
	C string
}

type ProductB struct {
	A string
	B string
	C string
}

type Director struct {
}

func (d *Director) ConstructProduct(b Builder) {
	b.SetOptionA("a")
	b.SetOptionB("b")
	b.SetOptionC("c")
}
