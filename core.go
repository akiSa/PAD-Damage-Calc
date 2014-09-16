package main

import (
	"fmt"
)
//dmg = multiplier for the elements, 0-5 fire-heal
//TODO: Factor in damage enhance AND row enhance, just crossreference the awakening list, hardcode what each one does basically >_>, have it auto set # of enhanced orbs on web front end, user can change.
func damageResolve (team []lookup, teamD []teamDamage, dmg []float64, msg *orbs) (res []teamDamage) {
	var comboMulti float64
	var comboCount float64
	
	comboMulti, comboCount=0,0
	
	for _,y := range msg.Fire {
		comboCount ++
		dmg[0] += 1 + (( y - 3)*0.25)
	}
	for _,y := range msg.Water {
		comboCount ++
		dmg[1] += 1 + (( y - 3)*0.25)
	}
	for _,y := range msg.Wood {
		comboCount ++
		dmg[2] += 1 + (( y - 3)*0.25)
	}
	for _,y := range msg.Light {
		comboCount ++
		dmg[3] += 1 + (( y - 3)*0.25)
	}
	for _,y := range msg.Dark {
		comboCount ++
		dmg[4] += 1 + (( y - 3)*0.25)
	}
	for _,y := range msg.Heart {
		comboCount ++
		dmg[5] += 1 + (( y - 3)*0.25)
	}

	comboMulti = 1 + (( comboCount - 1)*0.25)
	fmt.Println("Combo Multiplier:", comboMulti)
	//factor in leaderskill last
	var subMulti float64
	var lead, friend bool
	subMulti = 0
	for x,_ := range teamD {
		//Main attribute
		lead,friend = false, false
		teamD[x].Damage[0].Element = team[x].Element
		teamD[x].Damage[0].Value = dmg[team[x].Element] * float64(team[x].Stats.ATK) * comboMulti
		//Sub attribute
		teamD[x].Damage[1].Element = team[x].Element2
		if team[x].Element == team[x].Element2 { subMulti = 0.10 } else { subMulti = 0.30 }
		teamD[x].Damage[1].Value = dmg[team[x].Element2] * (float64(team[x].Stats.ATK)*subMulti) * comboMulti
		//Heal
		teamD[x].Damage[2].Element = 6
		teamD[x].Damage[2].Value = dmg[5] * float64(team[x].Stats.RCV) * comboMulti
		//teamD
		//team[x]
		
		//Leader Skill
		switch team[0].LeaderSkill.Conditional[0] {
		case "type":
			//Apply only to conditional[1] types
			if team[x].Type == team[0].LeaderSkill.Conditional[1] ||
				team[x].Type2 == team[0].LeaderSkill.Conditional[1] {
				if !lead {				
					teamD[x].Damage[0].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[1].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[2].Value *= team[0].LeaderSkill.RCV
					lead = true
				}
				//I THINK it works like this, need to check, but I remember d/l batman enhancing his own dark attacks too
			}
		case "elem":
			//Apply only to conditional[1] elements
			if team[x].Element == team[0].LeaderSkill.Conditional[1] ||
				team[x].Element2 == team[0].LeaderSkill.Conditional[1] {
				
				if !lead {
					teamD[x].Damage[0].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[1].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[2].Value *= team[0].LeaderSkill.RCV
					lead = true
				}
			}
		default:
			//Apply to errbody
			if !lead {
				teamD[x].Damage[0].Value *= team[0].LeaderSkill.ATK
				teamD[x].Damage[1].Value *= team[0].LeaderSkill.ATK
				teamD[x].Damage[2].Value *= team[0].LeaderSkill.RCV
				lead = true
			}
		}

		switch team[0].LeaderSkill.Conditional2[0] {
		case "type":
			//Apply only to conditional[1] types
			if team[x].Type == team[0].LeaderSkill.Conditional2[1] ||
				team[x].Type2 == team[0].LeaderSkill.Conditional2[1] {
				if !lead {				
					teamD[x].Damage[0].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[1].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[2].Value *= team[0].LeaderSkill.RCV
					lead = true
				}
				//I THINK it works like this, need to check, but I remember d/l batman enhancing his own dark attacks too
			}
		case "elem":
			//Apply only to conditional[1] elements
			if team[x].Element == team[0].LeaderSkill.Conditional2[1] ||
				team[x].Element2 == team[0].LeaderSkill.Conditional2[1] {
				
				if !lead {
					teamD[x].Damage[0].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[1].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[2].Value *= team[0].LeaderSkill.RCV
					lead = true
				}
			}
		default:
			//this'll never get called, but i'll leave it here just in case.
			if !lead {
				teamD[x].Damage[0].Value *= team[0].LeaderSkill.ATK
				teamD[x].Damage[1].Value *= team[0].LeaderSkill.ATK
				teamD[x].Damage[2].Value *= team[0].LeaderSkill.RCV
				lead = true
			}
		}

		//******************** FRIEND *********************

		switch team[5].LeaderSkill.Conditional[0] {
		case "type":
			//Apply only to conditional[1] types
			if team[x].Type == team[5].LeaderSkill.Conditional[1] ||
				team[x].Type2 == team[5].LeaderSkill.Conditional[1] {
				if !friend {				
					teamD[x].Damage[0].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[1].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[2].Value *= team[0].LeaderSkill.RCV
					friend = true
				}
			}
		case "elem":
			//Apply only to conditional[1] elements
			if team[x].Element == team[5].LeaderSkill.Conditional[1] ||
				team[x].Element2 == team[5].LeaderSkill.Conditional[1] {
				
				if !friend {
					teamD[x].Damage[0].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[1].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[2].Value *= team[0].LeaderSkill.RCV
					friend = true
				}
			}
		default:
			//Apply to errbody
			if !friend {
				teamD[x].Damage[0].Value *= team[0].LeaderSkill.ATK
				teamD[x].Damage[1].Value *= team[0].LeaderSkill.ATK
				teamD[x].Damage[2].Value *= team[0].LeaderSkill.RCV
				friend = true
			}
		}

		//FRIEND CONDITIONAL 2

		switch team[5].LeaderSkill.Conditional2[0] {
		case "type":
			//Apply only to conditional[1] types
			if team[x].Type == team[5].LeaderSkill.Conditional2[1] ||
				team[x].Type2 == team[5].LeaderSkill.Conditional2[1] {
				if !friend {				
					teamD[x].Damage[0].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[1].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[2].Value *= team[0].LeaderSkill.RCV
					friend = true
				}
			}
		case "elem":
			//Apply only to conditional[1] elements
			if team[x].Element == team[5].LeaderSkill.Conditional2[1] ||
				team[x].Element2 == team[5].LeaderSkill.Conditional2[1] {
				
				if !friend {
					teamD[x].Damage[0].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[1].Value *= team[0].LeaderSkill.ATK
					teamD[x].Damage[2].Value *= team[0].LeaderSkill.RCV
					friend = true
				}
			}
		default:
			//Apply to errbody
			if !friend {
				teamD[x].Damage[0].Value *= team[0].LeaderSkill.ATK
				teamD[x].Damage[1].Value *= team[0].LeaderSkill.ATK
				teamD[x].Damage[2].Value *= team[0].LeaderSkill.RCV
				friend = true
			}
		}
	}

	res = teamD
	return
}
