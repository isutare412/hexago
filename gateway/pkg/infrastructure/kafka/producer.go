package kafka

import (
	"context"
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
	pbPay "github.com/isutare412/hexago/common/pkg/pb/payment"
	"github.com/isutare412/hexago/gateway/pkg/config"
	"github.com/isutare412/hexago/gateway/pkg/logger"
	"google.golang.org/protobuf/proto"
)

type Producer struct {
	cli   sarama.AsyncProducer
	topic string
}

func NewProducer(
	ctx context.Context,
	kCfg *config.KafkaConfig,
	pCfg *config.KafkaProducerConfig,
) (*Producer, error) {
	sCfg := newSaslPlainConfig(pCfg.Username, pCfg.Password)
	sCfg.ClientID = kafkaClientId
	sCfg.Producer.RequiredAcks = sarama.WaitForLocal
	sCfg.Producer.Retry.Max = pCfg.MaxRetry
	sCfg.Producer.Return.Successes = false
	sCfg.Producer.Return.Errors = true

	done := make(chan *Producer)
	fail := make(chan error)
	go func() {
		defer close(done)
		defer close(fail)

		client, err := sarama.NewAsyncProducer(kCfg.Addrs, sCfg)
		if err != nil {
			fail <- fmt.Errorf("creating kafka async producer: %w", err)
			return
		}

		done <- &Producer{
			cli:   client,
			topic: pCfg.Topic,
		}
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("new kafka producer timeout")
	case err := <-fail:
		return nil, err
	case p := <-done:
		return p, nil
	}
}

func (p *Producer) Run(ctx context.Context) <-chan error {
	fails := make(chan error, 1)
	once := sync.Once{}
	go func() {
		defer close(fails)

		for prdErr := range p.cli.Errors() {
			err := fmt.Errorf("producing message: %w", prdErr.Err)
			logger.S().Error(err)
			once.Do(func() {
				fails <- err
			})
		}
	}()
	return fails
}

func (p *Producer) Shutdown(ctx context.Context) error {
	closeDone := make(chan struct{})
	closeFail := make(chan error)
	defer close(closeFail)
	go func() {
		defer close(closeDone)

		if err := p.cli.Close(); err != nil {
			closeFail <- fmt.Errorf("closing producer: %w", err)
			return
		}
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("producer shutdown time out")
	case err := <-closeFail:
		return fmt.Errorf("producer shutdown failure: %w", err)
	case <-closeDone:
	}
	return nil
}

func (p *Producer) SendDonationRequest(req *pbPay.DonationRequest) error {
	reqBytes, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshaling donation request: %w", err)
	}

	msg := sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(reqBytes),
	}
	p.cli.Input() <- &msg
	return nil
}

func newSaslPlainConfig(user, password string) *sarama.Config {
	cfg := sarama.NewConfig()
	cfg.Net.SASL.Enable = true
	cfg.Net.SASL.Handshake = true
	cfg.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	cfg.Net.SASL.User = user
	cfg.Net.SASL.Password = password
	return cfg
}
