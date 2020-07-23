
build:
	docker build -t flicker-api .

up:
	docker-compose up -d posgresdb_flicker

backup-db:
	docker-compose -f docker-compose.yml run --rm -T backup-db | pv > backup.tar.bz2
	# docker-compose -f docker-compose.yml run --rm -T backup-db

restore_db:
	cat backup.tar.bz2 | pv | docker-compose -f docker-compose.yml run --rm -T restore-db


clean:
	docker-compose stop
	docker-compose rm -f

backup:
	pg_dump -d 'micky' -U 'postgres' -h 'localhost' -W  > db.sql

restore:
	psql -h 'localhost' -U 'postgres' -d micky -1 -f db.sql

run:
	cd src && modd

git:
	git add .
	git commit -m "$m"
	git push -u origin master