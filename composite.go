package main

import (
	"fmt"
)

type Organizacao interface {
	CalcularSalarioTotal() float64
}

type FuncionarioIndividual struct {
	nome    string
	salario float64
}

func (f *FuncionarioIndividual) CalcularSalarioTotal() float64 {
	return f.salario
}

type Departamento struct {
	nome         string
	funcionarios []Organizacao
}

func (d *Departamento) AdicionarFuncionario(f Organizacao) {
	d.funcionarios = append(d.funcionarios, f)
}

func (d *Departamento) RemoverFuncionario(f Organizacao) {
	for i, funcionario := range d.funcionarios {
		if funcionario == f {
			d.funcionarios = append(d.funcionarios[:i], d.funcionarios[i+1:]...)
			break
		}
	}
}

func (d *Departamento) CalcularSalarioTotal() float64 {
	total := 0.0
	for _, funcionario := range d.funcionarios {
		total += funcionario.CalcularSalarioTotal()
	}
	return total
}

func (d *Departamento) ExibirEstrutura() {
	fmt.Printf("Departamento: %s, Número de Funcionários: %d\n", d.nome, len(d.funcionarios))
	for _, f := range d.funcionarios {
		switch f := f.(type) {
		case *FuncionarioIndividual:
			fmt.Printf("  - Funcionario: %s, Salario: %.2f\n", f.nome, f.salario)
		case *Departamento:
			f.ExibirEstrutura()
		}
	}
}

func main() {
	func1 := &FuncionarioIndividual{"Amanda", 2005}
	func2 := &FuncionarioIndividual{"Lucio", 2003}
	func3 := &FuncionarioIndividual{"Tufo", 2024}
	dptTI := &Departamento{"TI", []Organizacao{}}
	dptTI.AdicionarFuncionario(func2)

	dptModa := &Departamento{"Moda", []Organizacao{}}
	dptModa.AdicionarFuncionario(func1)
	dptModa.AdicionarFuncionario(func3)

	departamentoPrincipal := &Departamento{"Principal", []Organizacao{}}
	departamentoPrincipal.AdicionarFuncionario(dptTI)
	departamentoPrincipal.AdicionarFuncionario(dptModa)

	salarioTotal := departamentoPrincipal.CalcularSalarioTotal()
	fmt.Printf("Salário total da organização: %.2f\n", salarioTotal)

	departamentoPrincipal.ExibirEstrutura()
}
