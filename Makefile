createmigration:
	migrate create -ext=sql -dir=config/database/migrations -seq init

migrate: 
	migrate -path=config/database/migrations -database "postgresql://root:123@localhost:5432/picpay?sslmode=disable" -verbose up

migratedown:
	migrate -path=config/database/migrations -database "postgresql://root:123@localhost:5432/picpay?sslmode=disable" -verbose down

.PHONE: migrate migratedown createmigration