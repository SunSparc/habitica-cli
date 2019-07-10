package client

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"go-client/database"
)

// do live stuff here, try not to break the live account, there is no undo

type LiveMode struct {
	HabiticaUser database.HabiticaUser
}

func NewLiveMode(user database.HabiticaUser) *LiveMode {
	return &LiveMode{
		HabiticaUser: user,
	}
}

// TODO: get current list of mounts
//       compare current and local lists of mounts, update local list of mounts
//       pick a different mount to be current, one that has not been used before or in a long time
//   Initially, we can just randomly pick a mount that is not the current one
func (this *LiveMode) ManageMounts() {
	if this.HabiticaUser.Data.Items.CurrentMount == "" {
		log.Println("currentmount: none")
	} else {
		log.Println("currentmount:", this.HabiticaUser.Data.Items.CurrentMount)
	}
	mount, err := getRandomItemFromMapOfStringToBool(this.HabiticaUser.Data.Items.Mounts)
	if err != nil {
		log.Println(err)
	}
	log.Println("Your random mount is:", mount)
}

func getRandomItemFromMapOfStringToBool(mapSB map[string]bool) (string, error) {
	number := generateRandomNumber(len(mapSB))
	index := 0
	for key, value := range mapSB {
		index = index + 1
		if index == number && value {
			return key, nil
		}
	}
	return "", errors.New("[ERROR] map was empty")
}

func generateRandomNumber(size int) int {
	if size <= 0 {
		return 0
	}
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	return random.Intn(size)
}

// TODO: auto feed pets (in which order? preference?)
// TODO: auto hatch eggs, giving favor to quest and magic eggs
//        but only eggs that do not have a pet and/or mount yet, since you cannot hatch or feed if you have
//        both the pet and mount already

// TODO: If I have more than $x gold (1100?), do one armoire purchase:
//     - https://habitica.com/api/v3/user/buy-armoire
//      (keep track of last purchase? timestamp, what came out of the box)

// TODO: Switch my pet/mount every day
//     - random, sequential ?
// TODO: rotate background
// TODO: option to run this application as a daemon ?
