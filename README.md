# capital-gains

O objetivo desta aplicação é calcular o imposto a ser pago sobre lucros ou prejuízos de operações no mercado financeiro de ações.


## documentação funcional

As regras de negócio podem ser consultadas no arquivo [Business Rules – Capital Gains](docs/business.md)


## build e run com docker

Execute os seguintes comandos na raiz do projeto para construir a imagem e executá-la:

```bash
docker build -t capital-gains .
docker run -it capital-gains
```


## executar localmente

Requer a instalação do GO v1.26: https://go.dev/dl/

Download das dependências a partir da raiz do projeto e execução na pasta  cmd/cli

```bash
go mod tidy
go run main.go
```


## build local

Requer a instalação do GO v1.21: https://go.dev/dl/

Download das dependências a partir da raiz do projeto e execução pelo executável:

```bash
go mod tidy
go build -o capital-gains
./capital-gains
```


## teste

Unitários e de integração:

```bash
go test ./... -cover
```

Mutação: Rodar os comandos (a partir da raiz do projeto) para instalar o binário do mutante na maquina e posteriormente a execução.

```bash
go get -u github.com/zimmski/go-mutesting/...
go-mutesting ./...
```


## observações: testes

- Os testes unitários estão distribuídos dentro das pastas de seus respectivos pacotes.
- Os testes de integração estão centralizados na pasta cmd/cli e testam a integração a partir da main.go.
- Foi adotada a estratégia de centralizar o código dos testes integrados dentro de uma única função, que é chamada para os diversos casos de uso. Apesar de compreender que, em um cenário produtivo, isso pode trazer desvantagens, acredito que, para este caso, só há ganhos, visto a quantidade de código repetido e a baixa manutenabilidade que o arquivo teria.
- Há também uma pasta test/integration/data, na qual estão cenários de exemplo. O nome da pasta "test" é apenas para deixar mais evidente o objetivo dos arquivos ali presentes.
- Infelizmente, por causa de compromissos pessoais, não consegui investir mais tempo para cobrir mais ramificações lógicas.


### observações: Evolução

- O objetivo é ter uma solução simples e flexível, com poucas dependências. Nesse sentido, creio que ficaram como oportunidades melhorias como logs de debug, camada de repositório, injeção de dependência desacoplada, entre outras.
- A segregação da camada de input via CLI permite facilmente a integração com uma camada de API ou qualquer outro tipo de input, como mensageria por exemplo etc.


