# Pinboard to Markdown

Loop through an RSS feed, write markdown to stdout containing links to items in the feed that are new since the last run

Designed to be used as a weekly-links blog post from the likes of [Pinboard](https://pinboard.in) in a static site generator such as [Hugo](https://gohugo.io/)

It's deliberately not quite enough to automate this to run weekly while we're still testing the software. Future work will write to date-based files and include the relevant header material. Until then, `pinboard2markdown https://feeds.pinboard.in/rss/u:wheresalice/ > blog.md` is still a useful command