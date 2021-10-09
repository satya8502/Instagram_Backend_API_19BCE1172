# 19BCE1172_Instagram_API<br>
This is the work of S.M. Satya Sree Narayanan, Registration Number 19BCE1172. This Backend API was developed as a part of the task assigned by Appointy for internship selection.<br><br>

**The File Main.go contains the source code and testapi.exe is the associated exe file**<br><br>
**All the key and value in the JSON FORMAT must be given only in lower-case**
<br><br>
**Operation-1: Creating a User**<br><br>
When a "POST" request for the URL "/users" comes in the function user_handle gets fired up. Inside if there is a seperate if clause that deals with the "POST" request. The details of the user to be created is given in the Body of the JSON request. IT IS NECESSARY THAT THE "KEY" IN THE KEY VALUE PAIRS IN THE JSON REQUEST BE GIVEN IN LOWER-CASE AND THE KEYS BE THE EXACT SAME KEYS  GIVEN IN THE IMAGE 
<br><br>
![image](https://user-images.githubusercontent.com/68813690/136667778-8bb3448b-39af-4eab-bcb0-c4bc867b15fd.png)
<br><br>
**Operation-2: Getting user based on user ID**<br><br>
When a get request for the URL "/users" comes in the function user_handle gets fired up. Inside if there is a seperate if clause that deals with the "GET" request. The query is passed in the URL which is extracted and the user details corresponding to that ID is fetched. 
<br><br>
**Operation-3: Creating a Post**<br><br>
When a "POST" request for the URL "/posts" comes in the function post_handle gets fired up. Inside if there is a seperate if clause that deals with the "POST" request. The details of the user to be created is given in the Body of the JSON request. The "uid" is IT IS NECESSARY THAT THE KEY VALUE PAIRS IN THE JSON REQUEST BE GIVEN IN LOWER-CASE AND THE KEYS BE THE EXACT SAME KEYS  GIVEN IN THE IMAGE 
<br><br>
![image](https://user-images.githubusercontent.com/68813690/136668161-0b79ecf1-b1bf-4305-8703-740cb8898225.png)
<br><br>

**Operation-4: Getting a Post using the ID**<br><br>
When a get request for the URL "/users" comes in the function user_handle gets fired up. Inside if there is a seperate if clause that deals with the "GET" request. The query is passed in the URL which is extracted and the post details corresponding to that ID is fetched. 
<br><br>
**Operation-4: Getting a Post of a user using the USER_ID**<br><br>
When a get request for the URL "posts/users" comes in the function user_posts_handle gets fired up. Inside if there is a seperate if clause that deals with the "GET" request. The query is passed in the URL which is extracted and the posts details of all the posts of the user corresponding to that ID is fetched. 
