package service

import (
	"errors"
)

var (
	ErrInvalidRequestFormatOrParameters = errors.New("неверный формат запроса или параметры")
	ErrUserNotExistOrIncorrect          = errors.New("пользователь не существует или некорректен")
	ErrServerNotReadyToProcessRequests  = errors.New("сервер не готов обрабатывать запросы")
	ErrNotEnoughRights                  = errors.New("недостаточно прав для выполнения действия")
	ErrTenderNotFind                    = errors.New("тендер не найдено")
	ErrBidNotFind                       = errors.New("предложение не найдено")
	ErrVersionNotFind                   = errors.New("версия не найдены")
	ErrEmployeeNotFind                  = errors.New("сотрудник не найден")
	ErrOrganizationNotFind              = errors.New("компания не найден")
)
