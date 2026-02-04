# Business Rules — capital-gains-api

Este documento descreve **as regras de negócio** para cálculo de imposto sobre **ganho de capital** em operações de compra/venda de ações.

> Esta aplicação suporta **dois modos de entrada**:
> 1) **CLI (stdin)**, linha de comando
> 2) **API HTTP (POST)**, um endpoint simples que processa **a mesma estrutura de entrada** do CLI e devolve **o mesmo formato de saída** (detalhado abaixo).

---

## 1) Escopo e conceitos

### 1.1 Tipos de operação
Cada operação possui:
- `operation`: `"buy"` (compra) ou `"sell"` (venda)
- `unit-cost`: preço unitário com duas casas decimais
- `quantity`: quantidade de ações


### 1.2 Saída esperada
Para cada lista de operações recebida, retornamos uma lista de objetos contendo:
- `tax`: imposto devido em cada operação, em JSON

A lista de saída **deve ter o mesmo tamanho** da lista de operações da entrada.

---

## 2) Estado, simulações e isolamento

### BR-STATE-01 — Estado em memória
A aplicação mantém estado **apenas em memória** e **não depende de banco externo**.

### BR-STATE-02 — Simulações independentes
- **CLI:** cada **linha** de entrada é uma simulação independente; não se carrega estado entre linhas.
- **API:** cada **requisição** é uma simulação independente; não se carrega estado entre requests (equivalente ao comportamento por linha no CLI).

### BR-STATE-03 — Restrições assumidas
- Nunca haverá venda de mais ações do que o total atualmente em posse.
- Podemos assumir que não haverá erro de parsing/contrato no JSON de entrada.

---

## 3) Regras de cálculo

### 3.1 Variáveis de domínio mantidas durante a simulação
Durante o processamento de uma lista (linha do CLI / request da API), manter:
- `shares`: quantidade atual de ações em posse
- `avg_price`: preço médio ponderado atual
- `accumulated_loss`: prejuízo acumulado para dedução futura
s
---

## 4) Regras de negócio (imposto)

### BR-TAX-01 — Alíquota
O imposto é **20%** sobre o **lucro** obtido na operação de venda, quando houver lucro.

### BR-TAX-02 — Compra não paga imposto
Operações de **compra (`buy`)** sempre retornam `tax = 0`.

### BR-TAX-03 — Isenção por valor total da operação de venda
Em uma venda (`sell`), **não há imposto** se o **valor total da operação** for `<= 20000.00`, onde:
- `operation_total = unit-cost * quantity`

**Importante:** a isenção depende do **valor total da venda**, e não do lucro.
Ainda assim, **prejuízos** e **deduções** continuam valendo normalmente (ver BR-TAX-06).

### BR-TAX-04 — Como apurar lucro/prejuízo em uma venda
Para uma venda:
- `profit_or_loss = (sell_unit_cost - avg_price) * quantity`

Se `profit_or_loss > 0`, há lucro (pode haver imposto, dependendo das outras regras).
Se `profit_or_loss < 0`, há prejuízo (ver BR-TAX-05).
Se `profit_or_loss == 0`, não há imposto nem prejuízo.

### BR-TAX-05 — Prejuízo: não paga imposto e acumula para o futuro
Quando houver prejuízo em uma venda:
- `tax = 0`
- `accumulated_loss += abs(profit_or_loss)`

### BR-TAX-06 — Dedução de prejuízo acumulado em lucros futuros
Quando houver lucro em uma venda:
- Primeiro deduzir o prejuízo acumulado:  
  `net_profit = max(0, profit - accumulated_loss)`  
  `accumulated_loss = max(0, accumulated_loss - profit)`
- O imposto incide sobre `net_profit` (se não houver isenção e `net_profit > 0`).

### BR-TAX-07 — Cálculo final do imposto (venda tributável)
Uma venda é **tributável** quando:
1) `operation_total > 20000.00` (BR-TAX-03)
2) há lucro líquido após dedução (`net_profit > 0`) (BR-TAX-06)

Então:
- `tax = 0.20 * net_profit`

---

## 5) Regras do preço médio ponderado

### BR-AVG-01 — Atualização do preço médio em compras
Quando ocorrer uma compra (`buy`), recalcular o **preço médio ponderado**:

`new_avg = ((current_qty * current_avg) + (buy_qty * buy_unit_cost)) / (current_qty + buy_qty)`

Após a compra:
- `shares += buy_qty`
- `avg_price = new_avg`

### BR-AVG-02 — Vendas não alteram o preço médio
Em uma venda, o `avg_price` **não muda**; apenas reduzimos `shares`:
- `shares -= sell_qty`

### BR-AVG-03 — Recompras após zerar posição
É permitido comprar após vender todas as ações; o preço médio deve refletir as compras relevantes até a venda atual. (Ver FAQ e Caso #7 na spec.)

---

## 6) Arredondamento e precisão

### BR-ROUND-01 — Arredondar para 2 casas decimais
Valores decimais devem ser arredondados para **2 casas decimais**.
Isso vale para:
- `avg_price`
- `tax`

---

## 7) Contratos de entrada/saída (CLI e API)

### 7.1 CLI (stdin/stdout)
- A aplicação lê **uma lista JSON por linha** via `stdin`.
- A última linha é vazia (encerra a leitura).
- Para cada linha, imprime **apenas** o JSON de saída correspondente (sem mensagens adicionais).

### 7.2 API HTTP (POST) — extensão do projeto
- Endpoint único (exemplo): `POST /taxes`
- Entrada: uma lista JSON de operações, **idêntica ao CLI** (mesmos campos e semântica).
- Saída: lista JSON com `{ "tax": ... }` por operação, **idêntica ao CLI**.

---

## 8) Não objetivos
- Persistir simulações em banco de dados.
- Validar inputs malformados (a spec assume que não ocorrerão).
- Suportar ativos além de ações ou impostos com regras diferentes das descritas aqui.
