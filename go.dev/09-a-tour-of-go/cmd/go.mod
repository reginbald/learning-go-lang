module reginbald.com/cmd/main

go 1.19

replace reginbald.com/internal/loopsAndFunctions => ../internal/loopsAndFunctions
replace reginbald.com/internal/slices => ../internal/slices
replace reginbald.com/internal/maps => ../internal/maps
replace reginbald.com/internal/fibonacci => ../internal/fibonacci
replace reginbald.com/internal/stringers => ../internal/stringers
replace reginbald.com/internal/errors => ../internal/errors
replace reginbald.com/internal/readers => ../internal/readers
replace reginbald.com/internal/rot13Reader => ../internal/rot13Reader
replace reginbald.com/internal/images => ../internal/images
replace reginbald.com/internal/linkedList => ../internal/linkedList
replace reginbald.com/internal/tree => ../internal/tree
replace reginbald.com/internal/webCrawler => ../internal/webCrawler

require (
	reginbald.com/internal/loopsAndFunctions v0.0.0-00010101000000-000000000000
	reginbald.com/internal/slices v0.0.0-00010101000000-000000000000
	reginbald.com/internal/maps v0.0.0-00010101000000-000000000000
	reginbald.com/internal/fibonacci v0.0.0-00010101000000-000000000000
	reginbald.com/internal/stringers v0.0.0-00010101000000-000000000000
	reginbald.com/internal/errors v0.0.0-00010101000000-000000000000
	reginbald.com/internal/readers v0.0.0-00010101000000-000000000000
	reginbald.com/internal/rot13Reader v0.0.0-00010101000000-000000000000
	reginbald.com/internal/images v0.0.0-00010101000000-000000000000
	reginbald.com/internal/linkedList v0.0.0-00010101000000-000000000000
	reginbald.com/internal/tree v0.0.0-00010101000000-000000000000
	reginbald.com/internal/webCrawler v0.0.0-00010101000000-000000000000
)

require golang.org/x/tour v0.1.0 // indirect
