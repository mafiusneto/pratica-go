# pratica-go
Projeto para pratica em go

Descrição dos passos e notas de conhecimento/lembretes no decorrer da pratica

1ª criação de sub-packages para organização e separação do projeto
    Uso de função em arquivo na mesma pasta, o package precisa ser o mesmo.
    Para arquivo em sub-pasta, é preciso criar o mod na subpasta e também declarar no go.work na pasta raiz.





# Comandos úteis


 # Execução
 // comando para rodar multiplos arquivos, com terminal acessando a pasta main
 go run .
 subistitui go run main.go outro.go


 # Criação work
 // comando para criar o go.work no diretorio raiz
 go work init ./{nomepastainit}

 // comando para usar pasta local
 go work use ./{nomepasta}


 # Modulo

 // comando para inicializar o modulo
 go mod init {nomemodulo}

 // comando para atualizar o mod, busca e baixar dependencias
 go mod tidy

