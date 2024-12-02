package services

import (
	"github.com/Masterminds/squirrel"
	"log"
	"web4/internal/db"
	"web4/internal/models"
)

func AddAreaCheck(areaCheck models.AreaCheck) {
	log.Printf("Adding area check: %v\n", areaCheck)
	query := squirrel.Insert("requests").Values(areaCheck.X, areaCheck.Y, areaCheck.Radius, areaCheck.Result, areaCheck.SendTime, areaCheck.Author).PlaceholderFormat(squirrel.Dollar)

	_, err := query.RunWith(db.DB).Exec()
	if err != nil {
		log.Printf("Error adding area check: %v\n", err)
	}
}

func FindAllFor(user string) []models.AreaCheck {
	log.Printf("Getting area checks for %s\n", user)
	query := squirrel.Select("x", "y", "radius", "result", "send_time", "author").From("requests").Where(squirrel.Eq{"author": user}).PlaceholderFormat(squirrel.Dollar)

	rows, err := query.RunWith(db.DB).Query()
	if err != nil {
		log.Printf("Error getting area checks: %v\n", err)
		return nil
	}

	defer rows.Close()

	var areaChecks []models.AreaCheck
	for rows.Next() {
		var ac models.AreaCheck
		err := rows.Scan(&ac.X, &ac.Y, &ac.Radius, &ac.Result, &ac.SendTime, &ac.Author)
		if err != nil {
			log.Printf("Error scanning area check: %v\n", err)
			return nil
		}
		areaChecks = append(areaChecks, ac)
	}

	return areaChecks
}
