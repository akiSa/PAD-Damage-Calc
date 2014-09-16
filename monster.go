package main

//Elements: fire water wood dark light?
//Stats: HP Attack RCV
type Monster struct {
	Element int `json:"element"`
	Element2 int `json:"element2"`
	Image60Size int `json:"image60_size"`
	Name string `json:"name"`
	MaxLevel int `json:"max_level"`
	Awakenings []int `json:"awoken_skills"`
	Image60Href string `json:"image60_href"`
	Rarity int `json:"rarity"`
	HPMax int `json:"hp_max"`
	RCVMin int `json:"rcv_min"`
	RCVMax int `json:"rcv_max"`
	RCVScale float64 `json:"rcv_scale"`
	ATKScale float64 `json:"atk_scale"`
	ID int `json:"id"`
	Type2 int `json:"type2"`
	Image40Href string `json:"image40_href"`
	HPScale float64 `json:"hp_scale"`
	XPCurve int `json:"xp_curve"`
	LeaderSkill string `json:"leader_skill"`
	TeamCost int `json:"team_cost"`
	Type int `json:"type"`
	HPMin int `json:"hp_min"`
	Image40Size int `json:"image40_size"`
	ActiveSkill string `json:"active_skill"`
	ATKMin int `json:"atk_min"`
	FeedXP float64 `json:"feed_xp"`
	ATKMax int `json:"atk_max"`
	JPOnly bool `json:"jp_only"`
}

// func Lookup (ID int) Monster {
	
// }
