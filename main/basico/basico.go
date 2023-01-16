package basico

import (
	"fmt"
	"reflect"
	"time"
)

var vglobal string

const vConstante = "constante"

// executa funcionalidades basicas
func ExecucaoBasicos() {
	pulaLinhaLog("inicio basico")
	helloWord()
	variaveisDeclaracao()
	variaveisTipos()
	usandoStruct()
	outputFormat()
	usandoDatas()
}

func pulaLinhaLog(msg string) {
	fmt.Println("")
	fmt.Println("  #####>>>", msg, "<<<#####----------------------")
}

// executa hello word
func helloWord() {
	fmt.Println("Ola mundo! em sub pasta")
}

func variaveisDeclaracao() {
	pulaLinhaLog("variaveisDeclaracao")
	vglobal = "global"

	vLocal := "local" // implementação variavel local , sem "var" inferindo direto com ":=" o valor e tipo

	// vConstante = "mudando" // gera erro tentar mudar uma cosntante

	fmt.Println("Variaveis:", vglobal, vLocal, vConstante)
}

func variaveisTipos() {
	pulaLinhaLog("variaveisTipos")
	// exibindo tipo
	fmt.Println("Exibir tipo:", reflect.TypeOf(vglobal), reflect.TypeOf(0), reflect.TypeOf(0.0), reflect.TypeOf(true))

	// texto
	var s1 string = "text"

	// inteiros
	var i1 int      // int int8 int16 int32 int64
	var i2 uint = 2 // quando inteiro inicia com "u" o tipo numerico é apenas valores positivos

	// numericos
	var d1 float32
	d2 := 0.1

	// booleano
	var b1 bool
	b2 := true

	fmt.Println("String s1:", s1, reflect.TypeOf(s1))
	fmt.Println("Inteiro i1:", i1, reflect.TypeOf(i1))
	fmt.Println("Inteiro i2:", i2, reflect.TypeOf(i2))
	fmt.Println("Decimal d1:", d1, reflect.TypeOf(d1))
	fmt.Println("Decimal d2:", d2, reflect.TypeOf(d2))
	fmt.Println("Booleano b1:", b1, reflect.TypeOf(b1))
	fmt.Println("Booleano b2", b2, reflect.TypeOf(b2))

	var a, b, c, d int = 1, 3, 5, 7
	fmt.Println("Atribuição multiplas variaveis:", a, b, c, d)

	var e, f = "fulano", 10
	fmt.Println("Atribuição multiplas variaveis diferentes:", e, f)

}

func usandoStruct() {
	pulaLinhaLog("usandoStruct")

	type pessoa struct {
		nome  string
		idade int
	}

	var p1 = pessoa{"Mel", 2}
	fmt.Println(reflect.TypeOf(p1), p1)

	var p2 pessoa
	p2.nome = "Gustavo"
	p2.idade = 10
	fmt.Println("Nome pessoa:", p2.nome, "Idade:", p2.idade)

}

func outputFormat() {
	pulaLinhaLog("outputFormat")

	fmt.Print("Não pula", " linha")
	fmt.Print(", agora pula", "\n", "e pula\n")
	fmt.Println("Pula linha por padrão no fim")
	fmt.Printf("Formata: valor: '%v' e tipo: '%T' \n", vglobal, vglobal)

	var i = 1234.56789
	var txt = "Hello word"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%.2f\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
	fmt.Printf("%q\n", txt)

}

func usandoDatas() {
	pulaLinhaLog("usandoDatas")
	data := time.Now()
	fmt.Println("Data:", data)
	fmt.Println("Data formato BR", data.Format(("02/01/2006 15:04:05 ")))
	fmt.Println("Data formato EN", data.Format(("2006-01-02 15:04:05 ")))
	fmt.Println("Data com fuso", data.Format(("2006-01-02 15:04:05 MST")))
	fmt.Println("Data formato EN", data.Format(("2006-JAN-02 15:04:05 ")))
	fmt.Println("Dia:", data.Format(("02, 2, Monday, Mon")))
	fmt.Println("Mês:", data.Format(("01, 1, January, Jan")))
	fmt.Println("Ano:", data.Format(("2006, 06")))
	fmt.Println("time:", data.Format(("03:04:05 PM")))

	/*
		O que queremos editar		código
		dia (com zero)				02
		dia (sem o zero)			2
		dia da semana (inteiro)		Monday
		dia da semana abreviado		Mon
		mês com número (com zero)	01
		mês com número (sem zero)	1
		mês (nome inteiro)			January
		mês (nome abreviado)		Jan
		ano (inteiro)				2006
		ano (abreviado)				06
		hora (com zero)				03
		hora (sem zero)				3
		hora (formato 24 horas)		15
		minutos (com zero)			04
		minutos (sem zero)			4
		segundos (com zero)			05
		segundos (sem zero)			5
		MST (fuso horário)			MST

		exemplos

		2006-01-02
		20060102
		January 02, 2006
		02 January 2006
		02-Jan-2006
		01/02/2006
		01/02/06
		010206
		Jan-02-06
		Jan-02-2006
		06
		Mon
		Monday
		Jan-06
	*/
}

func usandoLoops() {
	pulaLinhaLog("usandoLoops")

}

func usandoSwitch() {
	pulaLinhaLog("usandoSwitch")

}

func usandoOperacoes() {
	pulaLinhaLog("usandoOperacoes")

}

func usandoCondicoes() {
	pulaLinhaLog("usandoCondicoes")

}

func usandoArray() {
	pulaLinhaLog("usandoArray")

}

func usandoSlice() {
	pulaLinhaLog("usandoSlice")

}

func usandoMaps() {
	pulaLinhaLog("usandoMaps")

}
