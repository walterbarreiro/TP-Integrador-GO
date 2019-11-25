package main

import "fmt"

type Pirata struct {
	items [3]string
	nivelDeEbriedad int
	monedas int
}

type Leyenda struct {
	itemObligatorio string
}

type Saqueo struct {
	monedasObligatorias int
}

func (persona *Pirata) BusquedaDelTesoro() bool {
	tieneElItem := false
	for i := 0; i < len(persona.items); i++ {
		if persona.items[i] == "Brujula" || persona.items[i] == "Mapa" || persona.items[i] == "BotellaDeGrogXD" {
			tieneElItem = true 
		}
	}
	
	if tieneElItem && persona.monedas <= 5 {	
		return true
	} else {
		return false
	}	
}

func(persona *Pirata) ConvertirseEnLeyenda() bool {
	var leyenda Leyenda
	leyenda.itemObligatorio = "Botella"
	
	tieneElItem := false
	for i := 0; i < len(persona.items); i++ {
		if persona.items[i] == leyenda.itemObligatorio {
			tieneElItem = true 
		}
	}
	
	if tieneElItem && persona.monedas >= 10 {	
		return true
	} else {
		return false
	}
}

func (persona *Pirata) Saquear() bool {
	var saqueo Saqueo
	saqueo.monedasObligatorias = 5
	
	if persona.monedas >= saqueo.monedasObligatorias {
		return true
	} else {
		return false
	}
}

func main() {
	var pirata Pirata
	
	pirata.items[0] = "Manzana"
	pirata.items[1] = "Botella"
	pirata.items[2] = "Mapa"
	pirata.nivelDeEbriedad = 5
	pirata.monedas = 3
	
	if pirata.BusquedaDelTesoro() {
		fmt.Println("El pirata es util para la busqueda del tesoro")
	} else {
		fmt.Println("El pirata no es util para la busqueda del tesoro")
	}
	
	if pirata.ConvertirseEnLeyenda() {
		fmt.Println("El pirata es util para convertirse en leyenda")
	} else {
		fmt.Println("El pirata no es util para convertirse en leyenda")
	}
	
	if pirata.Saquear() {	
		fmt.Println("El pirata es util para saquear")
	} else {
		fmt.Println("El pirata no es util para saquear")
	}
}
