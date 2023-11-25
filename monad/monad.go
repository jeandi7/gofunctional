package main

import (
	"fmt"
)

// ResultMonad représente la monade de résultat
type ResultMonad struct {
	value interface{}
	err   error
}

// NewResult crée une nouvelle instance de ResultMonad avec une valeur
func NewResult(value interface{}) ResultMonad {
	return ResultMonad{value: value, err: nil}
}

// Bind applique une fonction f à la valeur de la monade et retourne une nouvelle monade
func (r ResultMonad) Bind(f func(interface{}) ResultMonad) ResultMonad {
	if r.err != nil {
		return r
	}
	return f(r.value)
}

// Map applique une fonction f à la valeur de la monade et retourne une nouvelle monade avec la nouvelle valeur
func (r ResultMonad) Map(f func(interface{}) interface{}) ResultMonad {
	if r.err != nil {
		return r
	}
	return ResultMonad{value: f(r.value), err: nil}
}

// Err renvoie l'erreur associée à la monade
func (r ResultMonad) Err() error {
	return r.err
}

func main() {
	// Exemple d'utilisation de la monade de résultat
	result := NewResult(10).
		Bind(func(x interface{}) ResultMonad {
			return NewResult(x.(int) * 2)
		}).
		Map(func(x interface{}) interface{} {
			return x.(int) + 5
		})

	if result.Err() != nil {
		fmt.Println("Une erreur s'est produite:", result.Err())
	} else {
		fmt.Println("Résultat final:", result.value)
	}
}
