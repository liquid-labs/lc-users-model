module github.com/Liquid-Labs/lc-users-model

require (
	github.com/Liquid-Labs/lc-entities-model v1.0.0-alpha.0
	github.com/Liquid-Labs/lc-rdb-service v1.0.0-alpha.1
	github.com/Liquid-Labs/strkit v0.0.0-20190818184832-9e3e35dcfc9c
	github.com/Liquid-Labs/terror v1.0.0-alpha.0
	github.com/go-pg/pg v8.0.5+incompatible
	github.com/go-pg/pg/v9 v9.0.0-beta.7
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.4.0
)

replace github.com/Liquid-Labs/lc-entities-model => ../lc-entities-model

replace github.com/Liquid-Labs/terror => ../terror

replace github.com/Liquid-Labs/lc-rdb-service => ../lc-rdb-service
