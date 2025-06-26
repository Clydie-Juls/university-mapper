package main

import (
	"database/sql"
	"fmt"
	"log"
	"scraper/utils"
	"strings"

	_ "github.com/lib/pq"
)

func createDBConnection() (*sql.DB, error) {
	connStr := "postgres://postgres:12345678@localhost:5432/universitydb?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %s\n", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Unable to ping to database: %s\n", err)

	}

	return db, nil
}

func storeUniversityData(db *sql.DB, universities []University) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS universities (
    name                     VARCHAR(255) PRIMARY KEY,
    overall_rank             INT,
    country                  VARCHAR(255),
    arts_and_humanities_rank INT,
    business_and_economics_rank INT,
    clinical_and_health_rank INT,
    computer_science_rank    INT,
    education_rank           INT,
    engineering_rank         INT,
    law_rank                 INT,
    life_sciences_rank       INT,
    psychology_rank          INT,
    physical_sciences_rank   INT,
    social_sciences_rank     INT,
    latitude                 VARCHAR(50),
    longitude                VARCHAR(50)
);
    `)

	if err != nil {
		return fmt.Errorf("Unable to create university table: %s\n", err)
	}

	var values []string
	for _, uni := range universities {
		name := strings.ReplaceAll(uni.Name, "'", "''")
		latitude := strings.ReplaceAll(uni.Latitude, "'", "''")
		longitude := strings.ReplaceAll(uni.Longitude, "'", "''")
    country := strings.ReplaceAll(uni.Country, "'", "''")

    if uni.OverallRank == 1 {
      fmt.Println(uni)
    }

    insertValues := fmt.Sprintf(
			"('%s', %d, '%s', %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, '%s', '%s')",
			name, uni.OverallRank, country, uni.ArtsAndHumanitiesRank, uni.BusinessAndEconomicsRank,
			uni.ClinicalAndHealth, uni.ComputerScienceRank, uni.EducationRank, uni.EngineeringRank,
			uni.LawRank, uni.LifeSciencesRank, uni.PsychologyRank, uni.PhysicalSciencesRank,
			uni.SocialSciencesRank, latitude, longitude,
		)

		values = append(values, insertValues)
	}
  _, err = db.Exec(`DELETE FROM universities`)
	if err != nil {
		return fmt.Errorf("Unable to delete university data: %s\n", err)
	}


	insertQuery := fmt.Sprintf(`INSERT INTO universities VALUES %s;`, strings.Join(values, ",\n"))
	_, err = db.Exec(insertQuery)

	if err != nil {
		return fmt.Errorf("Unable to insert university data: %s\n", err)
	}

	return nil
}

func main() {
	uniRankFetcher := NewUniversityRankingsFetcher()
	log.Println("Fetching University data")
	universities, err := uniRankFetcher.GetUniversityRankings()

	if err != nil {
		log.Fatal("Unable to fetch university rankings data:", err)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(universities[i])
	}
  for i := range universities {
    utils.ConvertZeroToMax(&universities[i])
  }

	log.Printf("Obtained %d Universities", len(universities))
	NewUniversityCoordsFetcher(universities).GetUniversityCoords()

	db, err := createDBConnection()
	if err != nil {
		fmt.Print(err)
	}

	err = storeUniversityData(db, universities)
	if err != nil {
		fmt.Print(err)
	}
}
