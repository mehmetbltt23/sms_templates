package sms_template

import "time"

type SmsTemplate struct {
	Sms_template_id          int
	Company_id               int
	Branch_id                int
	Name                     string
	Subject                  string
	Content                  string
	Subscription_type_id     int
	Sms_template_category_id int
	Created_by               int
	Created_at               time.Time
	Updated_at               time.Time
	Deleted_at               time.Time
	Sms_template_type_id     int
	Activity_id              int
	Is_edited                bool
}
