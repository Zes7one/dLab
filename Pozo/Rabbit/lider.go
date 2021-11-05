package main

import(
	"log"
	// "os"
	amqp "github.com/rabbitmq/amqp091-go"
)

// helper function to check the return value for each amqp call:
func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}

func main() {
	//*** Recibir peticion de jugadores ***


	// connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://user:123@10.6.40.206:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Next we create a channel, which is where most of the API for getting things done resides:
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// To send, we must declare a queue for us to send to; then we can publish a message to the queue:
	q, err := ch.QueueDeclare(
	  "hello", // name
	  false,   // durable
	  false,   // delete when unused
	  false,   // exclusive
	  false,   // no-wait
	  nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	
	// Mensaje con numero de jugador, numero de ronda donde fue eliminado
	body := "Jugador 1 Ronda 1"



	
	err = ch.Publish(
	  "",     // exchange
	  q.Name, // routing key
	  false,  // mandatory
	  false,  // immediate
	  amqp.Publishing {
	    ContentType: "text/plain",
	    Body:        []byte(body),
	  })
	failOnError(err, "Failed to publish a message")


}


// cd "Desktop\UNI\-----TOMADOS-----\DISTRIBUIDOS\Tareas\Squid\lpozo"
//go mod init "Desktop/UNI/-----TOMADOS-----/DISTRIBUIDOS/Tareas/Squid/lpozo"
//go install "Desktop/UNI/-----TOMADOS-----/DISTRIBUIDOS/Tareas/Squid/lpozo"