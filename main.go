package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	fmt.Println("connecting")
	// these details match the docker-compose.yml file.
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres", 5432, "user", "mypassword", "testdb")
	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	start := time.Now()
	for db.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("failed to connect after 10 secs.")
			break
		}
	}
	fmt.Println("connected:", db.Ping() == nil)
	_, err = db.Exec(`DROP TABLE IF EXISTS SMS_TEMPLATES;`)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`CREATE TABLE SMS_TEMPLATES (SMS_TEMPLATE_ID INT PRIMARY KEY NOT NULL, COMPANY_ID INT, BRANCH_ID INT, NAME text, SUBJECT text, CONTENT text, SUBSCRIPTION_TYPE_ID INT, SMS_TEMPLATE_CATEGORY_ID INT, CREATED_BY INT, CREATED_AT DATE, UPDATED_AT DATE, DELETED_AT DATE, SMS_TEMPLATE_TYPE_ID INT, ACTIVITY_ID INT, IS_EDITED BOOL);`)
	if err != nil {
		panic(err)
	}
	fmt.Println("table sms_templates is created")

	_, err = db.Exec(`CREATE TABLE SMS_TEMPLATE_TYPES (SMS_TEMPLATE_TYPE_ID INT PRIMARY KEY NOT NULL, NAME text, DESCRIPTION text, KEY text,  CREATED_AT DATE, UPDATED_AT DATE, DELETED_AT DATE);`)
	if err != nil {
		panic(err)
	}

	fmt.Println("table sms_template_types is created")

	_, err = db.Exec(`CREATE TABLE SMS_TEMPLATE_CATEGORIES (SMS_TEMPLATE_CATEGORY_ID INT PRIMARY KEY NOT NULL, NAME text, DESCRIPTION text, COMPANY_ID INT, BRANCH_ID INT,  CREATED_AT DATE, UPDATED_AT DATE, DELETED_AT DATE);`)
	if err != nil {
		panic(err)
	}

	fmt.Println("table sms_template_categories is created")

}
