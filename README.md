## Desafio Stress Test

Objetivo: Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

O sistema deverá gerar um relatório com informações específicas após a execução dos testes.
Entrada de Parâmetros via CLI:

- --url: URL do serviço a ser testado.
- --requests: Número total de requests.
- --concurrency: Número de chamadas simultâneas.

Execução do Teste:
Realizar requests HTTP para a URL especificada.
Distribuir os requests de acordo com o nível de concorrência definido.
Garantir que o número total de requests seja cumprido.


Geração de Relatório:
Apresentar um relatório ao final dos testes contendo:
Tempo total gasto na execução
Quantidade total de requests realizados.
Quantidade de requests com status HTTP 200.
Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Exemplos

- `docker run stress-test --url=https://jsonplaceholder.typicode.com/todos/1 --requests=10000 --concurrency=10000`

- `docker run stress-test --url=https://jsonplaceholder.typicode.com/todos/exemplo-not-found --requests=10000 --concurrency=10000`

## Build

- `docker build -t stress-test .`
