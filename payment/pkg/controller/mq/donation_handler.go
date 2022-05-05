package mq

import (
	"context"
	"time"

	"github.com/Shopify/sarama"
	pbPay "github.com/isutare412/hexago/common/pkg/pb/payment"
	"github.com/isutare412/hexago/payment/pkg/config"
	"github.com/isutare412/hexago/payment/pkg/logger"
	"github.com/isutare412/hexago/payment/pkg/port"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type DonationHandler struct {
	dnxSvc  port.DonationService
	timeout time.Duration
}

func NewDonationHandler(
	cfg *config.KafkaConsumerConfig,
	dnxSvc port.DonationService,
) *DonationHandler {
	return &DonationHandler{
		dnxSvc:  dnxSvc,
		timeout: time.Duration(cfg.Timeout) * time.Second,
	}
}

func (h *DonationHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *DonationHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *DonationHandler) ConsumeClaim(
	sess sarama.ConsumerGroupSession,
	clm sarama.ConsumerGroupClaim,
) error {
	handleAndMark := func(msg *sarama.ConsumerMessage) {
		ctx, cancel := context.WithTimeout(sess.Context(), h.timeout)
		defer cancel()

		h.handleMessage(ctx, msg)
		sess.MarkMessage(msg, "")
	}

	for msg := range clm.Messages() {
		handleAndMark(msg)
	}
	return nil
}

func (h *DonationHandler) handleMessage(
	ctx context.Context,
	msg *sarama.ConsumerMessage,
) {
	var req pbPay.DonationRequest
	if err := proto.Unmarshal(msg.Value, &req); err != nil {
		logger.S().Error(err)
	}

	err := h.dnxSvc.Donate(ctx, req.DonatorId, req.DonateeId, req.Cents)
	if err != nil {
		logger.S().With(
			"donatorId", req.DonatorId,
			"donateeId", req.DonateeId,
			"cents", req.Cents,
		).Error(err)
	}

	logger.A().With(
		zap.String("donatorId", req.DonatorId),
		zap.String("donateeId", req.DonateeId),
		zap.Int64("cents", req.Cents),
	).Info("Handled donation request")
}
