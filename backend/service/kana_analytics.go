package service

import (
	"github.com/patt812/golang-nuxt-typing-analytics/domain"
	"gorm.io/gorm"
)

type KanaAnalyticsService struct {
	db *gorm.DB
}

func NewKanaAnalyticsService(db *gorm.DB) *KanaAnalyticsService {
	return &KanaAnalyticsService{db: db}
}

func (s *KanaAnalyticsService) KanaToPatterns(kanaString string) ([][]string, error) {
	allPatterns := [][]string{{}}

	// Split the input hiragana string into individual hiragana characters
	for _, kana := range kanaString {
		// Retrieve the corresponding romaji patterns for the hiragana character
		var kanaObj domain.Kana
		err := s.db.Where("kana = ?", string(kana)).First(&kanaObj).Error
		if err != nil {
			return nil, err
		}
		var patterns []domain.Pattern
		s.db.Where("kana_id = ?", kanaObj.ID).Find(&patterns)

		// Generate results for all romaji patterns
		var newAllPatterns [][]string
		for _, existingPattern := range allPatterns {
			for _, pattern := range patterns {
				newPattern := make([]string, len(existingPattern), len(existingPattern)+1)
				copy(newPattern, existingPattern)
				newPattern = append(newPattern, pattern.Roma)
				newAllPatterns = append(newAllPatterns, newPattern)
			}
		}
		allPatterns = newAllPatterns
	}

	return allPatterns, nil
}
