**Signup**
Creates a user on the database

**/signup** [POST]

**URL Parameters**
-

**Headers**
-

**Request Body**
{
	"username": "string",
	"password": "string",
	"refugeName": "string",
	"email": "string",
	"roleID": int,
	"lat": float,
	"lon": float
}

**Success 200 Response**
Signup successful

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
-

**Internal Server Error 500 Response**
Username already in use

-----------------------------------------------------
**Login**
Returns an access token for the provided user

**/login** [POST]

**URL Parameters**
-

**Headers**
-

**Request Body**
{
	"username": "string",
	"password": "string"
}

**Success 200 Response**
a63ab36162a4f4ee6622ccd787b0a048c26b93acfc05c6b1843659b253c3c00b //authentication token

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
-

**Internal Server Error 500 Response**
Wrong username or password

-----------------------------------------------------
**Create User**
Creates a user on the database

**/user** [POST]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"username": "string",
	"password": "string",
	"refugeName": "string",
	"email": "string",
	"roleID": int,
	"lat": float,
	"lon": float
}

**Success 200 Response**
User created successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get User list**
Gets a list of every user on the database

**/user/list** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"username": "string",
	"password": "string",
	"refugeName": "string",
	"email": "string",
	"roleID": int,
	"lat": float,
	"lon": float
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get User by Id**
Returns a user using the provided id

**/user/id/:id** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"username": "string",
	"password": "string",
	"refugeName": "string",
	"email": "string",
	"roleID": int,
	"lat": float,
	"lon": float
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Update User by Id**
Updates a user using the provided id

**/user/id/:id** [PUT]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"username": "string",
	"password": "string",
	"refugeName": "string",
	"email": "string",
	"roleID": int,
	"lat": float,
	"lon": float
}

**Success 200 Response**
User updated successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Delete User by Id**
Deletes a user using the provided id

**/user/id/:id** [DELETE]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
User deleted successfully

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get User by Username**
Returns a user using the provided username

**/user/username/:username** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"username": "string",
	"password": "string",
	"refugeName": "string",
	"email": "string",
	"roleID": int,
	"lat": float,
	"lon": float
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Update User by Username**
Updates a user using the provided username

**/user/username/:username** [PUT]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"username": "string",
	"password": "string",
	"refugeName": "string",
	"email": "string",
	"roleID": int,
	"lat": float,
	"lon": float
}

**Success 200 Response**
User updated successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Delete User by Username**
Deletes a user using the provided username

**/user/username/:username** [DELETE]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
User deleted successfully

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get User by Email**
Returns a user using the provided email

**/user/email/:email** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"username": "string",
	"password": "string",
	"refugeName": "string",
	"email": "string",
	"roleID": int,
	"lat": float,
	"lon": float
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Update User by Email**
Updates a user using the provided email

**/user/email/:email** [PUT]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"username": "string",
	"password": "string",
	"refugeName": "string",
	"email": "string",
	"roleID": int,
	"lat": float,
	"lon": float
}

**Success 200 Response**
User updated successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Delete User by Email**
Deletes a user using the provided email

**/user/email/:email** [DELETE]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
User deleted successfully

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Create Role**
Creates a role on the database

**/role** [POST]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"name": "string"
}

**Success 200 Response**
Role created successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get Role list**
Gets a list of every role on the database

**/role/list** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"name": "string"
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get Role by Id**
Returns a role using the provided id

**/role/id/:id** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"name": "string"
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Update Role by Id**
Updates a role using the provided id

**/role/id/:id** [PUT]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"name": "string"
}

**Success 200 Response**
Role updated successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Delete Role by Id**
Deletes a role using the provided id

**/role/id/:id** [DELETE]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
Role deleted successfully

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Create Photo**
Creates a photo on the database

**/photo** [POST]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"url": "string",
	"lat": float,
	"lon": float
}

**Success 200 Response**
Photo created successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get Photo list**
Gets a list of every photo on the database

**/photo/list** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"url": "string",
	"lat": float,
	"lon": float
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get Photo by Id**
Returns a photo using the provided id

**/photo/id/:id** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"url": "string",
	"lat": float,
	"lon": float
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Update Photo by Id**
Updates a photo using the provided id

**/photo/id/:id** [PUT]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"url": "string",
	"lat": float,
	"lon": float
}

