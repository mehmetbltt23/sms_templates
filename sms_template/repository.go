package sms_template

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]SmsTemplate, error)
	FindByID(ID int) (SmsTemplate, error)
	Create(sms_template SmsTemplate) (SmsTemplate, error)
	Update(sms_template SmsTemplate) (SmsTemplate, error)
	Delete(sms_template SmsTemplate) (SmsTemplate, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]SmsTemplate, error) {
	var sms_templates []SmsTemplate
	err := r.db.Find(&sms_templates).Error

	return sms_templates, err
}

func (r *repository) FindByID(ID int) (SmsTemplate, error) {
	var sms_templates SmsTemplate
	err := r.db.Find(&sms_templates, ID).Error

	return sms_templates, err
}

func (r *repository) Create(sms_template SmsTemplate) (SmsTemplate, error) {
	err := r.db.Create(&sms_template).Error

	return sms_template, err
}

func (r *repository) Update(sms_template SmsTemplate) (SmsTemplate, error) {
	err := r.db.Save(&sms_template).Error

	return sms_template, err
}

func (r *repository) Delete(sms_template SmsTemplate) (SmsTemplate, error) {
	err := r.db.Delete(&sms_template).Error

	return sms_template, err
}
