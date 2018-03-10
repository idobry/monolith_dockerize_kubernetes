FROM postgres:9.6

COPY ./postgresql.conf /usr/share/postgresql/postgresql.conf

COPY ./pg_hba.conf /var/lib/postgresql/pg_hba.conf

ENV POSTGRES_DB best_browser

COPY db.sql /docker-entrypoint-initdb.d/