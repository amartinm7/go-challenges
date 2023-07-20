package ad

type AdProducer interface {
	SendMessage(message BrokerMessage) (partition int32, offset int64, err error)
}

type BrokerMessage struct {
	Ad Ad
}

func NewBrokerMessage(ad Ad) BrokerMessage {
	return BrokerMessage{
		Ad: ad,
	}
}
