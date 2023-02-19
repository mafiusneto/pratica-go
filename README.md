# pratica-go
## Projeto para pratica em go

Descrição dos passos e notas de conhecimento/lembretes no decorrer da pratica

1º criação de sub-packages para organização e separação do projeto
    Uso de função em arquivo na mesma pasta, o package precisa ser o mesmo.
    Para arquivo em sub-pasta, é preciso criar o mod na subpasta e também declarar no go.work na pasta raiz.

2º dependendo ta etapa do projeto a execução da funcão pode estar comentada no main

## Lista exemplos basicos
    helloWord - exemplo de hello word
	variaveisDeclaracao - exemplos de declaração de variaveis
	variaveisTipos - exemplo de tipos de variaveis
	usandoStruct - exemplo de uso de struct
	outputFormat - exemplo de saidas output e formatação
	usandoDatas - exemplos de uso de datas e formatação
	usandoLoops - exemplos de loops (for)
	usandoSwitch - exemplos usando switch
	usandoOperacoes - exemplos e lista de tipos de operações
	usandoCondicoes - exemplos e lista de condições
	usandoArray - exemplos de uso do array
	usandoSlice - exemplos de uso de slice
	usandoMaps - exemplos de uso de maps
	usandoPonteiro - exemplos de uso de ponteiro

	usandoInterface - exemplo de uso de interface
	usandoGoroutine - exemplo de uso de goroutine
	usandoChannel - exemplo de uso de channels
	usandoRecursosEspeciais - uso de defer e time.sleep
	usandoConcorrencia - exemplo de concorrencia usando channels
	usandoMetodosDeStruct - exemplos de uso de struct
	usandoConversoes - exemplos de conversões de tipos basicos
	convertStructEmJson - exemplo de uso de json (encode e decode)
	usandoVariaveisDeAmbiente - exemplo e lista de uso de variaveis de ambiente
	usandoErrorEStringers - exemplo de uso de metodos como error e string

	usandoHttp - exemplo de uso de http para API


## Comandos úteis


 ### Execução
 // comando para rodar multiplos arquivos, com terminal acessando a pasta main
 go run .
 subistitui go run main.go outro.go


 ### Criação work
 // comando para criar o go.work no diretorio raiz
 go work init ./{nomepastainit}

 // comando para usar pasta local
 go work use ./{nomepasta}


 ### Modulo

 // comando para inicializar o modulo
 go mod init {nomemodulo}

 // comando para atualizar o mod, busca e baixar dependencias
 go mod tidy

