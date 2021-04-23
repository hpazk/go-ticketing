# Ticketing

Framework: Echo
RDBMS: Postgres
ORM: GORM
IN transactions ON events.id = transactions.id
    JOIN users ON transactions.participant_id = users.id;


1q




Cache: Redis
