package service

import (
	"Training/Assignment1/entity"
	"context"
	"encoding/json"
	"fmt"
)

// ISubmissionService mendefinisikan interface untuk layanan submission
type ISubmissionService interface {
	CreateSubmission_service(ctx context.Context, submission *entity.SubmissionRequest) error
	GetSubmissionByID_service(ctx context.Context, id int) (entity.Submission, error)
	GetSubmissionByUserID_service(ctx context.Context, id int) (entity.Submission, error)
	DeleteSubmission_service(ctx context.Context, id int) error
	GetAllSubmissions_service(ctx context.Context) ([]entity.Submission, error)
}

// ISubmissionRepository mendefinisikan interface untuk repository submission
type ISubmissionRepository interface {
	CreateSubmission_repo(ctx context.Context, submission *entity.Submission) (entity.Submission, error)
	GetSubmissionByID_repo(ctx context.Context, id int) (entity.Submission, error)
	GetSubmissionByUserID_repo(ctx context.Context, id int) (entity.Submission, error)
	DeleteSubmission_repo(ctx context.Context, id int) error
	GetAllSubmissions_repo(ctx context.Context) ([]entity.Submission, error)
}

// submissionService adalah implementasi dari ISubmissionService yang menggunakan ISubmissionRepository
type submissionService struct {
	submissionRepo ISubmissionRepository
}

// NewSubmissionService membuat instance baru dari submissionService
func NewSubmissionService(submissionRepo ISubmissionRepository) ISubmissionService {
	return &submissionService{submissionRepo: submissionRepo}
}

func (s *submissionService) CreateSubmission_service(ctx context.Context, submission *entity.SubmissionRequest) error {
	// Validasi input
	if submission.UserID <= 0 || len(submission.Answers) == 0 {
		return fmt.Errorf("validasi gagal: UserID dan Answers harus ada")
	}

	// Hitung skor risiko dan kategori berdasarkan jawaban
	score, category, definition := calculateProfileRiskFromAnswers(submission.Answers)

	// Cek apakah Answers adalah JSON
	jsonData, err := json.Marshal(submission.Answers)
	if err != nil {
		print("AAA")
		return fmt.Errorf("Answers bukan dalam format JSON: %v", err)
	}

	// Set nilai yang dihitung ke dalam submission
	submission.RiskScore = score
	submission.RiskCategory = category
	submission.RiskDefinition = definition
	body := &entity.Submission{
		UserID:         submission.UserID,
		RiskScore:      submission.RiskScore,
		RiskCategory:   submission.RiskCategory,
		RiskDefinition: submission.RiskDefinition,
		Answers:        jsonData,
	}
	// println(submission)

	// Simpan submission ke repository
	_, err = s.submissionRepo.CreateSubmission_repo(ctx, body)
	if err != nil {
		return fmt.Errorf("gagal menyimpan submission euy: %v", err)
	}

	return nil

}

func (s *submissionService) GetSubmissionByID_service(ctx context.Context, id int) (entity.Submission, error) {
	// Panggil repository untuk mendapatkan submission berdasarkan ID
	submission, err := s.submissionRepo.GetSubmissionByID_repo(ctx, id)
	if err != nil {
		if err.Error() == fmt.Sprintf("submission dengan ID %d tidak ditemukan", id) {
			return entity.Submission{}, fmt.Errorf("submission dengan ID %d tidak ditemukan", id)
		}
		return entity.Submission{}, fmt.Errorf("gagal mendapatkan submission untuk ID %d: %v", id, err)
	}

	return submission, nil
}

func (s *submissionService) GetSubmissionByUserID_service(ctx context.Context, userID int) (entity.Submission, error) {
	return s.submissionRepo.GetSubmissionByUserID_repo(ctx, userID)
}

func (s *submissionService) DeleteSubmission_service(ctx context.Context, id int) error {
	err := s.submissionRepo.DeleteSubmission_repo(ctx, id)
	if err != nil {
		// Tangani error spesifik pengguna tidak ditemukan
		if err.Error() == fmt.Sprintf("submission dengan ID %d tidak ditemukan", id) {
			return fmt.Errorf("tidak ada submission dengan ID %d yang ditemukan", id)
		}
		return fmt.Errorf("gagal menghapus submission dengan ID %d: %v", id, err)
	}
	return nil
}

func (s *submissionService) GetAllSubmissions_service(ctx context.Context) ([]entity.Submission, error) {
	return s.submissionRepo.GetAllSubmissions_repo(ctx)
}

// TODO: implement logic for profile risk calculation based on score mapping from entity.RiskMapping
// calculateProfileRiskFromAnswers will be used on submission creation
func calculateProfileRiskFromAnswers(answers []entity.Answer) (score int, category entity.ProfileRiskCategory, definition string) {
	// TODO: calculate total score from answers
	// TODO: get category and definition based on total score
	// Hitung total skor dari jawaban
	for _, answer := range answers {
		for _, question := range entity.Questions {
			if question.ID == answer.QuestionID {
				for _, option := range question.Options {
					if option.Answer == answer.Answer {
						score += option.Weight
					}
				}
			}
		}
	}

	// Tentukan kategori risiko berdasarkan total skor
	for _, riskProfile := range entity.RiskMapping {
		if score >= riskProfile.MinScore && score <= riskProfile.MaxScore {
			category = riskProfile.Category
			definition = riskProfile.Definition
			return score, category, definition
		}
	}

	return score, category, definition
}