**Success 200 Response**
Photo updated successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Delete Photo by Id**
Deletes a photo using the provided id

**/photo/id/:id** [DELETE]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
Photo deleted successfully

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Create RefugeAbility**
Creates a refugeAbility on the database

**/refugeAbility** [POST]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"name": "string"
}

**Success 200 Response**
RefugeAbility created successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get RefugeAbility list**
Gets a list of every refugeAbility on the database

**/refugeAbility/list** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"name": "string"
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get RefugeAbility by Id**
Returns a refugeAbility using the provided id

**/refugeAbility/id/:id** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"name": "string"
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Update RefugeAbility by Id**
Updates a refugeAbility using the provided id

**/refugeAbility/id/:id** [PUT]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"name": "string"
}

**Success 200 Response**
RefugeAbility updated successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Delete RefugeAbility by Id**
Deletes a refugeAbility using the provided id

**/refugeAbility/id/:id** [DELETE]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
RefugeAbility deleted successfully

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Create AuditorReview**
Creates a auditorReview on the database

**/auditorReview** [POST]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"idUser": int,
	"idRefuge": int,
	"idRefugeAbility": int,
	"score": int
}

**Success 200 Response**
AuditorReview created successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get AuditorReview list**
Gets a list of every auditorReview on the database

**/auditorReview/list** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"idUser": int,
	"idRefuge": int,
	"idRefugeAbility": int,
	"score": int
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get AuditorReview by Id**
Returns a auditorReview using the provided id

**/auditorReview/id/:id** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"idUser": int,
	"idRefuge": int,
	"idRefugeAbility": int,
	"score": int
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Update AuditorReview by Id**
Updates a auditorReview using the provided id

**/auditorReview/id/:id** [PUT]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"idUser": int,
	"idRefuge": int,
	"idRefugeAbility": int,
	"score": int
}

**Success 200 Response**
AuditorReview updated successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Delete AuditorReview by Id**
Deletes a auditorReview using the provided id

**/auditorReview/id/:id** [DELETE]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
AuditorReview deleted successfully

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Create Pet**
Creates a pet on the database

**/pet** [POST]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"refugeId": int,
	"familyId": int,
	"breed": "string",
	"color": "string",
	"age": int
}

**Success 200 Response**
Pet created successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get Pet list**
Gets a list of every pet on the database

**/pet/list** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"refugeId": int,
	"familyId": int,
	"breed": "string",
	"color": "string",
	"age": int
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get Pet by Id**
Returns a pet using the provided id

**/pet/id/:id** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"refugeId": int,
	"familyId": int,
	"breed": "string",
	"color": "string",
	"age": int
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Update Pet by Id**
Updates a pet using the provided id

**/pet/id/:id** [PUT]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"refugeId": int,
	"familyId": int,
	"breed": "string",
	"color": "string",
	"age": int
}

**Success 200 Response**
Pet updated successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Delete Pet by Id**
Deletes a pet using the provided id

**/pet/id/:id** [DELETE]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
Pet deleted successfully

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Create PetPhoto**
Creates a petPhoto on the database

**/petPhoto** [POST]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"idPet": int,
	"idPhoto": int
}

**Success 200 Response**
PetPhoto created successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get PetPhoto list**
Gets a list of every petPhoto on the database

**/petPhoto/list** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"idPet": int,
	"idPhoto": int
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Get PetPhoto by Id**
Returns a petPhoto using the provided id

**/petPhoto/id/:id** [GET]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
{
	"id": int,
	"idPet": int,
	"idPhoto": int
}

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Update PetPhoto by Id**
Updates a petPhoto using the provided id

**/petPhoto/id/:id** [PUT]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
{
	"idPet": int,
	"idPhoto": int
}

**Success 200 Response**
PetPhoto updated successfully

**Bad Request 400 Response**
Request body is wrong

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
**Delete PetPhoto by Id**
Deletes a petPhoto using the provided id

**/petPhoto/id/:id** [DELETE]

**URL Parameters**
-

**Headers**
Authorization: Token

**Request Body**
-

**Success 200 Response**
PetPhoto deleted successfully

**Bad Request 400 Response**
-

**Unauthorized 401 Response**
Unauthorized

**Internal Server Error 500 Response**
Server error

-----------------------------------------------------
