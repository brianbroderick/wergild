package mud

var (
	ServerInstance *Server
	// BcryptCost is set lower when testing to speed up tests. It determines how much CPU
	// is used when hashing passwords
	BcryptCost = 14
)
