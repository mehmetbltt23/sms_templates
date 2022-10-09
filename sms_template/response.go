package sms_template

type SmsTemplateResponse struct {
	Sms_template_id          int    `json:"sms_template_id" binding:"required"`
	Company_id               int    `json:"company_id" binding:"required"`
	Branch_id                int    `json:"branch_id" binding:"required"`
	Name                     string `json:"sms_template_name" binding:"required"`
	Subject                  string `json:"sms_template_subject" binding:"required"`
	Content                  string `json:"sms_template_content" binding:"required"`
	Subscription_type_id     int    `json:"subscription_type_id" binding:"required"`
	Sms_template_category_id int    `json:"sms_template_category_id" binding:"required"`
	Created_by               int    `json:"sms_template_created_by" binding:"required"`
	Sms_template_type_id     int    `json:"sms_template_type_id" binding:"required"`
	Activity_id              int    `json:"sms_template_activity_id" binding:"required"`
}
