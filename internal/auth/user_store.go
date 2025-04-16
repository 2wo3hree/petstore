package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

var (
	users = make(map[string]string) // username: hashed password
	mu    sync.RWMutex
)

func RegisterUser(username, password string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[username]; exists {
		return errors.New("такой пользователь уже существует")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("ошибка хеширования пароля")
	}

	users[username] = string(hash)
	return nil
}

func AuthenticateUser(username, password string) error {
	mu.RLock()
	hash, exists := users[username]
	mu.RUnlock()

	if !exists {
		return errors.New("пользователь не найден")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return errors.New("неверный пароль")
	}

	return nil
}
