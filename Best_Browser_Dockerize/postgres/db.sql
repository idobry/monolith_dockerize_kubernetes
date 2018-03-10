
CREATE ROLE bb WITH SUPERUSER;
ALTER ROLE bb WITH LOGIN;



CREATE TABLE votes(
    firefox integer,
    chrome integer,
    explorer integer
);

INSERT INTO votes (firefox, chrome, explorer) VALUES (0, 0, 0);

ALTER TABLE votes OWNER TO bb;

