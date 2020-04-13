#!/usr/bin/env python3
import os
import tensorflow_io as tfio

# Configure a connection to Postgres (env vars pre-populated)
postgres_endpoint = "postgresql://{}:{}@{}?port={}&dbname={}".format(
    os.environ['DB_USER_NAME'],
    os.environ['DB_PASSWORD'],
    os.environ['DB_HOSTNAME'],
    os.environ['DB_PORT'],
    os.environ['DB_NAME'],
)

# Load the dataset from Postgres
dataset = tfio.experimental.IODataset.from_sql(
    query = 'SELECT SUM(deaths) AS "deaths", SUM(confirmed) AS "confirmed", SUM(population) AS "population", date FROM daily_regions WHERE region_code = '' GROUP BY date ORDER BY date ASC;',
    endpoint = postgres_endpoint)

print(dataset.element_spec)
