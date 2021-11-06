package consulta

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) PeticionPozo(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Mensaje recibido desde cliente: %s ", message.Body)
	return &Message{Body: "Hola desde el server"}, nil
}
