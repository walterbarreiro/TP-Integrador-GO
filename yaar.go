package main

import (
	"fmt"
)

type Pirata struct {
	nombre string
	items [] string 
	nivelDeEbriedad int
	monedas int
}

type Barco struct{
	tripulantes [] Pirata 
	 //mision 
	capacidad int 
}

type CiudadCostera struct {
	habitantes int
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

func (barco *Barco) IncorporarALaTripulacion(persona Pirata) {
	if barco.PuedeFormarParteDeLaTripulacion(persona) {
		barco.tripulantes = append(barco.tripulantes, persona)
	}
}

func (barco *Barco) PuedeFormarParteDeLaTripulacion (persona Pirata) bool {
	if barco.capacidad > len(barco.tripulantes) {
		return true
	} else {
		return false
	}
}

func (barco *Barco) PirataMasEbrio() string {
	var pirataMasEbrio Pirata
	
	for i:=0; i<len(barco.tripulantes); i++ {
		if barco.tripulantes[i].nivelDeEbriedad > pirataMasEbrio.nivelDeEbriedad {
			pirataMasEbrio = barco.tripulantes[i]
		}
	}
	
	return pirataMasEbrio.nombre
}

func (barco *Barco) Anclar(ciudad CiudadCostera) {
	barco.TomarTragoDeGrogXD()
	ciudad.habitantes++
}

func (barco *Barco) TomarTragoDeGrogXD() {
	
	for i:=0; i<len(barco.tripulantes); i++ {
		barco.tripulantes[i].nivelDeEbriedad += 5
		
		if barco.tripulantes[i].monedas > 0 {
		barco.tripulantes[i].monedas -= 1 
		} else {
		 // lanzar un error
		}
	}
}

func main() {
	
	
	jorge := Pirata{"Jorge",[]string{"Manzana", "Botella", "Mapa"},5,3}
	pedro := Pirata{"Pedro",[]string{"Mapa", "Carne", "Brujula", "Cartas"},5,10}
	carlos := Pirata{"Carlos",[]string{"Parche","Mapa","Whisky"},13,1}
		
	var perlaNegra Barco
	perlaNegra.tripulantes = append(perlaNegra.tripulantes, jorge)
	perlaNegra.tripulantes = append(perlaNegra.tripulantes, carlos)
	perlaNegra.capacidad = 50

	var trianguloDeLasBermudas CiudadCostera
	trianguloDeLasBermudas.habitantes = 10
	
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
	
	//slicePrueba := []string{"hola","chau","cusbai"}
	//fmt.Println([hola chau cusbai])

        fmt.Println(perlaNegra)
	perlaNegra.IncorporarALaTripulacion(pedro)
	fmt.Println(perlaNegra)
	
	fmt.Println("Capacidad del Perla Negra:", len(perlaNegra.tripulantes))
	
	fmt.Println("Pirata mas ebrio del Perla Negra:" , perlaNegra.PirataMasEbrio())
	
	perlaNegra.Anclar(trianguloDeLasBermudas)
}