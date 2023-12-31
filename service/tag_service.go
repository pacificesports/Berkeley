package service

import (
	"berkeley/model"
)

func GetTagsForSchool(schoolId string) []model.Tag {
	var tags []model.Tag
	result := DB.Where("school_id = ?", schoolId).Find(&tags)
	if result.Error != nil {
	}
	return tags
}

func AddTagForSchool(tag model.Tag) error {
	result := DB.Create(&tag)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteTagForSchool(tag model.Tag) error {
	result := DB.Where("school_id = ? AND tag = ?", tag.SchoolID, tag.Tag).Delete(&model.Tag{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
