package plantillaservice_test

import (
	"fmt"
	"testing"

	m "github.com/udistrital/plantillas_crud/models"
	plantillaService "github.com/udistrital/plantillas_crud/services/plantilla.service"
)

func TestCreate(t *testing.T) {

	estiloFuente := m.EstiloFuente{
		Tamaño: 2,
		Estilo: "Times New Roman",
		Grosor: "",
		Altura: 2,
	}

	secciones := m.Seccion{
		Nombre:            "Seccion de prueba",
		Descripcion:       "Desc",
		Valor:             "Este es el texto 1",
		EstiloFuente:      estiloFuente,
		FechaCreacion:     "",
		FechaModificacion: "",
		Activo:            true,
	}

	minutas := m.Minuta{
		Nombre:            "Seccion de prueba",
		Descripcion:       "Desc",
		Valor:             "Este es el texto 2",
		EstiloFuente:      estiloFuente,
		FechaCreacion:     "",
		FechaModificacion: "",
		Activo:            true,
	}
	titulos := m.Titulo{
		Nombre:            "Titulo de prueba",
		Descripcion:       "Desc",
		Valor:             "Este es el texto 3",
		EstiloFuente:      estiloFuente,
		FechaCreacion:     "",
		FechaModificacion: "",
		Activo:            true,
	}
	imagenes := m.Imagen{
		Nombre:            "Titulo de prueba",
		Data:              "",
		FechaCreacion:     "",
		FechaModificacion: "",
		Activo:            true,
	}

	plantilla := m.Plantilla{
		Nombre:            "Plantilla_documento_1",
		Descripcion:       "Descripción 1",
		Secciones:         secciones,
		Minutas:           minutas,
		Titulos:           titulos,
		Imagenes:          imagenes,
		EnlaceDoc:         "",
		Version:           1.0,
		FechaCreacion:     "",
		FechaModificacion: "",
		Activo:            true,
	}

	fmt.Println("Plantilla:", plantilla)

	err := plantillaService.Create(plantilla)

	if err != nil {
		t.Error("La prueba de persistencia de datos de la plantilla falló")
		t.Fail()
		fmt.Println("ERR:", err)
	} else {
		t.Log("La prueba terminó con exito")
	}

}

func TestRead(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
