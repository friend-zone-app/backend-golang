// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package customTypes

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddAutomationInput struct {
	Title       string         `json:"title" bson:"title"`
	Location    []*float64     `json:"location" bson:"location"`
	Description string         `json:"description" bson:"description"`
	StartsAt    time.Time      `json:"startsAt" bson:"startsAt"`
	EndsAt      *time.Time     `json:"endsAt,omitempty" bson:"endsAt"`
	CreatedAt   time.Time      `json:"createdAt" bson:"createdAt"`
	Inviters    []*InviteInput `json:"inviters,omitempty" bson:"inviters"`
	Repeat      []WeekDays     `json:"repeat,omitempty" bson:"repeat"`
}

type AddBadgeInput struct {
	Type    BadgeType `json:"type" bson:"type"`
	Value   *string   `json:"value,omitempty" bson:"value"`
	Enabled bool      `json:"enabled" bson:"enabled"`
}

// Address
type Address struct {
	AddressLine       string        `json:"addressLine" bson:"addressLine"`
	AdminDistrict     string        `json:"adminDistrict" bson:"adminDistrict"`
	AdminDistrict2    string        `json:"adminDistrict2" bson:"adminDistrict2"`
	CountryRegion     string        `json:"countryRegion" bson:"countryRegion"`
	FormattedAddress  string        `json:"formattedAddress" bson:"formattedAddress"`
	Intersection      *Interception `json:"intersection" bson:"intersection"`
	Locality          string        `json:"locality" bson:"locality"`
	PostalCode        string        `json:"postalCode" bson:"postalCode"`
	CountryRegionIso2 string        `json:"countryRegionIso2" bson:"countryRegionIso2"`
}

type Badge struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Type    BadgeType          `json:"type" bson:"type"`
	Value   *string            `json:"value,omitempty" bson:"value"`
	Enabled bool               `json:"enabled" bson:"enabled"`
	Focused bool               `json:"focused" bson:"focused"`
}

type CreateEventInput struct {
	Title       string         `json:"title" bson:"title"`
	Description string         `json:"description" bson:"description"`
	Location    []*float64     `json:"location" bson:"location"`
	StartsAt    time.Time      `json:"startsAt" bson:"startsAt"`
	EndsAt      *time.Time     `json:"endsAt,omitempty" bson:"endsAt"`
	Inviters    []*InviteInput `json:"inviters,omitempty" bson:"inviters"`
	Type        EventType      `json:"type" bson:"type"`
}

type CreatePostInput struct {
	Title     string     `json:"title" bson:"title"`
	Time      float64    `json:"time" bson:"time"`
	Inviters  []string   `json:"inviters,omitempty" bson:"inviters"`
	Cancelled *bool      `json:"cancelled,omitempty" bson:"cancelled"`
	ExpireIn  *float64   `json:"expireIn,omitempty" bson:"expireIn"`
	Location  []*float64 `json:"location" bson:"location"`
}

// date
type Event struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Author      primitive.ObjectID `json:"author" bson:"author"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Location    *Point             `json:"location" bson:"location"`
	StartsAt    time.Time          `json:"startsAt" bson:"startsAt"`
	EndsAt      *time.Time         `json:"endsAt,omitempty" bson:"endsAt"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	Inviters    []*Invite          `json:"inviters,omitempty" bson:"inviters"`
	Type        EventType          `json:"type" bson:"type"`
	Privacy     *EventPrivacy      `json:"privacy,omitempty" bson:"privacy"`
}

