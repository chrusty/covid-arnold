covid-arnold
============
COVID-19 lockdown tracker powered by Arnold Technology


Components
----------

* Postgres DB: stores statistics & predictions
* Ingester: downloads data and inserts into the DB
* Forecast: makes predictions based on the current datasets
* Uploader: uploads generated JSON datasets to an S3 bucket
* Frontend: visualise datasets from the S3 bucket
