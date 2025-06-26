package main

import (
	"encoding/json"
	"fmt"
	"maps"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

type UniversityRankingsFetcher struct {
}

type University struct {
	Name                     string
	OverallRank              int
	Country                  string
	ArtsAndHumanitiesRank    int
	BusinessAndEconomicsRank int
	ClinicalAndHealth        int
	ComputerScienceRank      int
	EducationRank            int
	EngineeringRank          int
	LawRank                  int
	LifeSciencesRank         int
	PsychologyRank           int
	PhysicalSciencesRank     int
	SocialSciencesRank       int
	Latitude                 string
	Longitude                string
}

type UniversityRank struct {
	Name    string `json:"name"`
	Rank    string `json:"rank"`
	Country string `json:"location"`
}

func NewUniversityRankingsFetcher() *UniversityRankingsFetcher {
	return &UniversityRankingsFetcher{}
}

func getUniversitySpecificRanking(url string) ([]UniversityRank, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	body := struct {
		Data []UniversityRank `json:"data"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		return nil, err
	}

	return body.Data, nil
}

func CleanRank(rawRank string) int {
	rawRank = strings.TrimSpace(rawRank)
	rawRank = strings.TrimPrefix(rawRank, "=")
	rawRank = strings.TrimSuffix(rawRank, "+")

	if strings.Contains(rawRank, "–") {
		rangeParts := strings.Split(rawRank, "–")
		if len(rangeParts) > 0 {
			// Use the lower bound of the range
			cleanedRank, err := strconv.Atoi(rangeParts[0])
			if err == nil {
				return cleanedRank
			}
		}
	}

	if strings.ToLower(rawRank) == "reporter" {
		return 999999
	}

	cleanedRank, err := strconv.Atoi(rawRank)
	if err != nil {
		return 999999
	}

	return cleanedRank
}

func passRankValueToUniversity(url string, rankPassCallback func(UniversityRank)) error {

	ranks, err := getUniversitySpecificRanking(url)
	if err != nil {
		return err
	}
	for _, rank := range ranks {
		rank.Rank = fmt.Sprintf("%d", CleanRank(rank.Rank))
		rankPassCallback(rank)
	}

	return err
}

func (_ *UniversityRankingsFetcher) GetUniversityRankings() ([]University, error) {
	universities := map[string]University{}

	// Overall Rankings
	url := "https://www.timeshighereducation.com/sites/default/files/the_data_rankings/world_university_rankings_2025_0__ba2fbd3409733a83fb62c3ee4219487c.json"
	passRankValueToUniversity(url, func(uniRank UniversityRank) {
		uni := universities[uniRank.Name]
		rank, _ := strconv.Atoi(uniRank.Rank)
		uni.OverallRank = rank
		uni.Name = uniRank.Name
		uni.Country = uniRank.Country
		universities[uniRank.Name] = uni
	})

	// Arts and Humanities Rankings
	url = "https://www.timeshighereducation.com/sites/default/files/the_data_rankings/arts_humanities_rankings_2024_0__ccb001eff81d1137b1111a3a20ef5f32.json"
	passRankValueToUniversity(url, func(uniRank UniversityRank) {
		uni := universities[uniRank.Name]
		rank, _ := strconv.Atoi(uniRank.Rank)
		uni.ArtsAndHumanitiesRank = rank
		uni.Name = uniRank.Name
		uni.Country = uniRank.Country
		universities[uniRank.Name] = uni
	})

	// Business & Economics Rankings
	url = "https://www.timeshighereducation.com/sites/default/files/the_data_rankings/business_economics_rankings_2024_0__b3c196ad15cc3a840eea0e69a6b22323.json"
	passRankValueToUniversity(url, func(uniRank UniversityRank) {
		uni := universities[uniRank.Name]
		rank, _ := strconv.Atoi(uniRank.Rank)
		uni.BusinessAndEconomicsRank = rank
		uni.Name = uniRank.Name
		uni.Country = uniRank.Country
		universities[uniRank.Name] = uni
	})

	// Clinical and Health Rankings
	url = "https://www.timeshighereducation.com/sites/default/files/the_data_rankings/clinical_pre_clinical_health_ran_2024_0__b8881c632a7135c802c13201dc21d78b.json"
	passRankValueToUniversity(url, func(uniRank UniversityRank) {
		uni := universities[uniRank.Name]
		rank, _ := strconv.Atoi(uniRank.Rank)
		uni.ClinicalAndHealth = rank
		uni.Name = uniRank.Name
		uni.Country = uniRank.Country
		universities[uniRank.Name] = uni
	})

	// Computer Science Rankings
	url = "https://www.timeshighereducation.com/sites/default/files/the_data_rankings/computer_science_rankings_2024_0__fe5e1ecb8de7d97cac3213b7d3f5b05f.json"
	passRankValueToUniversity(url, func(uniRank UniversityRank) {
		uni := universities[uniRank.Name]
		rank, _ := strconv.Atoi(uniRank.Rank)
		uni.ComputerScienceRank = rank
		uni.Name = uniRank.Name
		uni.Country = uniRank.Country
		universities[uniRank.Name] = uni
	})

	// Education Rankings
	url = "https://www.timeshighereducation.com/sites/default/files/the_data_rankings/education_rankings_2024_0__5f2a8c685e8d7ac407b330777634adde.json"
	passRankValueToUniversity(url, func(uniRank UniversityRank) {
		uni := universities[uniRank.Name]
		rank, _ := strconv.Atoi(uniRank.Rank)
		uni.EducationRank = rank
		uni.Name = uniRank.Name
		uni.Country = uniRank.Country
		universities[uniRank.Name] = uni
	})

	// Engineering Rankings
	url = "https://www.timeshighereducation.com/sites/default/files/the_data_rankings/engineering_technology_rankings_2024_0__aa600f71232c516f6d241ee213f82c0f.json"
	passRankValueToUniversity(url, func(uniRank UniversityRank) {
		uni := universities[uniRank.Name]
		rank, _ := strconv.Atoi(uniRank.Rank)
		uni.EngineeringRank = rank
		uni.Name = uniRank.Name
		uni.Country = uniRank.Country
		universities[uniRank.Name] = uni
	})

	// Law Rankings
	url = "https://www.timeshighereducation.com/sites/default/files/the_data_rankings/law_rankings_2024_0__52d627e653b5e2db67688d415fb54533.json"
	passRankValueToUniversity(url, func(uniRank UniversityRank) {
		uni := universities[uniRank.Name]
		rank, _ := strconv.Atoi(uniRank.Rank)
		uni.LawRank = rank
		uni.Name = uniRank.Name
		uni.Country = uniRank.Country
		universities[uniRank.Name] = uni
	})

	// Life Sciences Rankings
	url = "https://www.timeshighereducation.com/sites/default/files/the_data_rankings/life_sciences_rankings_2024_0__891eecb7ad740bbf1497885899ac2eef.json"
	passRankValueToUniversity(url, func(uniRank UniversityRank) {
		uni := universities[uniRank.Name]
		rank, _ := strconv.Atoi(uniRank.Rank)
		uni.LifeSciencesRank = rank
		uni.Name = uniRank.Name
		uni.Country = uniRank.Country
		universities[uniRank.Name] = uni
	})

	// Psychology Rankings
	url = "https://www.timeshighereducation.com/sites/default/files/the_data_rankings/psychology_rankings_2024_0__a17810e20dbbfc018c4436a051eca33d.json"
	passRankValueToUniversity(url, func(uniRank UniversityRank) {
		uni := universities[uniRank.Name]
		rank, _ := strconv.Atoi(uniRank.Rank)
		uni.PsychologyRank = rank
		uni.Name = uniRank.Name
		uni.Country = uniRank.Country
		universities[uniRank.Name] = uni
	})

	// Physical Sciences Rankings
	url = "https://www.timeshighereducation.com/sites/default/files/the_data_rankings/physical_sciences_rankings_2024_0__10810ad22bbdaa7c2f980efc335f94f2.json"
	passRankValueToUniversity(url, func(uniRank UniversityRank) {
		uni := universities[uniRank.Name]
		rank, _ := strconv.Atoi(uniRank.Rank)
		uni.PhysicalSciencesRank = rank
		uni.Name = uniRank.Name
		uni.Country = uniRank.Country
		universities[uniRank.Name] = uni
	})

	// Social Sciences Rankings
	url = "https://www.timeshighereducation.com/sites/default/files/the_data_rankings/social_sciences_rankings_2024_0__345dd8efc1e3a81bfc7ec73595e8892e.json"
	passRankValueToUniversity(url, func(uniRank UniversityRank) {
		uni := universities[uniRank.Name]
		rank, _ := strconv.Atoi(uniRank.Rank)
		uni.SocialSciencesRank = rank
		uni.Name = uniRank.Name
		uni.Country = uniRank.Country
		universities[uniRank.Name] = uni
	})

	return slices.Collect(maps.Values(universities)), nil
}
