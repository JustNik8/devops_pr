migrateup:
	migrate -path migrations -database "postgresql://user:user@localhost:5440/ecom_db?sslmode=disable" --verbose up 

migratedown:
	migrate -path migrations -database "postgresql://user:user@localhost:5440/ecom_db?sslmode=disable" --verbose down