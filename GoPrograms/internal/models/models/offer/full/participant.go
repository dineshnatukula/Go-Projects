package offer

type ParticipantType int
type ParticipantSide int

type ParticipantTeamData struct {
	ParticipantSide `json:"participantSide,omitempty"` // enum type which contains information about participant side (UNKNOWN|HOME|AWAY)
}

type Participant struct {
	ID                    ParticipantID          `json:"id"`                              // unique identifier of the participant
	Name                  string                 `json:"name"`                            // name of the participant
	NameTid               string                 `json:"nameTid"`                         // participant name translation id
	ShortName             string                 `json:"shortName"`                       // participant short name
	Gender                string                 `json:"gender"`                          // participant's gender
	ParticipantType       ParticipantType        `json:"participantType,omitempty"`       // type of participant - team or player
	ParticipantTeamData   *ParticipantTeamData   `json:"participantTeamData,omitempty"`   // additional properties if participant is  a team
	ParticipantPlayerData *ParticipantPlayerData `json:"participantPlayerData,omitempty"` // additional properties if participant is a player
}

type ParticipantPlayerData struct {
	Team         ParticipantID `json:"teamId,omitempty"`
	JerseyNumber string        `json:"jerseyNum,omitempty"`
}
