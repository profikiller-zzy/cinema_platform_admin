package jwts

import (
	"fmt"
	"testing"
)

func TestVerifyToken(t *testing.T) {
	standClaims, err := VerifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcl9uYW1lIjoiIiwicm9sZSI6MywiZXhwIjoxNzA5ODc2MDAxLCJpc3MiOiJwcm9maWtpbGxlciJ9.jYdYBkx7tYTw8KtUvA2PokbMecld7rhoZB4oDoruEsw")
	if err != nil {
		fmt.Println("FALL")
		return
	}
	fmt.Println(standClaims.JwtPayLoad)
}