type EventAutomation struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Location    *Point             `json:"location" bson:"location"`
	Description string             `json:"description" bson:"description"`
	StartsAt    time.Time          `json:"startsAt" bson:"startsAt"`
	EndsAt      *time.Time         `json:"endsAt,omitempty" bson:"endsAt"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	Inviters    []*Invite          `json:"inviters,omitempty" bson:"inviters"`
	Repeat      []WeekDays         `json:"repeat,omitempty" bson:"repeat"`
}

type EventPrivacy struct {
	WhoCanJoin WhoCanJoin `json:"whoCanJoin" bson:"whoCanJoin"`
	WhoCanSee  WhoCanSee  `json:"whoCanSee" bson:"whoCanSee"`
}

type FriendReq struct {
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	Status    bool               `json:"status" bson:"status"`
	User      primitive.ObjectID `json:"user" bson:"user"`
}

// Geocodepoint
type GeocodePoint struct {
	Type              string    `json:"type" bson:"type"`
	Coordinates       []float64 `json:"coordinates" bson:"coordinates"`
	CalculationMethod string    `json:"calculationMethod" bson:"calculationMethod"`
	UsageTypes        []string  `json:"usageTypes" bson:"usageTypes"`
}

// Interception
type Interception struct {
	BaseStreet       string `json:"baseStreet" bson:"baseStreet"`
	SecondaryStreet1 string `json:"secondaryStreet1" bson:"secondaryStreet1"`
	IntersectionType string `json:"intersectionType" bson:"intersectionType"`
	DisplayName      string `json:"displayName" bson:"displayName"`
}

type Invite struct {
	Status    bool               `json:"status" bson:"status"`
	InviteTo  primitive.ObjectID `json:"inviteTo" bson:"inviteTo"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

type InviteInput struct {
	InviteTo string `json:"inviteTo" bson:"inviteTo"`
}

// Location
type Location struct {
	Bbox         []float64     `json:"bbox" bson:"bbox"`
	Name         string        `json:"name" bson:"name"`
	Point        *Point        `json:"point" bson:"point"`
	Address      *Address      `json:"address" bson:"address"`
	Confidence   string        `json:"confidence" bson:"confidence"`
	EntityType   string        `json:"entityType" bson:"entityType"`
	GeocodePoint *GeocodePoint `json:"geocodePoint" bson:"geocodePoint"`
	MatchCodes   []string      `json:"matchCodes" bson:"matchCodes"`
}

// Point
type Point struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type Post struct {
	ParentEvent    primitive.ObjectID `json:"parentEvent" bson:"parentEvent"`
	Author         primitive.ObjectID `json:"author" bson:"author"`
	Reactions      []*Reaction        `json:"reactions,omitempty" bson:"reactions"`
	ImageURI       string             `json:"imageUri" bson:"imageUri"`
	CurrentEmotion ReactionType       `json:"currentEmotion" bson:"currentEmotion"`
	Title          *string            `json:"title,omitempty" bson:"title"`
}

type Privacy struct {
	ShareLocation  UserPrivacy `json:"shareLocation" bson:"shareLocation"`
	ReactionOnPost UserPrivacy `json:"reactionOnPost" bson:"reactionOnPost"`
	JoinPost       UserPrivacy `json:"joinPost" bson:"joinPost"`
	FriendRequest  UserPrivacy `json:"friendRequest" bson:"friendRequest"`
}

type PrivacyInput struct {
	ShareLocation  *UserPrivacy `json:"shareLocation,omitempty" bson:"shareLocation"`
	ReactionOnPost *UserPrivacy `json:"reactionOnPost,omitempty" bson:"reactionOnPost"`
	JoinPost       *UserPrivacy `json:"joinPost,omitempty" bson:"joinPost"`
	FriendRequest  *UserPrivacy `json:"friendRequest,omitempty" bson:"friendRequest"`
}

