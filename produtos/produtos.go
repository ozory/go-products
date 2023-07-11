package produtos

import "main/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	allProducts, err := db.Query("SELECT * FROM produtos order by id")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for allProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = allProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CadastrarProduto(nome, descricao string, preco float64, quantidade int) {

	db := db.ConectaComBancoDeDados()

	insertData, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {

	db := db.ConectaComBancoDeDados()

	updateData, err := db.Prepare("update produtos set nome=$2, descricao=$3, preco=$4, quantidade=$5 where id=$1")
	if err != nil {
		panic(err.Error())
	}

	updateData.Exec(id, nome, descricao, preco, quantidade)
	defer db.Close()
}

func ExcluirProduto(id int) {

	db := db.ConectaComBancoDeDados()

	insertData, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(id)
	defer db.Close()
}

func BuscaProdutPorId(id string) Produto {
	db := db.ConectaComBancoDeDados()

	getProduto, err := db.Query("SELECT * FROM produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	for getProduto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = getProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
	}

	defer db.Close()
	return p
}
