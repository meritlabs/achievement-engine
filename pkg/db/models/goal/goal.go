package goal

type TaskSlug string

const (
	// Creator
	CreateWallet TaskSlug = "create-wallet"
	UnlockWallet          = "unlock-wallet"

	// FastStarter
	InviteFriends        = "invite-friends"
	ReceiveInviteRequest = "receive-invite-request"
	MineInvite           = "mine-invite"
	ConfirmInviteRequest = "confirm-invite-request"

	// MeritTycoon
	// TODO: add merit tycoon goals

	// GrowthMaster
	// TODO: add growth master goals
)

type Slug string

const (
	Creator      Slug = "creator"
	FastStarter       = "fast-starter"
	MeritTycoon       = "merit-tycoon"
	GrowthMaster      = "growth-master"
)

type Task struct {
	Slug        TaskSlug `json:"slug"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
}

type Goal struct {
	Slug        Slug   `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Tasks       []Task `json:"tasks"`
}

func GetGoals() []Goal {
	creatorTasks := []Task{
		Task{
			Name:        "Create wallet",
			Description: "",
			Slug:        CreateWallet,
		},
		Task{
			Name:        "Unlock wallet",
			Description: "Receive your activation token from a friend.",
			Slug:        UnlockWallet,
		},
	}

	creator := Goal{
		Slug:        Creator,
		Name:        "Creator",
		Description: "Create and unlock your first wallet.",
		Image:       "achi-creator",
		Tasks:       creatorTasks,
	}

	fastStarterTasks := []Task{
		Task{
			Name:        "Invite Your Friends to Merit!",
			Description: "Share your alias with your friends! Your alias can be used as their invite code when they create a wallet.",
			Slug:        InviteFriends,
		},
		Task{
			Name:        "Add a friend to your invite waitlist",
			Description: "You can populate your invite waitlist at any time, even when you don’t have available invites! Just share your alias with your friends and have them use it as their invite code when they create their Merit Wallet.",
			Slug:        ReceiveInviteRequest,
		},
		Task{
			Name:        "Receive an invite",
			Description: "Receive an invite from a friend, or mine an invite to complete this task. Invites are randomly distributed among the Merit community with every block that is mined. Keep an eye out for new invites in your wallet!",
			Slug:        MineInvite,
		},
		Task{
			Name:        "Confirm an invite request",
			Description: "Confirm the invite requests pending in your invite waitlist! Remember, invites are scarce!",
			Slug:        ConfirmInviteRequest,
		},
	}

	fastStarter := Goal{
		Slug:        FastStarter,
		Name:        "Fast Starter",
		Description: "Create and unlock your first wallet.",
		Image:       "achi-start",
		Tasks:       fastStarterTasks,
	}

	return []Goal{
		creator,
		fastStarter,
	}
}

func GetGoalForTask(slug TaskSlug) Goal {
	var goal *Goal

	for _, g := range GetGoals() {
		for _, t := range g.Tasks {
			if t.Slug == slug {
				goal = &g
				break
			}
		}

		if goal != nil {
			break
		}
	}

	return *goal
}
