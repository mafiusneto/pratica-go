package basico

import (
	"fmt"
	"reflect"
	"time"
)

var vglobal string

const vConstante = "constante"

type pessoa struct {
	nome  string
	idade int
}

type localizacao struct {
	endereco string
	numero   int
}

// executa funcionalidades basicas
func ExecucaoBasicos() {
	pulaLinhaLog("inicio basico")
	helloWord()
	variaveisDeclaracao()
	variaveisTipos()
	usandoStruct()
	outputFormat()
	usandoDatas()
	usandoLoops()
	usandoSwitch()
	usandoOperacoes()
	usandoCondicoes()

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

	//padrão
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Print(i, ",")
	}
	fmt.Print("\n")

	//loop in array/slice
	frutas := [...]string{"banana", "maçã", "uva"}
	for idx, value := range frutas {
		fmt.Println(idx, value)
	}
	fmt.Println("")

	for idx := range frutas {
		fmt.Println(idx, "-", frutas[idx])
	}
	fmt.Print("\n")

	for _, value := range frutas {
		fmt.Println(value)
	}

	//loop "infinito"
	i := 0
	for {
		fmt.Print(i)
		if i > 20 {
			fmt.Print("\n")
			break
		}
		i++
		fmt.Print(",")
	}

	for b := 0; ; b++ {
		fmt.Print(b)
		if b > 20 {
			fmt.Print("\n")
			break
		}
		fmt.Print(",")
	}

}

func usandoSwitch() {
	pulaLinhaLog("usandoSwitch")
	dia := 4

	switch dia {
	case 1:
		fmt.Println("Segunda")
	case 2:
		fmt.Println("Terça")
	case 3:
		fmt.Println("Quarta")
	case 4:
		fmt.Println("Quinta")
	case 5:
		fmt.Println("Sexta")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("dia errado")
	}

	mes := 4
	switch mes {
	case 1, 2, 3:
		fmt.Println("1º trimestre")
	case 4, 5, 6:
		fmt.Println("2º trimestre")
	case 7, 8, 9:
		fmt.Println("3º trimestre")
	case 10, 11, 12:
		fmt.Println("4º trimestre")
	default:
		fmt.Println("Mês invalido")
	}

	//switch de struct
	var p pessoa
	var l localizacao
	switchStructTeste(p)
	switchStructTeste(l)

	// switch por tip ode variavel
	switchTipoVariavel(10)
	switchTipoVariavel("oi")
	switchTipoVariavel(false)
	switchTipoVariavel(1.3)
	switchTipoVariavel(p)
}

func switchStructTeste(m interface{}) {

	switch m.(type) {
	case pessoa:
		fmt.Println("struct pessoa")
	case localizacao:
		fmt.Println("struct localizacao")
	default:
		fmt.Println("struct localizacao")
	}
}

func switchTipoVariavel(t interface{}) {
	fmt.Print(t, " = ")
	switch t.(type) {
	case int:
		fmt.Println("inteiro")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("booleano")
	case float32, float64:
		fmt.Println("float")
	default:
		fmt.Println("não identificado")
	}
}

func usandoOperacoes() {
	pulaLinhaLog("usandoOperacoes")

	// + - * / %
	var a int = 1 + 1
	b := 3 - 1
	c := 2 * 4
	d := 5 / 2
	e := 11 % 3
	f := 3
	f++
	g := 3
	g--
	fmt.Println(a, b, c, d, e, f, g)

	a = f
	b += 1
	c -= 1
	d *= 2
	e /= 3
	f %= 2
	g &= 3
	/*
		|=
		^=
		>>=
		<<=
	*/
	fmt.Println(a, b, c, d, e, f, g)

	/*
		comparação
		==
		!=
		>
		<
		>=
		<=


		logico
		&& 	and
		||	or
		!	negação

		bitwise
		&
		|
		^
		<<   Add 0  >>    12 << 2 = 1200
		>>  remove  >>    1234 >> 2 = 12
	*/

}

func usandoCondicoes() {
	pulaLinhaLog("usandoCondicoes")
	/*
		==
		>
		>=
		<
		<=
		!=

		&&
		||
		!
	*/

	a, b, c, d, e := 1, 2, 3, 1, 2

	if a == d {
		fmt.Println("igual")
	}

	if b > c {
		fmt.Println("maior")
	} else {
		fmt.Println("menor")
	}

	if (b > a) && (d < e) { // os 2 verdades
		fmt.Println("verdade")
	} else {
		fmt.Println("falso")
	}

	if (a > d) || (a == d) { // uma ou outra
		fmt.Println("verdade")
	} else {
		fmt.Println("falso")
	}

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
