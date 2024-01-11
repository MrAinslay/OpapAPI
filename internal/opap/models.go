package opap

type Draw struct {
	GameID      int    `json:"gameId"`
	DrawID      int    `json:"drawId"`
	DrawTime    int64  `json:"drawTime"`
	Status      string `json:"status"`
	DrawBreak   int    `json:"drawBreak"`
	VisualDraw  int    `json:"visualDraw"`
	PricePoints struct {
		AddOn []struct {
			Amount   float64 `json:"amount"`
			GameType string  `json:"gameType"`
		} `json:"addOn"`
		Amount float64 `json:"amount"`
	} `json:"pricePoints"`
	PrizeCategories []struct {
		ID           int    `json:"id"`
		Divident     int    `json:"divident"`
		Winners      int    `json:"winners"`
		Distributed  int    `json:"distributed"`
		Jackpot      int    `json:"jackpot"`
		Fixed        int    `json:"fixed"`
		CategoryType int    `json:"categoryType"`
		GameType     string `json:"gameType"`
	} `json:"prizeCategories"`
	WagerStatistics struct {
		Columns int   `json:"columns"`
		Wagers  int   `json:"wagers"`
		AddOn   []any `json:"addOn"`
	} `json:"wagerStatistics"`
}

type PasedDraws struct {
	Content []struct {
		GameID      int    `json:"gameId"`
		DrawID      int    `json:"drawId"`
		DrawTime    int64  `json:"drawTime"`
		Status      string `json:"status"`
		DrawBreak   int    `json:"drawBreak"`
		VisualDraw  int    `json:"visualDraw"`
		PricePoints struct {
			AddOn []struct {
				Amount   float64 `json:"amount"`
				GameType string  `json:"gameType"`
			} `json:"addOn"`
			Amount float64 `json:"amount"`
		} `json:"pricePoints"`
		WinningNumbers struct {
			List     []int `json:"list"`
			Sidebets struct {
				Symbol bool   `json:"symbol"`
				Color  string `json:"color"`
			} `json:"sidebets"`
		} `json:"winningNumbers"`
		ListWinningNumbers []struct {
			List     []int `json:"list"`
			Sidebets struct {
				Symbol bool   `json:"symbol"`
				Color  string `json:"color"`
			} `json:"sidebets"`
		} `json:"listWinningNumbers"`
		PrizeCategories []struct {
			ID           int     `json:"id"`
			Divident     float64 `json:"divident"`
			Winners      int     `json:"winners"`
			Distributed  float64 `json:"distributed"`
			Jackpot      float64 `json:"jackpot"`
			Fixed        float64 `json:"fixed"`
			CategoryType int     `json:"categoryType"`
			GameType     string  `json:"gameType"`
		} `json:"prizeCategories"`
		WagerStatistics struct {
			Columns int   `json:"columns"`
			Wagers  int   `json:"wagers"`
			AddOn   []any `json:"addOn"`
		} `json:"wagerStatistics"`
	} `json:"content"`
	TotalPages       int  `json:"totalPages"`
	TotalElements    int  `json:"totalElements"`
	Last             bool `json:"last"`
	NumberOfElements int  `json:"numberOfElements"`
	First            bool `json:"first"`
	Sort             []struct {
		Direction    string `json:"direction"`
		Property     string `json:"property"`
		IgnoreCase   bool   `json:"ignoreCase"`
		NullHandling string `json:"nullHandling"`
		Descending   bool   `json:"descending"`
		Ascending    bool   `json:"ascending"`
	} `json:"sort"`
	Size   int `json:"size"`
	Number int `json:"number"`
}

type DrawsByDate struct {
	Content []struct {
		GameID      int    `json:"gameId"`
		DrawID      int    `json:"drawId"`
		DrawTime    int64  `json:"drawTime"`
		Status      string `json:"status"`
		DrawBreak   int    `json:"drawBreak"`
		VisualDraw  int    `json:"visualDraw"`
		PricePoints struct {
			AddOn []struct {
				Amount   float64 `json:"amount"`
				GameType string  `json:"gameType"`
			} `json:"addOn"`
			Amount float64 `json:"amount"`
		} `json:"pricePoints"`
		WinningNumbers struct {
			List     []int `json:"list"`
			Sidebets struct {
				Symbol             bool   `json:"symbol"`
				Color              string `json:"color"`
				PowerSpinUnderOver string `json:"powerSpinUnderOver"`
			} `json:"sidebets"`
		} `json:"winningNumbers"`
		ListWinningNumbers []struct {
			List     []int `json:"list"`
			Sidebets struct {
				Symbol             bool   `json:"symbol"`
				Color              string `json:"color"`
				PowerSpinUnderOver string `json:"powerSpinUnderOver"`
			} `json:"sidebets"`
		} `json:"listWinningNumbers"`
		PrizeCategories []struct {
			ID           int     `json:"id"`
			Divident     float64 `json:"divident"`
			Winners      int     `json:"winners"`
			Distributed  float64 `json:"distributed"`
			Jackpot      float64 `json:"jackpot"`
			Fixed        float64 `json:"fixed"`
			CategoryType int     `json:"categoryType"`
			GameType     string  `json:"gameType"`
		} `json:"prizeCategories"`
		WagerStatistics struct {
			Columns int   `json:"columns"`
			Wagers  int   `json:"wagers"`
			AddOn   []any `json:"addOn"`
		} `json:"wagerStatistics"`
	} `json:"content"`
	TotalPages       int  `json:"totalPages"`
	TotalElements    int  `json:"totalElements"`
	Last             bool `json:"last"`
	NumberOfElements int  `json:"numberOfElements"`
	Sort             []struct {
		Direction    string `json:"direction"`
		Property     string `json:"property"`
		IgnoreCase   bool   `json:"ignoreCase"`
		NullHandling string `json:"nullHandling"`
		Descending   bool   `json:"descending"`
		Ascending    bool   `json:"ascending"`
	} `json:"sort"`
	First  bool `json:"first"`
	Size   int  `json:"size"`
	Number int  `json:"number"`
}
