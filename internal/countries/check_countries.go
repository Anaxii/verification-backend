package countries

import (
	log "github.com/sirupsen/logrus"
	"puffinverificationbackend/internal/externaldatabase"
	"strings"
)

func CheckCountries() {
	countries := externaldatabase.GetCountries()
	users, err := externaldatabase.GetAllUsers()
	if err != nil {
		log.Println(err)
	}
	for _, country := range countries.Countries {
		for _, v := range users {
			if strings.ToLower(v.Country) == country {
				// check if block if not block
			}
		}
	}

}
