package main

import "fmt"

type Pirata struct {
	items [] string 
	nivelDeEbriedad int
	monedas int
}

type Barco struct{
	tripulantes [] Pirata 
	 //mision 
	capacidad int 
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

func(persona *Pirata) ConvertirseEnLeyenda(itemObligatorio string) bool {
	var leyenda Leyenda
	leyenda.itemObligatorio = itemObligatorio
	
	tieneElItem := false
	for i := 0; i < len(persona.items); i++ {
		if persona.items[i] == leyenda.itemObligatorio {
			tieneElItem = true 
		}
	}
	
	if tieneElItem && len(persona.items) >= 10 {	
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
	var jorge Pirata
	jorge.items = append(jorge.items,"Manzana")
	jorge.items = append(jorge.items,"Botella")
	jorge.items = append(jorge.items,"Mapa")
	jorge.nivelDeEbriedad = 5
	jorge.monedas = 3


	pedro := Pirata{["Mapa", "Carne", "Brujula", "Cartas"],5,10}
	
	if jorge.BusquedaDelTesoro() {
		fmt.Println("El pirata es util para la busqueda del tesoro")
	} else {
		fmt.Println("El pirata no es util para la busqueda del tesoro")
	}
	
	if jorge.ConvertirseEnLeyenda("Botella") {
		fmt.Println("El pirata es util para convertirse en leyenda")
	} else {
		fmt.Println("El pirata no es util para convertirse en leyenda")
	}
	
	if jorge.Saquear() {	
		fmt.Println("El pirata es util para saquear")
	} else {
		fmt.Println("El pirata no es util para saquear")
	}
	slicePrueba := []string{"hola","chau","cusbai"}
	fmt.Println([hola chau cusbai])

}
