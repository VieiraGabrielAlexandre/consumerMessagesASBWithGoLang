package consumers

import (
	"bytes"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"golang.org/x/net/context"
	"origomicrosservices.com/consumer-contratos/domain/services"
	"os"
)

func ExecuteAll() {

	client := GetClient()
	fmt.Println("Consumindo contratos finalizados")
	//SendMessage("firstMessage", client)
	//GetMessage(10, client)

	fmt.Println("Consumindo contratos finalizados - Dead Letter")
	DeadLetterMessage(client)
}

func GetClient() *azservicebus.Client {
	asb, err := azservicebus.NewClientFromConnectionString(os.Getenv("TOPIC_CONTRATOS_FINALIZADOS_STRING_CONNECTION"), nil)

	if err != nil {
		panic(err)
	}

	return asb
}

func SendMessage(message string, client *azservicebus.Client) {
	sender, err := client.NewSender(os.Getenv("TOPIC_CONTRATOS_FINALIZADOS"), nil)
	if err != nil {
		panic(err)
	}
	defer sender.Close(context.TODO())

	sbMessage := &azservicebus.Message{
		Body: []byte(message),
	}
	err = sender.SendMessage(context.TODO(), sbMessage, nil)
	if err != nil {
		panic(err)
	}
}

func GetMessage(count int, client *azservicebus.Client) {
	receiver, err := client.NewReceiverForSubscription(os.Getenv("TOPIC_CONTRATOS_FINALIZADOS"), os.Getenv("TOPIC_CONTRATOS_FINALIZADOS_SUBSCRIPTION"), nil)

	if err != nil {
		panic(err)
	}
	defer receiver.Close(context.TODO())

	messages, err := receiver.ReceiveMessages(context.TODO(), count, nil)

	if err != nil {
		panic(err)
	}

	for _, message := range messages {
		body := message.Body
		fmt.Printf("%s\n", string(body))

		err = receiver.CompleteMessage(context.TODO(), message, nil)
		if err != nil {
			panic(err)
		}
	}
}

func DeadLetterMessage(client *azservicebus.Client) {
	receiver, err := client.NewReceiverForSubscription(
		os.Getenv("TOPIC_CONTRATOS_FINALIZADOS"), os.Getenv("TOPIC_CONTRATOS_FINALIZADOS_SUBSCRIPTION"),
		&azservicebus.ReceiverOptions{
			SubQueue: azservicebus.SubQueueDeadLetter,
		},
	)

	if err != nil {
		panic(err)
	}
	defer receiver.Close(context.TODO())

	for true {
		messages, err := receiver.ReceiveMessages(context.TODO(), 1, nil)
		if err != nil {
			panic(err)
		}

		for _, message := range messages {
			fmt.Printf("DeadLetter Reason: %s\nDeadLetter Description: %s\n", *message.DeadLetterReason, *message.DeadLetterErrorDescription) //change to struct an unmarshal into it

			services.SaveJSON(message.MessageID, bytes.NewBuffer(message.Body).String())
		}
	}
}
