# Server and Client with RabbitMQ message broker
## Dependencies
Be sure you have instaled `Golang v. 1.19` version in your machine and `Docker` for Server and RabbitMQ launching.
For Golang installation you can follow this [link](https://go.dev/doc/install)
For Docker installation you can follow this [link](https://docs.docker.com/get-docker/)

Project contain these libraries:
- [Cobra](https://github.com/spf13/cobra) for fetching commands through the terminal. Client side.
- [Rabbitmq](https://github.com/rabbitmq/amqp091-go) for message passing between Client and Server via RabbitMQ.
- [OrderdMap](https://github.com/elliotchance/orderedmap) for saving data to a variable type `map`. This library was chosen for simplifying and speeding up development. Server-side.

## Supported commands
- `addItem [item to add]` - add item to Storage.
- `getItem [item to get]` - show item on Server-side.
- `getAllItems` - show all items on Server-side.
- `removeItem [item to remove]` - remove item from Storage.

## Execntion
## Be sure to run `docker-compose up --build` before executing any command on Client-side.
Open folder `client` in your terminal and run `go run main.go addItem test` or any command listed above.

