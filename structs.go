package main

type orbs struct {
	Fire []float64 `json:"fire"`
	Water []float64 `json:"water"`
	Wood []float64 `json:"wood"`
	Dark []float64 `json:"dark"`
	Light []float64 `json:"light"`
	Heart []float64 `json:"heart"`
	Rows []float64 `json:"rows"` //# of rows per colour [ x x x x x x ]
	Enhance []float64 `json:"enhanced_orbs"` //# of enhanced orbs per colour [ x x x x x x ]
	Active []interface{} `json:"active_ability"` //[ "type", 2, 3] //Type 2, 3x
}
//Damage = [ main att, sub att, heal ]
type teamDamage struct {
	Damage [3]struct{
		Element int `json:"element"`
		Value float64 `json:"val"`
	} `json:"damage"`

}
type PADHMonster struct {
	ID int `json:"id"`
	Url string `json:"url"`
	Monster int `json:"monster"`//real id
	Note string `json:"note"`
	Priority int `json:"priority"`
	CurrXP int `json:"current_xp"`
	CurrSkill int `json:"current_skill"`
	CurrAwaken int `json:"current_awakening"`
	TargetLevel int `json:"target_level"`
	TargetEvo int `json:"target_evolution"`
	PlusHP int `json:"plus_hp"`
	PlusATK int `json:"plus_atk"`
	PlusRCV int `json:"plus_rcv"`
	
}
type PADHTeam struct {
	ID int `json:"id"`
	URL string `json:"url"`
	Name string `json:"name"`
	Desc string `json:"description"`
	Fav bool `json:"favourite"`
	Order int `json:"order"`
	TeamGroup int `json:"team_group"`
	Leader int `json:"leader"`
	Sub1 int `json:"sub1"`
	Sub2 int `json:"sub2"`
	Sub3 int `json:"sub3"`
	Sub4 int `json:"sub4"`
	FLead int `json:"friend_leader"`
	FLevel int `json:"friend_level"`
	FHP int `json:"friend_hp"`
	FATK int `json:"friend_atk"`
	FRCV int `json:"friend_rcv"`
	FSkill int `json:"friend_skill"`
	FAwaken int `json:"friend_awakening"`
}

type Awakenings struct {
	Desc string `json:"desc"` //irrelevant
	ID int `json:"id"` //lookup
	Name string `json:"name"` //irrelevant
}

type LeaderSkill struct {
	Data []interface{} `json:"data"` //either int's or int's and a []
	Effect string `json:"effect"` //metadata
	Name string `json:"name"` //unfortunately, necessary for lookup.. string comparison is so much worse than int, sigh
}

/*
Example
    {
        "data": [
            1, //hp
            2.5, //attack
            1, //rcv
            [
                "elem", //element
                4 //dark
            ]
        ],
        "effect": "2.5x ATK for Dark type monsters when reaching 4 combos.",
        "name": "Dark Harmony"
    },


Another example
    {
        "data": [
            1.5,
            1.5,
            1.5,
            [
                "type", //Type
                7 //Devil
            ]
        ],
        "effect": "1.5x HP, ATK, and RCV for Devil type monsters.",
        "name": "バーサークソウル"
    },
*/
