// Package shared defines the cross-platform Parser contract and snapshot types
// used by every w_popularity parser module and the backend.
package shared

import (
	"context"
	"errors"
	"time"
)

// Platform identifies a social network.
type Platform string

const (
	PlatformYouTube       Platform = "youtube"
	PlatformX             Platform = "x"
	PlatformTelegram      Platform = "telegram"
	PlatformFacebook      Platform = "facebook"
	PlatformInstagram     Platform = "instagram"
	PlatformLinkedIn      Platform = "linkedin"
	PlatformHabr          Platform = "habr"
	PlatformStackOverflow Platform = "stackoverflow"
	PlatformTBankPulse    Platform = "tbank_pulse"
	PlatformSmartLab      Platform = "smartlab"
)

// PostKind narrows a Post by its native form on the platform.
type PostKind string

const (
	PostKindPost  PostKind = "post"
	PostKindVideo PostKind = "video"
	PostKindShort PostKind = "short"
	PostKindStory PostKind = "story"
	PostKindReel  PostKind = "reel"
)

// ChannelSnapshot is one point-in-time measurement of a channel's audience.
// Backends persist these as append-only rows; deltas are computed downstream.
type ChannelSnapshot struct {
	Platform      Platform               `json:"platform"`
	Handle        string                 `json:"handle"`
	URL           string                 `json:"url"`
	FetchedAt     time.Time              `json:"fetched_at"`
	Followers     int64                  `json:"followers"`
	PostsCount    int64                  `json:"posts_count"`
	TotalLikes    int64                  `json:"total_likes,omitempty"`
	TotalViews    int64                  `json:"total_views,omitempty"`
	TotalComments int64                  `json:"total_comments,omitempty"`
	Raw           map[string]interface{} `json:"raw,omitempty"`
}

// PostSnapshot is one measurement of a single post's engagement at FetchedAt.
type PostSnapshot struct {
	Platform     Platform               `json:"platform"`
	ChannelHandle string                `json:"channel_handle"`
	PostID       string                 `json:"post_id"`
	URL          string                 `json:"url"`
	Kind         PostKind               `json:"kind"`
	PublishedAt  time.Time              `json:"published_at"`
	FetchedAt    time.Time              `json:"fetched_at"`
	Likes        int64                  `json:"likes"`
	Views        int64                  `json:"views"`
	Comments     int64                  `json:"comments"`
	Shares       int64                  `json:"shares,omitempty"`
	Raw          map[string]interface{} `json:"raw,omitempty"`
}

// Parser is the contract every per-platform module satisfies.
// Implementations should be stateless across calls; configuration is injected
// at construction time. All methods must respect ctx cancellation.
type Parser interface {
	Platform() Platform

	// FetchChannel returns the latest snapshot for handle.
	FetchChannel(ctx context.Context, handle string) (ChannelSnapshot, error)

	// FetchRecentPosts returns posts published since `since`. May return more
	// than requested by the caller; ordering is newest-first.
	FetchRecentPosts(ctx context.Context, handle string, since time.Time) ([]PostSnapshot, error)
}

// Sentinel errors. Parsers should wrap with fmt.Errorf("...: %w", err).
var (
	ErrNotImplemented = errors.New("parser: not implemented")
	ErrNotFound       = errors.New("parser: handle not found")
	ErrRateLimited    = errors.New("parser: rate limited")
	ErrAuth           = errors.New("parser: auth failed")
	ErrTransient      = errors.New("parser: transient failure")
)
