package main

func main() {
	data := []byte(`[{
  "person": {
    "name": {
      "first": "Leonid",
      "last": "Bugaev",
      "fullName": "Leonid Bugaev"
    },
    "github": {
      "handle": "buger",
      "followers": 109
    },
    "avatars": [
      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" },
      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
    ]
  },
  "company": {
    "name": "Acme"
  }
},{
"person": {
  "name": {
    "first": "Leonidd",
    "last": "Bugaevv",
    "fullName": "Leonidd Bugaevv"
  },
  "github": {
    "handle": "bugerr",
    "followers": 119
  },
  "avatars": [
    { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" },
    { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" },
    { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
  ]
},
"company": {
  "name": "Acmee"
}
},
{
"person": {
  "name": {
    "first": "Leoniddd",
    "last": "Bugaevvv",
    "fullName": "Leoniddd Bugaevvv"
  },
  "github": {
    "handle": "bugerrr",
    "followers": 129
  },
  "avatars": [
    { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" },
    { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" },
    { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" },
    { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
  ]
},
"company": {
  "name": "Acmeee"
}
}]`)

	jsonparsertest(data)
	gabstest(data)
	standard(data)
}
