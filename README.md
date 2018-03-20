# Data integration challenge


Welcome to Data API challenge.

Yawoen company has hired you to implement a Data API. 


**1 - Load company data in a database**

Read data from CSV file and load into the database to create an entity named **companies**.

This catalog should contain the following fields: id, company name and zip code. 

suport file: q1_catalog.csv


**2 - An API to integrate data using a database**

Yawoen company started to capture website data from another source and want to integrate it with the entity you've just created on the database.

This new source data has the following input protocol:

- Input file format: CSV

- Parameters:

    - Name: string
    - Zip: string 
    - Website: string

Build an API to integrate **website data** records into the entity records you've just created.

The data source doesn't provide the id field, so you'll have to use the available fields to aggregate the new attribute **website** and store it.

suport file: q2_clientData.csv


**Extra - Matching API to get data based on specified parameters**

Now a Yawoen client wants to access an API to get companies information. They have name and zip code fields.

You will need to have a matching strategy because the client only has part of the company name, for example "Yawoen Business Solutions" is captured as "Yawoen" in the source.

### Notes


- More than one instance of the application will serve HTTP requests at the same time.
- Make sure other programmers can easily run the application locally.
- Yawoen doesn't care about the language, the database and other tools that you might choose however make sure your decisions take notice of the market. For example, Brainfuck might not be a good decision.
- Automated tests are mandatory.
- Document your API if you are willing to.


### Deliverable


A git repo hosted wherever you like.
Make sure Yawoen folks will have access to the source code.
If you prefer, just compress this directory and send it back to us.

Have fun!
