package plantillaservice

import (
	m "github.com/udistrital/plantillas_crud/models"
	plantillaRepo "github.com/udistrital/plantillas_crud/repositories/plantilla.repository"
)

func Create(plantilla m.Plantilla) error {

	err := plantillaRepo.Create(plantilla)

	if err != nil {
		return err
	}
	return nil
}

func Read() (m.Plantillas, error) {

	plantillas, err := plantillaRepo.Read()

	if err != nil {
		return nil, err
	}

	return plantillas, nil
}

func ReadOne(id string) (m.Plantillas, error) {

	plantillas, err := plantillaRepo.ReadOne()

	if err != nil {
		return nil, err
	}

	return plantillas, nil
}

func Update(plantilla m.Plantilla, plantillaId string) error {

	err := plantillaRepo.Update(plantilla, plantillaId)

	if err != nil {
		return err
	}

	return nil
}

func Delete(plantillaId string) error {

	err := plantillaRepo.Delete(plantillaId)

	if err != nil {
		return err
	}

	return nil
}
