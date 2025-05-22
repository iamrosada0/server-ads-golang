
---

# Ads Server com Geolocalização e Filtragem por Navegador

Este projeto é um servidor HTTP em Go que realiza uma "leilão" simples de campanhas publicitárias com base na localização geográfica do usuário (país) e no navegador utilizado. A ideia é direcionar usuários para links de campanhas relevantes ao seu perfil.

---

## Funcionalidades

* Detecta o país do usuário a partir do IP usando a base de dados GeoLite2 Country (MaxMind).
* Identifica o navegador do usuário via User-Agent HTTP.
* Filtra campanhas disponíveis com base no país e navegador do usuário.
* Realiza um leilão simples, escolhendo a campanha com o menor preço entre as que se encaixam no perfil.
* Redireciona o usuário para o link da campanha vencedora.
* Caso não haja campanhas compatíveis, retorna erro 404.

---

## Estrutura do Projeto

* **ads/** - Pacote principal que contém as definições das campanhas, usuário, filtros e lógica do leilão (auction).
* **server.go** - Arquivo principal que inicializa o servidor HTTP, processa requisições, obtém IP, país e navegador, e chama a função de leilão.
* **GeoLite2-Country.mmdb** - Base de dados para geolocalização (não incluída no repositório, precisa ser baixada separadamente).

---

## Requisitos

* Go 1.18+ instalado
* Banco de dados GeoLite2-Country.mmdb (baixe em: [https://dev.maxmind.com/geoip/geolite2-free-geolocation-data](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data))
* Biblioteca Go `github.com/oschwald/geoip2-golang` para leitura do banco GeoIP

---

## Como rodar

1. Baixe o banco GeoLite2-Country:

```bash
wget https://geolite.maxmind.com/download/geoip/database/GeoLite2-Country.mmdb.gz
gunzip GeoLite2-Country.mmdb.gz
```
Podes ir nesse repo e baixar
```bash
https://github.com/P3TERX/GeoLite.mmdb
```

2. Clone este repositório e posicione o arquivo `GeoLite2-Country.mmdb` na raiz do projeto.

3. Rode o servidor:

```bash
go run server.go
```

4. Teste usando curl com simulação de IP e User-Agent:

```bash
curl -A "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/120.0.0.0" -H "X-Forwarded-For: 45.225.60.1" http://localhost:8080
```

---

## Exemplo de Campanhas

| País | Navegador | Preço | URL da Campanha                                  |
| ---- | --------- | ----- | ------------------------------------------------ |
| BR   | Chrome    | 1.1   | [https://uol.com.br](https://uol.com.br)         |
| BR   | Firefox   | 1.2   | [https://globo.com](https://globo.com)           |
| AO   | Chrome    | 0.9   | [https://zap.co.ao](https://zap.co.ao)           |
| DE   | Chrome    | 1.0   | [https://google.com](https://google.com)         |
| RU   | Chrome    | 1.0   | [https://yandex.ru](https://yandex.ru)           |
|      | Firefox   | 1.0   | [https://duckduckgo.com](https://duckduckgo.com) |

---

## Como funciona a filtragem

O servidor aplica filtros sequenciais:

* Primeiro, remove campanhas que não combinam com o país do usuário (exceto campanhas sem restrição de país).
* Depois, remove campanhas que não combinam com o navegador (exceto campanhas sem restrição de navegador).
* Se restar mais de uma campanha, escolhe a de menor preço.
* Se não sobrar nenhuma campanha, retorna erro.

---

## Futuras melhorias

* Suporte a múltiplas regras de targeting (ex: idade, idioma).
* Cache para GeoIP para melhorar performance.
* Interface web para gerenciamento de campanhas.
* Logs mais detalhados e métricas de uso.

---

## Contato

Desenvolvido por \[Luis de Água Rosada]
E-mail: [luisrosada@outlook.com](mailto:luisrosada@outlook.com)

---

