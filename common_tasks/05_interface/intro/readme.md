### Что такое интерфейс?

**Интерфейс** в языке Go — это специальный тип, который определяет набор сигнатур методов, но не содержит их реализацию. Интерфейсы позволяют описывать поведение типов, что делает код более гибким. 

#### Объявление интерфейса

Чтобы объявить интерфейс, используется ключевое слово `type`, затем указывается имя интерфейса и ключевое слово `interface`, после чего в фигурных скобках перечисляются сигнатуры методов, которые должны быть реализованы типами, удовлетворяющими этому интерфейсу. Например:

```go
type MyInterface interface {
	MyMethod()
}
```

### Встраивание интерфейсов
В Go интерфейсы поддерживает встраивание. Для этого нужно в объявлении нового интерфейса указать название какого-то другого интерфейса.

### [Значение интерфейса](https://go.dev/tour/methods/11)

Значение интерфейса можно рассматривать как кортеж из значения и конкретного типа:

`(value, type)`

Значение интерфейса содержит значение определенного базового конкретного типа.

Вызов метода для значения интерфейса выполняет метод с тем же именем для его базового типа.

### Пример создания значения интерфейсного типа:

Создадим интерфейс Animal, который требует реализации метода Speak().

``` go
type Animal interface {
    Speak() string
}
```

Затем мы создаем переменную animal интерфейсного типа Animal:

``` go
var animal Animal
```

Через конструкцию `fmt.Printf("Value %v, type %T\n", animal, animal)` выведем на экран значение и тип интерфейса `animal`. Увидим следующие сообщение:

``` bash
Value <nil>, type <nil>
```

Здесь мы видим, что `value` и `type` равны nil. Это означает, что переменная animal не содержит никакого значения и не указывает на конкретный тип. 

При попытке сравнить интерфейс с nil мы видим, что сообщение "animal is not nil" не распечатается.

``` go
if animal != nil {
    fmt.Println("animal is not nil")
}
```

Создадим указатель на новый объект Dog и присвоим dog интерфейсу animal.

``` go
dog := &Dog{}
animal = dog
```

Теперь, когда animal указывает на объект Dog, мы можем вызвать метод Speak. Это безопасно, потому что Dog реализует метод Speak, требуемый интерфейсом Animal.

``` go
animal.Speak() // OK
```

Снова воспользуемся конструкцией `fmt.Printf("Value %v, type %T\n", animal, animal)` выведем на экран значение и тип интерфейса. Увидим следующие сообщение:

``` bash
Value &{}, type *main.Dog
```

Тип: %T показывает, что animal — это указатель на Dog (*main.Dog).

Затем мы проверяем, не равна ли animal nil:

``` go
if animal != nil {
    fmt.Println("animal is not nil")
}
```

Поскольку animal указывает на объект Dog, условие выполняется, и выводится сообщение "animal is not nil".

Изменим поле Name у объекта Dog, на который указывает dog

``` go
dog.Name = "Шайтан"
```

Выводем снова через `Printf` тип и значение переменной animal:

``` go
Value &{Шайтан}, type *main.Dog
```

Теперь значение интерфейса поля Name обновлено до "Шайтан", что видно в выводе.

### Вызов метода интерфейсного типа

Значение интерфейсного типа != nil, когда
конкретный тип != nil. Мы можем безопасно вызвать метод у интерфейса только в случае, когда значение интерфейсного типа != nil. В противном случае при вызове метода мы словим панику.

``` go
var animal Animal
animal.Speak() // паника при попытке вызвать (interface == nil)

dog := &Dog{}
animal = dog // interface != nil
animal.Speak() // OK
```

### Описание структуры интерфейса 

#### **Структура `iface`**

``` go
type iface struct {
	tab  *itab          // это указатель на Interface Table или itable - структуру, которая хранит некоторые метаданные о типе и список методов, используемых для удовлетворения интерфейса. 
	data unsafe.Pointer // хранимые данные (указатель на значение)
}
```

