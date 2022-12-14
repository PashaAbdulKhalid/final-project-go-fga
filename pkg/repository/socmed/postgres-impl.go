package socmed

import (
	"context"
	"log"

	"github.com/PashaAbdulKhalid/final-project-go-fga/config/postgres"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/socmed"
	"gorm.io/gorm/clause"
)

type SocialMediaRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewSocialMediaRepo(pgCln postgres.PostgresClient) socmed.SocialMediaRepo {
	return &SocialMediaRepoImpl{pgCln: pgCln}
}

func (s *SocialMediaRepoImpl) CreateSocialMedia(ctx context.Context, inputSocialMedia *socmed.SocialMedia) (result socmed.SocialMedia, err error) {
	log.Printf("%T - CreateSocialMedia is invoked\n", s)
	defer log.Printf("%T - CreateSocialMedia executed\n", s)

	db := s.pgCln.GetClient()

	err = db.Model(&result).Create(&inputSocialMedia).Error

	if err != nil {
		log.Printf("error when creating socialMedia for photo id %v\n", inputSocialMedia)
	}

	result = *inputSocialMedia

	return result, err
}

func (s *SocialMediaRepoImpl) GetSocialMedias(ctx context.Context, userId uint64) (result []socmed.SocialMedia, err error) {
	log.Printf("%T - GetSocialMedias is invoked\n", s)
	defer log.Printf("%T - GetSocialMedias executed\n", s)

	db := s.pgCln.GetClient()

	err = db.Model(&socmed.SocialMedia{}).Where("user_id = ?", userId).Find(&result).Error

	if err != nil {
		log.Printf("error when getting social media by user id %v\n", userId)
	}

	return result, err
}

func (s *SocialMediaRepoImpl) GetSocialMediaById(ctx context.Context, socmedId uint64) (result socmed.SocialMedia, err error) {
	log.Printf("%T - GetSocialMediaById is invoked\n", s)
	defer log.Printf("%T - GetSocialMediaById executed\n", s)

	db := s.pgCln.GetClient()

	err = db.Table("social_medias").Where("id = ?", socmedId).Find(&result).Error

	if err != nil {
		log.Printf("error when getting social media by id %v\n", socmedId)
	}

	return result, err
}

func (s *SocialMediaRepoImpl) UpdateSocialMedia(ctx context.Context, inputSocialMedia socmed.SocialMedia) (result socmed.SocialMedia, err error) {
	log.Printf("%T - UpdateSocialMedia is invoked\n", s)
	defer log.Printf("%T - UpdateSocialMedia executed\n", s)

	socmedId := ctx.Value("socmedId").(uint64)

	db := s.pgCln.GetClient()

	err = db.Model(&result).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}, {Name: "name"}, {Name: "url"}, {Name: "user_id"}, {Name: "updated_at"}}}).Where("id = ?", socmedId).Updates(inputSocialMedia).Error

	if err != nil {
		log.Printf("error when updating socialMedia by id %v\n", socmedId)
	}

	return result, err
}

func (s *SocialMediaRepoImpl) DeleteSocialMedia(ctx context.Context, socmedId uint64) (err error) {
	log.Printf("%T - DeleteSocialMedia is invoked\n", s)
	defer log.Printf("%T - DeleteSocialMedia executed\n", s)

	db := s.pgCln.GetClient()

	err = db.Where("id = ?", socmedId).Delete(&socmed.SocialMedia{}).Error

	if err != nil {
		log.Printf("error when deleting social media by id %v \n", socmedId)
	}

	return err
}


