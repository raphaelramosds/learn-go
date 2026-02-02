package main

import (
	"fmt"
)

type Ser struct {
	idade int8
	nome  string
}

func (s *Ser) fazerAniversario() {
	s.idade++
}

type ISer interface {
	apresentar()
	fazerAniversario()
}

// Heranca em Go funciona com a Composicao de atributos entre structs
type Cachorro struct {
	Ser
	dono *Pessoa
}

func (p *Pessoa) apresentar() {
	fmt.Printf("Sou uma pessoa, me chamo %s\n", p.nome)
}

func (c *Cachorro) apresentar() {
	fmt.Printf("Au au\n")
}

func (c *Cachorro) apresentarDono() {
	fmt.Printf("Meu dono eh %s\n", c.dono.nome)
}

type Pessoa struct {
	Ser
}

// Interfaces sao enviadas por referencia por padrao
// Por isso seria redundante digitar "ser* ISer"
func apresentar(ser ISer) {
	ser.apresentar()
}

func crescer(ser ISer) {
	ser.fazerAniversario()
}

func main() {
	raphael := &Pessoa{Ser: Ser{idade: 24, nome: "Raphael"}}
	black := &Cachorro{Ser: Ser{idade: 10, nome: "Black"}, dono: raphael}
	black.apresentarDono()

	// Demonstracao de Polimorfismo: mudando o ponteiro, mudamos a implementacao de "ISer.apresentar"
	apresentar(raphael)
	apresentar(black)

	// Como passamos por referencia, a idade do objeto Pesso eh realmente modificada
	fmt.Println(raphael.idade)
	crescer(raphael)
	fmt.Println(raphael.idade)

	// NOTE: Pessoa nao implementa fazerAniversario, entao porque o metodo crescer nao quebra?
	// R: ha uma promocao do metodo fazerAniversario do struct Ser para o struct Pessoa
}
