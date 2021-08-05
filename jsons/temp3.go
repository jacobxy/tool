type Player6 struct {
     Ability0 Slot4 `json:"ability0,omitempty"`
     Ability1 KillList `json:"ability1,omitempty"`
     Ability2 Ability1 `json:"ability2,omitempty"`
     Ability3 KillList `json:"ability3,omitempty"`
     Ability4 KillList `json:"ability4,omitempty"`
     Ability5 Ability1 `json:"ability5,omitempty"`
}
 type Items struct {
     Team2 Team2 `json:"team2,omitempty"`
     Team3 Team3 `json:"team3,omitempty"`
}
 type Player5 struct {
     Neutral0 map[string]float64 `json:"neutral0,omitempty"`
     Slot0 Ability1 `json:"slot0,omitempty"`
}
 type Player1 struct {
     Ability0 Neutral0 `json:"ability0,omitempty"`
     Ability1 KillList `json:"ability1,omitempty"`
     Ability2 KillList `json:"ability2,omitempty"`
     Ability2 Neutral0 `json:"ability2,omitempty"`
     Ability3 KillList `json:"ability3,omitempty"`
     Ability4 KillList `json:"ability4,omitempty"`
     Ability5 Neutral0 `json:"ability5,omitempty"`
}
 type Previously struct {
     Abilities Abilities `json:"abilities,omitempty"`
     Hero map[string]map[string]map[string]float64 `json:"hero,omitempty"`
     Items Abilities `json:"items,omitempty"`
     Map map[string]float64 `json:"map,omitempty"`
     Player map[string]map[string]map[string]float64 `json:"player,omitempty"`
}
 type KillList struct {
     AbilityActive bool `json:"ability_active,omitempty"`
     CanCast bool `json:"can_cast,omitempty"`
     ChargeCooldown float64 `json:"charge_cooldown,omitempty"`
     Charges float64 `json:"charges,omitempty"`
     Cooldown float64 `json:"cooldown,omitempty"`
     Level float64 `json:"level,omitempty"`
     MaxCharges float64 `json:"max_charges,omitempty"`
     Name string `json:"name,omitempty"`
     Passive bool `json:"passive,omitempty"`
     Ultimate bool `json:"ultimate,omitempty"`
}
 type Slot5 struct {
     AbilityActive bool `json:"ability_active,omitempty"`
     CanCast bool `json:"can_cast,omitempty"`
     ChargeCooldown float64 `json:"charge_cooldown,omitempty"`
     Charges float64 `json:"charges,omitempty"`
     ContainsRune string `json:"contains_rune,omitempty"`
     Cooldown float64 `json:"cooldown,omitempty"`
     Level float64 `json:"level,omitempty"`
     MaxCharges float64 `json:"max_charges,omitempty"`
     Name string `json:"name,omitempty"`
     Passive bool `json:"passive,omitempty"`
     Purchaser float64 `json:"purchaser,omitempty"`
     Ultimate bool `json:"ultimate,omitempty"`
}
 type Team2 struct {
     Player0 Player0 `json:"player0,omitempty"`
     Player1 Player1 `json:"player1,omitempty"`
     Player2 Player2 `json:"player2,omitempty"`
     Player3 Player3 `json:"player3,omitempty"`
     Player4 Player1 `json:"player4,omitempty"`
}
 type Player3 struct {
     Ability0 KillList `json:"ability0,omitempty"`
     Ability1 Neutral0 `json:"ability1,omitempty"`
     Ability2 KillList `json:"ability2,omitempty"`
     Ability3 Slot4 `json:"ability3,omitempty"`
     Ability4 KillList `json:"ability4,omitempty"`
     Ability5 Neutral0 `json:"ability5,omitempty"`
}
 type Player7 struct {
     Ability0 KillList `json:"ability0,omitempty"`
     Ability1 Neutral0 `json:"ability1,omitempty"`
     Ability2 Neutral0 `json:"ability2,omitempty"`
     Ability3 Neutral0 `json:"ability3,omitempty"`
     Ability4 Neutral0 `json:"ability4,omitempty"`
     Ability5 Neutral0 `json:"ability5,omitempty"`
}
 type Neutral0 struct {
     AbilityActive bool `json:"ability_active,omitempty"`
     CanCast bool `json:"can_cast,omitempty"`
     Charges float64 `json:"charges,omitempty"`
     ContainsRune string `json:"contains_rune,omitempty"`
     Cooldown float64 `json:"cooldown,omitempty"`
     Level float64 `json:"level,omitempty"`
     Name string `json:"name,omitempty"`
     Passive bool `json:"passive,omitempty"`
     Purchaser float64 `json:"purchaser,omitempty"`
     Ultimate bool `json:"ultimate,omitempty"`
}
 type Map struct {
     ClockTime float64 `json:"clock_time,omitempty"`
     Customgamename string `json:"customgamename,omitempty"`
     Daytime bool `json:"daytime,omitempty"`
     DireWardPurchaseCooldown float64 `json:"dire_ward_purchase_cooldown,omitempty"`
     GameState string `json:"game_state,omitempty"`
     GameTime float64 `json:"game_time,omitempty"`
     Matchid string `json:"matchid,omitempty"`
     Name string `json:"name,omitempty"`
     NightstalkerNight bool `json:"nightstalker_night,omitempty"`
     Paused bool `json:"paused,omitempty"`
     RadiantWardPurchaseCooldown float64 `json:"radiant_ward_purchase_cooldown,omitempty"`
     RoshanState string `json:"roshan_state,omitempty"`
     RoshanStateEndSeconds float64 `json:"roshan_state_end_seconds,omitempty"`
     WinTeam string `json:"win_team,omitempty"`
}
 type Player2 struct {
     Ability0 KillList `json:"ability0,omitempty"`
     Ability1 KillList `json:"ability1,omitempty"`
     Ability2 Neutral0 `json:"ability2,omitempty"`
     Ability3 KillList `json:"ability3,omitempty"`
     Ability4 Slot4 `json:"ability4,omitempty"`
     Ability5 Slot4 `json:"ability5,omitempty"`
     Ability6 Ability1 `json:"ability6,omitempty"`
     Ability7 KillList `json:"ability7,omitempty"`
}
 type Player8 struct {
     Ability0 Neutral0 `json:"ability0,omitempty"`
     Ability1 KillList `json:"ability1,omitempty"`
     Ability2 Neutral0 `json:"ability2,omitempty"`
     Ability3 Neutral0 `json:"ability3,omitempty"`
     Ability4 Slot4 `json:"ability4,omitempty"`
     Ability5 KillList `json:"ability5,omitempty"`
}
 type Ability1 struct {
     AbilityActive bool `json:"ability_active,omitempty"`
     CanCast bool `json:"can_cast,omitempty"`
     Cooldown float64 `json:"cooldown,omitempty"`
     Level float64 `json:"level,omitempty"`
     Name string `json:"name,omitempty"`
     Passive bool `json:"passive,omitempty"`
     Ultimate bool `json:"ultimate,omitempty"`
}
 type Slot3 struct {
     Charges float64 `json:"charges,omitempty"`
     ContainsRune string `json:"contains_rune,omitempty"`
}
 type Draft struct {
     Abilities Items `json:"abilities,omitempty"`
     Buildings map[string]map[string]map[string]float64 `json:"buildings,omitempty"`
     Draft Draft `json:"draft,omitempty"`
     Hero map[string]map[string]Player0 `json:"hero,omitempty"`
     Items Items `json:"items,omitempty"`
     Map Map `json:"map,omitempty"`
     Player map[string]map[string]Player0 `json:"player,omitempty"`
     Previously Previously `json:"previously,omitempty"`
     Provider Provider `json:"provider,omitempty"`
}
 type Provider struct {
     Appid float64 `json:"appid,omitempty"`
     Name string `json:"name,omitempty"`
     Timestamp float64 `json:"timestamp,omitempty"`
     Version float64 `json:"version,omitempty"`
}
 type Slot4 struct {
     AbilityActive bool `json:"ability_active,omitempty"`
     CanCast bool `json:"can_cast,omitempty"`
     Cooldown float64 `json:"cooldown,omitempty"`
     Level float64 `json:"level,omitempty"`
     Name string `json:"name,omitempty"`
     Passive bool `json:"passive,omitempty"`
     Purchaser float64 `json:"purchaser,omitempty"`
     Ultimate bool `json:"ultimate,omitempty"`
}
 type Team3 struct {
     Player5 Player9 `json:"player5,omitempty"`
     Player6 Player6 `json:"player6,omitempty"`
     Player7 Player7 `json:"player7,omitempty"`
     Player7 Player6 `json:"player7,omitempty"`
     Player8 Player8 `json:"player8,omitempty"`
     Player9 Player9 `json:"player9,omitempty"`
}
 type Abilities struct {
     Team2 Team2 `json:"team2,omitempty"`
     Team2 map[string]map[string]map[string]float64 `json:"team2,omitempty"`
     Team3 map[string]Player5 `json:"team3,omitempty"`
}
 type Player0 struct {
     Ability0 KillList `json:"ability0,omitempty"`
     Ability1 KillList `json:"ability1,omitempty"`
     Ability2 Slot5 `json:"ability2,omitempty"`
     Ability3 KillList `json:"ability3,omitempty"`
     Ability4 KillList `json:"ability4,omitempty"`
     Ability5 Slot5 `json:"ability5,omitempty"`
}
 type Player9 struct {
     Ability0 KillList `json:"ability0,omitempty"`
     Ability1 Slot4 `json:"ability1,omitempty"`
     Ability1 Ability1 `json:"ability1,omitempty"`
     Ability2 KillList `json:"ability2,omitempty"`
     Ability3 KillList `json:"ability3,omitempty"`
     Ability4 KillList `json:"ability4,omitempty"`
     Ability5 KillList `json:"ability5,omitempty"`
     Ability6 Neutral0 `json:"ability6,omitempty"`
}
 