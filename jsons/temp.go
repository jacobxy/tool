package main

type PicksBans struct {
	HeroId float64 `json:"hero_id"`
	IsPick bool    `json:"is_pick"`
	Order  float64 `json:"order"`
	Team   float64 `json:"team"`
}
type AdditionalUnits struct {
	AccountId       float64           `json:"account_id"`
	AdditionalUnits []AdditionalUnits `json:"additional_units"`
	Assists         float64           `json:"assists"`
	Backpack0       float64           `json:"backpack_0"`
	Backpack1       float64           `json:"backpack_1"`
	Backpack2       float64           `json:"backpack_2"`
	Deaths          float64           `json:"deaths"`
	Denies          float64           `json:"denies"`
	GoldPerMin      float64           `json:"gold_per_min"`
	HeroId          float64           `json:"hero_id"`
	Item0           float64           `json:"item_0"`
	Item1           float64           `json:"item_1"`
	Item2           float64           `json:"item_2"`
	Item3           float64           `json:"item_3"`
	Item4           float64           `json:"item_4"`
	Item5           float64           `json:"item_5"`
	ItemNeutral     float64           `json:"item_neutral"`
	Kills           float64           `json:"kills"`
	LastHits        float64           `json:"last_hits"`
	LeaverStatus    float64           `json:"leaver_status"`
	Level           float64           `json:"level"`
	PlayerSlot      float64           `json:"player_slot"`
	Unitname        string            `json:"unitname"`
	XpPerMin        float64           `json:"xp_per_min"`
}
type Result struct {
	BarracksStatusDire    float64           `json:"barracks_status_dire"`
	BarracksStatusRadiant float64           `json:"barracks_status_radiant"`
	Cluster               float64           `json:"cluster"`
	DireCaptain           float64           `json:"dire_captain"`
	DireLogo              float64           `json:"dire_logo"`
	DireName              string            `json:"dire_name"`
	DireScore             float64           `json:"dire_score"`
	DireTeamComplete      float64           `json:"dire_team_complete"`
	DireTeamId            float64           `json:"dire_team_id"`
	Duration              float64           `json:"duration"`
	Engine                float64           `json:"engine"`
	FirstBloodTime        float64           `json:"first_blood_time"`
	Flags                 float64           `json:"flags"`
	GameMode              float64           `json:"game_mode"`
	HumanPlayers          float64           `json:"human_players"`
	Leagueid              float64           `json:"leagueid"`
	LobbyType             float64           `json:"lobby_type"`
	MatchId               float64           `json:"match_id"`
	MatchSeqNum           float64           `json:"match_seq_num"`
	NegativeVotes         float64           `json:"negative_votes"`
	PicksBans             []PicksBans       `json:"picks_bans"`
	Players               []AdditionalUnits `json:"players"`
	PositiveVotes         float64           `json:"positive_votes"`
	PreGameDuration       float64           `json:"pre_game_duration"`
	RadiantCaptain        float64           `json:"radiant_captain"`
	RadiantLogo           float64           `json:"radiant_logo"`
	RadiantName           string            `json:"radiant_name"`
	RadiantScore          float64           `json:"radiant_score"`
	RadiantTeamComplete   float64           `json:"radiant_team_complete"`
	RadiantTeamId         float64           `json:"radiant_team_id"`
	RadiantWin            bool              `json:"radiant_win"`
	StartTime             float64           `json:"start_time"`
	TowerStatusDire       float64           `json:"tower_status_dire"`
	TowerStatusRadiant    float64           `json:"tower_status_radiant"`
}
