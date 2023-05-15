package schemas

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"testGoGraph/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
)

var secretKey = []byte("my-secret-key")

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"userName": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"passwordHash": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
		},
		"lastTimeSignedIn": &graphql.Field{
			Type: graphql.DateTime,
		},
		"role": &graphql.Field{
			Type: roleType,
		},
	},
})

var authPayload = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AuthPayload",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"user": &graphql.Field{
				Type: userType,
			},
		},
	},
)

var users = []models.User{
	{ID: 1, UserName: "john_doe", Email: "john@example.com", PasswordHash: "5f4dcc3b5aa765d61d8327deb882cf99", CreatedAt: time.Date(2023, 05, 07, 11, 05, 51, 854880000, time.UTC), Role: &roles[0]},
	{ID: 2, UserName: "jane_smith", Email: "jane@example.com", PasswordHash: "e10adc3949ba59abbe56e057f20f883e", CreatedAt: time.Date(2023, 05, 07, 11, 05, 51, 854880000, time.UTC), Role: &roles[0]},
	{ID: 3, UserName: "joe_bloggs", Email: "joe@example.com", PasswordHash: "fcea920f7412b5da7be0cf42b8c93759", CreatedAt: time.Date(2023, 05, 07, 11, 05, 51, 854880000, time.UTC), Role: &roles[0]},
	{ID: 4, UserName: "alice_wonderland", Email: "alice@example.com", PasswordHash: "96e79218965eb72c92a549dd5a330112", CreatedAt: time.Date(2023, 05, 07, 11, 05, 51, 854880000, time.UTC), Role: &roles[1]},
}

func getUsers(params graphql.ResolveParams) (interface{}, error) {
	return users, nil
}

// GenerateJWT generates a JWT for the given user
func GenerateJWT(user *models.User) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.UserName
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Expires in 24 hours

	// Sign the token with the secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func logInUser(params graphql.ResolveParams) (interface{}, error) {
	userName, _ := params.Args["userName"].(string)
	password, _ := params.Args["password"].(string)
	hash := md5.Sum([]byte(password))
	// Convert the hash to a hexadecimal string
	hashStr := hex.EncodeToString(hash[:])
	for _, user := range users {
		if user.UserName == userName && user.PasswordHash == hashStr {
			token, err := GenerateJWT(&user)
			if err == nil {

				authPayload := &models.AuthPayload{
					Token: token, User: &user,
				}
				return authPayload, nil
			}
			return nil, errors.New("something went wrong couldn't generte token")

		}
	}
	return nil, errors.New("Author not found")
}

func regester(params graphql.ResolveParams) (interface{}, error) {
	userName, _ := params.Args["userName"].(string)
	password, _ := params.Args["password"].(string)
	email, _ := params.Args["email"].(string)

	/// Calculate the MD5 hash of a password
	hash := md5.Sum([]byte(password))
	// Convert the hash to a hexadecimal string
	hashStr := hex.EncodeToString(hash[:])
	//newUser :=
	newUser := models.User{ID: len(users) + 1, UserName: userName, Email: email, PasswordHash: hashStr, CreatedAt: time.Now().UTC()}
	users = append(users, newUser)
	// log.Fatalln(newUser)
	fmt.Println(newUser)
	return newUser, nil

}

func verifayJwt(tokenString string) (bool, error) {
	// Parse the JWT and verify the signature
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})
	if err != nil {
		print("i was here")
		//
		return false, errors.New("your token is invalid")
	}

	// Get the user ID from the JWT
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		// http.Error(w, "Invalid token", http.StatusUnauthorized)

		return false, errors.New("your token is invalid")
	}
	userID := int(claims["id"].(float64))
	exp := int64(claims["exp"].(float64))
	fmt.Print("this is the user id in the data base userID ", userID)
	if exp < time.Now().Unix() {
		return false, errors.New("token is expired")
	}
	return true, nil

}
