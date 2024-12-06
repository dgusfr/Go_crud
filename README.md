# Gerenciador de Produtos em Go

Um projeto CRUD simples desenvolvido em Go, utilizando a arquitetura MVC e PostgreSQL para gerenciamento de produtos.

## Interface

<div align="center">
  <img src="img/logo.png" alt="Imagem do Projeto" width="900">
</div>

## Sumário

- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Status](#status)
- [Descrição](#descrição)
- [Funcionalidades](#funcionalidades)
- [Como Usar](#como-usar)
- [Autor](#autor)

## Tecnologias Utilizadas

<div style="display: flex; flex-direction: row;">
  <div style="margin-right: 20px; display: flex; justify-content: flex-start;">
    <img src="img/go.png" alt="Logo Go" width="100"/>
  </div>
  <div style="margin-right: 20px; display: flex; justify-content: flex-start;">
    <img src="img/postgres.png" alt="Logo PostgreSQL" width="100"/>
  </div>
</div>

## Status

![Em Desenvolvimento](http://img.shields.io/static/v1?label=STATUS&message=EM%20DESENVOLVIMENTO&color=RED&style=for-the-badge)

## Descrição

Este projeto implementa operações básicas de CRUD (Create, Read, Update, Delete) para gerenciar produtos, incluindo funcionalidades para criar, listar, editar e excluir produtos armazenados em um banco de dados PostgreSQL. O projeto utiliza templates HTML para exibição e manipulação de dados.

## Funcionalidades

- Criar novos produtos.
- Listar todos os produtos cadastrados.
- Editar informações de produtos existentes.
- Excluir produtos do sistema.
- Interface web simples e funcional.

## Como Usar

1. **Pré-requisitos**:

   - Go instalado em sua máquina.
   - Banco de dados PostgreSQL configurado e rodando.

2. **Configurar o banco de dados**:

   - Crie um banco de dados chamado `go_crud` e configure a tabela `produtos`:
     ```sql
     CREATE TABLE produtos (
         id SERIAL PRIMARY KEY,
         nome VARCHAR(100),
         descricao TEXT,
         preco NUMERIC(10, 2),
         quantidade INT
     );
     ```

3. **Configurar o projeto**:

   - No arquivo `db/db.go`, atualize a string de conexão com as credenciais do seu banco de dados:
     ```go
     conexao := "user=SEU_USUARIO dbname=go_crud password=SUA_SENHA host=localhost sslmode=disable"
     ```

4. **Rodar a aplicação**:

   - Navegue até o diretório do projeto e execute:
     ```bash
     go run main.go
     ```
   - Acesse a aplicação em seu navegador pelo endereço: [http://localhost:8000](http://localhost:8000).

5. **Testar as funcionalidades**:
   - Use as opções da interface para adicionar, listar, editar ou excluir produtos.

## Autor

Desenvolvido por Diego Franco