- **`tab *itab`:** Это указатель на таблицу интерфейса (`itable`), которая содержит информацию о типе и методах, необходимых для реализации интерфейса. Эта таблица помогает Go определить, какие методы доступны для данного интерфейсного значения и как их вызывать. (описание см. ниже)

- **`data unsafe.Pointer`:** Это указатель на конкретные данные или значение, которые реализуют интерфейс. Использование `unsafe.Pointer` позволяет интерфейсу ссылаться на данные произвольного типа, сохраняя при этом информацию о том, как к ним обращаться через `itab`.

#### **Структура `itab`**

``` go
type itab struct {       // 40 bytes on a 64bit arch
	inter *interfacetype // тип интерфейса
	_type *_type         // все, что мы знаем про тип из которого образован элемент интерфейса
	hash  uint32         // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr     // методы, которые должна описывать структура, чтобы релизовывать интерфейс
}
```

- **`inter *interfacetype`:** Метаданные интерфейса.
- **`_type *_type`:** Указатель на информацию о конкретном типе, который реализует интерфейс. Это позволяет Go знать, как обращаться с данными, которые реализуют интерфейс.

- **`hash uint32`:** Хеш типа, который используется для оптимизации операций с интерфейсами, таких как type switches.

- **`fun [1]uintptr`:** Массив указателей на функции, которые должны быть реализованы для удовлетворения интерфейса. Это позволяет динамически вызывать методы на интерфейсных значениях. uintptr - целочисленное представление адреса в памяти, указатель на первый элемент массива, который содержит указатели на методы. Размер массива [1], чтобы сохранить указатель на первый элемент массива.

### Иллюстрация хранения значений интерфейсного типа в структуре интерфейса:

Создадим свой пользовательский тип Binary с двумя методами `String() string` и `Get() uint64`.

```go
type Binary uint64

func (i Binary) String() string {
    return strconv.Uitob64(i.Get(), 2)
}

func (i Binary) Get() uint64 {
    return uint64(i)
}
```

Создадим экземпляр структуры `Binary` и присвоим ему значение:

```go
b := Binary(200)
```

Значение интерфейса представлены в виде пары из двух машинных слов, дающей указатель на информацию о типе, хранящемся в интерфейсе, и указатель на связанные данные.

![gointer2](images/gointer2.png)

Первое слово в значении интерфейса указывает на таблицу интерфейсов `itable`. В нем хранится информация о конкретном типе `type` и списке указателей на методы `fun[0]`. В нашем случае `type`- `Binary`, методы `String() string` и `Get() uint64`.

Второе слово указывает на значение `data`. В нашем случае `data` - 200.

Через конструкцию `fmt.Printf("Value %v, type %T\n", num, num)` выведем на экран значение и тип интерфейса. Получим:

``` bash
Value 11001000, type main.Binary
```

### Вопрос на собеседовании

