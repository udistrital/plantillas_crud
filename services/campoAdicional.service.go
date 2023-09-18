package services

import (
	m "github.com/udistrital/plantillas_crud/models"
	imagenRepo "github.com/udistrital/plantillas_crud/repositories/campoAdicional.repository"
)

func CreateCampoAdicional(campoAdicional m.CampoAdicional) error {

	err := imagenRepo.Create(campoAdicional)

	if err != nil {
		return err
	}
	return nil
}

func ReadCampoAdicional() (m.CamposAdicionales, error) {

	camposAdicionales, err := imagenRepo.Read()

	if err != nil {
		return nil, err
	}

	return camposAdicionales, nil
}

func UpdateCampoAdicional(campoAdicional m.CampoAdicional, campoAdicionalId string) error {

	err := imagenRepo.Update(campoAdicional, campoAdicionalId)

	if err != nil {
		return err
	}

	return nil
}

func DeleteCampoAdicional(campoAdicionalId string) error {

	err := imagenRepo.Delete(campoAdicionalId)

	if err != nil {
		return err
	}

	return nil
}
