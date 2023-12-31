package service

import (
	"berkeley/model"
)

func GetEmailsForSchool(schoolId string) []model.Email {
	var emails []model.Email
	result := DB.Where("school_id = ?", schoolId).Find(&emails)
	if result.Error != nil {
	}
	return emails
}

func AddEmailForSchool(email model.Email) error {
	result := DB.Create(&email)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteEmailForSchool(email model.Email) error {
	result := DB.Where("school_id = ? AND email = ?", email.SchoolID, email.Email).Delete(&model.Email{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
