# go_api

## 🖥️ Tecnologias Utilizadas

-  Linguagem: Go (Golang)
- Banco de Dados: PostgreSQL
-  Framework: Gin (ou padrão net/http para simplicidade)


## 📂 Documentação da API

#### Retorna todos os produtos

```http
  GET /products
```

#### Retorna um produto

```http
  GET /products/${id} -> /products/1
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `int` | **Obrigatório**. O ID do item que você quer |


#### Cadastrar/Adicionar um produto
```http
  POST /products
```

#### Atuaizar um produto

```http
  PUT /products/${id} -> /products/1
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `int` | **Obrigatório**. O ID do item que você quer |

#### Excluir um produto

```http
  DELETE /products/${id} -> /products/1
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `int` | **Obrigatório**. O ID do item que você quer |



## Autores

- [@Elias](https://www.github.com/EliasBRodrigues)

