type App struct {
	Router *mux.Router
	DB 		 *sql.DB
}

func (a *App) Initialize(user, password, dbname string){}
func (a *App) Run(addr string){}