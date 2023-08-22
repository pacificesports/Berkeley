package service

import (
	"berkeley/config"
	"berkeley/model"
	"berkeley/utils"
)

func GetAllSchools() []model.School {
	var schools []model.School
	result := DB.Find(&schools)
	if result.Error != nil {
	}
	for i := range schools {
		schools[i].Tags = GetTagsForSchool(schools[i].ID)
		schools[i].Emails = GetEmailsForSchool(schools[i].ID)
	}
	return schools
}

func GetSchoolByID(id string) model.School {
	var school model.School
	result := DB.Where("id = ?", id).First(&school)
	if result.Error != nil {
	}
	school.Tags = GetTagsForSchool(school.ID)
	school.Emails = GetEmailsForSchool(school.ID)
	return school
}

func CreateSchool(school model.School) error {
	if DB.Where("id = ?", school.ID).Select("*").Updates(&school).RowsAffected == 0 {
		utils.SugarLogger.Infoln("New school created with id: " + school.ID)
		if result := DB.Create(&school); result.Error != nil {
			return result.Error
		}
		go DiscordLogNewSchool(school)
	} else {
		utils.SugarLogger.Infoln("School with id: " + school.ID + " has been updated!")
		go DiscordLogUpdatedSchool(school)
	}
	return nil
}

func VerifySchool(id string) error {
	school := GetSchoolByID(id)
	if school.ID != "" {
		school.Verified = true
		if err := CreateSchool(school); err != nil {
			return err
		}
		go Discord.ChannelMessageSend(config.DiscordChannel, ":white_check_mark: "+school.Name+" has been marked as verified!")
	}
	return nil
}

func UnverifySchool(id string) error {
	school := GetSchoolByID(id)
	if school.ID != "" {
		school.Verified = false
		if err := CreateSchool(school); err != nil {
			return err
		}
		go Discord.ChannelMessageSend(config.DiscordChannel, ":x:  "+school.Name+" has been marked as unverified!")
	}
	return nil
}
