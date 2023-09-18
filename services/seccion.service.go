package services

import (
	m "github.com/udistrital/plantillas_crud/models"
	SeccionRepo "github.com/udistrital/plantillas_crud/repositories/seccion.repository"
)

func CreateSeccion(seccion m.Seccion) error {

	err := SeccionRepo.Create(seccion)

	if err != nil {
		return err
	}
	return nil
}

func ReadSeccion() (m.Secciones, error) {

	secciones, err := SeccionRepo.Read()

	if err != nil {
		return nil, err
	}

	return secciones, nil
}

func ReadOneSeccion(id string) (m.Seccion, error) {

	seccion, err := SeccionRepo.ReadOne(id)

	if err != nil {
		return seccion, err
	}

	return seccion, nil
}

func UpdateSeccion(seccion m.Seccion, seccionId string) error {

	err := SeccionRepo.Update(seccion, seccionId)

	if err != nil {
		return err
	}

	return nil
}

func DeleteSeccion(seccionId string) error {

	err := SeccionRepo.Delete(seccionId)

	if err != nil {
		return err
	}

	return nil
}
