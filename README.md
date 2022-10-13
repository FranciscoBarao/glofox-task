# glofox-task
Repository for glofox task

## Project Structure
Based on the [Repository Pattern](https://dakaii.medium.com/repository-pattern-in-golang-d22d3fa76d91) which intends to asbtract the different layers.  
In this specific implementation of the repository pattern, the repository serves as a bridge between the database logic and the controller which due to the simplicity of the task also includes the business logic (Service).    

Flow: 
- User makes requests to Controller
- Controller calls Repository
- Repository calls specific Database implementation
- Database Implementation connects to Database

### Improvements:
Implementation of a Service layer that would abstract business logic creating the flow ```controller -> service -> repository -> dbImplementation -> db```
