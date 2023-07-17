// ref: https://www.digitalocean.com/community/tutorials/how-to-use-variables-and-constants-in-go-pt
// ref: https://diegomoal.medium.com/golang-tipos-de-dados-3774fc0e0b91

package main

import "fmt"

func main() {

	// Declaração e Inicialização de Variáveis

	fmt.Println("Declaração e Inicialização de Variáveis")

	var a int // declaração com o tipo
	a = 1     // inicialização
	fmt.Printf("var a int\t// declaração com o tipo\na = %d\t\t// inicialização\n\n", a)

	var b int = 2 // declaração com o tipo e inicialização na mesma linha
	fmt.Printf("var b int = %d\t// declaração com o tipo e inicialização na mesma linha\n\n", b)

	var c = 3 // declaração e inicialização sem definir o tipo
	fmt.Printf("var c = %d\t// declaração e inicialização sem definir o tipo\n\n", c)

	d := 4 // declaração e inicialização curta
	fmt.Printf("d := %d\t\t// declaração e inicialização curta\n", d)

	fmt.Println("")

	// Tipos Primitivos e Valor Zero

	var e string
	var f rune
	var g int8
	var h int16
	var i int32
	var j int64
	var k float32
	var l float64
	var m bool
	var n byte
	var o uint8
	var p uint16
	var q uint32
	var r uint64

	fmt.Println("Tipos Primitivos e Valor Zero")
	fmt.Printf("var e %T, valor padrão: %s\n", e, e)
	fmt.Printf("var f %T, valor padrão: %d\n", f, f)
	fmt.Printf("var g %T, valor padrão: %d\n", g, g)
	fmt.Printf("var h %T, valor padrão: %d\n", h, h)
	fmt.Printf("var i %T, valor padrão: %d\n", i, i)
	fmt.Printf("var j %T, valor padrão: %d\n", j, j)
	fmt.Printf("var k %T, valor padrão: %f\n", k, k)
	fmt.Printf("var l %T, valor padrão: %f\n", l, l)
	fmt.Printf("var m %T, valor padrão: %t\n", m, m)
	fmt.Printf("var n %T(byte), valor padrão: %d\n", n, n)
	fmt.Printf("var o %T, valor padrão: %d\n", o, o)
	fmt.Printf("var p %T, valor padrão: %d\n", p, p)
	fmt.Printf("var q %T, valor padrão: %d\n", q, q)
	fmt.Printf("var r %T, valor padrão: %d\n", r, r)
	fmt.Println()

	// Nomeando Variáveis

	// Os nomes das variáveis devem consistir em apenas uma palavra (ou seja, sem nenhum espaço).
	// Os nomes das variáveis devem ser constituídos apenas de letras, números e sublinhados (_).
	// Os nomes das variáveis não podem começar com um número.

	fmt.Println("Nomeando Variáveis")
	fmt.Println("// Os nomes das variáveis devem consistir em apenas uma palavra (ou seja, sem nenhum espaço).")
	fmt.Println("// Os nomes das variáveis devem ser constituídos apenas de letras, números e sublinhados (_).")
	fmt.Println("// Os nomes das variáveis não podem começar com um número.")
	fmt.Println()
	fmt.Println("válido: userName,	inválido: user-name		Hifens não são permitidos")
	fmt.Println("válido: name4,		inválido: 4name			Não é possível começar com um número")
	fmt.Println("válido: user,		inválido: $user			Não é possível usar símbolos")
	fmt.Println("válido: userName,	inválido: user name		Não é possível haver mais de uma palavra")
	fmt.Println()

	// Nível de Acesso

	var Publica string = "É possível acessar fora do pacote"
	var privada string = "Acesso somento dentro do pacote"

	fmt.Println("Nível de Acesso")
	fmt.Printf("Variávies com a primeira letra maíscula são públicas e com a primeira letra minúscula são privadas\n\n")
	fmt.Println("var Publica string - ", Publica)
	fmt.Println("var privada string - ", privada)
	fmt.Println()

}