type Reaction struct {
	ReactionType ReactionType       `json:"reactionType" bson:"reactionType"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	Author       primitive.ObjectID `json:"author" bson:"author"`
}

type ReactionTypeInput struct {
	Type ReactionType `json:"type" bson:"type"`
}

type RemoveAutomationInput struct {
	InputID string `json:"inputId" bson:"inputId"`
}

type RemoveEventInput struct {
	EventID string `json:"eventId" bson:"eventId"`
}

type RemovePostInput struct {
	Username string `json:"username" bson:"username"`
	ID       string `json:"id" bson:"_id"`
}

type Setting struct {
	Privacy         *Privacy           `json:"privacy" bson:"privacy"`
	ColorMode       ColorMode          `json:"colorMode" bson:"colorMode"`
	EventAutomation []*EventAutomation `json:"eventAutomation,omitempty" bson:"eventAutomation"`
}

type SettingInput struct {
	Privacy   *PrivacyInput `json:"privacy,omitempty" bson:"privacy"`
	ColorMode *ColorMode    `json:"colorMode,omitempty" bson:"colorMode"`
}

// Tokens
type Tokens struct {
	AccessToken  string `json:"accessToken" bson:"accessToken"`
	RefreshToken string `json:"refreshToken" bson:"refreshToken"`
}

type UpdateBadgeInput struct {
	Query   string  `json:"query" bson:"query"`
	Value   *string `json:"value,omitempty" bson:"value"`
	Enabled *bool   `json:"enabled,omitempty" bson:"enabled"`
}

type UpdateEventInput struct {
	EventID     string         `json:"eventId" bson:"eventId"`
	Title       *string        `json:"title,omitempty" bson:"title"`
	Description *string        `json:"description,omitempty" bson:"description"`
	Location    []*float64     `json:"location,omitempty" bson:"location"`
	StartsAt    *time.Time     `json:"startsAt,omitempty" bson:"startsAt"`
	EndsAt      *time.Time     `json:"endsAt,omitempty" bson:"endsAt"`
	Inviters    []*InviteInput `json:"inviters,omitempty" bson:"inviters"`
	Type        *EventType     `json:"type,omitempty" bson:"type"`
}

type UpdatePostInput struct {
	Query     string   `json:"query" bson:"query"`
	Author    *string  `json:"author,omitempty" bson:"author"`
	Title     *string  `json:"title,omitempty" bson:"title"`
	Time      *float64 `json:"time,omitempty" bson:"time"`
	Inviters  []string `json:"inviters,omitempty" bson:"inviters"`
	Cancelled *bool    `json:"cancelled,omitempty" bson:"cancelled"`
	ExpireIn  *float64 `json:"expireIn,omitempty" bson:"expireIn"`
}

type UpdateSettingInput struct {
	Privacy   *PrivacyInput `json:"privacy,omitempty" bson:"privacy"`
	ColorMode *ColorMode    `json:"colorMode,omitempty" bson:"colorMode"`
}

type UpdateUserArgs struct {
	Query        string  `json:"query" bson:"query"`
	Username     *string `json:"username,omitempty" bson:"username"`
	Password     *string `json:"password,omitempty" bson:"password"`
	Email        *string `json:"email,omitempty" bson:"email"`
	Biography    *string `json:"biography,omitempty" bson:"biography"`
	Avatar       *string `json:"avatar,omitempty" bson:"avatar"`
	RefreshToken *string `json:"refreshToken,omitempty" bson:"refreshToken"`
}

// user
type User struct {
	ID          primitive.ObjectID   `json:"_id" bson:"_id"`
	DisplayName string               `json:"displayName" bson:"displayName"`
	Username    string               `json:"username" bson:"username"`
	Email       string               `json:"email" bson:"email"`
	Biography   *string              `json:"biography,omitempty" bson:"biography"`
	Avatar      *string              `json:"avatar,omitempty" bson:"avatar"`
	CreatedAt   time.Time            `json:"createdAt" bson:"createdAt"`
	Events      []primitive.ObjectID `json:"events,omitempty" bson:"events"`
	Posts       []primitive.ObjectID `json:"posts,omitempty" bson:"posts"`
	FriendsReq  []*FriendReq         `json:"friendsReq,omitempty" bson:"friendsReq"`
	FriendList  []primitive.ObjectID `json:"friendList,omitempty" bson:"friendList"`
	Badges      []*Badge             `json:"badges,omitempty" bson:"badges"`
	Level       int                  `json:"level" bson:"level"`
	Setting     *Setting             `json:"setting" bson:"setting"`
}

type UserRes struct {
	User  *User   `json:"user" bson:"user"`
	Token *Tokens `json:"token" bson:"token"`
}

type BadgeType string

const (
	BadgeTypeSpecial BadgeType = "SPECIAL"
	BadgeTypeRole    BadgeType = "ROLE"
	BadgeTypeDaily   BadgeType = "DAILY"
)

var AllBadgeType = []BadgeType{
	BadgeTypeSpecial,
	BadgeTypeRole,
	BadgeTypeDaily,
}

func (e BadgeType) IsValid() bool {
	switch e {
	case BadgeTypeSpecial, BadgeTypeRole, BadgeTypeDaily:
		return true
	}
	return false
}

func (e BadgeType) String() string {
	return string(e)
}

func (e *BadgeType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BadgeType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BadgeType", str)
	}
	return nil
}

func (e BadgeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ColorMode string

const (
	ColorModeLight ColorMode = "LIGHT"
	ColorModeDark  ColorMode = "DARK"
)

var AllColorMode = []ColorMode{
	ColorModeLight,
	ColorModeDark,
}

func (e ColorMode) IsValid() bool {
	switch e {
	case ColorModeLight, ColorModeDark:
		return true
	}
	return false
}

func (e ColorMode) String() string {
	return string(e)
}

func (e *ColorMode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ColorMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ColorMode", str)
	}
	return nil
}

func (e ColorMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EventType string

const (
	EventTypeSchool      EventType = "SCHOOL"
	EventTypeWork        EventType = "WORK"
	EventTypeParty       EventType = "PARTY"
	EventTypeFestival    EventType = "FESTIVAL"
	EventTypeDate        EventType = "DATE"
	EventTypeBirthday    EventType = "BIRTHDAY"
	EventTypeCelebration EventType = "CELEBRATION"
)

var AllEventType = []EventType{
	EventTypeSchool,
	EventTypeWork,
	EventTypeParty,
	EventTypeFestival,
	EventTypeDate,
	EventTypeBirthday,
	EventTypeCelebration,
}

func (e EventType) IsValid() bool {
	switch e {
	case EventTypeSchool, EventTypeWork, EventTypeParty, EventTypeFestival, EventTypeDate, EventTypeBirthday, EventTypeCelebration:
		return true
	}
	return false
}

func (e EventType) String() string {
	return string(e)
}

func (e *EventType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventType", str)
	}
	return nil
}

func (e EventType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ReactionType string

const (
	ReactionTypeHappy     ReactionType = "HAPPY"
	ReactionTypeHaha      ReactionType = "HAHA"
	ReactionTypeOmg       ReactionType = "OMG"
	ReactionTypeWhat      ReactionType = "WHAT"
	ReactionTypeConfuse   ReactionType = "CONFUSE"
	ReactionTypeSurprised ReactionType = "SURPRISED"
)

var AllReactionType = []ReactionType{
	ReactionTypeHappy,
	ReactionTypeHaha,
	ReactionTypeOmg,
	ReactionTypeWhat,
	ReactionTypeConfuse,
	ReactionTypeSurprised,
}

func (e ReactionType) IsValid() bool {
	switch e {
	case ReactionTypeHappy, ReactionTypeHaha, ReactionTypeOmg, ReactionTypeWhat, ReactionTypeConfuse, ReactionTypeSurprised:
		return true
	}
	return false
}

func (e ReactionType) String() string {
	return string(e)
}

func (e *ReactionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReactionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReactionType", str)
	}
	return nil
}

func (e ReactionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserPrivacy string

const (
	UserPrivacyFriends  UserPrivacy = "FRIENDS"
	UserPrivacyEveryone UserPrivacy = "EVERYONE"
	UserPrivacyNobody   UserPrivacy = "NOBODY"
)

var AllUserPrivacy = []UserPrivacy{
	UserPrivacyFriends,
	UserPrivacyEveryone,
	UserPrivacyNobody,
}

func (e UserPrivacy) IsValid() bool {
	switch e {
	case UserPrivacyFriends, UserPrivacyEveryone, UserPrivacyNobody:
		return true
	}
	return false
}

func (e UserPrivacy) String() string {
	return string(e)
}

func (e *UserPrivacy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserPrivacy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserPrivacy", str)
	}
	return nil
}

func (e UserPrivacy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WeekDays string

const (
	WeekDaysMonday    WeekDays = "MONDAY"
	WeekDaysTuesday   WeekDays = "TUESDAY"
	WeekDaysWednesday WeekDays = "WEDNESDAY"
	WeekDaysThursday  WeekDays = "THURSDAY"
	WeekDaysFriday    WeekDays = "FRIDAY"
	WeekDaysSatursday WeekDays = "SATURSDAY"
	WeekDaysSunday    WeekDays = "SUNDAY"
)

var AllWeekDays = []WeekDays{
	WeekDaysMonday,
	WeekDaysTuesday,
	WeekDaysWednesday,
	WeekDaysThursday,
	WeekDaysFriday,
	WeekDaysSatursday,
	WeekDaysSunday,
}

func (e WeekDays) IsValid() bool {
	switch e {
	case WeekDaysMonday, WeekDaysTuesday, WeekDaysWednesday, WeekDaysThursday, WeekDaysFriday, WeekDaysSatursday, WeekDaysSunday:
		return true
	}
	return false
}

func (e WeekDays) String() string {
	return string(e)
}

func (e *WeekDays) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WeekDays(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WeekDays", str)
	}
	return nil
}

func (e WeekDays) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WhoCanJoin string

const (
	WhoCanJoinLocal    WhoCanJoin = "LOCAL"
	WhoCanJoinInvite   WhoCanJoin = "INVITE"
	WhoCanJoinApproval WhoCanJoin = "APPROVAL"
)

var AllWhoCanJoin = []WhoCanJoin{
	WhoCanJoinLocal,
	WhoCanJoinInvite,
	WhoCanJoinApproval,
}

func (e WhoCanJoin) IsValid() bool {
	switch e {
	case WhoCanJoinLocal, WhoCanJoinInvite, WhoCanJoinApproval:
		return true
	}
	return false
}

func (e WhoCanJoin) String() string {
	return string(e)
}

func (e *WhoCanJoin) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WhoCanJoin(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WhoCanJoin", str)
	}
	return nil
}

func (e WhoCanJoin) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WhoCanSee string

const (
	WhoCanSeeFriends              WhoCanSee = "FRIENDS"
	WhoCanSeeFriendsOfFriends     WhoCanSee = "FRIENDS_OF_FRIENDS"
	WhoCanSeeEveryone             WhoCanSee = "EVERYONE"
	WhoCanSeeParticipants         WhoCanSee = "PARTICIPANTS"
	WhoCanSeeOnlyFriends          WhoCanSee = "ONLY_FRIENDS"
	WhoCanSeeOnlyFriendsOfFriends WhoCanSee = "ONLY_FRIENDS_OF_FRIENDS"
)

var AllWhoCanSee = []WhoCanSee{
	WhoCanSeeFriends,
	WhoCanSeeFriendsOfFriends,
	WhoCanSeeEveryone,
	WhoCanSeeParticipants,
	WhoCanSeeOnlyFriends,
	WhoCanSeeOnlyFriendsOfFriends,
}

func (e WhoCanSee) IsValid() bool {
	switch e {
	case WhoCanSeeFriends, WhoCanSeeFriendsOfFriends, WhoCanSeeEveryone, WhoCanSeeParticipants, WhoCanSeeOnlyFriends, WhoCanSeeOnlyFriendsOfFriends:
		return true
	}
	return false
}

func (e WhoCanSee) String() string {
	return string(e)
}

func (e *WhoCanSee) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WhoCanSee(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WhoCanSee", str)
	}
	return nil
}

func (e WhoCanSee) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
