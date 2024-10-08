package service

import (
	"conducting-tenders/internal/entity"
	authorType "conducting-tenders/internal/entity/author-type"
	"conducting-tenders/internal/entity/statusBid"
	"conducting-tenders/internal/entity/statusTender"
	"conducting-tenders/internal/repo"
	"conducting-tenders/internal/repo/repoerrs"
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type BidService struct {
	bidRepo                     repo.Bid
	employeeRepo                repo.Employee
	tenderRepo                  repo.Tender
	organizationRepo            repo.Organization
	organizationResponsibleRepo repo.OrganizationResponsible
}

func NewBidService(bidRepo repo.Bid, employeeRepo repo.Employee, tenderRepo repo.Tender, organizationRepo repo.Organization, organizationResponsibleRepo repo.OrganizationResponsible) *BidService {
	return &BidService{
		bidRepo:                     bidRepo,
		employeeRepo:                employeeRepo,
		tenderRepo:                  tenderRepo,
		organizationRepo:            organizationRepo,
		organizationResponsibleRepo: organizationResponsibleRepo,
	}
}

func (b *BidService) CreateBid(ctx context.Context, input CreateBidInput) (entity.Bid, error) {
	var bid entity.Bid

	bid.Id = uuid.New()
	bid.Name = input.Name
	bid.Description = input.Description
	bid.Status = statusBid.StatusCreated

	_, err := b.tenderRepo.GetTenderById(ctx, input.TenderId)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrTenderNotFind
		}
		return entity.Bid{}, err
	}

	bid.TenderId = input.TenderId

	bid.AuthorType = input.AuthorType
	if input.AuthorType == authorType.AuthorUser {
		_, err := b.employeeRepo.GetEmployeeById(ctx, input.AuthorId)
		if err != nil {
			if errors.Is(err, repoerrs.ErrNotFound) {
				return entity.Bid{}, ErrEmployeeNotFind
			}
			return entity.Bid{}, err
		}
	}

	if input.AuthorType == authorType.AuthorOrganization {
		_, err := b.organizationRepo.GetOrganizationById(ctx, input.AuthorId)
		if err != nil {
			if errors.Is(err, repoerrs.ErrNotFound) {
				return entity.Bid{}, ErrOrganizationNotFind
			}
			return entity.Bid{}, err
		}
	}

	bid.AuthorId = input.AuthorId
	bid.Version = 1
	bid.CreatedAt = time.Now()
	bid.Tag = uuid.New()

	_, err = b.bidRepo.CreateBid(ctx, bid)

	if err != nil {
		return entity.Bid{}, err
	}
	log.Println(";lsjkaf")
	return bid, nil
}

func (b *BidService) GetBidsByUsername(ctx context.Context, input GetBidsByUsernameInput) ([]entity.Bid, error) {
	employee, err := b.employeeRepo.GetEmployeeByUsername(ctx, input.Username)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return []entity.Bid{}, ErrEmployeeNotFind
		}
		return []entity.Bid{}, err
	}

	if input.Limit == 0 {
		input.Limit = 5
	}

	bids, err := b.bidRepo.GetBidsByAuthorId(ctx, employee.Id, input.Limit, input.Offset)

	if err != nil {
		return []entity.Bid{}, err
	}

	return bids, nil
}

func (b *BidService) GetBidsByTenderId(ctx context.Context, input GetBidsByTenderIdInput) ([]entity.Bid, error) {
	employee, err := b.employeeRepo.GetEmployeeByUsername(ctx, input.Username)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return []entity.Bid{}, ErrEmployeeNotFind
		}
		return []entity.Bid{}, err
	}

	organizationEmpId, err := b.organizationResponsibleRepo.GetOrganizationIdByEmployeeId(ctx, employee.Id)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return []entity.Bid{}, ErrEmployeeNotFind
		}
		return []entity.Bid{}, err
	}

	if input.Limit == 0 {
		input.Limit = 5
	}

	bids, err := b.bidRepo.GetBidsByTenderIdAndUserIdAndOrganizationId(ctx, input.TenderId, organizationEmpId, employee.Id, input.Limit, input.Offset)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return []entity.Bid{}, ErrEmployeeNotFind
		}
		return []entity.Bid{}, err
	}

	return bids, nil
}

func (b *BidService) GetBidStatusById(ctx context.Context, input GetBidStatusByIdInput) (statusBid.Status, error) {
	bid, err := b.bidRepo.GetBidById(ctx, input.BidId)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return "", ErrBidNotFind
		}
		return "", err
	}

	employee, err := b.employeeRepo.GetEmployeeByUsername(ctx, input.Username)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return "", ErrEmployeeNotFind
		}
		return "", err
	}

	orgaEmpId, err := b.organizationResponsibleRepo.GetOrganizationIdByEmployeeId(ctx, employee.Id)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return "", ErrOrganizationNotFind
		}
		return "", err
	}

	if bid.AuthorId != employee.Id && bid.AuthorId != orgaEmpId {
		return "", ErrNotEnoughRights
	}

	return bid.Status, nil
}

func (b *BidService) UpdateBidStatusById(ctx context.Context, input UpdateBidStatusByIdInput) (entity.Bid, error) {
	bid, err := b.bidRepo.GetBidById(ctx, input.BidId)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrBidNotFind
		}
		return entity.Bid{}, err
	}

	employee, err := b.employeeRepo.GetEmployeeByUsername(ctx, input.Username)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrEmployeeNotFind
		}
		return entity.Bid{}, err
	}

	orgaEmpId, err := b.organizationResponsibleRepo.GetOrganizationIdByEmployeeId(ctx, employee.Id)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrOrganizationNotFind
		}
		return entity.Bid{}, err
	}

	if bid.AuthorId != employee.Id && orgaEmpId != bid.AuthorId {
		return entity.Bid{}, ErrNotEnoughRights
	}

	bid.Status = input.Status
	maxVersion, err := b.bidRepo.GetBidVersionMaxByTag(ctx, bid.Tag)
	if err != nil {
		return entity.Bid{}, err
	}

	bid.Version = maxVersion + 1
	bid.Id = uuid.New()
	bidId, err := b.bidRepo.CreateBid(ctx, bid)

	if err != nil {
		return entity.Bid{}, err
	}

	bid.Id = bidId

	return bid, nil
}

