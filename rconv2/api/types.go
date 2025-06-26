package api

// KickPlayer defines the request to kick a player from the server.
type KickPlayer struct {
	PlayerId string
	Reason   string
}

// TemporaryBanPlayer defines the request to temporarily ban a player.
type TemporaryBanPlayer struct {
	PlayerId  string
	Duration  int32
	Reason    string
	AdminName string
}

// RemoveTemporaryBan defines the request to remove a temporary ban.
type RemoveTemporaryBan struct {
	PlayerId string
}

// PermanentBanPlayer defines the request to permanently ban a player.
type PermanentBanPlayer struct {
	PlayerId  string
	Reason    string
	AdminName string
}

// SetVoteToKick defines the request to enable or disable vote-to-kick functionality.
type SetVoteToKick struct {
	Enabled bool
}

// ResetVoteToKickThreshold defines the request to reset the vote-to-kick threshold.
type ResetVoteToKickThreshold struct{}

// SetVoteToKickThreshold defines the request to set the vote-to-kick threshold.
type SetVoteToKickThreshold struct {
	ThresholdValue string
}

// AddBannedWords defines the request to add words to the banned words list.
type AddBannedWords struct {
	BannedWords string
}

// RemoveBannedWords defines the request to remove words from the banned words list.
type RemoveBannedWords struct {
	BannedWords string
}

// AddVipPlayer defines the request to add a VIP player.
type AddVipPlayer struct {
	PlayerId    string
	Description string
}
