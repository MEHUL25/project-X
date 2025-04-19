package lld

type AuthenticationManager struct {
	timeToLive int
	tokens     map[string]int // tokenId -> expiryTime
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		timeToLive: timeToLive,
		tokens:     make(map[string]int),
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.tokens[tokenId] = currentTime + this.timeToLive
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	expiryTime, exists := this.tokens[tokenId]
	if exists && expiryTime > currentTime {
		this.tokens[tokenId] = currentTime + this.timeToLive
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	count := 0
	for tokenId, expiryTime := range this.tokens {
		if expiryTime > currentTime {
			count++
		} else {
			// optional: clean up expired tokens for memory efficiency
			delete(this.tokens, tokenId)
		}
	}
	return count
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
