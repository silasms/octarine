package repository

import "github.com/silasms/octarine/domain"

type ImageRepository {
	GetAll() ([]domain.Image, error)
}