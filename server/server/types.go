


type Server struct {
	Db *sql.DB
	Router *mux.Router
	Logger *zerolog.Logger
}