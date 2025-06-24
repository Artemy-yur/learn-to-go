package main
import(
	"fmt"
	"time"
	
)
var new_city string
var cities = []string{
	"New York",
	"Los Angeles",
	"Chicago",
	"Houston",
	"Miami",
	"Moscow",
	"Paris",
	"Berlin",
	"Madrid",
	"Rome",
	"London",
	"Tokyo",
}
func main() {
	now := time.Now()
	time_cites(now.Format("2006-01-02"))
	fmt.Println("Введите город для получения времени: ")
	var city_for_time string
	fmt.Scanln(&city_for_time)

	switch city_for_time {
	case "New York":
		fmt.Printf("Время в %s: %s\n", city_for_time, now.In(time.FixedZone("EST", -5*3600)).Format("15:04:05"))
	case "Los Angeles":
		fmt.Printf("Время в %s: %s\n", city_for_time, now.In(time.FixedZone("PST", -8*3600)).Format("15:04:05"))
	case "Chicago":
		fmt.Printf("Время в %s: %s\n", city_for_time, now.In(time.FixedZone("CST", -6*3600)).Format("15:04:05"))
	case "Houston":
		fmt.Printf("Время в %s: %s\n", city_for_time, now.In(time.FixedZone("CST", -6*3600)).Format("15:04:05"))
	case "Miami":
		fmt.Printf("Время в %s: %s\n", city_for_time, now.In(time.FixedZone("EST", -5*3600)).Format("15:04:05"))
	case "Moscow":
		fmt.Printf("Время в %s: %s\n", city_for_time, now.In(time.FixedZone("MSK", 3*3600)).Format("15:04:05"))
	case "Paris":
		fmt.Printf("Время в %s: %s\n", city_for_time, now.In(time.FixedZone("CET", 1*3600)).Format("15:04:05"))
	case "Berlin":
		fmt.Printf("Время в %s: %s\n", city_for_time, now.In(time.FixedZone("CET", 1*3600)).Format("15:04:05"))
	case "Madrid":
		fmt.Printf("Время в %s: %s\n", city_for_time, now.In(time.FixedZone("CET", 1*3600)).Format("15:04:05"))
	case "Rome":
		fmt.Printf("Время в %s: %s\n", city_for_time, now.In(time.FixedZone("CET", 1*3600)).Format("15:04:05"))
	case "London":
		fmt.Printf("Время в %s: %s\n", city_for_time, now.In(time.FixedZone("GMT", 0)).Format("15:04:05"))
	case "Tokyo":
		fmt.Printf("Время в %s: %s\n", city_for_time, now.In(time.FixedZone("JST", 9*3600)).Format("15:04:05"))
	default:
		fmt.Println("Вы можете добавить город в список: ")
		fmt.Scanln(&new_city)
		new_cites(now.Format("2006-01-02"),new_city)
	
	}
	
	
	
	
}
func time_cites(name string) {
	for _, city := range cities {
		fmt.Println(city)
	}
}
func new_cites(name string,new_city string) {

	cities = append(cities, new_city)
	switch {
	case new_city == "New York":
		fmt.Println("Есть такой")
	case new_city == "Los Angeles":
		fmt.Println("Есть такой")
	case new_city == "Chicago":
		fmt.Println("Есть такой")
	case new_city == "Houston":
		fmt.Println("Есть такой")
	case new_city == "Miami":
		fmt.Println("Есть такой")
	case new_city == "Moscow":
		fmt.Println("Есть такой")
	case new_city == "Paris":
		fmt.Println("Есть такой")
	case new_city == "Berlin":
		fmt.Println("Есть такой")
	case new_city == "Madrid":
		fmt.Println("Есть такой")
	case new_city == "Rome":
		fmt.Println("Есть такой")
	case new_city == "London":
		fmt.Println("Есть такой")
	case new_city == "Tokyo":
		fmt.Println("Есть такой")
	default:
		fmt.Println("Город не найден")
	}
	// fmt.Println(new)
}
