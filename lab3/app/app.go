package app

type DB interface {
	CreateTable() error
	Insert(text string) error
	GetFirst() (string, error)
}

type App struct {
	db DB
}

func New(db DB) *App {
	return &App{db: db}
}

func (a *App) Run() (string, error) {
	if err := a.db.CreateTable(); err != nil {
		return "", err
	}

	if err := a.db.Insert("hello fyne!"); err != nil {
		return "", err
	}

	message, err := a.db.GetFirst()
	if err != nil {
		return "", err
	}

	return message, nil
}
