package service

import (
	"Assignment4/entity"
	"Assignment4/repository"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

const redisUrlShortKey = "url:%v"

type URLService interface {
	ShortenURL(ctx context.Context, originalURL string) (string, error)
	GetOriginalURL(shortURL string) (string, error)
}

type urlService struct {
	repo repository.URLRepository
	rdb  *redis.Client
}

func NewURLService(repo_ repository.URLRepository, rdb_ *redis.Client) URLService {
	return &urlService{repo: repo_, rdb: rdb_}
}

func generateShortID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (s *urlService) ShortenURLOLD(originalURL string) (string, error) {
	shortID, err := shortid.Generate()
	if err != nil {
		return "", err
	}

	url := &entity.URL{OriginalURL: originalURL, ShortURL: shortID}
	if err := s.repo.CreateURL(url); err != nil {
		return "", err
	}
	return shortID, nil
}

func (s *urlService) ShortenURL(ctx context.Context, originalURL string) (string, error) {
	const shortIDLength = 8
	var shortID string
	var err error

	// Coba menghasilkan unique short ID beberapa kali
	for attempts := 0; attempts < 10; attempts++ {
		shortID = generateShortID(shortIDLength)
		_, err = s.repo.GetURLByShort(shortID)
		if err != nil {
			// Jika tidak ditemukan, lanjutkan untuk menyimpan
			if errors.Is(err, gorm.ErrRecordNotFound) {
				break
			}
			// Jika ada error lain, return error
			return "", err
		}
	}

	if err == nil {
		// Jika tidak berhasil mendapatkan unique short ID setelah beberapa kali percobaan
		return "", errors.New("could not generate a unique short ID")
	}

	url := &entity.URL{OriginalURL: originalURL, ShortURL: shortID}
	if err := s.repo.CreateURL(url); err != nil {
		return "", err
	}

	// bagian insert ke redis
	rediskey := fmt.Sprintf(redisUrlShortKey, shortID)
	createdUrlJSON, err := json.Marshal(url)
	if err != nil {
		log.Println("gagal marshal json")
	}
	if err := s.rdb.Set(ctx, rediskey, createdUrlJSON, 0).Err(); err != nil {
		log.Println("error when set redis key", rediskey)
	}

	return shortID, nil
}

func (s *urlService) GetOriginalURL(shortURL string) (string, error) {
	url, err := s.repo.GetURLByShort(shortURL)
	if err != nil {
		return "", err
	}
	return url.OriginalURL, nil
}
