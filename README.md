# Library API excercise

**Library API excercise** is a golang implement(using go-kit, gorm) of library API

Submitted by: **{{github_user_name}}**

Time spent: **{{time_spent}}** hours spent in total

## Prerequisite

* Must Read about [git flows](https://nvie.com/posts/a-successful-git-branching-model)
* Should read about Postgres database, [go-kit](https://gokit.io/examples/stringsvc.html), [GORM](http://doc.gorm.io/)
* Already Setup: `git`, `golang`, [docker](https://docs.docker.com/install), [docker-compose](https://docs.docker.com/compose/install/#install-compose)

## How to start

* `make local-env`: To create local DB(port: 5432), test DB(port:5439), adminer(tools to view database, port:8080)
* `make clean-local-env`: Turn off `local-env` (be careful it will also clear DB)
* `make dev`: To start server (default port is 3000)
* `make test`: To run test (both integration and unit test)

## User Stories

### Required:
* NOTE:  Each main point should follow `git flows` (each step should have a feature branch,...)

#### Stage 1:

* [x] API must be able to CRUD users:
  * [x] Each user should have the following fields:
    * id (uuid)
    * name (string)
    * email (string)
  * [x] API must be able to get detail of a user.
  * [x] API must be able to get list of users
  * [x] API must be able to create a user
  * [x] API must be able to update a user
  * [x] API must be able to delete a user
* [x] Validate user is correct before (Create/Update):
  * [x] validate name of user is not empty.
  * [x] validate email of user is not empty.
  * [x] validate email of user is a valid email.
  
* [x] API must be able to CRUD category of books:
  * [x] Each category should have the following fields:
    * id (uuid)
    * name (string)
  * [x] API must be able to get detail of a category.
  * [x] API must be able to get list of categories.
  * [x] API must be able to create a category.
  * [x] API must be able to update a category.
  * [x] API must be able to delete a category.
  * [x] when delete categories all book belongs to that categories should deleted too.
  
* [x] Validate category is correct before (Create/Update):
  * [x] validate name of category is not empty and length > 5 characters.
  * [x] validate name of category not existed yet (for both create and update).

* [x] API must be able to CRUD books:
  * [x] Each book should have the following fields:
    * id (uuid)
    * name (string)
    * category_id (uuid)
    * author (string)
    * description (string)
  * [x] API must be able to get detail of a book.
  * [x] API must be able to get list of books.
  * [x] API must be able to create a book.
  * [x] API must be able to update a book.
  * [x] API must be able to delete a book.
  
* [x] Validate books is correct before (Create/Update):
  * [x] validate category of a book is exist, if not reject it with error message
  * [x] validate name of a book is not empty and length > 5 characters. if not reject it with error message
  * [x] validate description of a book is not empty and length > 5 characters. if not reject it with error message
  
* [x] API must be able to make action: user lend a book:
  * [x] API should have the following fields:
    * id (uuid)
    * book_id (uuid)
    * user_id (uuid)
    * from (datetime)
    * to (datetime)
  * [x] validate book_id of a book is exist, if not reject it with error message
  * [x] validate user_id of a user is exist, if not reject it with error message
  * [x] validate books is available to lend, if not available reject with error message
  
  
#### Stage 2:

* [x] Add testing for CRUD users:
  * [x] Validation testing (unit test)
  * [x] Endpoint testing (unit test)
  * [x] Database testing (integration)
* [ ] Add testing for CRUD category:
  * [ ] Validation testing
  * [ ] Database testing
* [ ] Add testing for CRUD books:
  * [ ] Validation testing
  * [ ] Database testing
* [ ] Add testing for lending books:
  * [ ] Validation testing
  * [ ] Database testing

### Optional:

*  [ ] List books and filter by name, availables status.
*  [ ] batch create books (create multiple book with 1 API)
*  [ ] batch lending books (user can lending multiple books with 1 api)
*  [ ] implement feature add a tags to books can search book by tag name
*  [ ] implement multiple errors return by an array


The following **additional** features are implemented:

* [ ] .....

The following **known issues**:

* ...

## Video Walkthrough

Here's a walkthrough of implemented user stories:


## Notes

Notes about current git.

## License

    Copyright [2018] [{{github_user_name}}]

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
