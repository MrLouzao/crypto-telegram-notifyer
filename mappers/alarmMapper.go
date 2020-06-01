package mappers

import (
	"crypto-telegram-notifyer/models"
	"crypto-telegram-notifyer/repositories"
)

// Cast from entities to dto
func CastToAlarmDtoList(alarmEntities *repositories.AlarmResults) *[]models.AlarmDto {
	entities := alarmEntities.Items
	dtos := make([]models.AlarmDto, len(entities))

	for i := range entities {
		dtos[i] = models.AlarmDto(entities[i])
	}

	return &dtos
}
