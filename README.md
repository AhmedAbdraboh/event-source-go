## Event Source Example

This is an example of how to use event source in a very easy application like create account, deposit and withdrawal from it. 

###How to use: 
- Install docker
- run 

    ```make up```
    
A server will start with exposed port 8080

to use it: 

- Send a POST request to 

    ``localhost:8080/accounts`` 
    
    
   to create an account and it will respond with the account id 
   
- Send a PUT request to 

    ``localhost:8080/accounts/1``
    
    with body
    
    ``{
      	"owner" : "ahmed"
      }``
      
     to update the owner of the account.
- Send a POST request to 

    ``localhost:8080/account/deposit``
    
    with body
    
    ``{
      	"accountId":1,
      	"amount": 150
      }``
      
    this will deposit an amount of 150 to the created account
    
- Send a GET request to 
    
    ``localhost:8080/accounts/1``
    
    this will get the current balance of the account with id=1
    
 
- Send a POST request to 

    ``localhost:8080/account/withdrawal``
    
    with body
    
    ``{
      	"accountId":1,
      	"amount": 50
      }``
      
    this will withdrawal an amount of 50 from the account with id=1
    
    
    
###Todo:

   - Make Unit Test
   - Make API Test