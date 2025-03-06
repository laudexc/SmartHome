package main

import "fmt"

// Структура устройств с слайсов девайсов
type smartHome struct {
	Devices    []AllDevices
	Light      *Light
	Thermostat *Thermostat
	SecSys     *SecSys
}

// Для обновления статуса после изменений
func (sH *smartHome) updateStatus() {
	sH.Devices = []AllDevices{sH.Light, sH.Thermostat, sH.SecSys}
}

// Интерфейс всех девайсов, с методом статуса устройств
type AllDevices interface {
	Status()
}

// Все для света:
type Light struct {
	isOn bool
}

func (l *Light) Status() {
	if l.isOn {
		fmt.Printf("Свет:\n Устройство активно\n")
	} else {
		fmt.Printf("Свет:\n Устройство неактивно\n")
	}
}

func (l *Light) turnOn() {
	l.isOn = true
}

func (l *Light) turnOff() {
	l.isOn = false
}

// Все для термостата:
type Thermostat struct {
	isOn        bool
	Temperature int
}

func (t *Thermostat) Status() {
	if t.isOn {
		fmt.Printf("Термостат:\n Устройство активно \n  Температура: %d°C\n", t.Temperature)
	} else {
		fmt.Printf("Термостат:\n Устройство неактивно\n")
	}
}

func (t *Thermostat) setTemperature(temp int) {
	t.Temperature = temp
	t.isOn = true
}

func (t *Thermostat) turnOff() {
	t.isOn = false
}

// Все для системы безопасности:
type SecSys struct {
	isOn bool
}

func (s *SecSys) Status() {
	if s.isOn {
		fmt.Printf("Система безопасности:\n Устройство активно\n")
	} else {
		fmt.Printf("Система безопасности:\n Устройство неактивно\n")
	}
}

func (s *SecSys) turnOn() {
	s.isOn = true
}

func (s *SecSys) turnOff() {
	s.isOn = false
}

// Метод для просмотра статуса всех устройств сразу
func (sH *smartHome) Status() {
	fmt.Printf("\nСтатус всех устройств умного дома:\n")
	for _, devices := range sH.Devices {
		devices.Status()
	}
}

// Методы управления состояниями устройств
func (sH *smartHome) turnOnLight() {
	sH.Light.turnOn()
}

func (sH *smartHome) turnOffLight() {
	sH.Light.turnOff()
}

func (sH *smartHome) SetTemperature(temp int) {
	sH.Thermostat.setTemperature(temp)
}

func (sH *smartHome) turnOffThermostat() {
	sH.Thermostat.turnOff()
}

func (sH *smartHome) turnOnSec() {
	sH.SecSys.turnOn()
}

func (sH *smartHome) turnOffSec() {
	sH.SecSys.turnOff()
}

// Универсальный контроллер управления устройствами
func (sH *smartHome) ControlPanel(action string, value ...int) { // зачем тут ...?
	switch action {
	case "light on":
		sH.turnOnLight()

	case "light off":
		sH.turnOffLight()

	case "set temperature":
		if len(value) > 0 {
			if value[0] > 20 {
				sH.SetTemperature(20)
			} else if value[0] < -10 {
				sH.SetTemperature(-10)
			} else {
				sH.SetTemperature(value[0])
			}
		} else {
			fmt.Println("Ошибка, не указана температура!")
		}

	case "thermostat off":
		sH.turnOffThermostat()

	case "SecSys on":
		sH.turnOnSec()

	case "SecSys off":
		sH.turnOffSec()

	default:
		fmt.Println("Неизвестная команда")
	}
	sH.updateStatus() // Обнвление статуса устройств
}

func main() {
	home := smartHome{
		Devices: []AllDevices{
			&Light{isOn: true},
			&Thermostat{isOn: true, Temperature: 10},
			&SecSys{isOn: false},
		},
		Light:      &Light{isOn: true},
		Thermostat: &Thermostat{isOn: true, Temperature: 10},
		SecSys:     &SecSys{isOn: false},
	}

	home.Status()

	home.ControlPanel("light off")
	home.ControlPanel("SecSys on")

	home.Status()
}
