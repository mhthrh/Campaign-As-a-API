# Dashboard Services

Dashboard services has been designed and implemented for Adding Campaigns and Events in run time mode without changing code. If one of the top ten user' scores changes, a request will fire to dashboard UI with a DASHBOARD new object.

## Platforms
Golang 1.17  
PostgreSQL 14.1  
HTTP REST service

## Running
Change APP config in [Here](.\Utilitys\ConfigUtil\ConfigUtil.go) Change Database configuration and set IP, Port for listener
Use bellow command for a run or build project in the main directory



Import [SQL](.\Script\SQl-Scripts.sql)
```bash
go run ./main.go
go build ./main.go
```
## Solution description
To design a scoreboard and display ten users with the highest scores Dashboard Services has been independently designed and implemented.
In this regard, the following databases can be used.  

   1- SQL database  
   2- No SQL Database

SQL database has been used in this project for the following reasons:

   1- Maintain transaction records consistently.  
   2- Ease of aggregation operations  

In this project, for each transaction registration operation, the customer's score is calculated based on the campaign, total, month and two weeks, and the result is stored in memory as a hash. After the next transaction, the hash score of 10 people with the highest score is calculated and compared with the previous hash. In the event of inequality, the UI service will be called with new values.

Another method can be used to save transactions and points.
It Saves scores as transactions and keeps the total points per person in each campaign, total, month and two weeks. This method is better in terms of performance, but it may cause a DB LOCK as a result.

      First, campaigns as well as events should be registered in the Database, with API or directly.
      Total periods, months and years can be taken as parameters from the database. This issue has not been implemented in this version.
      If we want to change the length of the course for example changing from a month course to a week course, the program can easily calculate the points. In fact, the time is configurable.

## Services Description
1. Add Score
   * Description: Core must call this method according score policy
   * URL: /scores
   * Method: POST
   * Request: [request](.\Campaign\Model\Input\Input.go)
   * Response: [response](.\Campaign\Model\Result\result.go)
2. Add Campaign
   * Description: system admin must call this method for create new Campaign
   * URL: /campaign
   * Method: POST
   * Request: [request](.\Campaign\Model)
   * Response: [response](.\Campaign\Model\Result\result.go)
3. Add Event
    * Description: system admin must call this method for create new Event
    * URL: /event
    * Method: POST
    * Request: [request](.\Campaign\Model)
    * Response: [response](.\Campaign\Model\Result\result.go)
4. Update Campaigns and events
    * Description: If admin declare a campaign or event directly in database, services must be updated
    * URL: / update
    * Method: Get

We're imaging that `USER` was `Authenticate`,`Authorized`, and every policy require in core.

Some of these services must protect with `User, pass` or `another way such as routing.`


## Hint
Some methods are not implemented



## License
[Open Source](https://en.wikipedia.org/wiki/Open_source)