package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type SafkaResponse struct {
	Restaurants map[string]Restaurant `json:"restaurants_tty"`
}

type Meal struct {
	Category    string `json:"kok"`
	PriceString string `json:"price"`
	Price       float64
	Food        []map[string]string `json:"mo"`
}

type Restaurant struct {
	Name  string `json:"restaurant"`
	Open  bool   `json:"open_today"`
	Meals []Meal `json:"meals"`
}

type Restaurants []Restaurant

var weekDays = []string{"su", "ma", "ti", "ke", "to", "pe", "la"}

func FetchRestaurants() (*Restaurants, error) {
	t := time.Now()

	_, weekNumber := t.ISOWeek()
	weekDay := weekDays[t.Weekday()]
	year := t.Year()

	addr := fmt.Sprintf("https://unisafka.fi/static/json/%d/%d/1/%s.json", year, weekNumber, weekDay)

	res, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	var sr SafkaResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&sr); err != nil {
		return nil, err
	}

	restaurants := make(Restaurants, 0)

	for _, v := range sr.Restaurants {
		for i, m := range v.Meals {
			v.Meals[i].Price, err = strconv.ParseFloat(
				strings.ReplaceAll(strings.Split(m.PriceString, " ")[0], ",", "."),
				64)
		}
		restaurants = append(restaurants, v)
	}

	return &restaurants, nil
}

func (m Meal) String() string {
	var parts []string

	for _, part := range m.Food {
		parts = append(parts, part["mpn"])
	}

	food := strings.Join(parts, ", ")

	//return fmt.Sprintf("%s (%.2f€)", food, m.Price)
	if m.Price == float64(2.6) {
		return food
	}

	return ""

}

func (rs *Restaurants) String() string {
	result := ""

	for _, restaurant := range *rs {
		if !restaurant.Open {
			continue
		}
		result += "\n*" + restaurant.Name + "*\n"
		for _, meal := range restaurant.Meals {
			if meal.String() != "" {
				result += "• " + meal.String() + "\n"
			}
		}
	}

	return result
}

func (rs *Restaurants) Filter(allowed []string) {
	filtered := make(Restaurants, 0)
	for _, restaurant := range *rs {
		for _, name := range allowed {
			if restaurant.Name == name {
				filtered = append(filtered, restaurant)
			}
		}
	}

	*rs = filtered
}
