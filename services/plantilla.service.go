package services

import (
	m "github.com/udistrital/plantillas_crud/models"
	plantillaRepo "github.com/udistrital/plantillas_crud/repositories/plantilla.repository"
)

func CreatePlantilla(plantilla m.Plantilla) error {

	err := plantillaRepo.Create(plantilla)

	if err != nil {
		return err
	}
	return nil
}

func ReadPlantilla() (m.Plantillas, error) {

	plantillas, err := plantillaRepo.Read()

	if err != nil {
		return nil, err
	}

	return plantillas, nil
}

func ReadOnePlantilla(id string) (m.Plantillas, error) {

	plantillas, err := plantillaRepo.ReadOne()

	if err != nil {
		return nil, err
	}

	return plantillas, nil
}

func UpdatePlantilla(plantilla m.Plantilla, plantillaId string) error {

	err := plantillaRepo.Update(plantilla, plantillaId)

	if err != nil {
		return err
	}

	return nil
}

func DeletePlantilla(plantillaId string) error {

	err := plantillaRepo.Delete(plantillaId)

	if err != nil {
		return err
	}

	return nil
}
