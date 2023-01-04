package countries

import (
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffinverificationbackend/internal/blockchain"
	"puffinverificationbackend/internal/externaldatabase"
	"strings"
)

func CheckCountries() {
	countries := externaldatabase.GetCountries()
	users, err := externaldatabase.GetAllUsers()
	if err != nil {
		log.Println(err)
	}

	toSet := map[string]*big.Int{}
	for _, country := range countries.Countries {
		for _, v := range users {
			tier, isKYC := blockchain.GetTier(v.WalletAddress)
			country = strings.ToLower(country)
			v.Country = strings.ToLower(v.Country)
			if v.Country != country && tier == "0" && isKYC {
				continue
			}  else if v.Country == country && tier == "0" {
				toSet[v.WalletAddress] = big.NewInt(1)
				continue
			} else if v.Country != country && !isKYC {
				if toSet[v.WalletAddress] != nil {
					if toSet[v.WalletAddress].Cmp(big.NewInt(1)) == 0 {
						continue
					}
				}
				toSet[v.WalletAddress] = big.NewInt(0)
				continue
			}
		}
	}

	for walletAddress, tier := range toSet {
		err = blockchain.SetTier(walletAddress, tier)
		if err != nil {
			log.Error("error settings tier for user")
		}
	}

}
