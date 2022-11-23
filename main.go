package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5836228683:AAGIeREyyrz5A-m1eHEFOM4_VoBgS12YttQ")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := bot.GetUpdatesChan(u)

	var nameOfVictim string
	var lastComamnd string
	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {
		var reply string
		// универсальный ответ на любое сообщение
		if update.Message.Command() == "" && lastComamnd == "/play" {
			reply = "Откусить голову?!\nЖми /kill"
			nameOfVictim = update.Message.Text
		}

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// свитч на обработку комманд
		// комманда - сообщение, начинающееся с "/"
		switch update.Message.Command() {
		case "start":
			reply = "Привет. Я телеграм-бот, твой кровожадный лев в телефоне.\nИграем /play"
			lastComamnd = update.Message.Text
		case "play":
			reply = "Имя того, кто тебя обидел:"
			lastComamnd = update.Message.Text
		case "kill":
			reply = nameOfVictim + " убит(а).\nПовторить /play\nПосмотреть фотографию льва /photo"
			lastComamnd = update.Message.Text
		case "photo":
			photoBytes, _ := os.ReadFile("img.png")
			photoFileBytes := tgbotapi.FileBytes{
				Name:  "picture",
				Bytes: photoBytes,
			}
			bot.Send(tgbotapi.NewPhotoUpload(update.Message.Chat.ID, photoFileBytes))
			reply = "Играть заново /play"

		}

		// создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// отправляем
		bot.Send(msg)
	}
}
