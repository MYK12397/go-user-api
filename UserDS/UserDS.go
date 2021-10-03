package UserDS

import (
	"net/http"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty" `
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty" ` //validate: "required,fl" `
	CreatedOn time.Time          `json:"createdon" bson:"createdon"`
	UpdateOn  time.Time          `json:"updateon" bson:"updateon"`
	Mobile    string             `json:"mobile,omitempty" bson:"mobile,omitempty" ` //validate: "required,mob"`
	Active    bool               `json:"active,omitempty" bson:"active,omitempty"`
	Age       AgeDS              `json:"age,omitempty"  bson:"age,omitempty"`
}

type AgeDS struct {
	Value    int    `json:"age,omitempty" bson:"age,omitempty"`
	Interval string `json:"interval,omitempty" bson:"interval,omitempty"`
}

func (u *User) Validate() error {
	validate := validator.New()

	return validate.Struct(u)
}

func (u *User) ValidateInput(rw http.ResponseWriter) bool {
	validate := validator.New()

	validate.RegisterValidation("fl", validateName)
	validate.RegisterValidation("mob", validatMobile)

	err := validate.Struct(u)

	if err != nil {
		http.Error(rw, "Invalid user data", http.StatusInternalServerError)
		return false
	}
	return true
}

func validateName(fl validator.FieldLevel) bool {

	isAlpha := regexp.MustCompile(`[a-zA-Z]`).MatchString
	name := fl.Field().String()

	return isAlpha(name)
}

func validatMobile(fl validator.FieldLevel) bool {
	isNumber := regexp.MustCompile(`[0-9]+`).MatchString

	number := fl.Field().String()

	return len(number) == 10 || isNumber(number)

}
