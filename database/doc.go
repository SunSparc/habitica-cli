package database

import "time"

type HabiticaUser struct {
	AppVersion string   `json:"appVersion"` // "4.100.0"
	Data       UserData `json:"data"`       // {}
	//Notifications []UserNotification `json:"notifications"` // []
	Success bool `json:"success"` // true
	UserV   int  `json:"userV"`   // 8723
}

type UserData struct {
	//_ABTests ABTests `json:"_ABTests"` // {"onboardingPushNotification": "Onboarding-Step7-Phasea-VersionC"},
	//_ID string `json:"_id"` // "639ec312-3ddb-4c11-a903-4914b72efeb9"
	//_V int `json:"_v"` // 8723
	//Achievements Achievements `json:"achievements"` // {}
	//Auth Auth `json:"auth"` // {}
	//Backer Backer `json:"backer"` // {}
	Balance float32 `json:"balance"` // 1.25
	//Challenges []Challenge "challenges" // []
	//Contributor Contributor "contributor" // {}
	//Extra Extra `json:"extra"` // {}
	//Flags Flags `json:"flags"` // {} - eg "cronCount": 176 // Total Check Ins
	Guilds []string `json:"guilds"` // ["9ddbb860-67ff-4e7b-81aa-23f81de910b3"]
	//History History `json:"history"` // {}
	ID string `json:"id"` // "639ec312-3ddb-4c11-a903-4914b72efeb9",
	//Inbox Inbox `json:"inbox"` // {}
	//Invitations Invitations `json:"invitations"` // {}
	InvitesSent     int       `json:"invitesSent"`     // 0
	Items           Items     `json:"items"`           // {}
	LastCron        time.Time `json:"lastCron"`        // "2019-06-12T12:17:32.713Z",
	LoginIncentives int       `json:"loginIncentives"` // 176
	Migration       string    `json:"migration"`       // "20190530_halfmoon_glasses",
	NeedsCron       bool      `json:"needsCron"`       // false
	//NewMessages string `json:"newMessages"` // {},
	//Notifications []UserDataNotification `json:"notifications"` // []
	//Party UserDataParty `json:"party"` // {}
	//PinnedItems []PinnedItems `json:"pinnedItems"` // []
	//PinnedItemsOrder []pinnedItemsOrder `json:"pinnedItemsOrder"` // []
	//Preferences Preferences `json:"preferences"` // {}
	//Profile Profile `json:"profile"` // map[string]string or {"blurb": "string", "imageUrl": "string", "name": "string"}
	//Purchased Purchased `json:"purchased"` // {}
	//PushDevices string `json:"pushDevices"` // []map[string]string or BigOldStruct
	//Stats Stats `json:"stats"` // BigOldStruct
	//Tags string `json:"tags"` // []map[string]string
	//TasksOrder string `json:"tasksOrder"` // map[string][]string
	//UnpinnedItems string `json:"unpinnedItems"` // []
	//Webhooks string `json:"webhooks"` // []
}

type UserNotification struct {
	Data map[string]string `json:"data"` // {"bodyText": "someuser accepted your invitation to guild!","headerText": "Your Invitation has been Accepted"}
	ID   string            `json:"id"`   // "5d3d8618-25bc-450a-9bf7-171bc1cce3fa",
	Seen bool              `json:"seen"` // false
	Type string            `json:"type"` // "GROUP_INVITE_ACCEPTED"
}

type Items struct {
	CurrentMount    string         `json:"currentMount"` // "TigerCub-Floral",
	CurrentPet      string         `json:"currentPet"`   // "Egg-Red",
	Eggs            map[string]int `json:"eggs"`         // {"BearCub": 10}
	Food            map[string]int `json:"food"`         // {"Cake_Base": 0}
	Gear            Gear           `json:"gear"`
	HatchingPotions map[string]int `json:"hatchingPotions"` // {"Base": 20}
	//LastDrop LastDrop `json:"lastDrop"` // {}
	Mounts map[string]bool `json:"mounts"` // {"BearCub-Shade": true}
	Pets   map[string]int  `json:"pets"`   // {"Dragon-Celestial": -1}
	Quests map[string]int  `json:"quests"` // {"atom1": 0}
	//Special Special `json:"special"` // {}
}

type Gear struct {
	Costume  map[string]string `json:"costume"`  // {"armor": "armor_mystery_201512"}
	Equipped map[string]string `json:"equipped"` // {"armor": "armor_wizard_5"}
	Owned    map[string]bool   `json:"owned"`    // {"armor_armoire_candlestickMakerOutfit": true}
}

///////////////////////////////////////////

type Mount struct {
	Name    string
	Mounted time.Time
}
type Pet struct{}

type MountManager struct {
}

// every 24 hours, switch currentmount to a different mount
// if there are new mounts in the list of mounts, use one of them
// otherwise, use a mount that has not been used
// otherwise, use the mount that was used the longest time ago
// Special cases: there is no currentmount, there are no mounts, there is only one mount available
