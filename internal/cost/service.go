package cost

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gofrs/uuid"
	"time"
)

type Service interface {
	CheckParticipantAmount(amount int, participants []entity.GroupCostItemParticipantRequestArg) bool
	CreateGroupCostAmdParticipants(requestArg *entity.GroupCostItemRequestArg, group *entity.Group, auth *entity.User) error
}

type service struct {
	repo   Repository
	logger *log.Logger
}

func NewService(repo Repository, logger *log.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) CheckParticipantAmount(amount int, participants []entity.GroupCostItemParticipantRequestArg) bool {
	var countParticipantTotal int
	for _, participant := range participants {
		countParticipantTotal = countParticipantTotal + participant.Amount
	}
	return amount == countParticipantTotal
}

func (s *service) CreateGroupCostAmdParticipants(requestArg *entity.GroupCostItemRequestArg, group *entity.Group, auth *entity.User) error {
	currentTime := time.Now()
	u1, err := uuid.NewV1()
	if err != nil {
		return err
	}

	groupCostItem := &entity.GroupCostItem{
		GroupId:     group.Id,
		UUID:        u1.String(),
		Name:        requestArg.Name,
		TotalAmount: requestArg.TotalAmount,
		PayerId:     requestArg.PayerId,
		CreatorId:   auth.Id,
		PayAt:       requestArg.PayAt,
		Remark:      requestArg.Remark,
	}

	for _, participantArg := range requestArg.Participants {
		groupCostItem.Participants = append(groupCostItem.Participants, entity.GroupCostItemParticipant{
			UserId: participantArg.UserId,
			Amount: participantArg.Amount,
		})
	}
	if groupCostItem.PayAt.IsZero() {
		groupCostItem.PayAt = currentTime
	}
	return s.repo.CreateCostItemByUser(groupCostItem)
}
