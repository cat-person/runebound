## Домашнее задание для лекции 5
* Откройте файл main.go
* Добавьте имплементацию в метод levelUp таким образом, чтобы:
  * При каждом вызове levelUp уроень игрока увеличивался на 1
  * При достижении 3 - го уровня игрок имел возможность выбрать новую профессию
  * При достижении 5 - го уровня игрок имел возможность выбрать продвинутое происхождение (см. ancestry.Progression)
  * В любом другом случае игроку будет дана возможность выбрать свойство (Feat) из списка (в файле feat.go)
  * При достижении 10-го уровня выводится надпись "You Won" и игра прекращается

## Задание со звездочкой 
Превратите Feat / Profession в колоды с неповторяющимися "картами" каждый раз при выборе сокращайте количество доступных вариантов

## Задание с двумя звездочками 
Придумайте пару-тройку новых Свойств и Профессий каждый раз при выборе вытаскивайте 3 карты из перетасованной колоды (Свойств или Профессий) выбираете одну из них и остальные сбрасывайте (Будет интересно если вы добавите также колоды сброса) 
