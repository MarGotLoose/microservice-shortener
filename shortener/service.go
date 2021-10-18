package shortener

type RedirectService interface {
	Find(code string) (*Redirect, error) //Pasamos un códido y devuelve la dirección o error

	Store(redirect *Redirect) error //Pasamos una dirección y en caso de devolver algo sería un error
}
