package sms_template_types

import "time"

type SmsTemplate struct {
	Sms_template_type_id int
	Name                 string
	Description          string
	Key                  string
	Created_at           time.Time
	Updated_at           time.Time
	Deleted_at           time.Time
}
