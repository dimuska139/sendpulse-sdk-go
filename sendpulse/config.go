package sendpulse

type Config struct {
	UserID string
	Secret string
	Rps    int // Max allowed count of requests per second (default: 10)
}
