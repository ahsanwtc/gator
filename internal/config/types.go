package config

const configFileName = ".gatorconfig.json"

type Config struct {
	DB_URL string `json:"db_url"`
	CURRENT_USER string `json:"current_user_name"`
}