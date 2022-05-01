package port

import pbPay "github.com/isutare412/hexago/common/pkg/pb/payment"

type PaymentMessageQueue interface {
	SendDonationRequest(req *pbPay.DonationRequest) error
}
