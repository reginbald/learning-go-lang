module reginbald.com/cmd/main/v2

go 1.19

replace reginbald.com/internal/rot13Reader => ../internal/rot13Reader
replace reginbald.com/internal/images => ../internal/images
replace reginbald.com/internal/linkedList => ../internal/linkedList

require (
	reginbald.com/internal/images v0.0.0-00010101000000-000000000000
	reginbald.com/internal/rot13Reader v0.0.0-00010101000000-000000000000
	reginbald.com/internal/linkedList v0.0.0-00010101000000-000000000000
)

require golang.org/x/tour v0.1.0 // indirect
