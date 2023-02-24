<pre style="font-size: 1.4vw;">
<p align="center">
             _                                     _ _                                    
            | |                                   | (_)               _              _    
  ____  ___ | | ____ ____   ____     ____ ___   _ | |_ ____   ____   | |_  ____  ___| |_  
 / _  |/ _ \| |/ _  |  _ \ / _  |   / ___) _ \ / || | |  _ \ / _  |  |  _)/ _  )/___)  _) 
( ( | | |_| | ( ( | | | | ( ( | |  ( (__| |_| ( (_| | | | | ( ( | |  | |_( (/ /|___ | |__ 
 \_|| |\___/|_|\_||_|_| |_|\_|| |   \____)___/ \____|_|_| |_|\_|| |   \___)____|___/ \___)
(_____|                   (_____|                           (_____|                       
</p>
</pre>
<p align="center">
<a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg">
</a>
<a href="/LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-green.svg">
</a>
</p>
<p align="center">
RESTful API of <b>Golang Coding Test</b>
</p>

# Golang Coding Test of API Guide

## 🔀 Compatible Route Endpoint
| NO | Use                           | Endpoint           | Example                                 | Action
|----|-------------------------------|--------------------|-----------------------------------------|------------
| 1  | Get All Member                | /api/members       | http://localhost:3000/api/members       | GET
| 2  | Create Member                 | /api/members       | http://localhost:3000/api/members       | POST
| 3  | Get Detail Member             | /api/members       | http://localhost:3000/api/members/:id   | Get
| 4  | Update Member                 | /api/members       | http://localhost:3000/api/members/:id   | PUT
| 5  | Delete Member                 | /api/members       | http://localhost:3000/api/members/:id   | DELETE

---

## 📖 Compatible JSON Payload Golang Coding Test of API
This is the JSON payload that's sended to Golang Conding Test of API

### 💲 Create Member JSON Payload (Raw JSON)
```js
http://localhost:3000/api/members => POST
{
    "username": "developer Laravel",
    "gender": "L",
    "skintype": "B",
    "skincolor": "Red asda"
}
```

### 💸 Get Member by ID
```js
http://localhost:3000/api/members/1 => GET
```
### 💸 Update Member JSON Payload (Raw JSON)

```js
http://localhost:3000/api/members/1 => UPDATE

{
    "username": "Member JS",
    "gender": "P",
    "skintype": "B",
    "skincolor": "Red asda"
}
```

### 💸 Delete Member by ID
```js
http://localhost:3000/api/members/1  => DELETE 
```