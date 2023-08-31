package repository

type IDonationRepository interface {
}

type DonationRepository struct {
	opt Option
}

func NewDonationRepository(opt Option) IDonationRepository {
	return &DonationRepository{
		opt: opt,
	}
}