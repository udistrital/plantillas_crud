package services

import (
	m "github.com/udistrital/plantillas_crud/models"
	imagenRepo "github.com/udistrital/plantillas_crud/repositories/estiloFuente.repository"
)

func CreateEstiloFuente(estiloFuente m.EstiloFuente) error {

	err := imagenRepo.Create(estiloFuente)

	if err != nil {
		return err
	}
	return nil
}

func ReadEstiloFuente() (m.EstilosFuente, error) {

	estilosFuente, err := imagenRepo.Read()

	if err != nil {
		return nil, err
	}

	return estilosFuente, nil
}

func UpdateEstiloFuente(estiloFuente m.EstiloFuente, estiloFuenteId string) error {

	err := imagenRepo.Update(estiloFuente, estiloFuenteId)

	if err != nil {
		return err
	}

	return nil
}

func DeleteEstiloFuente(estiloFuenteId string) error {

	err := imagenRepo.Delete(estiloFuenteId)

	if err != nil {
		return err
	}

	return nil
}
