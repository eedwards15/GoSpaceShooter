package definitions

type EnemyConfig struct {
	Key      string  `json:"Key"`
	Chance   int     `json:"Chance"`
	CoolDown float64 `json:"CoolDown"`
	Assets   []struct {
		LocationKey string `json:"LocationKey"`
		ImageKey    string `json:"ImageKey"`
	} `json:"Assets"`
	CanShoot      bool    `json:"CanShoot"`
	Life          int     `json:"Life"`
	ScoreAmount   int     `json:"ScoreAmount"`
	FireRate      int64   `json:"FireRate"`
	MovementSpeed float64 `json:"MovementSpeed"`
}