Обладая этими знаниями, вы сможете ответить на популярный [вопрос из собеседования](https://github.com/alivewel/sber-tasks/tree/main/10_interface):

#### Что выведет программа?

```go
func main() {
	var ptr *struct{}
	var iface interface{}
	iface = ptr
	if iface == nil {
		println("It's nil!")
	}
} 
```

### Пустой интерфейс

Пустой интерфейс - это интерфейс, у которого отсутствуют методы. Для имплементации интерфейса нужно реализовать всего его методы. Для имплементации пустого интерфейса не нужно реализовывать никаких методов. Соответственно, любой тип в Go имплементирует пустой интерфейс. В других языках программирования такое называется `any`. В Go тоже есть `any`, это алиас (пользовательский тип) на пустой интерфейс.

При создании переменной пустого интерфейса в дальнейшем мы можем присвоить ему любой тип. 

```go
var emptyInterface interface{}

emptyInterface = dog

emptyInterface = 123

emptyInterface = true
```

### Цитата про пустой интерфейс из статьи про постулаты Go ([Go proverbs](https://habr.com/ru/articles/272383/))

#### Пустой интерфейс ни о чём не говорит (interface{} says nothing)

Этот постулат говорит о том, что интерфейсы — «поведенческие типы» — должны что-то означать. Если вы создаёте интерфейс, это что-то означает и служит конкретной цели. Пустой же интерфейс (`interface{}`) ничего не означает и ни о чём не говорит. 

Есть ситуации, когда его нужно использовать, но они чаще исключение — **не используйте `interface{}` без повода**. Новички часто переиспользуют пустые интерфейсы, и масса вопросов на Stack Overflow именно о них.

### Неявная имплементация интерфейсов

В Go используется неявная имплементация интерфейсов. В других языках программирования требуется использование ключевого слово `implements`. В Go используется концепция утиной типизации. Для того, чтобы имплементировать интерфейс типу необходимо реализовать все его методы. При этом можно реализовать больше методов чем это требуется, но меньше нельзя. Один тип может имплементировать несколько интерфейсов.

### Полиморфизм

Полиморфизм — это концепция, позволяющая объектам разных типов быть обработанными через единый интерфейс. В данном примере полиморфизм проявляется в том, что функция MakeAnimalSpeak может принимать любой тип, который реализует интерфейс Animal, и вызывать метод Speak, не зная конкретного типа объекта.
Это позволяет писать более гибкий и расширяемый код, так как вы можете добавлять новые типы, реализующие интерфейс Animal, без необходимости изменять существующий код, который работает с этим интерфейсом.
С помощью интерфейсов мы можем обстрагироваться от конкретных типов.

### Пример

Обе структуры Dog и Cat реализуют метод Speak, что делает их совместимыми с интерфейсом Animal. Метод Speak возвращает строку, описывающую звук, который издает животное.

Функция MakeAnimalSpeak принимает параметр типа Animal. Поскольку Dog и Cat реализуют интерфейс Animal, они могут быть переданы в эту функцию. Внутри функции вызывается метод Speak, который возвращает строку, и эта строка выводится на экран.
Aункция MakeAnimalSpeak принимает параметр типа Animal. Поскольку Dog и Cat реализуют интерфейс Animal, они могут быть переданы в эту функцию. Внутри функции вызывается метод Speak, который возвращает строку, и эта строка выводится на экран.

```go
type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("Собака %s лает", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("Кошка %s мяукает", c.Name)
}

// Функция, принимающая интерфейс Animal и вызывающая метод Speak
func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	dog := Dog{Name: "Шайтан"}
	cat := Cat{Name: "Тайсон"}

	// Вызов функции MakeAnimalSpeak для каждого животного
	MakeAnimalSpeak(dog)
	MakeAnimalSpeak(cat)
}
```

### Type Assertion

Type Assertion позволяет нам вызывать только те методы, которые существуют у данного конкретного типа. Она возвращает две переменные: значение конкретного типа и флаг, указывающий на успешность преобразования значения интерфейсного типа в конкретный тип. После успешного преобразования мы можем вызывать методы, специфичные для этого конкретного типа.

Для каждой структуры `Dog` и `Cat` мы добавили уникальные методы:
```go
func (d Dog) Bark() string {
	return fmt.Sprintf("%s громко лает!", d.Name)
}

func (c Cat) Purr() string {
	return fmt.Sprintf("%s мурлычет.", c.Name)
}
```

Создадим функцию `processAnimalTypeAssertion`, которая принимает на вход интерфейс `Animal` и выполняет `type assertion` для проверки конкретного типа, чтобы вызвать его уникальные методы:

```go
func processAnimalTypeAssertion(animal Animal) {
	if dog, ok := animal.(*Dog); ok {
		fmt.Printf("Type: %T Value: %#v\n", dog, dog)
		fmt.Println(dog.Bark())
	}
	if cat, ok := animal.(*Cat); ok {
		fmt.Printf("Type: %T Value: %#v\n", cat, cat)
		fmt.Println(cat.Purr())
	}
}
```

Вызовем функцию `processAnimalTypeAssertion` в `main`.

```go
func main() {
	dog := &Dog{Name: "Шайтан"}
	cat := &Cat{Name: "Тайсон"}

	processAnimalTypeAssertion(dog)
	processAnimalTypeAssertion(cat)
}
```

При выполнении этого кода получим следующий вывод:

```go
Type: *main.Dog Value: &main.Dog{Name:"Шайтан"}
Шайтан громко лает!
Type: *main.Cat Value: &main.Cat{Name:"Тайсон"}
Тайсон мурлычет.
```

### Type Switch

Type Switch предоставляет синтаксический сахар для работы с Type Assertion. Таким образом можем заменить функцию processAnimalTypeAssertion функцией processAnimalTypeSwitch:

```go
func processAnimalTypeSwitch(animal Animal) {
	switch v := animal.(type) {
	case *Dog:
		fmt.Printf("Type: %T Value: %#v\n", v, v)
		fmt.Println(v.Bark())
	case *Cat:
		fmt.Printf("Type: %T Value: %#v\n", v, v)
		fmt.Println(v.Purr())
	default:
		fmt.Printf("Type: %T Value: %#v\n", v, v)
	}
}
```

### [Где лучше размещать интерфейс?](https://www.youtube.com/watch?v=eYHCCht8eX4)

Небольший спойлер: **Интерфейсы лучше размещать в месте их использования.**

Рекомендации по использованию интерфейсов:
1. **Интерфейсы должны быть минималистичными.**
2. **Интерфейс ничего не должен знать о типах, которые его реализуют.**

Рассмотрим пример некого сервиса.
В этом сервисе нас интересует два слоя: `storage` и `handlers`. В слое `storage` есть пакет `users`, в котором указаны методы для различных БД: `Postgres`, `Redis`, `MySQL` и т.д.

#### Структура проекта

- **`some_service/`**: Главная директория сервиса.
  - **`handlers/`**: Папка для хэндлеров (обработчиков).
    - **`createuser/`**: Подпапка для создания пользователей.
    - **`userinfo/`**: Подпапка для работы с информацией о пользователях.
  - **`lib/`**: Библиотеки и вспомогательные модули.
  - **`services/`**: Логика и бизнес-слой сервиса.
  - **`storage/`**: Хранилище данных.
    - **`users/`**: Папка для управления пользователями.
      - Подпапки для различных реализаций хранения, такие как:
        - **`cache/`**
        - **`mysql/`**
        - **`postgres/`**
        - **`redis/`**
      - **`users.go`**: Файл с реализацией интерфейсов для работы с пользователями. 

Чтобы не зависеть от типа реализации мы решили описать общий интерфейс, который называется `Storage`. В общем интерфейсе содержаться все методы, которые необходимы для взаимодействия со всеми БД.

```go
package users

type User struct {
    ID   int
    Name string
    Age  int
}

type Storage interface {
    Users() ([]User, error)
    UsersByAge(age int) ([]User, error)
    User(id int) (User, error)
    Create(user User) error
    Update(user User) error
    Delete(id int) error
    // другие методы...
}
```

Кажется, что мы соблюдаем правило №2 - интерфейс ничего не должен знать о типах, которые его реализуют, но это не совсем так.

Мы используем данный интерфейс в слой `handlers`. В частности, в нем у нас есть функция `New`, которая принимает интерфейс `Storage`, в котором множество методов.

```go
func New(userRepo users.Storage) handlers.Handler {
    return func(ctx context.Context) {
        // Получение UID из запроса
        uid := 1
        user, err := userRepo.User(uid)
        if err != nil {
            // Обработка ошибки
        }
    }
}
```

Чтобы не тащить за собой огромный интерфейс с кучей методов, мы можем описать интерфейс в месте его использования, прямо в этом хендлере. В этом хендлере мы используем один единственный метод `User()`. Это значит, что мы можем создать здесь интерфейс, в котором будет необходимый метод. Создадим интерфейс `UserProvider`, в котором будет метод `User()`.

```go
type UserProvider interface {
    User(int) users.User
}
```

#### Что нам это дает?

- **Минималистичный интерфейс**: В нашем методе нет намека на какую-либо базу данных. Метод `User()` просто каким-то образом возвращает пользователя.
- **Уменьшение связности**: Пакет `handlers` никак не зависит от пакета `storage`. Связность компонентов системы должна быть как можно меньше.
- **Понятность кода**: Мы сделали ясными ожидания и потребности разных частей системы. При чтении кода в пакете `handlers`, мы видим, какой интерфейс ожидает функция, и этот интерфейс описан в этом же пакете. При использовании большого интерфейса `Storage`, мы видим, что у него много методов, и не сразу понятно, для чего они нужны. Также, чтобы прочитать описание интерфейса, нужно перейти в другой пакет.
- **Гибкость системы**: Допустим, мы хотим передать вместе с сущностью `Postgres` сущность `Redis`. Чтобы соответствовать данному интерфейсу, нам придется реализовать все его методы, даже если они не используются.
- **Тестирование**: При написании юнит-тестов для тестирования логики функции нам необходимо изолироваться от какой-либо базы данных. Mock'и позволяют нам этого добиться. Поскольку интерфейс описан в пакете `handlers`, то и сгенерировать mock мы можем в этом же пакете.

#### Минусы подхода

- Дублирование описания интерфейса по всем частям сервиса. Если мы захотим изменить сигнатуры методов, то нам придется это сделать во всех частях системы. Если бы у нас был один общий интерфейс, достаточно было бы изменить его только в одном месте.

- Новичкам из других языков не всегда понятен такой подход, который связан с утиной типизацией и неявной имплементацией интерфейсов в Go.

#### Принципы SOLID

Также хочется отметить, что придерживаясь такого подхода, мы соответствуем следующим принципам SOLID:

- **Принцип разделения интерфейсов (I)**: **Программные сущности не должны зависеть от методов, которые они не используют.**  
Разделение одного большого интерфейса на несколько мелких.

- **Принцип инверсии зависимостей (D)**: **Модули верхних уровней не должны зависеть от модулей нижних уровней. Оба типа модулей должны зависеть от абстракций. Абстракции не должны зависеть от деталей, детали должны зависеть от абстракций.**  
Создание абстракций, которые позволяют модулям взаимодействовать без прямой зависимости друг от друга.

### Цитата про необходимость разделения одного большого интерфейса на несколько мелких из статьи про постулаты Go ([Go proverbs](https://habr.com/ru/articles/272383/))

#### Чем больше интерфейс, тем слабее абстракция (The bigger the interface, the weaker the abstraction)

Новички в Go, особенно пришедшие с Java, часто считают, что интерфейсы должны быть большими и содержать много методов. Также часто их смущает неявное удовлетворение интерфейсов. Но самое важное в интерфейсах не это, а культура вокруг них, которая отражена в этом постулате. Чем меньше интерфейс, тем более он полезен. Пайк шутит, что три самых полезных интерфейса, которые он написал — `io.Reader`, `io.Writer` и `interface{}` — на троих в среднем имеют 0.666 метода.

### Полезные материалы

[The Go Programming Language Specification](https://go.dev/ref/spec#Interface_types)

[Effective Go - interfaces](https://go.dev/doc/effective_go#interfaces_and_types)

[Почему интерфейсы лучше размещать в месте использования - GoLang best practices | Николай Тузов](https://www.youtube.com/watch?v=eYHCCht8eX4)

[Go Data Structures: Interfaces](https://research.swtch.com/interfaces)

[Go Proverbs](https://go-proverbs.github.io/)

[Practical SOLID in Golang](https://github.com/MaksimDzhangirov/practicalSolid)
