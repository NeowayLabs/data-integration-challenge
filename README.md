# Data integration challenge


Welcome to Data Integration challenge.

Yawoen company has hired you to implement a Data API for Data Integration team.

Data Integration team is focused on combining data from different heterogeneous sources and providing it to an unified view into entities.

## The challenge

### 1 - Load company data in a database

Read data from CSV file and load into the database to create an entity named **companies**.

This entity should contain the following fields: id, company name and zip code. 

support file: q1_catalog.csv


### 2 - An API to integrate data using a database

Yawoen now wants to get website data from another source and integrate it with the entity you've just created on the database. When the requirements are met, it's **mandatory** that the **data are merged**.

This new source data must meet the following requirements:

- Input file format: CSV
- Data treatment
    - **Name:** upper case text
    - **zip:** a five digit text
    - **website:** lower case text
- Parameters
    - Name: string
    - Zip: string 
    - Website: string

Build an API to **integrate** `website` data field into the entity records you've just created using **HTTP protocol**.

An id field is non existent on the data source, so you'll have to use the available fields to aggregate the new attribute **website** and store it. If the record doesn't exist, discard it.

support file: q2_clientData.csv


### Extra - Matching API to retrieve data

Now Yawoen wants to create an API to provide information getting companies information from the entity to a client. 
The parameters would be `name` and `zip` fields. To query on the database an **AND** logic operator must be used between the fields.

You will need to have a matching strategy because the client might only have a part of the company name. 
Example: "Yawoen" string from "Yawoen Business Solutions".

Output example: 
 ```
 {
 	"id": "abc-1de-123fg",
 	"name": "Yawoen Business Solutions",
 	"zip":"10023",
 	"website": "www.yawoen.com"
 }
 ```

## Notes


- Make sure other developers can easily run the application locally.
- Yawoen isn't picky about the programming language, the database and other tools that you might choose. Just take notice of the market before making your decision.
- Automated tests are mandatory.
- Document your API if you are willing to.


## Deliverable


- It would be REALLY nice if it was hosted in a git repo of your own. You can create a new empty project, create a branch and Pull Request it to the master branch. Provide the PR URL for us. BUT if you'd rather, just compress this directory and send it back to us.
- Make sure Yawoen folks will have access to the source code.
- Fill the **Makefile** targets with the apropriated commands (**TODO** tags).

Have fun!