func (b *BidService) EditBidById(ctx context.Context, input EditBidByIdInput) (entity.Bid, error) {
	bid, err := b.bidRepo.GetBidById(ctx, input.BidId)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrBidNotFind
		}
		return entity.Bid{}, err
	}

	employee, err := b.employeeRepo.GetEmployeeByUsername(ctx, input.Username)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrEmployeeNotFind
		}
		return entity.Bid{}, err
	}

	orgaEmpId, err := b.organizationResponsibleRepo.GetOrganizationIdByEmployeeId(ctx, employee.Id)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrOrganizationNotFind
		}
		return entity.Bid{}, err
	}

	if bid.AuthorId != employee.Id && orgaEmpId != bid.AuthorId {
		return entity.Bid{}, ErrNotEnoughRights
	}

	if input.Name == "" && input.Description == "" {
		return entity.Bid{}, ErrInvalidRequestFormatOrParameters
	}

	if input.Name != "" {
		bid.Name = input.Name
	}

	if input.Description != "" {
		bid.Description = input.Description
	}

	maxVersion, err := b.bidRepo.GetBidVersionMaxByTag(ctx, bid.Tag)
	if err != nil {
		return entity.Bid{}, err
	}

	bid.Version = maxVersion + 1
	bid.Id = uuid.New()
	bidId, err := b.bidRepo.CreateBid(ctx, bid)

	if err != nil {
		return entity.Bid{}, err
	}

	bid.Id = bidId

	return bid, nil
}

func (b *BidService) UpdateVersionBid(ctx context.Context, input UpdateVersionBidInput) (entity.Bid, error) {
	bid, err := b.bidRepo.GetBidById(ctx, input.BidId)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrBidNotFind
		}
		return entity.Bid{}, err
	}

	employee, err := b.employeeRepo.GetEmployeeByUsername(ctx, input.Username)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrEmployeeNotFind
		}
		return entity.Bid{}, err
	}

	orgaEmpId, err := b.organizationResponsibleRepo.GetOrganizationIdByEmployeeId(ctx, employee.Id)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrOrganizationNotFind
		}
		return entity.Bid{}, err
	}

	if bid.AuthorId != employee.Id && orgaEmpId != bid.AuthorId {
		return entity.Bid{}, ErrNotEnoughRights
	}

	bidVer, err := b.bidRepo.GetBidByTagAndVersion(ctx, bid.Tag, input.Version)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrBidNotFind
		}
		return entity.Bid{}, err
	}

	maxVersion, err := b.bidRepo.GetBidVersionMaxByTag(ctx, bid.Tag)
	if err != nil {
		return entity.Bid{}, err
	}
	bidVer.Version = maxVersion + 1
	bidVer.Id = uuid.New()

	bidVerId, err := b.bidRepo.CreateBid(ctx, bidVer)

	if err != nil {
		return entity.Bid{}, err
	}

	bidVer.Id = bidVerId

	return bidVer, nil
}

func (b *BidService) UpdateSubmitDecision(ctx context.Context, input UpdateSubmitDecisionInput) (entity.Bid, error) {
	bid, err := b.bidRepo.GetBidById(ctx, input.BidId)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrBidNotFind
		}
		return entity.Bid{}, err
	}

	employee, err := b.employeeRepo.GetEmployeeByUsername(ctx, input.Username)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrEmployeeNotFind
		}
		return entity.Bid{}, err
	}

	orgaEmpId, err := b.organizationResponsibleRepo.GetOrganizationIdByEmployeeId(ctx, employee.Id)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrOrganizationNotFind
		}
		return entity.Bid{}, err
	}

	tender, err := b.tenderRepo.GetTenderById(ctx, bid.TenderId)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return entity.Bid{}, ErrTenderNotFind
		}
		return entity.Bid{}, err
	}

	if tender.OrganizationId != orgaEmpId {
		return entity.Bid{}, ErrNotEnoughRights
	}

	if input.Decision == "Rejected" {
		bid.Status = statusBid.StatusCanceled
		bid.Id = uuid.New()
		maxVerl, err := b.bidRepo.GetBidVersionMaxByTag(ctx, bid.Tag)
		bid.Version = maxVerl + 1
		if err != nil {
			return entity.Bid{}, err
		}

		_, err = b.bidRepo.CreateBid(ctx, bid)
		if err != nil {
			return entity.Bid{}, err
		}

		return bid, nil
	} else {
		tender.Id = uuid.New()
		tender.Status = statusTender.StatusClosed
		maxVer, err := b.tenderRepo.GetTenderVersionMaxByTag(ctx, tender.Tag)
		if err != nil {
			return entity.Bid{}, err
		}

		tender.Version = maxVer + 1
		_, err = b.tenderRepo.CreateTender(ctx, tender)

		bid.TenderId = tender.Id
		bid.Id = uuid.New()
		bid.Status = statusBid.StatusPublished

		maxVerl, err := b.bidRepo.GetBidVersionMaxByTag(ctx, bid.Tag)
		bid.Version = maxVerl + 1
		if err != nil {
			return entity.Bid{}, err
		}

		_, err = b.bidRepo.CreateBid(ctx, bid)
		if err != nil {
			return entity.Bid{}, err
		}

		return bid, nil
	}
}
