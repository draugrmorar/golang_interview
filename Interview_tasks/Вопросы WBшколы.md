# 📚 СВОДНЫЙ ПЛАН ПОДГОТОВКИ К ЭКЗАМЕНАМ

## 📌 Навигация
- [ЭКЗАМЕН 1](#экзамен-1)
- [ЭКЗАМЕН 2](#экзамен-2)
- [Практические задачи](#практические-задачи)

---

# ЭКЗАМЕН 1

## 🧠 1. Go Runtime & Планировщик (GMP)

### Модель GMP
- Что такое G (Goroutine), M (Machine/Thread), P (Processor)
- Сколько горутин можно взять из одной очереди
- Как процессы могут передавать сообщения
- Различие handoff и netpoller
- Размер локальной очереди P (тип очереди)
- Что такое `GOMAXPROCS`

### Планировщик
- Статусы горутин (Running, Runnable, Waiting, etc.)
- Netpoller — что это и как работает
- Как понять, что горутина готова к выполнению
- `runtime.Gosched()` — что делает
- Safe point — что это и зачем нужен

### Sys.Monitor
- Методы Sys.Monitor
- Что делает системный монитор
- Как он влияет на планировщик

### Stack
- Stack frame — что это
- Stack trace — зачем нужен
- Размер стека горутины (начальный и динамический)

### Escape Analysis
- Что такое escape analysis
- Как получить информацию об escape analysis
- Влияние на производительность

### Spans & Arena
- Span в arena (экспериментальная функция)
- Что такое arena в Go

---

## 🔄 2. Garbage Collector (GC)

### Основы
- Как работает GC в Go
- Что может замедлить сборку мусора
- GC, call optimization, timeout, limit
- Safe point

### Stop-The-World (STW)
- Сколько длится STW
- Какие фазы GC требуют STW
- Как уменьшить STW

---

## 🔧 3. Примитивы синхронизации

### Mutex
- Устройство под капотом (state, sema)
- Что такое state и sema
- Почему не рекомендуется TryLock
- Что делает `noCopy` и почему он первым полем

### RWMutex
- Отличие от Mutex
- Когда использовать

### WaitGroup
- Устройство и принцип работы
- Правила использования (Add до go)

### Cond
- Что делает `sync.Cond`
- Когда использовать
- Какие проблемы решает
- В чем отличие от каналов

### Pool (`sync.Pool`)
- Что такое и зачем нужен
- Как устроен (local, victim)
- Механизм Victim Cache
- Проблемы с растущими объектами

### Once
- `sync.Once` — что делает
- `sync.OnceValue` (Go 1.21+)

### Map (`sync.Map`)
- Как устроен `sync.Map`
- Почему принимает `interface{}`
- Что такое `readOnly` map и `dirty` map
- Когда происходит promotion (продвижение)
- Что изменилось в новой реализации

### ErrGroup (`errgroup`)
- Что такое `errgroup`
- Отличие от WaitGroup

### Semaphore
- Что такое семафор
- Когда использовать

### Каналы (Channels)
- Устройство канала
- Буферизированные vs небуферизированные
- select

---

## 🧵 4. Concurrency Patterns

### Базовые паттерны
- Worker pool
- Fan-in (схождение)
- Fan-out (разветвление)
- Pipeline
- Future
- Singleflight

### Примитивы
- Реализовать SpinLock (busy-wait)
- Реализовать `sync.Once`
- Реализовать Future
- Реализовать Singleflight

---

## 🛡️ 5. Race Conditions

- Race condition vs Data race — в чём отличие
- Что делает флаг `-race`
- Как работает `-race` (инструментация кода)
- Почему бинарник с `-race` медленный
- Почему внутри mutex есть проверки race
- Как предотвращать data race

---

## 🌐 6. Сети и протоколы

### OSI / TCP/IP
- Модель OSI — 7 уровней
- Где используется OSI
- Что происходит на каждом уровне
- Чем отличается OSI от TCP/IP
- Как летят данные по сети

### HTTP
- HTTP 1.1 — особенности
- HTTP/2 — сколько потоков
- Заголовки

### gRPC
- Как в gRPC передавать context
- Мультиплексирование в gRPC
- Интерсепторы (middleware) в gRPC
- Поверх какого протокола работает gRPC
- Кому нужен protobuf — клиенту или серверу
- Как дальше сериализуются данные

---

## 🏗️ 7. Архитектура и Паттерны

### SOLID
- Interface segregation — разделение на интерфейсы
- Принципы SOLID в Go
- Экспортируемость: с каких букв должны начинаться названия структур и полей

### Паттерны проектирования
- **Circuit Breaker** — что это, состояния (Closed, Open, Half-Open)
- **Retry** — с Backoff
- **Timeout**
- **Bulkhead** — изоляция ресурсов
- **Rate Limiter** — отличие от Throttling
- **Dead Letter Queue** (DLQ)
- **Fallback** — что делать при ошибке
- **Hedging** — отправка запроса на несколько серверов

### Структурные паттерны (Go)
- Adapter
- Decorator
- Facade

### Дополнительно
- Middleware в Go — что это
- AI Health Check
- Event Sourcing — что это

---

## 📦 8. Context

- `context.Context` — интерфейс и набор методов
- `Done()` — канал завершения
- `Err()` — ошибка завершения
- `Deadline()` — дедлайн
- Как в gRPC передавать context
- Какие данные передаются через context
- Scope данные (контекстные значения)
- Ограничения на передачу данных в context

---

## 📝 9. Память и Zero-Value

- Zero-value в Go (для всех типов)
- `nil` в структурах
- `for` в структурах + nil (вероятно про range по slice с nil)
- Escape analysis (повтор)
- Stack frame (повтор)

---

# ЭКЗАМЕН 2

## 🗄️ 1. Базы данных (Общие вопросы)

### Основы
- Что такое БД?
- Что такое СУБД?
- Чем БД отличается от СУБД?

### SQL vs NoSQL
- Сравнение
- Когда использовать SQL?
- Когда использовать NoSQL?

### Типы NoSQL БД
- Графовые БД
- Документные БД
- Key-Value БД
- Column Family БД

---

## 🔗 2. Связи и Модели данных

### Типы связей
- Один к одному (1:1)
- Один ко многим (1:N)
- Многие к одному (N:1)
- Многие ко многим (M:N)

### Реализация
- Как реализуются отношения в SQL?
- Для чего нужен внешний ключ (Foreign Key)?
- Зачем нужен surrogate key (id)?

---

## 📊 3. Нормализация

- Что такое нормализация?
- Что такое денормализация?
- Когда денормализация оправдана?

### Нормальные формы
- Первая нормальная форма (1NF)
- Вторая нормальная форма (2NF)
- Третья нормальная форма (3NF)

---

## 📈 4. Индексы

### Основы
- Что такое индекс?
- Как работает B-Tree?
- Как PostgreSQL ищет данные по индексу?
- Почему индекс ускоряет поиск?

### Типы индексов (PostgreSQL)
- B-Tree
- Hash
- GIN
- GiST
- BRIN
- INCLUDE Index
- Covering Index
- Multi-column Index

### Оптимизация
- Что такое селективность?
- Как оптимизатор выбирает индекс?
- Когда индекс НЕ используется?
- Когда Seq Scan быстрее индекса?

### Побочные эффекты
- Почему индекс может ухудшить INSERT?
- Почему индекс может ухудшить UPDATE?
- Почему индекс может ухудшить DELETE?

### Создание индексов
- Как добавить индекс без блокировки таблицы?
- Что делает `CREATE INDEX CONCURRENTLY`?

---

## 🐘 5. PostgreSQL (Внутреннее устройство)

### WAL и согласованность
- Что такое WAL (Write-Ahead Log)?
- Почему WAL обеспечивает согласованность?
- Что происходит при COMMIT?
- Что происходит после сбоя?

### Фоновые процессы
- Что такое Checkpoint?
- Что делает Background Writer?

### MVCC
- Что такое MVCC?
- Почему MVCC лучше блокировок?
- Что такое tuple?
- Что такое dead tuple?
- Почему появляются dead tuple?

### VACUUM и ANALYZE
- Что делает VACUUM?
- Что делает VACUUM FULL?
- Что делает ANALYZE?
- Что делает AUTOVACUUM?

### Explain
- Чем Explain отличается от Explain Analyze?
- Какие полезные флаги есть у Explain Analyze?
- Что показывает стоимость (cost)?

### Типы сканирования
- Что такое Seq Scan?
- Что такое Index Scan?
- Что такое Bitmap Scan?
- Что такое Index Only Scan?

### Template
- Что такое Template0?
- Что такое Template1?

---

## 🔄 6. Транзакции

### ACID
- Что такое транзакция?
- Atomicity (Атомарность)
- Consistency (Согласованность)
- Isolation (Изоляция)
- Durability (Долговечность)

### Уровни изоляции
- Read Uncommitted
- Read Committed
- Repeatable Read
- Serializable

### Аномалии
- Dirty Read (Грязное чтение)
- Non-repeatable Read (Неповторяемое чтение)
- Phantom Read (Фантомное чтение)

### Snapshot Isolation
- Что такое Snapshot Isolation?
- Чем отличаются снапшоты Read Committed и Repeatable Read?

### Блокировки
- Что делает `SELECT FOR UPDATE`?

---

## 📊 7. Масштабирование

### Партиционирование
- Что такое партиционирование?
- Горизонтальное партиционирование
- Вертикальное партиционирование

### Шардирование
- Что такое шардирование?
- Виртуальное шардирование
- Когда выбирать шардирование?
- Когда выбирать партиционирование?

### Connection Pool
- Что такое Connection Pool?
- PgBouncer — зачем нужен

---

## 📝 8. SQL

### JOIN
- INNER JOIN
- LEFT JOIN
- RIGHT JOIN
- FULL JOIN
- CROSS JOIN

### Подзапросы
- CTE (Common Table Expression) — `WITH`
- Подзапросы
- Correlated Subquery (коррелированный подзапрос)
- EXISTS / NOT EXISTS

### Агрегация
- GROUP BY
- HAVING
- Window Functions (Оконные функции)

### Задачи
- Написать SQL без оконных функций
- Написать SQL с CTE
- Написать тяжёлый SQL-запрос

---

## 📨 9. Kafka

### Архитектура
- Что такое Kafka?
- Broker
- Topic
- Partition
- Offset
- Consumer Group
- Producer
- Consumer

### Producer
- Продюсер умный или глупый?
- Как выбирается partition?
- Что можно настроить у producer?
- Idempotent Producer
- Transactions

### Consumer
- Что можно настроить у consumer?
- Что происходит при чтении без Group ID?
- Что происходит при чтении с Group ID?

### Delivery Guarantees
- At Most Once
- At Least Once
- Exactly Once
- Как реализован Exactly Once?

### Rebalance
- Что такое Rebalance?
- Когда начинается Rebalance?
- Какие есть алгоритмы?
- Sticky Assignor
- Cooperative Rebalance

### Партиции
- Что если consumer больше partition?
- Что если partition больше consumer?
- Что если один consumer умер?

---

## 🛡️ 10. Отказоустойчивость

### Circuit Breaker
- Что такое Circuit Breaker?
- Состояния: Closed, Open, Half-Open
- Реализации:
    - Circuit Breaker на счётчике ошибок
    - Circuit Breaker на скользящем окне

### Другие паттерны
- Retry with Backoff
- Timeout
- Bulkhead
- Rate Limiter (отличие от Throttling)
- Throttling
- Dead Letter Queue
- Fallback
- Hedging

---

## 📤 11. Transaction Outbox

- Что такое Transaction Outbox?
- Какие проблемы решает?
- Почему нельзя писать в Kafka прямо из транзакции?
- Как работает Outbox?

### CDC
- Что такое CDC (Change Data Capture)?
- Что такое Debezium?
- Когда использовать CDC вместо polling?

---

## 🧵 12. Go Concurrency (повтор из Экзамена 1)

- Mutex, RWMutex, WaitGroup, Once, Cond, Pool, Map
- `sync.Map` — устройство
- atomic — Atomic Pointer, Atomic Value, CompareAndSwap, Load, Store, Swap, Add
- Mutex — устройство под капотом
- Cond — когда использовать

---

## 🔍 13. Race Detector (повтор)

- Race Condition vs Data Race
- Флаг `-race`
- Как работает `-race`
- Бинарник с `-race` медленный

---

## ⚙️ 14. Runtime (повтор)

- Scheduler GMP
- GOMAXPROCS
- Что изменилось в Go 1.25?
- `runtime.Gosched()`
- SpinLock
- Escape Analysis

---

## 🏗️ 15. Структуры данных

- `container/heap`
- Priority Queue (Очередь с приоритетом)
- Реализовать собственную кучу
- Стриминговое слияние нескольких отсортированных каналов

---

## 📡 16. Сети (повтор)

- Модель OSI
- TCP/IP
- Чем отличается OSI от TCP/IP

---

## 🔴 17. Redis

- Что такое Redis?
- Когда использовать Redis?
- Какие структуры данных поддерживает Redis?
- Где Redis лучше PostgreSQL?

---

# ПРАКТИЧЕСКИЕ ЗАДАЧИ

## SQL (написать)
- SQL по условиям
- SQL без оконных функций
- SQL с подзапросами
- SQL с CTE
- Оптимизировать запрос по EXPLAIN

## Go (реализовать)
- `sync.Once`
- Circuit Breaker
- Rate Limiter
- Transaction Outbox
- Future
- Singleflight
- SpinLock
- HTTP-ручка
- Использовать WaitGroup
- Использовать Mutex
- Использовать ErrGroup
