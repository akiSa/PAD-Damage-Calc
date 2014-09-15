package main

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
	PlusAtk int `json:"plus_atk"`
	PlusRCV int `json:"plus_rcv"`
	
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
