package generate_password

import (
	"crypto/rand"
	"fmt"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GeneratePassword(n int) (string, error) {
	if n == 0 {
		return "", nil
	}
	if n <= 0 {
		return "", fmt.Errorf("длина пароля должна быть больше 0")
	}

	bytes := make([]byte, n)
	lettersLen := byte(len(letters))

	for i := 0; i < n; i++ {
		var b [1]byte
		_, err := rand.Read(b[:])
		if err != nil {
			return "", fmt.Errorf("ошибка генерации случайного числа: %v", err)
		}
		bytes[i] = letters[b[0]%lettersLen]
	}

	return string(bytes), nil
}

// Что можно улучшить:
// 1. добавить crypto/rand так, чтобы каждый символ выбирался с равной вероятностью, без остатка от деления.
// 2. хранить историю сгенерированных паролей, чтобы не допустить повторений

// Можно ли угадать генерируемую строку?
// - Если используется crypto/rand, угадать пароль практически невозможно.
// - Даже если человек знает алгоритм и набор символов, для длины 16 символов вероятность подобрать пароль случайно: 4.77*10^−29
// - но возможны коллизии 