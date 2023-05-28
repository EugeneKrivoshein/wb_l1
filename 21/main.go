package main

// Реализовать паттерн «адаптер» на любом примере.

import "fmt"

// Интерфейс для системы управления пользователями
type UserManager interface {
	GetUserByID(userID int) string
}

// Реализация системы управления пользователями
type UserManagementSystem struct{}

func (ums *UserManagementSystem) GetUserByID(userID int) string {
	// Здесь была бы реализация получения информации о пользователе по ID
	return fmt.Sprintf("User with ID %d", userID)
}

// Старая система с несовместимой функцией
type LegacySystem struct{}

func (ls *LegacySystem) GetUserInfo(userKey string) string {
	// Здесь была бы реализация получения информации о пользователе по ключу
	return fmt.Sprintf("User info for key: %s", userKey)
}

// Адаптер для использования функции старой системы в новой системе управления пользователями
type LegacyAdapter struct {
	LegacySystem *LegacySystem
}

func (la *LegacyAdapter) GetUserByID(userID int) string {
	// Адаптер вызывает функцию GetUserInfo старой системы и преобразует результат
	userKey := fmt.Sprintf("ID:%d", userID)
	return la.LegacySystem.GetUserInfo(userKey)
}

func main() {
	// Создаем экземпляры системы управления пользователями и старой системы
	userManager := &UserManagementSystem{}
	legacySystem := &LegacySystem{}

	// Создаем адаптер для старой системы
	adapter := &LegacyAdapter{
		LegacySystem: legacySystem,
	}

	// Используем адаптер вместо старой системы
	userInfo := adapter.GetUserByID(123)
	fmt.Println(userInfo)

	// Используем новую систему управления пользователями
	userInfo = userManager.GetUserByID(456)
	fmt.Println(userInfo)
}
