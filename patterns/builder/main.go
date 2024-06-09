package main

import (
	"fmt"
)

/*
	Builder - порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово
	Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов

	Use-cases:
	Когда вы хотите избавиться от «телескопического конструктора»
	Когда ваш код должен создавать разные представления какого-то объекта. Например, деревянные и железобетонные дома

	Props:
	Позволяет создавать объект пошагов
	Позволяет использовать один и тот же код для создания различных представлений объектов
	Изолирует код создания класса/структуры от его основной бизнес-логики.

	Cons:
	Усложняет код программы из-за введения дополнительных классов.
	Builder classes must be mutable
*/

//Моя программа будет создавать персонаад с разными характеристиками

// Обозначим методы создания
type CreateHero interface {
	SetHeroName()
	SetSkinColor()
	SetHairColor()
	SetHeroGrowth()
	SetHeroWeight()
	GetHero() Hero
}

type Hero struct {
	Name       string
	SkinColor  string
	HairColor  string
	HeroGrowth float32
	HeroWeight float32
}

// Функция вывода инфы о новом герое
func (h *Hero) NewHeroInfo() {
	fmt.Printf("Create new Hero: %s!\n Its parameters:\n SkinColor: %s\n HairColor: %s\n Growth: %.1f\n Weight: %.1f\n", h.Name, h.SkinColor, h.HairColor, h.HeroGrowth, h.HeroWeight)
}

// Обозначи константы дефолтных героев
const (
	Man   = "Boy"
	Woman = "Girl"
)

// Конструктор для возвращения создания героя
func GetNewHero(HeroType string) CreateHero {
	switch HeroType {
	default:
		return nil
	case Man:
		return &Boy{}
	case Woman:
		return &Girl{}
	}
}

// Создадим дефолтных героев
type Boy struct {
	Name       string
	SkinColor  string
	HairColor  string
	HeroGrowth float32
	HeroWeight float32
}

func (Character *Boy) SetHeroName() {
	Character.Name = "Aboba"
}

func (Character *Boy) SetSkinColor() {
	Character.SkinColor = "Green"
}

func (Character *Boy) SetHairColor() {
	Character.HairColor = "Black"
}

func (Character *Boy) SetHeroGrowth() {
	Character.HeroGrowth = 185.4
}

func (Character *Boy) SetHeroWeight() {
	Character.HeroWeight = 100.7
}

func (Character *Boy) GetHero() Hero {
	return Hero{
		Name:       Character.Name,
		SkinColor:  Character.SkinColor,
		HairColor:  Character.HairColor,
		HeroGrowth: Character.HeroGrowth,
		HeroWeight: Character.HeroWeight,
	}
}

type Girl struct {
	Name       string
	SkinColor  string
	HairColor  string
	HeroGrowth float32
	HeroWeight float32
}

func (Character *Girl) SetHeroName() {
	Character.Name = "Biba"
}

func (Character *Girl) SetSkinColor() {
	Character.SkinColor = "White"
}

func (Character *Girl) SetHairColor() {
	Character.HairColor = "Red"
}

func (Character *Girl) SetHeroGrowth() {
	Character.HeroGrowth = 165.4
}

func (Character *Girl) SetHeroWeight() {
	Character.HeroWeight = 55.7
}

func (Character *Girl) GetHero() Hero {
	return Hero{
		Name:       Character.Name,
		SkinColor:  Character.SkinColor,
		HairColor:  Character.HairColor,
		HeroGrowth: Character.HeroGrowth,
		HeroWeight: Character.HeroWeight,
	}
}

// Создадим инициализатор нового базового героя
type DefaultHero struct {
	CreateHero CreateHero
}

func NewDefaultHero(createHero CreateHero) *DefaultHero {
	return &DefaultHero{CreateHero: createHero}
}

// СОздадим функцию для изменения героев
func (defaultHero *DefaultHero) SetHero(createHero CreateHero) {
	defaultHero.CreateHero = createHero
}

// Основная функция создания которая создает уже готового героя, в ней можно управлять созданием героя
func (defaultHero *DefaultHero) CreateNewHero() Hero {
	defaultHero.CreateHero.SetHeroName()
	defaultHero.CreateHero.SetSkinColor()
	defaultHero.CreateHero.SetHairColor()
	defaultHero.CreateHero.SetHeroGrowth()
	defaultHero.CreateHero.SetHeroWeight()
	return defaultHero.CreateHero.GetHero()
}

// Проверяем
func main() {
	boy := GetNewHero("Boy")
	girl := GetNewHero("Girl")

	newHero := NewDefaultHero(boy)
	boyHero := newHero.CreateNewHero()
	boyHero.NewHeroInfo()

	newHero.SetHero(girl)
	girlHero := newHero.CreateNewHero()
	girlHero.NewHeroInfo()
}
