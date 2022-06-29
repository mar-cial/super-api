package main

import "net/http"

type App struct {
	h http.Handler
}

func (a *App) ServeMeatProducts(w http.ResponseWriter, r *http.Request) {

}

func (a *App) ServeSingleMeatProduct(w http.ResponseWriter, r *http.Request) {

}

func (a *App) ServeGreenProducts(w http.ResponseWriter, r *http.Request) {

}

func (a *App) ServeSingleGreenProduct(w http.ResponseWriter, r *http.Request) {

}

func (a *App) ServeMealProducts(w http.ResponseWriter, r *http.Request) {

}

func (a *App) ServeSingleMealProduct(w http.ResponseWriter, r *http.Request) {

}
