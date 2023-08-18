package migrations

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/patt812/golang-nuxt-typing-analytics/domain"
	"gorm.io/gorm"
)

func InsertKanaPattern() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20230818_01",
		Migrate: func(tx *gorm.DB) error {
			tx = tx.Begin()
			if tx.Error != nil {
				return tx.Error
			}

			file, err := os.Open("app/resources/patterns.json")
			if err != nil {
				return err
			}

			patterns := map[string][]string{}
			decoder := json.NewDecoder(file)
			// check json format
			if err := decoder.Decode(&patterns); err != nil {
				return err
			}

			// duplicate check
			seen := make(map[string]bool)
			for kana := range patterns {
				if seen[kana] {
					return fmt.Errorf("duplicate kana: %s", kana)
				}
				seen[kana] = true
			}

			// insert kana and pattern
			for kana, patternList := range patterns {
				kanaRecord := &domain.Kana{Kana: kana}
				if err := tx.Create(kanaRecord).Error; err != nil {
					tx.Rollback()
					return err
				}

				for _, pattern := range patternList {
					patternRecord := &domain.Pattern{Roma: pattern, KanaID: kanaRecord.ID}
					if err := tx.Create(patternRecord).Error; err != nil {
						tx.Rollback()
						return err
					}
				}
			}

			if err := tx.Commit().Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}
}
