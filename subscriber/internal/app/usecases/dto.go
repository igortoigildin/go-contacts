package usecases

type (
	FriendRequestDTO struct {
		Username       string
		TargetUsername string
	}

	FriendAcceptDTO struct {
		Username       string
		TargetUsername string
	}

	FriendRejectDTO struct {
		Username       string
		TargetUsername string
	}

	RemoveFriendDTO struct {
		Username       string
		FriendUsername string
	}

	ListFriendsDTO struct {
		Username string
	}
)
