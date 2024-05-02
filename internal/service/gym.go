package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/entity"
)

var _ IGymService = &GymService{}  

type GymService struct {
	// Repos
	gymRepo IGymRepository
	// Services
	// RegisterNewGym(gym entity.Gym) error
	// ChangeGym(gym entity.Gym) error
	// DeleteGym(gymID uuid.UUID) error
	// GetGymByID(gymID uuid.UUID) (error, entity.Gym)
	// ListGyms() (error, []entity.Gym)
}

func NewGymService(gymRepo IGymRepository) IGymService{
	return &GymService{
		gymRepo: gymRepo,
	}
}

func (s *GymService) RegisterNewGym(ctx context.Context, gym entity.Gym) error {
	if !gym.Validate() {
		return ErrValidation
	}
	
	err := s.gymRepo.RegisterNewGym(ctx, gym)

	return err
}

func (s *GymService) ChangeGym(ctx context.Context, gym entity.Gym) error {
	if !gym.Validate() {
		return ErrValidation
	}

	err := s.gymRepo.ChangeGym(ctx, gym)

	return err
}

func (s *GymService) DeleteGym(ctx context.Context, gymID uuid.UUID) error {
	err := s.gymRepo.DeleteGym(ctx, gymID)

	return err
}

func (s *GymService) GetGymByID(ctx context.Context, gymID uuid.UUID) (entity.Gym, error) {
	gym, err := s.gymRepo.GetGymByID(ctx, gymID)

	return gym, err
}

func (s *GymService) ListGyms(ctx context.Context) ([]entity.Gym, error) {
	gym, err := s.gymRepo.ListGyms(ctx)

	return gym, err
}