package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Stats struct {
	LenghthOfDen int // = 10
	Health       int // = 100
	Respect      int // = 20
	Weight       int // = 30
}

var stats = Stats{10, 100, 20, 30}
var lod, hp, resp, wgh, wtd int

func main() {
	for {
		fmt.Println("Чем заняться? Пойти покопать нору(1), поспать(2), поесть травы(3), подраться(4)?")
		fmt.Print("Ваш ввод: ")
		fmt.Scanln(&wtd)
		if wtd <= 0 || wtd > 4 {
			fmt.Println("Неверный ввод")
			continue
		}

		switch wtd {
		case 1:
			lod, hp, resp, wgh = stats.Dig(wtd)
			stats.LenghthOfDen = lod
			stats.Health = hp
			stats.Respect = resp
			stats.Weight = wgh
			fmt.Println("Вы копали нору, ваши статы:", "Длина норы:", lod, "Здоровье:", hp, "Уважение:", resp, "Вес:", wgh)
			stats.DeathCheck()
			lod, hp, resp, wgh = stats.Night()
			fmt.Println("После того, как вы выполнили действие, вы пошли спать, ваши статы:", "Длина норы:", lod, "Здоровье:", hp, "Уважение:", resp, "Вес:", wgh)
			stats.DeathCheck()
			stats.LenghthOfDen = lod
			stats.Health = hp
			stats.Respect = resp
			stats.Weight = wgh
			main()
		case 2:
			lod, hp, resp, wgh = stats.Night()
			stats.LenghthOfDen = lod
			stats.Health = hp
			stats.Respect = resp
			stats.Weight = wgh
			fmt.Println("Вы решили проспать весь день, ваши статы:", "Длина норы:", lod, "Здоровье:", hp, "Уважение:", resp, "Вес:", wgh)
			stats.DeathCheck()
			lod, hp, resp, wgh = stats.Night()
			fmt.Println("После того, как вы выполнили действие, вы пошли спать, ваши статы:", "Длина норы:", lod, "Здоровье:", hp, "Уважение:", resp, "Вес:", wgh)
			stats.DeathCheck()
			stats.LenghthOfDen = lod
			stats.Health = hp
			stats.Respect = resp
			stats.Weight = wgh
			main()
		case 3:
			lod, hp, resp, wgh = stats.LetsEat(wtd)
			stats.LenghthOfDen = lod
			stats.Health = hp
			stats.Respect = resp
			stats.Weight = wgh
			fmt.Println("Вы поели, ваши статы:", "Длина норы:", lod, "Здоровье:", hp, "Уважение:", resp, "Вес:", wgh)
			stats.DeathCheck()
			lod, hp, resp, wgh = stats.Night()
			fmt.Println("После того, как вы выполнили действие, вы пошли спать, ваши статы:", "Длина норы:", lod, "Здоровье:", hp, "Уважение:", resp, "Вес:", wgh)
			stats.DeathCheck()
			stats.LenghthOfDen = lod
			stats.Health = hp
			stats.Respect = resp
			stats.Weight = wgh
			main()
		case 4:
			lod, hp, resp, wgh = stats.Fight(wtd)
			stats.LenghthOfDen = lod
			stats.Health = hp
			stats.Respect = resp
			stats.Weight = wgh
			fmt.Println("Вы подрались, ваши статы:", "Длина норы:", lod, "Здоровье:", hp, "Уважение:", resp, "Вес:", wgh)
			stats.DeathCheck()
			stats.WinCheck()
			lod, hp, resp, wgh = stats.Night()
			fmt.Println("После того, как вы выполнили действие, вы пошли спать, ваши статы:", "Длина норы:", lod, "Здоровье:", hp, "Уважение:", resp, "Вес:", wgh)
			stats.DeathCheck()
			stats.LenghthOfDen = lod
			stats.Health = hp
			stats.Respect = resp
			stats.Weight = wgh
			main()
		}

	}
}

// Ночь наступает после каждого действия.
func (st *Stats) Night() (int, int, int, int) {
	lod := st.LenghthOfDen
	hp := st.Health
	resp := st.Respect
	wgh := st.Weight
	lod -= 2
	hp += 20
	resp -= 2
	wgh -= 5
	return lod, hp, resp, wgh

}
func (st *Stats) Dig(wtd int) (int, int, int, int) {
	lod := st.LenghthOfDen
	hp := st.Health
	resp := st.Respect
	wgh := st.Weight
	fmt.Println("Покопать нору интенсивно(1) или лениво(2)?")
	fmt.Print("Ваш ввод: ")
	fmt.Scanln(&wtd)
	switch wtd {
	case 1:
		lod += 5
		hp -= 30
	case 2:
		lod += 2
		hp -= 10
	}
	return lod, hp, resp, wgh

}

func (st *Stats) LetsEat(wtd int) (int, int, int, int) {
	lod := st.LenghthOfDen
	hp := st.Health
	resp := st.Respect
	wgh := st.Weight
	fmt.Println("Поесть травы жухлой(1) или зелёной(2)?")
	fmt.Print("Ваш ввод: ")
	fmt.Scanln(&wtd)
	switch wtd {
	case 1:
		fmt.Println("Здоровье увеличилось на 10, вес увеличился на 15")
		hp += 10
		wgh += 15
	case 2:
		if resp < 30 {
			fmt.Println("Здоровье уменьшилось на 30")
			hp -= 30
		} else {
			fmt.Println("Здоровье уменьшилось на 30, вес увеличился на 30")
			wgh += 30
			hp -= 30
		}
	}
	return lod, hp, resp, wgh
}
func (st *Stats) DeathCheck() (int, int, int, int) {
	lod := st.LenghthOfDen
	hp := st.Health
	resp := st.Respect
	wgh := st.Weight
	if (lod <= 0) || (hp <= 0) || (resp <= 0) || (wgh <= 0) {
		fmt.Println("Одна из ваших характеристик упала до нуля, игра окончена")
		os.Exit(1)
	}
	return lod, hp, resp, wgh
}
func (st *Stats) WinCheck() int {
	resp := st.Respect
	if resp >= 100 {
		fmt.Println("Вы добились уважения равное или превышающее 100. Победа!")
		os.Exit(1)
	}
	return resp
}
func (st *Stats) Fight(wtd int) (int, int, int, int) {
	lod := st.LenghthOfDen
	hp := st.Health
	resp := st.Respect
	wgh := st.Weight
	var enemyWeight, WieghtSum int
	fmt.Println("Подраться со слабым(1), средним(2) или сильным противником(3)?")
	fmt.Print("Ваш ввод: ")
	fmt.Scanln(&wtd)
	switch wtd {
	case 1:
		enemyWeight = 30
	case 2:
		enemyWeight = 50
	case 3:
		enemyWeight = 70
	}
	WieghtSum = enemyWeight + wgh
	var chanceToWin float64 = float64(wgh) / float64(WieghtSum)
	var random float64 = rand.Float64()
	rand.Seed(time.Now().UnixNano())
	if random <= chanceToWin {
		fmt.Println("Вы победили и получили за это уважение")
		var plusRep = enemyWeight - wgh
		if plusRep <= 0 {
			plusRep += 10
		}
		resp += plusRep
		fmt.Println("Уважение:", resp)
	} else {
		fmt.Println("Вы проиграли, ваше здоровье уменьшилось")
		var minusHp = enemyWeight - wgh
		if minusHp <= 0 {
			minusHp += 10
		}
		hp -= minusHp
		fmt.Println("Здоровье:", hp)
	}

	return lod, hp, resp, wgh
}
