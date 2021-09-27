package main

import "fmt"

// Objects can suscribed to an event
type Topic interface {
	register(observer Observer)
	//Añadir observadores al objeto
	broadcast() // Notificar a todos los observadores
}

type Observer interface {
	getId() string      // Get the id of the observer
	updateValue(string) // Update the value of the observer, trigger the event
}

// Item -> No disponible
// Cuando tenga disponiblidad, avise a los observadores
type Item struct {
	Observers []Observer // Lista de observadores
	name      string     //Nombre del item
	available bool       //Disponibilidad
}

// NewItem -> Crear un nuevo item
func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

//Mandar el evento
func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available'\n", i.name)
	i.available = true
	i.broadcast()
}

//Registrar un observador
func (i *Item) register(observer Observer) {
	i.Observers = append(i.Observers, observer)
}

// Mandar el evento a todos los observadores
func (i *Item) broadcast() {
	for _, observer := range i.Observers {
		observer.updateValue(i.name)

	}
}

type EmailClient struct {
	id string
}

func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %savailable from client %s\n", value, eC.id)

}

func (e *EmailClient) getId() string {
	return e.id
}

func main() {
	nvidiaItem := NewItem("RTC 3080")
	firstObserver := &EmailClient{
		id: "34dc",
	}
	secondObserver := &EmailClient{
		id: "34dc",
	}
	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.UpdateAvailable()
}
