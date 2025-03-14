package common

import (
	"io/ioutil"
	"os"
)

func In(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func Out(path, str string) error {
	// Открываем файл для записи, создаем его, если он не существует
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close() // Закрываем файл после завершения работы

	// Записываем строку в файл
	_, err = file.WriteString(str)
	if err != nil {
		return err
	}

	return nil
}

func OutBytes(path string, str []byte) error {
	// Открываем файл для записи, создаем его, если он не существует
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close() // Закрываем файл после завершения работы

	// Записываем байты в файл
	_, err = file.Write(str)
	if err != nil {
		return err
	}

	return nil
}

func CountFilesInDirectory(dir string) (int, error) {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, entry := range entries {
		if !entry.IsDir() { // Проверяем, что это не директория
			count++
		}
	}
	return count, nil
}
