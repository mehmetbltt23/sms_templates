package sms_template_categories

import "time"

type SmsTemplate struct {
	Sms_template_category_id int
	Company_id               int
	Branch_id                int
	Name                     string
	Description              string
	Created_by               int
	Created_at               time.Time
	Updated_at               time.Time
	Deleted_at               time.Time
}
