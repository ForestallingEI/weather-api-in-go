package structs

// 天気データ構造
type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	LocalTime string  `json:"localtime"`
}

type Current struct {
	LastUpdated string     `json:"last_updated"`
	TempC       float64    `json:"temp_c"`
	Condition   Condition  `json:"condition"`
	WindKph     float64    `json:"wind_kph"`
	WindDir     string     `json:"wind_dir"`
	PressureMB  float64    `json:"pressure_mb"`
	PrecipMm    float64    `json:"precip_mm"`
	Humidity    int64      `json:"humidity"`
	Cloud       int64      `json:"cloud"`
	FeelslikeC  float64    `json:"feelslike_c"`
	VisKM       float64    `json:"vis_km"`
	UV          float64    `json:"uv"`
	GustKph     float64    `json:"gust_kph"`
	AirQuality  AirQuality `json:"air_quality"`
}

type AirQuality struct {
	Co           float64 `json:"co"`
	No2          float64 `json:"no2"`
	O3           float64 `json:"o3"`
	So2          float64 `json:"so2"`
	Pm25         float64 `json:"pm2_5"`
	Pm10         float64 `json:"pm10"`
	UsEpaIndex   int     `json:"us-epa-index"`
	GbDefraIndex int     `json:"gb-defra-index"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
}
