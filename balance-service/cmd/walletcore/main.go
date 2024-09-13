package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/database"
	get_account_balance "github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/get-account-balance"
	update_account_balance "github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/update-account-balance"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/web"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/web/webserver"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

type BalanceUpdatedKafkaDto struct {
	Name    string                     `json:"Name"`
	Payload BalanceUpdatedKafkaPayload `json:"Payload"`
}

type BalanceUpdatedKafkaPayload struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql-balance", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	accountDb := database.NewAccountDB(db)
	updateAccountBalanceUseCase := update_account_balance.NewUpdateAccountBalanceUseCase(accountDb)
	getAccountBalanceUseCase := get_account_balance.NewGetAccountBalanceUseCase(accountDb)

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	topics := []string{"balances"}
	kafkaConsumer := kafka.NewConsumer(&configMap, topics)
	msgChan := make(chan *ckafka.Message)
	go func() {
		err := kafkaConsumer.Consume(msgChan)
		if err != nil {
			fmt.Println("Erro no consumidor:", err)
		}
	}()

	go func() {
		for msg := range msgChan {
			msgHandle, err := handleMessage(msg)
			if err != nil {
				fmt.Printf("Erro handle message kafka: %s\n", err.Error())
				continue
			}

			updateBalanceFrom := update_account_balance.UpdateAccountBalanceInputDto{
				AccountId: msgHandle.Payload.AccountIDFrom,
				Amount:    msgHandle.Payload.BalanceAccountIDFrom,
			}
			updateBalanceTo := update_account_balance.UpdateAccountBalanceInputDto{
				AccountId: msgHandle.Payload.AccountIDTo,
				Amount:    msgHandle.Payload.BalanceAccountIDTo,
			}

			go updateAccountBalanceUseCase.Execute(updateBalanceFrom)
			go updateAccountBalanceUseCase.Execute(updateBalanceTo)
		}
	}()

	webserver := webserver.NewWebServer(":3003")
	accountHandler := web.NewWebAccountHandler(*getAccountBalanceUseCase)
	webserver.AddHandler("/balance/{account_id}", accountHandler.GetBalance)
	fmt.Println("Server is running")
	webserver.Start()
}

func handleMessage(msg *ckafka.Message) (*BalanceUpdatedKafkaDto, error) {
	fmt.Println("Mensagem Kafka recebida:", string(msg.Value))
	var dto BalanceUpdatedKafkaDto
	err := json.Unmarshal(msg.Value, &dto)
	if err != nil {
		return nil, err
	}

	return &dto, nil
}
