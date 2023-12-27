package cryptorg_api

type BotDetail struct {
	ID        int      `json:"id"`
	UserID    int      `json:"userId"`
	PairTitle string   `json:"pairTitle"`
	AccessID  int      `json:"accessId"`
	Title     string   `json:"title"`
	Status    string   `json:"status"`
	Strategy  string   `json:"strategy"`
	Leverage  int      `json:"leverage"`
	Settings  Settings `json:"settings"`
	Updated   Updated  `json:"updated"`
	Exchange  string   `json:"exchange"`
}

type Settings struct {
	MarginMode         string      `json:"margin_mode"`
	TypeTakeProfit     string      `json:"type_take_profit"`
	TypeFirstOrder     string      `json:"type_first_order"`
	PercentTakeProfit  string      `json:"percent_take_profit"`
	PercentSafetyOrder string      `json:"percent_safety_order"`
	VolumeFirstOrder   string      `json:"volume_first_order"`
	VolumeSafetyOrder  string      `json:"volume_safety_order"`
	CountOrderWorking  string      `json:"count_order_working"`
	CountOrderMax      string      `json:"count_order_max"`
	Martingale         string      `json:"martingale"`
	MartingaleScale    string      `json:"martingale_scale"`
	DynamicStep        string      `json:"dynamic_step"`
	DynamicStepScale   string      `json:"dynamic_step_scale"`
	CycleLimit         string      `json:"cycle_limit"`
	CycleLimitScale    string      `json:"cycle_limit_scale"`
	TbGroupIndicator   string      `json:"tb_group_indicator"`
	TbIndicator        string      `json:"tb_indicator"`
	TbInterval         string      `json:"tb_interval"`
	PriceLevel         string      `json:"price_level"`
	PriceLevelUp       interface{} `json:"price_level_up"`
	PriceLevelDown     interface{} `json:"price_level_down"`
}

type Updated struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}

type BotDetailsResponse struct {
	Data         []BotDetail `json:"data"`
	Success      bool        `json:"isSuccess"`
	ErrorMessage string      `json:"errorMessage"`
}
