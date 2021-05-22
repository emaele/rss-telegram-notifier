package types

type ConfigurationParameters struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string

	AuthorizationToken string

	BindAddress string

	TelegramToken  string
	TelegramChatID int64
}
