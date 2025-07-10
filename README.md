# CacheCluster

CacheCluster - это открытая класстерная система, разработанная на Golang, которая обеспечивает единую систему хранения данных в микросервисной архитектуре.
## Запуск
### Нативный запуск
1. `git clone https://github.com/thisisfunt/CacheCluster.git`
2. `cd CacheCluster/storage`
3. Создайте и зарегистрируйте свои структуры данных согласно примерам в файле "model/structs.go"
4. При необходимости, создайте и зарегистрируйте свои функции согласно примерам в файле "model/operations.go"
5. `go run main.go`