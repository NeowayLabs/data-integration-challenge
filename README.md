# Data integration challenge


Welcome to Data API challenge.

Yawoen company has hired you to implement a Data API. 


**1 - Load company data in a database**

Read data from CSV file and load into the database to create a catalog named **companies**.

This catalog should contain the following fields: id, company name and zip code. 

suport file: q1_catalog.csv


**2 - API to integrate data using a database**

Yawoen company started to capture website data and want to integrate with your database.

This new source data has the following protocol:

- Response format: JSON

- Parameters:

    - Name: string
    - Zip: string 
    - Website: string

Example: 

```
{
	"name": "Yawoen Business Solutions",
	"zip":"10023",
	"website": "www.yawoen.com"
}
``` 

Build an API to integrate **website data** records using the catalog you just created.

The data source doesn't provide the id field, so you'll have to use the available fields to aggregate the new attribute **website** and store it.

suport file: q2_clientData.csv


**Extra - Matching API to get data based on specified parameters**

Now a Yawoen client wants to access a API to get companies information. They have name and zip code fields.

You will need to have a matching strategy because the client have only part of the company name, for example "Yawoen Business Solutions" is captured as "Yawoen" in the source.


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

