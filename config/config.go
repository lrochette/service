// Server holds all config needed to start a support analytics service instance
import "log"

type Server struct {
	ListenPort int
	Logger     log.Logger
}
