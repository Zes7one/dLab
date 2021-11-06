package main

import (
  "log"
  "os"
  "fmt"
  "strconv"

  amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}


func main(){

  
  
  conn, err := amqp.Dial("amqp://admin:123@10.6.40.206:5672/")
  failOnError(err, "Failed to connect to RabbitMQ")
  defer conn.Close()

  ch, err := conn.Channel()
  failOnError(err, "Failed to open a channel")
  defer ch.Close()

  q, err := ch.QueueDeclare(
    "hello", // name
    false,   // durable
    false,   // delete when unused
    false,   // exclusive
    false,   // no-wait
    nil,     // arguments
  )
  failOnError(err, "Failed to declare a queue")


  // Definicion nombre archivo txt y variable con pozo actual
  path := "eliminados.txt"
  pozo := 0

  // var _, err1 = os.Stat(path)
  // // create file if not exists
  // if os.IsNotExist(err1) {
  //   f, err := os.Create(path)
  //   failOnError(err, "Failed to create txt file")
  //   defer f.Close()
  // }
  
  msgs, err := ch.Consume(
	  q.Name, // queue
	  "",     // consumer
	  true,   // auto-ack
	  false,  // exclusive
	  false,  // no-local
	  false,  // no-wait
	  nil,    // args
  )
  failOnError(err, "Failed to register a consumer")

  forever := make(chan bool)

  f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
  failOnError(err, "Failed to open txt file")
  
  go func() {
    for d := range msgs {
      fmt.Printf("var1 = %T, %d, %s\n", pozo, pozo, strconv.Itoa(pozo))
      pozo = pozo + 100000000
      fmt.Printf("var2 = %T, %d, %s\n", pozo, pozo, strconv.Itoa(pozo))
      // val := "old falcon"
      // fmt.Printf("var1 = %T\n", string(d.Body))
      inicio:= string(d.Body) + " "+ strconv.Itoa(pozo)
      fin :=  inicio +"\n"
      data := []byte(fin)
      _, err2 := f.Write(data)
      // _, err := f.WriteString("old falcon\n")
      failOnError(err2, "Failed to write on txt file")

      log.Printf("Received a message: %s", d.Body)
    }
  }()

  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  <-forever





}