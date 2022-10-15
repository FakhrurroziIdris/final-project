package configs

type Config struct {
	WebServer       Server
	Database        Database
	CRON            CRON
	WebPresentation WebPresentation
	Authentication  Authentication
}

type Server struct {
	Port string
}

type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	JsonFile string
}

type CRON struct {
	SensorInterval string
}

type WebPresentation struct {
	TemplateFolder string
	PageRefresh    int
}

type Authentication struct {
	ExpiresInMinute int
}
