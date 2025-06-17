remigrate:
	@goose sqlite3 ./test_db.sql -dir "./internal/database/migrations" reset
	@goose sqlite3 ./test_db.sql -dir "./internal/database/migrations" up
