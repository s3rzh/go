package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Cond может быть полезен при решении проблемы производитель-потребитель, классической проблемы синхронизации, включающей два типа процессов, производителей и потребителей, которые совместно используют общий буфер или очередь фиксированного размера.

const MaxMessageChannelSize = 5

type MessageChannel struct {
	maxBufferSize int
	buffer        []string
}

func NewMessageChannel(size int) *MessageChannel {
	return &MessageChannel{
		maxBufferSize: size,
		buffer:        make([]string, 0, size),
	}
}

func (mc *MessageChannel) IsEmpty() bool {
	return len(mc.buffer) == 0

}

func (mc *MessageChannel) IsFull() bool {
	return len(mc.buffer) == mc.maxBufferSize
}

func (mc *MessageChannel) Add(message string) {
	mc.buffer = append(mc.buffer, message)
}

func (mc *MessageChannel) Get() string {
	message := mc.buffer[0]
	mc.buffer = mc.buffer[1:]
	return message
}

type Producer struct {
	cond           *sync.Cond
	messageChannel *MessageChannel
}

func NewProducer(cond *sync.Cond, messageChannel *MessageChannel) *Producer {
	return &Producer{
		cond:           cond,
		messageChannel: messageChannel,
	}
}

func (p *Producer) Produce(message string) {
	time.Sleep(500 * time.Millisecond) // Simulating some work
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	for p.messageChannel.IsFull() {
		fmt.Println("Producer is waiting because the message channel is full")
		p.cond.Wait()
	}
	p.messageChannel.Add(message)
	fmt.Println("Producer produced the message:", message)
	p.cond.Signal()
}

type Consumer struct {
	id             int
	cond           *sync.Cond
	messageChannel *MessageChannel
}

func NewConsumer(cond *sync.Cond, messageChannel *MessageChannel) *Consumer {
	return &Consumer{
		cond:           cond,
		messageChannel: messageChannel,
	}
}

func (c *Consumer) Consume() {
	time.Sleep(1 * time.Second) // Simulating some work
	c.cond.L.Lock()
	defer c.cond.L.Unlock()
	for c.messageChannel.IsEmpty() {
		fmt.Println("Consumer is waiting because the message channel is empty")
		c.cond.Wait()
	}
	message := c.messageChannel.Get()
	fmt.Println("Consumer consumed the message:", message)
	c.cond.Signal()
}

func main() {
	var wg sync.WaitGroup

	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	messageChannel := NewMessageChannel(MaxMessageChannelSize)

	producer := NewProducer(cond, messageChannel)
	consumer := NewConsumer(cond, messageChannel)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := range 10 {
			producer.Produce(fmt.Sprintf("Message %d", i))
		}
	}()

	go func() {
		defer wg.Done()
		for range 10 {
			consumer.Consume()
		}
	}()
	wg.Wait()
}


// В этом примере у нас есть горутина-производитель, которая создает сообщения и добавляет их в канал сообщений, и горутина-потребитель, которая потребляет сообщения. Канал сообщений имеет максимальный размер, определенный MaxMessageChannelSize .

// Горутина-производитель добавляет сообщения в канал сообщений и использует cond.Signal() для уведомления горутины-потребителя о появлении новых данных. 
// Если канал сообщений заполнен, горутина производителя ожидает с помощью cond.Wait() , пока потребитель не израсходует некоторые данные и не освободит место в канале сообщений.

// Аналогично, горутина-потребитель потребляет сообщения из канала сообщений и использует cond.Signal() для уведомления горутины-производителя, когда в канале сообщений становится доступным место. 
// Если он пуст, горутина-потребитель ждет, используя cond.Wait() , пока производитель не создаст некоторые данные и не добавит их в канал сообщений.

// Здесь sync.Cond обеспечивает координацию и синхронизацию между горутинами производителя и потребителя. 
// Это гарантирует, что потребитель ждет, когда канал сообщений станет пустым, а производитель ждет, когда он заполнится, тем самым решая проблему производитель-потребитель.
