type Message struct {
     Eventtype string `json:"eventtype,omitempty"`
     Gametime float64 `json:"gametime,omitempty"`
     Iskillerleft bool `json:"iskillerleft,omitempty"`
     Killer float64 `json:"killer,omitempty"`
     Victim float64 `json:"victim,omitempty"`
}
 type Stats struct {
     Account []float64 `json:"account,omitempty"`
     Assists []float64 `json:"assists,omitempty"`
     Buybackcooldown []float64 `json:"buybackcooldown,omitempty"`
     Deaths []float64 `json:"deaths,omitempty"`
     Denies []float64 `json:"denies,omitempty"`
     Direassists float64 `json:"direassists,omitempty"`
     Direbarrack float64 `json:"direbarrack,omitempty"`
     Diredeath float64 `json:"diredeath,omitempty"`
     Direheros []float64 `json:"direheros,omitempty"`
     Direkill float64 `json:"direkill,omitempty"`
     Diretowers float64 `json:"diretowers,omitempty"`
     Draft Draft `json:"draft,omitempty"`
     Exp []float64 `json:"exp,omitempty"`
     Expleft []float64 `json:"expleft,omitempty"`
     Gametime float64 `json:"gametime,omitempty"`
     Gold []float64 `json:"gold,omitempty"`
     HealthPercent []float64 `json:"health_percent,omitempty"`
     Items map[string][]float64 `json:"items,omitempty"`
     KillDiff float64 `json:"kill_diff,omitempty"`
     Kills []float64 `json:"kills,omitempty"`
     Lasthits []float64 `json:"lasthits,omitempty"`
     Leftscore float64 `json:"leftscore,omitempty"`
     Level []float64 `json:"level,omitempty"`
     Message []Message `json:"message,omitempty"`
     Networth []float64 `json:"networth,omitempty"`
     NetworthDiff float64 `json:"networth_diff,omitempty"`
     Playername []string `json:"playername,omitempty"`
     Radiantassists float64 `json:"radiantassists,omitempty"`
     Radiantbarrack float64 `json:"radiantbarrack,omitempty"`
     Radiantdeath float64 `json:"radiantdeath,omitempty"`
     Radiantheros []float64 `json:"radiantheros,omitempty"`
     Radiantkill float64 `json:"radiantkill,omitempty"`
     Radianttowers float64 `json:"radianttowers,omitempty"`
     Respawntime []float64 `json:"respawntime,omitempty"`
     Rightscore float64 `json:"rightscore,omitempty"`
     Roshanstate string `json:"roshanstate,omitempty"`
     RoshanstateEndSeconds float64 `json:"roshanstate_end_seconds,omitempty"`
     Xpos []float64 `json:"xpos,omitempty"`
     Xposv []float64 `json:"xposv,omitempty"`
     Ypos []float64 `json:"ypos,omitempty"`
     Yposv []float64 `json:"yposv,omitempty"`
}
 type Base struct {
     Game string `json:"game,omitempty"`
     Gamestate string `json:"gamestate,omitempty"`
     Gametime float64 `json:"gametime,omitempty"`
     Ispaused bool `json:"ispaused,omitempty"`
     LeftName float64 `json:"left_name,omitempty"`
     MatchId float64 `json:"match_id,omitempty"`
     RightName float64 `json:"right_name,omitempty"`
     SeriesCode float64 `json:"series_code,omitempty"`
     StartTime float64 `json:"start_time,omitempty"`
     Steamtime float64 `json:"steamtime,omitempty"`
     StreamDelay float64 `json:"stream_delay,omitempty"`
}
 type Key struct {
     Base Base `json:"base,omitempty"`
     Stats Stats `json:"stats,omitempty"`
}
 type Draft struct {
     ActiveTeam string `json:"active_team,omitempty"`
     Direban []float64 `json:"direban,omitempty"`
     Direpick []float64 `json:"direpick,omitempty"`
     Radiantban []float64 `json:"radiantban,omitempty"`
     Radiantpick []float64 `json:"radiantpick,omitempty"`
}
 