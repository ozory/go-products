package controllers

import (
	"html/template"
	"log"
	"main/produtos"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := produtos.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println(err.Error())
		}

		quantidadeConv, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println(err.Error())
		}

		produtos.CadastrarProduto(nome, descricao, precoConv, quantidadeConv)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	idProdutoConv, err := strconv.Atoi(idProduto)

	if err != nil {
		log.Println(err.Error())
	}

	produtos.ExcluirProduto(idProdutoConv)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := produtos.BuscaProdutPorId(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		intConv, err := strconv.Atoi(id)

		if err != nil {
			log.Println(err.Error())
		}

		precoConv, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println(err.Error())
		}

		quantidadeConv, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println(err.Error())
		}

		produtos.AtualizaProduto(intConv, nome, descricao, precoConv, quantidadeConv)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
