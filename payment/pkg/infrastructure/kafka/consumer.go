package kafka

import (
	"context"
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/isutare412/hexago/payment/pkg/config"
)

type Consumer struct {
	cli     sarama.ConsumerGroup
	handler sarama.ConsumerGroupHandler
	wg      *sync.WaitGroup
	groupId string
	topics  []string
}

func NewConsumer(
	ctx context.Context,
	kCfg *config.KafkaConfig,
	csmCfg *config.KafkaConsumerConfig,
	handler sarama.ConsumerGroupHandler,
) (*Consumer, error) {
	sCfg := newSaslPlainConfig(csmCfg.Username, csmCfg.Password)
	sCfg.ClientID = kafkaClientId
	sCfg.Consumer.Offsets.Initial = sarama.OffsetNewest

	done := make(chan *Consumer)
	fail := make(chan error)
	defer close(done)
	defer close(fail)
	go func() {
		client, err := sarama.NewConsumerGroup(kCfg.Addrs, csmCfg.Group, sCfg)
		if err != nil {
			fail <- fmt.Errorf("creating kafka consumer group[%s]: %w",
				csmCfg.Group, err)
			return
		}

		done <- &Consumer{
			cli:     client,
			handler: handler,
			wg:      &sync.WaitGroup{},
			groupId: csmCfg.Group,
			topics:  []string{csmCfg.Topic},
		}
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("new kafka consumer timeout")
	case err := <-fail:
		return nil, err
	case c := <-done:
		return c, nil
	}
}

func (c *Consumer) Run(ctx context.Context) <-chan error {
	fails := make(chan error, 1)
	c.wg.Add(1)
	go func() {
		defer close(fails)
		defer c.wg.Done()

		for {
			if err := c.cli.Consume(ctx, c.topics, c.handler); err != nil {
				fails <- fmt.Errorf(
					"consuming from topics[%v] with group[%s]: %w",
					c.topics, c.groupId, err,
				)
				return
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()
	return fails
}

func (c *Consumer) Shutdown(ctx context.Context) error {
	closeDone := make(chan struct{})
	closeFail := make(chan error)
	defer close(closeFail)
	go func() {
		defer close(closeDone)

		c.wg.Wait()

		if err := c.cli.Close(); err != nil {
			closeFail <- fmt.Errorf("closing consumer group[%s]: %w",
				c.groupId, err)
			return
		}
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("consumer group[%s] shutdown time out", c.groupId)
	case err := <-closeFail:
		return fmt.Errorf("consumer group[%s] shutdown failure: %w",
			c.groupId, err)
	case <-closeDone:
	}
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
