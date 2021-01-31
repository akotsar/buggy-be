package auth

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat/go-jwx/jwk"
)

// CognitoPoolID contains the User Pool ID.
var CognitoPoolID string

// CognitoClientID contains the User Pool Client ID.
var CognitoClientID string

// CognitoTokenKeys contains keys used for signing tokens.
var CognitoJWK *jwk.Set

// RegisterUserInput contains information necessary for registering a new user.
type RegisterUserInput struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
}

// LoginUserInput contains user credentials for authentication.
type LoginUserInput struct {
	Username string
	Password string
}

// LoginUserOutput contains results of logging in a user.
type LoginUserOutput struct {
	AccessToken  string
	TokenType    string
	ExpiresIn    int64
	RefreshToken string
}

// ValidateTokenOutput contains information about the provided token.
type ValidateTokenOutput struct {
	UserID   string
	Username string
	Token    string
}

// ChangePasswordInput contains information for changing password.
type ChangePasswordInput struct {
	Username        string
	CurrentPassword string
	NewPassword     string
	Token           string
}

type cognitoClaims struct {
	Subject  string `json:"sub,omitempty"`
	Username string `json:"username,omitempty"`
}

func (c *cognitoClaims) Valid() error {
	return nil
}

func init() {
	CognitoClientID = os.Getenv("COGNITO_POOL_CLIENT_ID")
	CognitoPoolID = os.Getenv("COGNITO_POOL_ID")

	const region = "ap-southeast-2"
	jwkURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, CognitoPoolID)
	set, err := jwk.Fetch(jwkURL)
	if err != nil {
		panic(err)
	}
	CognitoJWK = set
}

// LoginUser logs a user in.
func LoginUser(session *session.Session, request LoginUserInput) (*LoginUserOutput, error) {
	cognito := cognitoidentityprovider.New(session)
	authResult, err := cognito.AdminInitiateAuth(&cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow:   aws.String("ADMIN_USER_PASSWORD_AUTH"),
		ClientId:   aws.String(CognitoClientID),
		UserPoolId: aws.String(CognitoPoolID),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": request.Username,
			"PASSWORD": request.Password,
		}),
	})
	if err != nil {
		return nil, err
	}

	return &LoginUserOutput{
		AccessToken:  *authResult.AuthenticationResult.AccessToken,
		TokenType:    *authResult.AuthenticationResult.TokenType,
		ExpiresIn:    *authResult.AuthenticationResult.ExpiresIn,
		RefreshToken: *authResult.AuthenticationResult.RefreshToken,
	}, nil
}

// RegisterUser registers a new user.
func RegisterUser(session *session.Session, request RegisterUserInput) (string, error) {
	cognito := cognitoidentityprovider.New(session)
	signUpResponse, err := cognito.SignUp(&cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(CognitoClientID),
		Username: aws.String(request.Username),
		Password: aws.String(request.Password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("family_name"),
				Value: aws.String(request.LastName),
			},
			{
				Name:  aws.String("given_name"),
				Value: aws.String(request.FirstName),
			},
			{
				Name:  aws.String("custom:is_admin"),
				Value: aws.String("false"),
			},
		},
	})
	if err != nil {
		return "", err
	}

	_, err = cognito.AdminConfirmSignUp(&cognitoidentityprovider.AdminConfirmSignUpInput{
		UserPoolId: aws.String(CognitoPoolID),
		Username:   aws.String(request.Username),
	})
	if err != nil {
		return "", err
	}

	return *signUpResponse.UserSub, nil
}

// ValidateToken validates an authentication token.
func ValidateToken(tokenString string) (*ValidateTokenOutput, error) {
	var claims cognitoClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		keys := CognitoJWK.LookupKeyID(fmt.Sprintf("%v", token.Header["kid"]))
		if len(keys) == 0 {
			log.Printf("Invalid key in the token: %v\n", token.Header["kid"])
			return nil, nil
		}

		key, err := keys[0].Materialize()
		if err != nil {
			log.Printf("Failed to generate public key: %v\n", err)
			return nil, nil
		}

		return key, nil
	})

	if err != nil {
		log.Println("Error!", err)
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("the token is invalid")
	}

	return &ValidateTokenOutput{UserID: claims.Subject, Username: claims.Username, Token: tokenString}, nil
}

// ChangePassword changes user password.
func ChangePassword(session *session.Session, input *ChangePasswordInput) error {
	cognito := cognitoidentityprovider.New(session)
	_, err := cognito.ChangePassword(&cognitoidentityprovider.ChangePasswordInput{
		AccessToken:      aws.String(input.Token),
		PreviousPassword: aws.String(input.CurrentPassword),
		ProposedPassword: aws.String(input.NewPassword),
	})
	if err != nil {
		return err
	}

	return nil
}
