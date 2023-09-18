package services

import (
	m "github.com/udistrital/plantillas_crud/models"
	minutaRepo "github.com/udistrital/plantillas_crud/repositories/minuta.repository"
)

func CreateMinuta(minuta m.Minuta) error {

	err := minutaRepo.Create(minuta)

	if err != nil {
		return err
	}
	return nil
}

func ReadMinuta() (m.Minutas, error) {

	minutas, err := minutaRepo.Read()

	if err != nil {
		return nil, err
	}

	return minutas, nil
}

func UpdateMinuta(minuta m.Minuta, minutaId string) error {

	err := minutaRepo.Update(minuta, minutaId)

	if err != nil {
		return err
	}

	return nil
}

func DeleteMinuta(minutaId string) error {

	err := minutaRepo.Delete(minutaId)

	if err != nil {
		return err
	}

	return nil
}
