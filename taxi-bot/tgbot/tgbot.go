package tgbot

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"taxi-bot/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			switch update.Message.Command() {
			case "Добавить машину такси":
				_, taskText, _ := strings.Cut(update.Message.Text, update.Message.Command()+" ")

				if len(taskText) == 0 {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы не ввели значение.")
					bot.Send(msg)
					continue
				}

				err := database.SetCar(update.Message.From.ID, taskText)
				message := ""

				if err != nil {
					message = "Ошибка"
				} else {
					message = "Готово"
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(msg)
				continue
			case "delete":
				_, TaxiInfo, _ := strings.Cut(update.Message.Text, update.Message.Command()+" ")

				if len(TaxiInfo) == 0 {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пусто.")
					bot.Send(msg)
					continue
				}

				err := database.DeleteCar(update.Message.From.ID, TaxiInfo)

				message := ""

				if err != nil {
					message = "Ошибка"
				} else {
					message = "Готово"
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(msg)
				continue
			case "list":
				cars, err := database.GetListCars(update.Message.From.ID)

				message := ""

				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка")
					bot.Send(msg)
					continue
				}

				if len(cars) > 0 {

					message = "Машины такси:\n"

					for i := 0; i < len(cars); i++ {
						message += strconv.Itoa(i+1) + ") " + cars[i] + "\n"
					}
				} else {
					message = "Нет доступных машин"
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(msg)
				continue
			case "help":
				message := "/add - Добавить машину такси.\n/delete - Убрать машину такси.\n/list - Показать все машины такси.\n/help - Доступные команды."
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(msg)
				continue
			default:
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда."))
			}
		}
	}
}
