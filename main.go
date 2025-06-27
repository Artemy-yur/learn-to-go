package main

import (
	"fmt"
	"time"
)

type city struct {
	name string
	time time.Time
	timezone string
	offset   int
}

var cities = []city{
	{name: "New York", time: time.Now(), timezone: "EST", offset: -5},
	{name: "Los Angeles", time: time.Now(), timezone: "PST", offset: -8},
	{name: "Chicago", time: time.Now(), timezone: "CST", offset: -6},
	{name: "Houston", time: time.Now(), timezone: "CST", offset: -6},
	{name: "Miami", time: time.Now(), timezone: "EST", offset: -5},
	{name: "Moscow", time: time.Now(), timezone: "MSK", offset: 3},
	{name: "Paris", time: time.Now(), timezone: "CET", offset: 1},
	{name: "Berlin", time: time.Now(), timezone: "CET", offset: 1},
	{name: "Madrid", time: time.Now(), timezone: "CET", offset: 1},
	{name: "Rome", time: time.Now(), timezone: "CET", offset: 1},
	{name: "London", time: time.Now(), timezone: "GMT", offset: 0},
	{name: "Tokyo", time: time.Now(), timezone: "JST", offset: 9},
}
var new_city string

func main() {
	now := time.Now()
	time_cites(now.Format("2006-01-02"))
	fmt.Println("Введите город для получения времени: ")
	var city_for_time string
	fmt.Scanln(&city_for_time)

	// Получаем текущее время в разных часовых поясах
	if city_for_time == "New York" || city_for_time == "Los Angeles" || city_for_time == "Chicago" || city_for_time == "Houston" || city_for_time == "Miami" || city_for_time == "Moscow" || city_for_time == "Paris" || city_for_time == "Berlin" || city_for_time == "Madrid" || city_for_time == "Rome" || city_for_time == "London" || city_for_time == "Tokyo" {
		fmt.Printf("Время в %s: %s\n", city_for_time, now.In(time.FixedZone("EST", -5*3600)).Format("15:04:05"))
	} else {
		fmt.Println("Вы можете добавить город в список: ")
		fmt.Scanln(&new_city)
		new_cites(now.Format("2006-01-02"), new_city)
	}

}
func time_cites(name string) {
	for _, city := range cities {
		fmt.Println(city)
	}
}
func new_cites(name string, new_city string) {

	cities = append(cities, city{name: new_city, time: time.Now()})
	if new_city == "New York" || new_city == "Los Angeles" || new_city == "Chicago" || new_city == "Houston" || new_city == "Miami" || new_city == "Moscow" || new_city == "Paris" || new_city == "Berlin" || new_city == "Madrid" || new_city == "Rome" || new_city == "London" || new_city == "Tokyo" {
		fmt.Println("Есть такой")
	} else {
		fmt.Println("Город не найден")
	}
	// fmt.Println(new)
}
func delete_city(name string, new_city string) {
	for i, city := range cities {
		if city.name == new_city {
			cities = append(cities[:i], cities[i+1:]...)
		}
	}
}
