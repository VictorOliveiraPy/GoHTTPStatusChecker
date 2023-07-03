
```
# Go URL Checker

Este é um programa simples escrito em Go que verifica o status de URLs usando goroutines.

## Pré-requisitos

Certifique-se de ter o Go instalado em sua máquina. Você pode fazer o download e instalá-lo em https://golang.org/dl/.

## Instalação

1. Clone este repositório:

```shell
git clone https://github.com/VictorOliveiraPy/GoHTTPStatusChecker
```

2. Navegue até o diretório do projeto:

```shell
cd go-url-checker
```

3. Baixe as dependências do projeto usando o Go Modules:

```shell
go mod download
```

## Execução

Para executar o programa, use o seguinte comando:

```shell
go run process.go
```

## Fluxo e Uso de Goroutines

1. O programa inicia com uma lista de URLs pré-definida no código.


2. As URLs são enviadas para um canal (channel) chamado `urlChannel`.


3. A função `readURLs` lê as URLs do canal `urlChannel` e dispara goroutines para fazer as requisições HTTP usando a função `requestURL`.


4. A função `requestURL` cria uma solicitação HTTP para cada URL e faz a verificação de status. O resultado do status é enviado para o canal `responseChannel`.


5. A função `readResponses` lê as respostas do canal `responseChannel` e imprime os resultados.


6. O programa aguarda a conclusão de todas as goroutines usando `sync.WaitGroup` para garantir que todas as requisições sejam processadas antes de encerrar.


