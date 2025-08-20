package handler

import (
	"api/db"
	"api/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type University struct {
	Name                     string
	OverallRank              int
	Country                  string
	ArtsAndHumanitiesRank    int
	BusinessAndEconomicsRank int
	ClinicalAndHealthRank    int
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

type UniversityHandler struct{}

func NewUniversityHandler() *UniversityHandler {
	return &UniversityHandler{}
}

func (_ *UniversityHandler) GetRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getUniversities)
	mux.HandleFunc("/fields", getUniversityFields)
	mux.HandleFunc("/countries", getCountries)

	return mux
}

func getCountries(w http.ResponseWriter, r *http.Request) {
	db, err := db.NewPostgresDB()
	log.Println("1")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.DB.Query(`SELECT DISTINCT(country) FROM universities ORDER BY country;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	countries := []string{}
	for rows.Next() {
		var country string
		if err = rows.Scan(&country); err != nil {
			log.Fatal(err)
		}

		countries = append(countries, country)
	}

	json.NewEncoder(w).Encode(countries)
}

func getUniversityFields(w http.ResponseWriter, r *http.Request) {
	log.Println("3")
	uniType := reflect.TypeOf(University{})
	w.Header().Set("Content-Type", "application/json")

	uniFields := []string{}
	for i := 0; i < uniType.NumField()-2; i++ {
		uniFields = append(uniFields, uniType.Field(i).Name)
	}

	json.NewEncoder(w).Encode(uniFields)
}

func getUniversities(w http.ResponseWriter, r *http.Request) {
	sortedFields := r.URL.Query().Get("sort")
	limit := r.URL.Query().Get("limit")
	country := r.URL.Query().Get("country")
	name := r.URL.Query().Get("name")

	if limit != "" {
		limit = "LIMIT " + limit
	}

	if sortedFields != "" {
		sortedFields = "ORDER BY " + sortedFields
	}

	country = utils.ConvertToSQLWhereCondition(country)
	name = utils.ConvertToSQLWhereCondition(name)

	whereClause := fmt.Sprintf("WHERE country %s AND name %s", country, name)
	fmt.Println(whereClause)

	if sortedFields != "" {
		sortedFields = strings.ReplaceAll(sortedFields, ":asc", " ASC")
		sortedFields = strings.ReplaceAll(sortedFields, ":desc", " DESC")
	}

	query := fmt.Sprintf("SELECT * FROM universities %s %s %s", whereClause, sortedFields, limit)
	fmt.Println(query)
	db, err := db.NewPostgresDB()
	if err != nil {
		w.WriteHeader(500)
		log.Println(err.Error())
		w.Write([]byte("Server cannot connect to db: " + err.Error()))
		return
	}

	rows, err := db.DB.Query(query)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err.Error())
		w.Write([]byte("Server cannot execute sorting" + err.Error()))
		return
	}

	unis := []University{}
	for rows.Next() {
		uni := University{}

		dest := make([]any, reflect.TypeOf(uni).NumField())
		structValue := reflect.ValueOf(&uni).Elem()
		for i := 0; i < structValue.NumField(); i++ {
			dest[i] = structValue.Field(i).Addr().Interface()
		}

		if err = rows.Scan(dest...); err != nil {
			log.Fatal("Uhhhh", err)
		}

		unis = append(unis, uni)
	}

	err = json.NewEncoder(w).Encode(unis)
	if err != nil {
		w.Write([]byte("Cannot convert data to json" + err.Error()))
		w.WriteHeader(500)
		return
	}
}
