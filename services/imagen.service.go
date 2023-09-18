package services

import (
	m "github.com/udistrital/plantillas_crud/models"
	imagenRepo "github.com/udistrital/plantillas_crud/repositories/imagen.repository"
)

func CreateImagen(imagen m.Imagen) error {

	err := imagenRepo.Create(imagen)

	if err != nil {
		return err
	}
	return nil
}

func ReadImagen() (m.Imagenes, error) {

	imagenes, err := imagenRepo.Read()

	if err != nil {
		return nil, err
	}

	return imagenes, nil
}

func UpdateImagen(imagen m.Imagen, imagenId string) error {

	err := imagenRepo.Update(imagen, imagenId)

	if err != nil {
		return err
	}

	return nil
}

func DeleteImagen(imagenId string) error {

	err := imagenRepo.Delete(imagenId)

	if err != nil {
		return err
	}

	return nil
}
