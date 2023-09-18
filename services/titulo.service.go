package services

import (
	m "github.com/udistrital/plantillas_crud/models"
	SeccionRepo "github.com/udistrital/plantillas_crud/repositories/titulo.repository"
)

func CreateTitulo(titulo m.Titulo) error {

	err := SeccionRepo.Create(titulo)

	if err != nil {
		return err
	}
	return nil
}

func ReadTitulo() (m.Titulos, error) {

	titulos, err := SeccionRepo.Read()

	if err != nil {
		return nil, err
	}

	return titulos, nil
}

func ReadOneTitulo(id string) (m.Titulo, error) {

	titulo, err := SeccionRepo.ReadOne(id)

	if err != nil {
		return titulo, err
	}

	return titulo, nil
}

func UpdateTitulo(titulo m.Titulo, tituloId string) error {

	err := SeccionRepo.Update(titulo, tituloId)

	if err != nil {
		return err
	}

	return nil
}

func DeleteTitulo(tituloId string) error {

	err := SeccionRepo.Delete(tituloId)

	if err != nil {
		return err
	}

	return nil
}
