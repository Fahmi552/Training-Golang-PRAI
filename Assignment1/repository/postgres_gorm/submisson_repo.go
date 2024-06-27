package postgresgorm

import (
	"Training/Assignment1/entity"
	"Training/Assignment1/service"
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type submissionRepository struct {
	db GormDBIface
}

// NewSubmissionRepository membuat instance baru dari submissionRepository
func NewSubmissionRepository(db GormDBIface) service.ISubmissionRepository {
	return &submissionRepository{db: db}
}

func (s *submissionRepository) CreateSubmission_repo(ctx context.Context, submission *entity.Submission) (entity.Submission, error) {

	if err := s.db.WithContext(ctx).Create(submission).Error; err != nil {
		log.Printf("Error saat create submission: %v\n", err)
		return entity.Submission{}, err
	}
	return *submission, nil
}

func (s *submissionRepository) GetSubmissionByID_repo(ctx context.Context, id int) (entity.Submission, error) {
	// Validasi ID
	if id <= 0 {
		return entity.Submission{}, fmt.Errorf("ID submission tidak valid: %d", id)
	}

	// Cari submission berdasarkan ID
	var submission entity.Submission
	if result := s.db.WithContext(ctx).First(&submission, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Jika tidak ditemukan, kembalikan error spesifik
			return entity.Submission{}, fmt.Errorf("submission dengan ID %d tidak ditemukan", id)
		}
		// Kembalikan error lainnya
		return entity.Submission{}, result.Error
	}

	// Kembalikan submission yang ditemukan
	return submission, nil
}

func (s *submissionRepository) GetSubmissionByUserID_repo(ctx context.Context, userID int) (entity.Submission, error) {
	var submissions entity.Submission

	// Cari semua submission berdasarkan UserID
	if err := s.db.WithContext(ctx).Where("user_id = ?", userID).Find(&submissions).Error; err != nil {
		// Jika ada error saat pencarian, log error dan kembalikan error tersebut
		log.Printf("Error saat mencari submissions dengan UserID %d: %v\n", userID, err)
		//return nil, err
	}

	return submissions, nil
}

func (s *submissionRepository) GetAllSubmissions_repo(ctx context.Context) ([]entity.Submission, error) {
	var submissions []entity.Submission

	if err := s.db.WithContext(ctx).Find(&submissions).Error; err != nil {
		// Jika ada error saat mengambil data, log error dan kembalikan error tersebut
		log.Printf("Error saat mengambil semua submissions: %v\n", err)
		return nil, err
	}

	return submissions, nil
}

func (s *submissionRepository) DeleteSubmission_repo(ctx context.Context, id int) error {
	// Pertama, cek apakah submission dengan ID tersebut ada
	var submission entity.Submission
	if result := s.db.First(&submission, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Jika tidak ditemukan, kembalikan error spesifik
			return fmt.Errorf("submission dengan ID %d tidak ditemukan", id)
		}
		return result.Error
	}

	// Jika ditemukan, lanjutkan untuk menghapus submission tersebut
	if err := s.db.WithContext(ctx).Delete(&entity.Submission{}, id).Error; err != nil {
		log.Printf("Error deleting submission: %v\n", err)
		return err
	}
	return nil
}
