package sms_template

type Service interface {
	FindAll() ([]SmsTemplate, error)
	FindByID(ID int) (SmsTemplate, error)
	Create(smsTemplateRequest SmsTemplateRequest) (SmsTemplate, error)
	Update(ID int, smsTemplateRequest SmsTemplateRequest) (SmsTemplate, error)
	Delete(ID int) (SmsTemplate, error)
}
