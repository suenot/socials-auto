# socials-auto

Cross-cutting Go types and the `Parser` contract for [w_popularity](https://github.com/suenot/w-popularity).

Every parser module imports this. Keep it tiny — only types/interfaces consumed by both backend and parsers belong here.

```go
import shared "github.com/suenot/socials-auto"

type Parser interface {
    Platform() shared.Platform
    FetchChannel(ctx context.Context, handle string) (shared.ChannelSnapshot, error)
    FetchRecentPosts(ctx context.Context, handle string, since time.Time) ([]shared.PostSnapshot, error)
}
```

## License

MIT
